![Github CI/CD](https://img.shields.io/github/workflow/status/evt/blockchain-api/Go)
![Repository Top Language](https://img.shields.io/github/languages/top/evt/blockchain-api)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/evt/blockchain-api/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/evt/blockchain-api)
![Codacy Grade](https://img.shields.io/codacy/grade/68c42124afcf456ab0c28e2d6da7e534)
![Github Repository Size](https://img.shields.io/github/repo-size/evt/blockchain-api)
![Github Open Issues](https://img.shields.io/github/issues/evt/blockchain-api)
![Lines of code](https://img.shields.io/tokei/lines/github/evt/blockchain-api)
![GitHub last commit](https://img.shields.io/github/last-commit/evt/blockchain-api)
![License](https://img.shields.io/badge/license-Apache%202-blue)
![GitHub contributors](https://img.shields.io/github/contributors/evt/blockchain-api)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

<img align="right" width="50%" src="./images/big-gopher.jpg">

# Blockchain API

API-server for blockchain indexes

## Solution notes

- :trident: clean architecture (handler->service->repository)
- :book: standard Go project layout (well, more or less :blush:)
- :cd: github CI/CD + docker compose + Makefile included
- :white_check_mark: handler tests with mocks included
- :heart: Swagger auto-generated on `make generate` included

## HOWTO

- :running_man: run app in docker-compose with `make dc`
- :test_tube: run tests with `make test`
- :coin: generate contract from `abi` file with `make abigen`
- :sunflower: run linter with `make lint`

## A picture is worth a thousand words

<img src="./images/make-run.png">

