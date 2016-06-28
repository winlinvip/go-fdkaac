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

// The aac encoder, to encode the PCM samples to aac frame.
package enc

/*
#cgo CFLAGS: -I${SRCDIR}/../fdk-aac/objs/include/fdk-aac
#cgo LDFLAGS: ${SRCDIR}/../fdk-aac/objs/lib/libfdk-aac.a
#include "aacenc_lib.h"

typedef struct {
	// the encoder handler.
	HANDLE_AACENCODER enc;

	// encoder info.
	int frame_size;

	// user specified params.
	int aot;
	int channels;
	int sample_rate;
	int bitrate;
} aacenc_t;

int aacenc_init(aacenc_t* h, int aot, int channels, int sample_rate, int bitrate) {
	AACENC_ERROR err = AACENC_OK;

	h->aot = aot;
	h->channels = channels;
	h->sample_rate = sample_rate;
	h->bitrate = bitrate;

    // AACENC_TRANSMUX
    // Transport type to be used. See ::TRANSPORT_TYPE in FDK_audio.h.
    // Following types can be configured in encoder library:
    //         - 0: raw access units
    //         - 1: ADIF bitstream format
    //         - 2: ADTS bitstream format
    //         - 6: Audio Mux Elements (LATM) with muxConfigPresent = 1
    //         - 7: Audio Mux Elements (LATM) with muxConfigPresent = 0, out of band StreamMuxConfig
    //         - 10: Audio Sync Stream (LOAS)
	int trans_mux = 2; // adts

	int signaling = 0; // Implicit backward compatible signaling (default for ADIF and ADTS)
	int afterburner = 1; // 1 or 0(default)

	// -------------------------------------------------------------------------------
	//  ChannelMode           | ChCfg  | front_El      | side_El  | back_El  | lfe_El
	// -----------------------+--------+---------------+----------+----------+--------
	// MODE_1                 |      1 | SCE           |          |          |
	// MODE_2                 |      2 | CPE           |          |          |
	// MODE_1_2               |      3 | SCE, CPE      |          |          |
	// MODE_1_2_1             |      4 | SCE, CPE      |          | SCE      |
	// MODE_1_2_2             |      5 | SCE, CPE      |          | CPE      |
	// MODE_1_2_2_1           |      6 | SCE, CPE      |          | CPE      | LFE
	// MODE_1_2_2_2_1         |      7 | SCE, CPE, CPE |          | CPE      | LFE
	// -----------------------+--------+---------------+----------+----------+--------
	// MODE_7_1_REAR_SURROUND |      0 | SCE, CPE      |          | CPE, CPE | LFE
	// MODE_7_1_FRONT_CENTER  |      0 | SCE, CPE, CPE |          | CPE      | LFE
	// -------------------------------------------------------------------------------
	//  - SCE: Single Channel Element.
	//  - CPE: Channel Pair.
	//  - LFE: Low Frequency Element.
    CHANNEL_MODE mode = MODE_INVALID;
    int sce = 0, cpe = 0;
    switch (channels) {
        case 1: mode = MODE_1;       sce = 1; cpe = 0; break;
        case 2: mode = MODE_2;       sce = 0; cpe = 1; break;
        case 3: mode = MODE_1_2;     sce = 1; cpe = 1; break;
        case 4: mode = MODE_1_2_1;   sce = 2; cpe = 1; break;
        case 5: mode = MODE_1_2_2;   sce = 1; cpe = 2; break;
        case 6: mode = MODE_1_2_2_1; sce = 2; cpe = 2; break;
        default:
            return -1;
    }

    if ((err = aacEncOpen(&h->enc, 0, channels)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_AOT, aot)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_SAMPLERATE, sample_rate)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_CHANNELMODE, mode)) != AACENC_OK) {
        return err;
    }

    // Input audio data channel ordering scheme:
    //      - 0: MPEG channel ordering (e. g. 5.1: C, L, R, SL, SR, LFE). (default)
    //      - 1: WAVE file format channel ordering (e. g. 5.1: L, R, C, LFE, SL, SR).
    if ((err = aacEncoder_SetParam(h->enc, AACENC_CHANNELORDER, 1)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_BITRATE, bitrate)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_TRANSMUX, trans_mux)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_SIGNALING_MODE, signaling)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncoder_SetParam(h->enc, AACENC_AFTERBURNER, afterburner)) != AACENC_OK) {
        return err;
    }

    if ((err = aacEncEncode(h->enc, NULL, NULL, NULL, NULL)) != AACENC_OK) {
        return err;
    }

    AACENC_InfoStruct info;
    if ((err = aacEncInfo(h->enc, &info)) != AACENC_OK) {
        return err;
    }

    h->frame_size = info.frameLength;

	return err;
}

void aacenc_close(aacenc_t* h) {
	aacEncClose(&h->enc);
}

int aacenc_encode(aacenc_t*h, char* pcm, int nb_pcm, int nb_samples, char* aac, int* pnb_aac) {
	AACENC_ERROR err = AACENC_OK;

    INT iidentify = IN_AUDIO_DATA;
    INT oidentify = OUT_BITSTREAM_DATA;

	INT ibuffer_element_size = 2; // 16bits.
	INT ibuffer_size = 2 * h->channels * nb_samples;

	// The intput pcm must be resampled to fit the encoder,
	// for example, the intput is 2channels but encoder is 1channels,
	// then we should resample the intput pcm to 1channels
	// to make the intput pcm size equals to the encoder calculated size(ibuffer_size).
	if (ibuffer_size != nb_pcm) {
		return -1;
	}

	AACENC_BufDesc ibuf = {0};
	if (pcm) {
		ibuf.numBufs = 1;
		ibuf.bufs = (void**)&pcm;
		ibuf.bufferIdentifiers = &iidentify;
		ibuf.bufSizes = &ibuffer_size;
		ibuf.bufElSizes = &ibuffer_element_size;
	}

	AACENC_InArgs iargs = {0};
	if (pcm) {
		iargs.numInSamples = h->channels * nb_samples;
	} else {
		iargs.numInSamples = -1;
	}

	INT obuffer_element_size = 1;
	INT obuffer_size = *pnb_aac;

	AACENC_BufDesc obuf = {0};
	obuf.numBufs = 1;
	obuf.bufs = (void**)&aac;
	obuf.bufferIdentifiers = &oidentify;
	obuf.bufSizes = &obuffer_size;
	obuf.bufElSizes = &obuffer_element_size;

	AACENC_OutArgs oargs = {0};

	if ((err = aacEncEncode(h->enc, &ibuf, &obuf, &iargs, &oargs)) != AACENC_OK) {
		// Flush ok, no bytes to output anymore.
		if (!pcm && err == AACENC_ENCODE_EOF) {
			*pnb_aac = 0;
			return AACENC_OK;
		}
		return err;
	}

	*pnb_aac = oargs.numOutBytes;

	return err;
}

int aacenc_frame_size(aacenc_t* h) {
	return h->frame_size;
}

int aacenc_max_output_buffer_size(aacenc_t* h) {
	// The maximum packet size is 8KB aka 768 bytes per channel.
	INT obuffer_size = 8192;
	if (h->channels * 768 > obuffer_size) {
		obuffer_size = h->channels * 768;
	}
	return obuffer_size;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// The AAC encoder, input is PCM samples, output AAC frame in ADTS.
type AacEncoder struct {
	m C.aacenc_t
}

func NewAacEncoder() *AacEncoder {
	return &AacEncoder{}
}

// Initialize the encoder in LC profile.
// @remark the encoder use sampleRate and channels, user should resample the PCM to fit it,
//		that is, the channels and sampleRate of PCM should always equals to encoder's.
// @remark for the fdkaac always use 16bits sample, so the bits of pcm always 16,
//		which must be: [SHORT PCM] [SHORT PCM] ... ...
func (v *AacEncoder) InitLc(channels, sampleRate, bitrateBps int) (err error) {
	aot := int(2) // AOT_AAC_LC
	r := C.aacenc_init(&v.m, C.int(aot), C.int(channels), C.int(sampleRate), C.int(bitrateBps))
	if int(r) != 0 {
		return fmt.Errorf("Initialize encoder failed, code=%v", int(r))
	}
	return
}

func (v *AacEncoder) Close() {
	C.aacenc_close(&v.m)
}

// Encode the pcm to aac
// User must specifies the nbSamples of pcm.
// @remark fdkaac always use 16bits pcm, so the bits of pcm always 16.
// @remark user should resample the pcm to fit the encoder, so the channels of pcm equals to encoder's.
// @remark user should resample the pcm to fit the encoder, so the sampleRate of pcm equals to encoder's.
// @return when aac is nil, encoded completed(the Flush() return nil also),
//		because we will flush the encoder automatically to got the last frames.
func (v *AacEncoder) Encode(pcm []byte, nbSamples int) (aac []byte, err error) {
	// The maximum packet size is 8KB aka 768 bytes per channel.
	nbAac := int(C.aacenc_max_output_buffer_size(&v.m))
	aac = make([]byte, nbAac)

	pAac := (*C.char)(unsafe.Pointer(&aac[0]))
	pAacSize := C.int(nbAac)

	pPcm := (*C.char)(unsafe.Pointer(&pcm[0]))
	pPcmSize := C.int(len(pcm))
	pNbSamples := C.int(nbSamples)

	r := C.aacenc_encode(&v.m, pPcm, pPcmSize, pNbSamples, pAac, &pAacSize)
	if int(r) != 0 {
		return nil, fmt.Errorf("Encode failed, code=%v", int(r))
	}

	valid := int(pAacSize)

	// when got nil packet, flush encoder.
	if valid == 0 {
		return v.Flush()
	}

	return aac[0:valid],nil
}

// Flush the encoder to get the cached aac frames.
// @return when aac is nil, flush ok, should never flush anymore.
func (v *AacEncoder) Flush() (aac []byte, err error) {
	// The maximum packet size is 8KB aka 768 bytes per channel.
	nbAac := int(C.aacenc_max_output_buffer_size(&v.m))
	aac = make([]byte, nbAac)

	pAac := (*C.char)(unsafe.Pointer(&aac[0]))
	pAacSize := C.int(nbAac)

	r := C.aacenc_encode(&v.m, nil, 0, 0, pAac, &pAacSize)
	if int(r) != 0 {
		return nil, fmt.Errorf("Flush failed, code=%v", int(r))
	}

	valid := int(pAacSize)
	if valid == 0 {
		return nil, nil
	}

	return aac[0:valid],nil
}

// Get the frame size of encoder.
func (v *AacEncoder) FrameSize() int {
	return int(C.aacenc_frame_size(&v.m))
}
