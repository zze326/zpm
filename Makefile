build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o zpm-darwin-arm64

build-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o zpm-darwin-amd64

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o zpm-linux-amd64

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o zpm-windows-amd64.exe