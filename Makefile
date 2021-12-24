build: migrate clean
	mkdir bin
	go build -mod=vendor ./cmd/srv-item
	mv srv-item bin/
	cp -r ./cmd/srv-item/config ./bin
	cp ./cmd/srv-item/openapi.json ./bin

run:
	cd ./cmd/srv-item && go run .

migrate:
	cd ./cmd/srv-item && go run . migrate_expr
	cd ./cmd/srv-item && go run . migrate

clean:
	rm -rf bin

