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
git clone https://github.com/winlinvip/fdk-aac.git fdk-aac-lib &&
cd fdk-aac-lib/ && bash autogen.sh && ./configure --prefix=`pwd`/objs && make && make install &&
cd ..
```

Done, import and use the package:

* [ExampleAacDecoder_RAW](codec/example_test.go), decode the aac frame to PCM samples.
* [ExampleAacEncoder_LC](codec/example_test.go), encode the PCM samples to aac frame.

There are an example of AAC audio packets in ADTS:

* [avatar aac over ADTS](doc/adts_data.go), user can use this file to decode to PCM then encode.

To run all examples:

```
cd $GOPATH/src/github.com/winlinvip/go-fdkaac && go test ./...
```

Winlin 2016
