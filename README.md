# go-fdkaac

Golang binding for lib-fdkaac(https://github.com/winlinvip/fdk-aac)

## Usage

First, get the source code:

```
go get -d github.com/winlinvip/go-fdkaac
```

Then, compile the fdk-aac:

```
cd $GOPATH/src/github.com/winlinvip/go-fdkaac &&
git clone https://github.com/winlinvip/fdk-aac.git &&
cd fdk-aac/ && bash autogen.sh && ./configure --prefix=`pwd`/objs && make && make install &&
cd ..
```

Done, for example:

*. [aac decoder](blob/master/dec/example_test.go)

Winlin 2016