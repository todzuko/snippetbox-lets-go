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

### Tests
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
```sh
go test -v -run="^TestHumanDate$/^UTC$" ./cmd/web
```

check test coverage
```sh
  go test -cover ./...
```
check detailed test coverage
```sh
  go test -coverprofile=/tmp/profile.out ./...   # -covermode=count/atomic 
#  Check results: 
  go tool cover -func=/tmp/profile.out
#  or
  go tool cover -html=/tmp/profile.out
```