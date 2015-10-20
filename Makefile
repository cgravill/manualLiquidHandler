test:
	go test -v `go list ./... | grep -v internal`

.PHONY: all
