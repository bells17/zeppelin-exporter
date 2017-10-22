# zeppelin-exporter

The export tool for all zeppelin notebooks.

## Installation

```
go get github.com/bells17/go-zeppelin-exporter
cd $GOPATH/src/github.com/bells17/go-zeppelin-exporter
make && make install
```

## Usage

```
go-zeppelin-exporter --host 127.0.0.1 -p 8080 > notebooks.json
cat notebooks.json | jq .[0] > notebook1.json
content=`cat notebook1.json`
curl -v -H "Accept: application/json" -H "Content-type: application/json" -d $content http://127.0.0.1:8080/api/notebook/import
```

## Help

```
zeppelin-exporter --help
Usage:
  zeppelin-exporter [OPTIONS]

Application Options:
      --host=     Zeppelin host (default: 127.0.0.1)
  -p, --port=     port (default: 8080)
      --protocol= protocol (default: http)

Help Options:
  -h, --help      Show this help message
```

## Reference

- https://zeppelin.apache.org/docs/0.6.1/rest-api/rest-notebook.html#export-a-notebook
- https://zeppelin.apache.org/docs/0.6.1/rest-api/rest-notebook.html#import-a-notebook

## Development

### Requirements

- [ghr](https://github.com/tcnksm/ghr)
  - And need set your Github Token(export GITHUB_TOKEN="...")
- [zopfli](https://github.com/google/zopfli)

### Initialize

```
make init
```

### Run Testing

```
make test
```

### Build

```
make build
```

### Release

```
make tag
make release
```
