Snippetbox 

https://lets-go.alexedwards.net/sample/00.00-front-matter.html

### Setup
- copy `.env.dist` to `.env`
- generate self-signed tls with standard go library 
```sh
  go run {GO_ROOT_PATH}/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
- start the server
```sh
  go run ./cmd/web
```

- run tests
- `-count=1` - run without cache
```sh
  go test ./cmd/web -v
```
run specific test
```sh
  go test -v -run="^TestPing$" ./cmd/web/
```
run specific sub-test
```she
go test -v -run="^TestHumanDate$/^UTC$" ./cmd/web
```