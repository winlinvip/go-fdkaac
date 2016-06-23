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
	h->dec = NULL;
	h->info = NULL;

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

static void aacdec_close(aacdec_t* h) {
	if (h->dec) {
		aacDecoder_Close(h->dec);
	}
	h->dec = NULL;
}

static int aacdec_fill(aacdec_t* h, char* data, int nb_data, int* pnb_left) {
	UCHAR* udata = (UCHAR*)data;
	UINT unb_data = (UINT)nb_data;
	UINT unb_left = 0;
	AAC_DECODER_ERROR err = aacDecoder_Fill(h->dec, &udata, &unb_data, &unb_left);
	if (err != AAC_DEC_OK) {
		return err;
	}

	if (pnb_left) {
		*pnb_left = (int)unb_left;
	}

	return 0;
}

static int aacdec_sample_rate(aacdec_t* h) {
	return h->info->sampleRate;
}

static int aacdec_frame_size(aacdec_t* h) {
	return h->info->frameSize;
}

static int aacdec_num_channels(aacdec_t* h) {
	return h->info->numChannels;
}

static int aacdec_aac_sample_rate(aacdec_t* h) {
	return h->info->aacSampleRate;
}

static int aacdec_profile(aacdec_t* h) {
	return h->info->profile;
}

static int aacdec_audio_object_type(aacdec_t* h) {
	return h->info->aot;
}

static int aacdec_channel_config(aacdec_t* h) {
	return h->info->channelConfig;
}

static int aacdec_bitrate(aacdec_t* h) {
	return h->info->bitRate;
}

static int aacdec_aac_samples_per_frame(aacdec_t* h) {
	return h->info->aacSamplesPerFrame;
}

static int aacdec_aac_num_channels(aacdec_t* h) {
	return h->info->aacNumChannels;
}

static int aacdec_extension_audio_object_type(aacdec_t* h) {
	return h->info->extAot;
}

static int aacdec_extension_sampling_rate(aacdec_t* h) {
	return h->info->extSamplingRate;
}

static int aacdec_num_lost_access_units(aacdec_t* h) {
	return h->info->numLostAccessUnits;
}

static int aacdec_num_total_bytes(aacdec_t* h) {
	return h->info->numTotalBytes;
}

static int aacdec_num_bad_bytes(aacdec_t* h) {
	return h->info->numBadBytes;
}

static int aacdec_num_total_access_units(aacdec_t* h) {
	return h->info->numTotalAccessUnits;
}

static int aacdec_num_bad_access_units(aacdec_t* h) {
	return h->info->numBadAccessUnits;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
	"bytes"
)

type AacDecoder struct {
	m C.aacdec_t
}

func NewAacDecoder() *AacDecoder {
	return &AacDecoder{}
}

// Explicitly configure the decoder by passing a raw AudioSpecificConfig (ASC)
// contained in a binary buffer. This is required for MPEG-4 and Raw Packets file format bitstreams
// as well as for LATM bitstreams with no in-band SMC. If the transport format is LATM with or without
// LOAS, configuration is assumed to be an SMC, for all other file formats an ASC.
func (v *AacDecoder) InitAdts(asc []byte) (err error) {
	p := (*C.char)(unsafe.Pointer(&asc[0]))
	pSize := C.int(len(asc))

	r := C.aacdec_init_adts(&v.m, p, pSize)

	if int(r) != 0 {
		return fmt.Errorf("init aac decoder failed, code is %d", int(r))
	}

	return nil
}

// De-allocate all resources of an AAC decoder instance.
func (v *AacDecoder) Close() {
	C.aacdec_close(&v.m)
}

// Fill AAC decoder's internal input buffer with bitstream data from the external input buffer.
// The function only copies such data as long as the decoder-internal input buffer is not full.
// So it grabs whatever it can from pBuffer and returns information (bytesValid) so that at a
// subsequent call of Fill(), the right position in pBuffer can be determined to
// grab the next data.
// @remark we will consume the input buffer and there maybe left bytes in buffer to parsed next time.
func (v *AacDecoder) Fill(input *bytes.Buffer) (err error) {
	b := input.Bytes()
	p := (*C.char)(unsafe.Pointer(&b[0]))
	pSize := C.int(input.Len())
	leftSize := C.int(0)

	r := C.aacdec_fill(&v.m, p, pSize, &leftSize)

	if int(r) != 0 {
		return fmt.Errorf("fill aac decoder failed, code is %d", int(r))
	}

	if int(leftSize) >= 0 {
		input.Next(input.Len() - int(leftSize))
	}

	return
}

// The samplerate in Hz of the fully decoded PCM audio signal (after SBR processing).
// @remark The only really relevant ones for the user.
func (v *AacDecoder) SampleRate() int {
	return int(C.aacdec_sample_rate(&v.m))
}

// The frame size of the decoded PCM audio signal.
//		1024 or 960 for AAC-LC
//		2048 or 1920 for HE-AAC (v2)
//		512 or 480 for AAC-LD and AAC-ELD
// @remark The only really relevant ones for the user.
func (v *AacDecoder) FrameSize() int {
	return int(C.aacdec_frame_size(&v.m))
}

// The number of output audio channels in the decoded and interleaved PCM audio signal.
// @remark The only really relevant ones for the user.
func (v *AacDecoder) NumChannels() int {
	return int(C.aacdec_num_channels(&v.m))
}

// sampling rate in Hz without SBR (from configuration info).
// @remark Decoder internal members.
func (v *AacDecoder) AacSampleRate() int {
	return int(C.aacdec_aac_sample_rate(&v.m))
}

// MPEG-2 profile (from file header) (-1: not applicable (e. g. MPEG-4)).
// @remark Decoder internal members.
func (v *AacDecoder) Profile() int {
	return int(C.aacdec_profile(&v.m))
}

// Audio Object Type (from ASC): is set to the appropriate value for MPEG-2 bitstreams (e. g. 2 for AAC-LC).
// @remark Decoder internal members.
func (v *AacDecoder) AudioObjectType() int {
	return int(C.aacdec_audio_object_type(&v.m))
}

// Channel configuration (0: PCE defined, 1: mono, 2: stereo, ...
// @remark Decoder internal members.
func (v *AacDecoder) ChannelConfig() int {
	return int(C.aacdec_channel_config(&v.m))
}

// Instantaneous bit rate.
// @remark Decoder internal members.
func (v *AacDecoder) Bitrate() int {
	return int(C.aacdec_bitrate(&v.m))
}

// Samples per frame for the AAC core (from ASC).
//		1024 or 960 for AAC-LC
//		512 or 480 for AAC-LD and AAC-ELD
// @remark Decoder internal members.
func (v *AacDecoder) AacSamplesPerFrame() int {
	return int(C.aacdec_aac_samples_per_frame(&v.m))
}

// The number of audio channels after AAC core processing (before PS or MPS processing).
//		CAUTION: This are not the final number of output channels!
// @remark Decoder internal members.
func (v *AacDecoder) AacNumChannels() int {
	return int(C.aacdec_aac_num_channels(&v.m))
}

// Extension Audio Object Type (from ASC)
// @remark Decoder internal members.
func (v *AacDecoder) ExtensionAudioObjectType() int {
	return int(C.aacdec_extension_audio_object_type(&v.m))
}

// Extension sampling rate in Hz (from ASC)
// @remark Decoder internal members.
func (v *AacDecoder) ExtensionSamplingRate() int {
	return int(C.aacdec_extension_sampling_rate(&v.m))
}

// This integer will reflect the estimated amount of lost access units in case aacDecoder_DecodeFrame()
// returns AAC_DEC_TRANSPORT_SYNC_ERROR. It will be < 0 if the estimation failed.
// @remark Statistics.
func (v *AacDecoder) NumLostAccessUnits() int {
	return int(C.aacdec_num_lost_access_units(&v.m))
}

// This is the number of total bytes that have passed through the decoder.
// @remark Statistics.
func (v *AacDecoder) NumTotalBytes() int {
	return int(C.aacdec_num_total_bytes(&v.m))
}

// This is the number of total bytes that were considered with errors from numTotalBytes.
// @remark Statistics.
func (v *AacDecoder) NumBadBytes() int {
	return int(C.aacdec_num_bad_bytes(&v.m))
}

// This is the number of total access units that have passed through the decoder.
// @remark Statistics.
func (v *AacDecoder) NumTotalAccessUnits() int {
	return int(C.aacdec_num_total_access_units(&v.m))
}

// This is the number of total access units that were considered with errors from numTotalBytes.
// @remark Statistics.
func (v *AacDecoder) NumBadAccessUnits() int {
	return int(C.aacdec_num_bad_access_units(&v.m))
}
