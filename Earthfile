build:
    FROM DOCKERFILE --target build .
    SAVE ARTIFACT /fdkaac-objs/include/fdk-aac /fdkaac-include
    SAVE ARTIFACT /fdkaac-objs/lib /fdkaac-lib

test:
    FROM golang:1.16
    COPY +build/fdkaac-include /usr/include
    COPY +build/fdkaac-lib /usr/lib/fdk-aac
    COPY fdkaac /go-fdkaac
    WORKDIR /go-fdkaac
    RUN GO111MODULE=off go test ./...
