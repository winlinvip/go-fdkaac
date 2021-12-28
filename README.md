# go-fdkaac

Golang binding for lib-fdkaac(https://github.com/mstorsjo/fdk-aac)

Only decoders available.

## Setup
```bash
$ git clone https://github.com/IzumiSy/go-fdkaac
$ git submodule update --init --recursive
```

#### Build and install fdkaac locally
```bash
$ apt install build-essential autoconf libtool
$ cd fdkaac-lib
$ ./autogen.sh
$ ./configure --prefix=/usr
$ make
$ make install
```

#### Build with buildkit and install locally
```bash
$ docker build --target artifacts --output type=local,dest=. .

# Manually install files under /usr
$ sudo cp artifacts/include /usr/include
$ sudo cp artifacts/lib /usr/lib
```

## Usage
Fdkaac Installation required beforehand.
```bash
$ go get github.com/IzumiSy/go-fdkaac
```

## Tests
```
$ earthly +test

# or just run go test (fdkaac installation on local required beforehand)
$ GO111MODULE=off go test ./...
```
