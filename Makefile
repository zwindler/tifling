prepare:
	go mod tidy

build: prepare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tifling -ldflags "-X main.Version=$$VERSION" main.go

dockerbuild:
	docker build -t zwindler/tifling:$$TAG --build-arg VERSION=$$TAG . && docker build -t zwindler/tifling:latest --build-arg VERSION=$$TAG .

dockerpush: dockerbuild
	docker push zwindler/tifling:$$TAG && docker push zwindler/tifling:latest
