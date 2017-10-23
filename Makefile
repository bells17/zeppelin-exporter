VERSION=0.1.6

all: gom ghr bundle build

init: gom ghr

gom:
	go get -u github.com/mattn/gom

ghr:
	go get -u github.com/tcnksm/ghr

bundle:
	gom install

build:
	gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/zeppelin-exporter

install:
	install zeppelin-exporter /usr/local/bin/zeppelin-exporter

fmt:
	gom exec go fmt ./...

test:
	gom exec go test -v .

build-cross:
	GOOS=linux GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/zeppelin-exporter-linux-amd64
	GOOS=darwin GOARCH=amd64 gom build -ldflags '-X main.BuildVersion=${VERSION}' -o bin/zeppelin-exporter-darwin-amd64

dist: build-cross
	cd bin && \
		tar cvf release/zeppelin-exporter-linux-amd64-${VERSION}.tar zeppelin-exporter-linux-amd64 && \
		zopfli release/zeppelin-exporter-linux-amd64-${VERSION}.tar && \
		rm release/zeppelin-exporter-linux-amd64-${VERSION}.tar
	cd bin && \
		tar cvf release/zeppelin-exporter-darwin-amd64-${VERSION}.tar zeppelin-exporter-darwin-amd64 && \
		zopfli release/zeppelin-exporter-darwin-amd64-${VERSION}.tar && \
		rm release/zeppelin-exporter-darwin-amd64-${VERSION}.tar

clean:
	rm -f bin/zeppelin-exporter*
	rm -f bin/release/zeppelin-exporter*

tag:
	git checkout master
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master

release: clean dist
	rm -f bin/release/.gitkeep && \
		ghr ${VERSION} bin/release && \
		touch bin/release/.gitkeep
