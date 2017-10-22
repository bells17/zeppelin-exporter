FROM golang:1.9.1

WORKDIR /go/src/app
COPY . /go/src/app
RUN make init
