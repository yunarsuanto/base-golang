pkgs = $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)
datetime = $(shell date +%s)

test:
	@echo " >> running tests"
	@go test  -cover $(pkgs)

race:
	@echo " >> running tests with race"
	@go test  -cover -race $(pkgs)

run:
	gin -p 9000 -a 8000 serve-http

install:
	@go mod download

.PHONY: test clean

grpc-go-gen: ## folder=profile make grpc-gen-feature
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/user/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/general/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/permission/*.proto
	protoc --proto_path=./proto --go-grpc_out=./ --go_out=:./ ./proto/role/*.proto
	ls handlers/*/*/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

generate: 
	go run main.go make:handler $(name)

dbStatus: 
	goose -dir migrations postgres "user=postgres password=balakutak dbname=umkm sslmode=disable" status

dbUp: 
	goose -dir migrations postgres "user=postgres password=balakutak dbname=umkm sslmode=disable" up

dbDown: 
	goose -dir migrations postgres "user=postgres password=balakutak dbname=umkm sslmode=disable" down

dbCreate:
	goose -dir migrations create create_$(name)_table sql

dbSeed: 
	go run main.go seed

genCrud: 
	go run main.go crud $(name)

deploy:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o yokila-be
	zip -r ./migrations.zip ./migrations
	zip -r ./templates.zip ./templates
	scp ./yokila-be majid@143.198.89.101:/home/majid/yokila/yokila-be/yokila-be-$(datetime)
	scp ./migrations.zip majid@143.198.89.101:/home/majid/yokila/yokila-be
	scp ./templates.zip majid@143.198.89.101:/home/majid/yokila/yokila-be
	ssh majid@143.198.89.101 "cd /home/majid/yokila/yokila-be && mkdir -p temp && unzip -o migrations.zip && unzip -o templates.zip && sudo service yokila-be stop && sudo unlink yokila-be && sudo ln -s yokila-be-$(datetime) yokila-be && sudo service yokila-be start"
