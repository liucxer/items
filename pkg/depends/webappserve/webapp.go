package webappserve

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

type WebappServerOpt struct {
	WebAppConfig string
	AppEnv       string
	WebAppRoot   string
	Port         string
}

var opt = &WebappServerOpt{}

var App = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if err := Serve(opt); err != nil {
			panic(err)
		}
	},
}

func init() {
	port := os.Getenv("SRV_ITEM_PORT")
	if port == "" {
		port = "9999"
	}
	if _, err := strconv.Atoi(port); err != nil {
		port = "9999"
	}
	config := "SRV_ITEM=,isOnline=true,port=" + port
	root := os.Getenv("SRV_ITEM_WEB_PATH")
	if root == "" {
		root, _ = os.Getwd()
	}
	root = path.Join(root, "web")
	fmt.Printf("webapp root: %v\n", root)
	App.Flags().StringVarP(&opt.WebAppRoot, "root", "", root, "app root")
	App.Flags().StringVarP(&opt.WebAppConfig, "config", "c", config, "app config")
}

func Serve(opt *WebappServerOpt) error {
	gzipHandler, err := gziphandler.GzipHandlerWithOpts(
		gziphandler.CompressionLevel(gzip.BestSpeed),
		gziphandler.ContentTypes([]string{
			"application/json",
			"application/javascript",
			"image/svg+xml",
			"text/html",
			"text/xml",
			"text/plain",
			"text/css",
			"text/*",
		}),
	)
	if err != nil {
		return err
	}

	if opt.Port == "" {
		opt.Port = "80"
	}

	srv := &http.Server{Addr: ":" + opt.Port, Handler: gzipHandler(WebappServer(opt))}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Printf("\nwebapp serve on %s \n", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println(err)
			} else {
				log.Fatalln(err)
			}
		}
	}()

	<-stopCh

	log.Printf("shutdowning in %s\n", 10*time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func WebappServer(opt *WebappServerOpt) http.Handler {
	indexHTML, err := ioutil.ReadFile(path.Join(opt.WebAppRoot, "./index.html"))
	if err != nil {
		panic(err)
	}

	appConfig := ParseAppConfig(opt.WebAppConfig)

	indexHTML = bytes.ReplaceAll(indexHTML, []byte("__ENV__"), []byte(opt.AppEnv))
	indexHTML = bytes.ReplaceAll(indexHTML, []byte("__APP_CONFIG__"), []byte(appConfig.String()))

	return &webappServer{
		indexHTML:   indexHTML,
		fileHandler: http.FileServer(http.Dir(opt.WebAppRoot)),
		corsHandler: cors.Default(),
		appConfig:   appConfig,
	}
}

type webappServer struct {
	appConfig   AppConfig
	indexHTML   []byte
	corsHandler *cors.Cors
	fileHandler http.Handler
}

func (s *webappServer) responseFromIndexHTML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", mime.TypeByExtension(".html"))

	w.Header().Set("X-Frame-Options", "sameorigin")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, bytes.NewBuffer(s.indexHTML)); err != nil {
		writeErr(w, http.StatusNotFound, err)
	}
}

func (s *webappServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	p := r.URL.Path

	if p == "/__built__/favicon.ico" {
		expires(w.Header(), 24*time.Hour)

		s.fileHandler.ServeHTTP(w, r)
		return
	}

	if p == "/sw.js" {
		s.fileHandler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(p, "/__built__/") {
		if p == "/__built__/config.json" {
			s.corsHandler.HandlerFunc(w, r)
			w.Header().Set("Content-Type", mime.TypeByExtension(".json"))
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(s.appConfig)
			return
		}

		expires(w.Header(), 30*24*time.Hour)

		s.fileHandler.ServeHTTP(w, r)
		return
	}

	s.responseFromIndexHTML(w)
}

func expires(header http.Header, d time.Duration) {
	header.Set("Cache-Control", fmt.Sprintf("max-age=%d", d/time.Second))
}

func writeErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(err.Error()))
}

func ParseAppConfig(s string) AppConfig {
	parts := strings.Split(s, ",")

	c := AppConfig{}

	for i := range parts {
		kv := strings.Split(parts[i], "=")

		if kv[0] == "" {
			continue
		}

		if len(kv) == 2 {
			c[kv[0]] = kv[1]
		} else {
			c[kv[0]] = ""
		}
	}

	return c
}

type AppConfig map[string]string

func (c AppConfig) String() string {
	keys := make([]string, 0)

	for k := range c {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	buf := bytes.NewBuffer(nil)

	for i, k := range keys {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(c[k])
	}

	return buf.String()
}
