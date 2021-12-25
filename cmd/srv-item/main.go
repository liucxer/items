package main

import (
	"sync"
	"time"

	"github.com/go-courier/courier"
	"github.com/saitofun/items/cmd/srv-item/apis"
	"github.com/saitofun/items/cmd/srv-item/global"
	"github.com/saitofun/items/pkg/depends/webappserve"
)

func main() {
	global.App.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})
	global.App.AddCommand("migrate_expr", func(args ...string) {
		global.MigrateExprs()
	})
	global.App.Execute(
		func(args ...string) {
			BatchRun(
				func() {
					courier.Run(apis.Root, global.Server())
				},
				func() {
					if err := webappserve.App.Execute(); err != nil {
						panic(err)
					}
				},
			)
		},
	)
}

func BatchRun(funcs ...func()) {
	wg := &sync.WaitGroup{}

	for i := range funcs {
		fn := funcs[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
