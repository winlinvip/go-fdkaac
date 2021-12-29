VERSION 0.6

build:
  FROM DOCKERFILE --target build .
  SAVE ARTIFACT /fdkaac-objs/include/fdk-aac /fdkaac-include AS LOCAL ./artifact/include
  SAVE ARTIFACT /fdkaac-objs/lib /fdkaac-lib AS LOCAL ./artifact/lib
  SAVE IMAGE --push gcr.io/izumisy/fdkaac:build

test:
  FROM golang:1.16
  COPY +build/fdkaac-include /usr/include
  COPY +build/fdkaac-lib /usr/lib/fdk-aac
  COPY fdkaac /go-fdkaac
  WORKDIR /go-fdkaac
  RUN GO111MODULE=off go test ./...
