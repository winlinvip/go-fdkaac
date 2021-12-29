# go-fdkaac
Golang binding for lib-fdkaac(https://github.com/mstorsjo/fdk-aac)

Only decoders available.

## Usage
Fdkaac Installation required beforehand.
```bash
$ go get github.com/IzumiSy/go-fdkaac
```

Docker image of fdkaac is available for source of multi-stage build.
```dockerfile
COPY --from=ghcr.io/izumisy/fdkaac:latest /fdkaac-include /usr/include/fdk-aac
COPY --from=ghcr.io/izumisy/fdkaac:latest /fdkaac-lib /usr/lib/fdk-aac
```

## Installation
Earthly is required if you don't want to install build tools on your local.

```bash
$ git clone https://github.com/IzumiSy/go-fdkaac
$ sudo earthly +install # Linux only
```

### Build manually and install locally
```bash
$ apt install build-essential autoconf libtool
$ cd fdkaac-lib
$ ./autogen.sh
$ ./configure --prefix=/usr
$ make
$ make install
```

## Tests
```bash
$ earthly +test

# or just run go test (fdkaac installation on local required beforehand)
$ GO111MODULE=off go test ./...
```
