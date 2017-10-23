FROM golang:1.9.1

WORKDIR /go/src/app
COPY . /go/src/app
RUN apt-get update && \
			apt-get install zopfli && \
			make init
