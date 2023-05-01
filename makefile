GIT_VERSION = $(shell git rev-list -1 HEAD)
ifdef RELEASE
	SAMPLE_VERSION := $(RELEASE)
else
	SAMPLE_VERSION := local
endif
LDFLAGS:="-X main.commit=$(GIT_VERSION) -X main.version=$(SAMPLE_VERSION)"

lint:
	revive -set_exit_status ./...

vet:
	go vet

fmt:
	gofmt -l -w -s .

.PHONY: build
build: conf
	GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o gogpt-linux-amd64 -buildvcs=false
	GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -o gogpt-linux-arm64 -buildvcs=false
	GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o gogpt-darwin-amd64 -buildvcs=false
	GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -o gogpt-windows-amd64.exe -buildvcs=false

conf:
ifeq (,$(wildcard gogpt.conf))
	cp ./gogpt.conf.gogpt ./gogpt.conf
endif

test-ci: conf
	cp ./gogpt.conf /tmp
	go test -v -coverprofile=coverage.out -covermode=atomic ./... | tee test.log | go-junit-report > report.xml
	gocov convert coverage.out | gocov-xml > coverage.xml
	go tool cover -func=coverage.out | tee -a test.log
	go tool cover -html=coverage.out -o coverage.html
	# remove  from report.xml
	sed -i 's///g' report.xml
	cat test.log

test: conf
	cp ./gogpt.conf /tmp
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

prepare-checkin: lint test
	go mod tidy

install:
	go build -ldflags $(LDFLAGS) -o gogpt
	chmod a+x gogpt
	sudo cp gogpt /usr/bin
