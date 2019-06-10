docker-build:
	docker build -t wof-readwrited .

fmt:
	go fmt cmd/wof-fs-reader/*.go
	go fmt http/*.go
	go fmt pruner/*.go
	go fmt reader/*.go
	go fmt utils/*.go
	go fmt writer/*.go

tools:
	go build -mod vendor -o bin/wof-fs-reader cmd/wof-fs-reader/main.go
