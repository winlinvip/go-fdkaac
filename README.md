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

Done, use the `go-fdkaac/dec`:

```
package main

import (
	"fmt"
	"github.com/winlinvip/go-fdkaac/dec"
)

func main() {
	fmt.Println("AAC to PCM")

	d := dec.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitAdts(asc); err != nil {
		fmt.Println("init adts aac decoder failed, err is", err)
		return
	}

	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
}
```

Winlin 2016