# Cron Expression Parser

## How to build and run app?

```bash
go buld -o app
./app "*/15 0 1,15 * 1-5 /usr/bin/find"
```

## How to run tests

```bash
go test ./...
or with coverage
go test ./... --cover
```

## Tested with

```
go version go1.21.4 darwin/arm64
```
