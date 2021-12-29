FROM debian:bullseye-slim

RUN apt-get update && apt-get -y install build-essential autoconf libtool

build:
  GIT CLONE --branch master https://github.com/mstorsjo/fdk-aac.git /fdkaac-lib
  WORKDIR /fdkaac-lib
  RUN ./autogen.sh
  RUN ./configure --prefix=/fdkaac-objs
  RUN make
  RUN make install
  SAVE ARTIFACT /fdkaac-objs/include/fdk-aac /fdkaac-include AS LOCAL ./artifact/include
  SAVE ARTIFACT /fdkaac-objs/lib /fdkaac-lib AS LOCAL ./artifact/lib

image:
  FROM +build
  SAVE IMAGE --push ghcr.io/izumisy/fdkaac:latest

install:
  LOCALLY
  COPY +build/fdkaac-include /usr/include/fdk-aac
  COPY +build/fdkaac-lib /usr/lib/fdk-aac

test:
  FROM golang:1.16
  COPY +build/fdkaac-include /usr/include
  COPY +build/fdkaac-lib /usr/lib/fdk-aac
  COPY fdkaac /go-fdkaac
  WORKDIR /go-fdkaac
  RUN GO111MODULE=off go test ./...
