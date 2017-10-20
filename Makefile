VERSION=0.1.2

all: gom bundle build

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

build:
	gom build -o zeppelin-exporter

linux:
	GOOS=linux GOARCH=amd64 gom build -o zeppelin-exporter

fmt:
	gom exec go fmt ./...

test:
	gom exec go test -v .

dist:
	git archive --format tgz HEAD -o zeppelin-exporter-$(VERSION).tar.gz --prefix zeppelin-exporter-$(VERSION)/

clean:
	rm -rf zeppelin-exporter zeppelin-exporter-*.tar.gz

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
