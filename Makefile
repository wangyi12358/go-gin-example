start-service:
	go run ./cmd/main.go
build-service:
	go build -o ./bin/$(SERVICE) ./cmd/$(SERVICE)/main.go
gen-db-model:
	go run ./script/gen_db_model.go