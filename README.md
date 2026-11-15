mkdir open-gateway
cd open-gateway
go mod init example.com/open-gateway

npx docsify-cli serve .
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5
go mod tidy
go run .
go run main.go
brew install go
brew upgrade go
go version

lsof -nP -iTCP:8080 -sTCP:LISTEN

http://localhost:8080/healthz

go tool pprof http://127.0.0.1:6060/debug/pprof/heap

安装：go install github.com/google/gops@latest

http://localhost:6060/debug/pprof/heap?debug=1
go tool pprof http://127.0.0.1:6060/debug/pprof/heap