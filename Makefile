pkgs 					= $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)
datetime			= $(shell date +%s)

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
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	# zip -r ./migrations.zip ./data/migrations
	# zip -r ./templates.zip ./templates
	# scp ./eakademik-be itb-dev@10.218.15.75:/home/itb-dev/eakademik/eakademik-be/pupr-backend-$(datetime)
	# scp ./migrations.zip itb-dev@10.218.15.75:/home/itb-dev/eakademik/eakademik-be
	# scp ./templates.zip itb-dev@10.218.15.75:/home/itb-dev/eakademik/eakademik-be
	# ssh itb-dev@10.218.15.75 "cd /home/itb-dev/eakademik/eakademik-be && mkdir -p temp && unzip -o migrations.zip && unzip -o templates.zip && sudo service eakademik-be stop && sudo unlink pupr-backend && sudo ln -s pupr-backend-$(datetime) pupr-backend && sudo service eakademik-be start"

deploy-prod:
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	# zip -r ./migrations.zip ./data/migrations
	# zip -r ./templates.zip ./templates
	# scp ./eakademik-be sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be/eakademik-be-$(datetime)
	# scp ./migrations.zip sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be
	# scp ./templates.zip sccic-app@10.218.15.71:/home/sccic-app/eakademik/eakademik-be
	# echo "cd /home/sccic-app/eakademik/eakademik-be && mkdir -p temp && unzip -o templates.zip && unzip -o migrations.zip && sudo service eakademik-be stop && sudo unlink eakademik-be && sudo ln -s eakademik-be-$(datetime) eakademik-be && sudo service eakademik-be start"
	# echo $(PUPROD_PASS)
