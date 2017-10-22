VERSION=0.1.4

all: gom bundle build

init: gom bundle

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

build:
	gom build -ldflags '-X main.BuildVersion=${VERSION}' -o zeppelin-exporter

install:
	install zeppelin-exporter /usr/local/bin/zeppelin-exporter

fmt:
	gom exec go fmt ./...

test:
	gom exec go test -v .

build-cross: build-cross
	GOOS=linux GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/linux/amd64/zeppelin-exporter
	GOOS=darwin GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/darwin/amd64/zeppelin-exporter

dist:
	cd bin/linux/amd64/ && tar cvf zeppelin-exporter-linux-amd64-${VERSION}.tar zeppelin-exporter && zopfli zeppelin-exporter-linux-amd64-${VERSION}.tar
	cd bin/darwin/amd64/ && tar cvf zeppelin-exporter-darwin-amd64-${VERSION}.tar zeppelin-exporter && zopfli zeppelin-exporter-darwin-amd64-${VERSION}.tar

clean:
	rm -f zeppelin-exporter
	rm -f ./bin/*/*/*

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
