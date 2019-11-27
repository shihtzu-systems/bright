version := $(shell cat app.version)
datestamp := $(shell cat app.datestamp)
timestamp := $(shell cat app.timestamp)

githubProject := github.com/shihtzu-systems
imageRepo := shihtzu
app := bright

org := shihtzu

stamp: out-dir
	printf `/bin/date "+%Y%m%d"` > app.datestamp
	printf `/bin/date "+%H%M%S"` > app.timestamp
	printf "$(version)" > app.version

build: fmt vet test build-binary build-version-container archive

build-binary:
	GOOS=linux  GOARCH=amd64 go build -o bin/linux_amd64/$(app)  main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin_amd64/$(app) main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

archive: out-dir
	zip -r out/$(app).zip templates static bin app.datestamp app.timestamp app.version

download-destiny: out-dir
	curl -s https://raw.githubusercontent.com/Bungie-net/api/master/openapi.json > out/openapi.json
	curl -s https://raw.githubusercontent.com/Bungie-net/api/master/openapi-2.json > out/openapi-2.json

out-dir:
	mkdir -p out

build-version-container:
	docker build . \
		-t local/$(app) \
		-t $(imageRepo)/$(app):$(version)-on.$(datestamp).at.$(timestamp) \
		-t $(imageRepo)/$(app):$(version) \
		-t $(imageRepo)/$(app):latest

git-develop-branch:
	git checkout develop

git-master-branch:
	git checkout master

git-commit:
	git add --all
	git commit

git-tag:
	git tag v$(version)

git-push-tags:
	git push origin --tags

git-push:
	git push origin

push-version-container: build-version-container
	docker push $(imageRepo)/$(app):$(version)-on.$(datestamp).at.$(timestamp)
	docker push $(imageRepo)/$(app):$(version)
	docker push $(imageRepo)/$(app):latest