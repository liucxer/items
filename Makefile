build: clean
	mkdir bin
	go build -mod=vendor ./cmd/srv-item
	mv srv-item bin/
	cp -r ./cmd/srv-item/config ./bin
	cp ./cmd/srv-item/openapi.json ./bin
	cp -r ./web bin/

run:
	cd ./cmd/srv-item && go run .

migrate:
	go run ./cmd/srv-item/main.go migrate_expr
	go run ./cmd/srv-item/main.go migrate

clean:
	rm -rf bin

