build:
	go-bindata -o data.go ./typescript
	go build -o done .
