// The MIT License (MIT)
//
// Copyright (c) 2016 winlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// The aac decoder, to decode the encoded aac frame to PCM samples.
package dec

/*
#cgo CFLAGS: -I${SRCDIR}/../fdk-aac/objs/include/fdk-aac
#cgo LDFLAGS: ${SRCDIR}/../fdk-aac/objs/lib/libfdk-aac.a
#include "aacdecoder_lib.h"

typedef struct {
	HANDLE_AACDECODER dec;
	CStreamInfo* info;
} aacdec_t;

static int aacdec_init_adts(aacdec_t* h, char* asc, int nb_asc) {
	h->dec = aacDecoder_Open(TT_MP4_ADTS, 1);
	if (!h->dec) {
		return -1;
	}

	UCHAR* uasc = (UCHAR*)asc;
	UINT unb_asc = (UINT)nb_asc;
	AAC_DECODER_ERROR err = aacDecoder_ConfigRaw(h->dec, &uasc, &unb_asc);
	if (err != AAC_DEC_OK) {
		return err;
	}

	h->info = aacDecoder_GetStreamInfo(h->dec);

	return 0;
}

static int aacdec_sample_rate(aacdec_t* h) {
	return h->info->sampleRate;
}

static int aacdec_aac_sample_rate(aacdec_t* h) {
	return h->info->aacSampleRate;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type AacDecoder struct {
	m C.aacdec_t
}

func NewAacDecoder() *AacDecoder {
	return &AacDecoder{}
}

func (v *AacDecoder) InitAdts(asc []byte) (err error) {
	p := (*C.char)(unsafe.Pointer(&asc[0]))
	pSize := C.int(len(asc))

	r := C.aacdec_init_adts(&v.m, p, pSize)

	if int(r) != 0 {
		return fmt.Errorf("init aac decoder failed, code is %d", int(r))
	}

	return nil
}

// These three members are the only really relevant ones for the user.
// The samplerate in Hz of the fully decoded PCM audio signal (after SBR processing).
func (v *AacDecoder) SampleRate() int {
	return int(C.aacdec_sample_rate(&v.m))
}

// Decoder internal members.
// sampling rate in Hz without SBR (from configuration info).
func (v *AacDecoder) AacSampleRate() int {
	return int(C.aacdec_aac_sample_rate(&v.m))
}
