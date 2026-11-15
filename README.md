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