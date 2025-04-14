Snippetbox 

https://lets-go.alexedwards.net/sample/00.00-front-matter.html

### Setup
- copy `.env.dist` to `.env`
- generate self-signed tls with standard go library 
```sh
  go run {GO_ROOT_PATH}/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
-start the server
```sh
  go run ./cmd/web
```