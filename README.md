# tg-example

## About
Simple example usage tg.
## Update transport
```bash
go generate ./...
```
## Run servcie
```bash
go run ./cmd/adder/main.go
```
## API Example
Request:
```bash
curl -X POST "http://localhost:9000/" -H "Content-Type: application/json" -d '[{"id":"1", "jsonrpc":"2.0", "method":"adder.sumTwoNumbers", "params":{"first":1, "second":2}}]'
```
Response:
```bash
[{"id":"1","jsonrpc":"2.0","result":{"result":3}}]
```