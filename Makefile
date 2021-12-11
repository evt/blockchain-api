scrutinizer:
# default target that does nothing to make scrutinizer happy

generate:
	go generate -x ./...

build:
	go build -race cmd/app/main.go

test:
	go test -v ./...

abigen:
	abigen --abi=./contract/contract.abi --pkg contract --out ./contract/contract.go

dc:
	docker-compose up  --remove-orphans --build

run:
	./run_app.sh

lint:
	golangci-lint run --fix
