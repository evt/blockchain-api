#!/bin/bash
HTTP_ADDR=:9090 \
INFURA_ENDPOINT=https://ropsten.infura.io/v3/101df4c66df249d1ae14489c8af62eb3 \
CONTRACT_ADDRESS=0x4f7f1380239450AAD5af611DB3c3c1bb51049c29 \
go run -race cmd/app/main.go
