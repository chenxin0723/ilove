SHELL = /bin/bash

build:
	@GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -tags 'bindatafs' -o ilove main.go
	@mkdir -p admin_auth
	@cp -r $$GOPATH/src/github.com/theplant/ec/admin_auth/views admin_auth/views
	@docker build -t ilovedocker .
	@rm -r admin_auth
	@rm ilove

push: build
	@$(eval REV := $(shell git rev-parse HEAD|cut -c 1-6))
	@docker tag ilovedocker xindocker/ilove:$(REV)
	@docker push xindocker/ilove:$(REV)
