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

package fdkaac

import (
	"bytes"
	"fmt"
)

func ExampleAacDecoder_RAW() {
	var err error
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	// directly decode the frame to pcm.
	result := new(bytes.Buffer)
	if err = d.Decode([]byte{
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}, result); err != nil {
		fmt.Println("decode failed, err is", err)
		return
	}

	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("FrameSize:", d.FrameSize())
	fmt.Println("NumChannels:", d.NumChannels())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
	fmt.Println("Profile:", d.Profile())
	fmt.Println("AudioObjectType:", d.AudioObjectType())
	fmt.Println("ChannelConfig:", d.ChannelConfig())
	fmt.Println("Bitrate:", d.Bitrate())
	fmt.Println("AacSamplesPerFrame:", d.AacSamplesPerFrame())
	fmt.Println("AacNumChannels:", d.AacNumChannels())
	fmt.Println("ExtensionAudioObjectType:", d.ExtensionAudioObjectType())
	fmt.Println("ExtensionSamplingRate:", d.ExtensionSamplingRate())
	fmt.Println("NumLostAccessUnits:", d.NumLostAccessUnits())
	fmt.Println("NumTotalBytes:", d.NumTotalBytes())
	fmt.Println("NumBadBytes:", d.NumBadBytes())
	fmt.Println("NumTotalAccessUnits:", d.NumTotalAccessUnits())
	fmt.Println("NumBadAccessUnits:", d.NumBadAccessUnits())
	fmt.Println("SampleBits:", d.SampleBits())
	fmt.Println("PCM:", result.Len())

	// Output:
	// SampleRate: 44100
	// FrameSize: 1024
	// NumChannels: 2
	// AacSampleRate: 44100
	// Profile: 1
	// AudioObjectType: 2
	// ChannelConfig: 2
	// Bitrate: 31352
	// AacSamplesPerFrame: 1024
	// AacNumChannels: 2
	// ExtensionAudioObjectType: 0
	// ExtensionSamplingRate: 0
	// NumLostAccessUnits: 0
	// NumTotalBytes: 91
	// NumBadBytes: 0
	// NumTotalAccessUnits: 1
	// NumBadAccessUnits: 0
	// SampleBits: 16
	// PCM: 4096
}

func ExampleAacDecoder_ADTS() {
	var err error
	d := fdkaac.NewAacDecoder()

	if err := d.InitAdts(); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	result := new(bytes.Buffer)
	if err = d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}, result); err != nil {
		fmt.Println("decode failed, err is", err)
		return
	}

	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("FrameSize:", d.FrameSize())
	fmt.Println("NumChannels:", d.NumChannels())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
	fmt.Println("Profile:", d.Profile())
	fmt.Println("AudioObjectType:", d.AudioObjectType())
	fmt.Println("ChannelConfig:", d.ChannelConfig())
	fmt.Println("Bitrate:", d.Bitrate())
	fmt.Println("AacSamplesPerFrame:", d.AacSamplesPerFrame())
	fmt.Println("AacNumChannels:", d.AacNumChannels())
	fmt.Println("ExtensionAudioObjectType:", d.ExtensionAudioObjectType())
	fmt.Println("ExtensionSamplingRate:", d.ExtensionSamplingRate())
	fmt.Println("NumLostAccessUnits:", d.NumLostAccessUnits())
	fmt.Println("NumTotalBytes:", d.NumTotalBytes())
	fmt.Println("NumBadBytes:", d.NumBadBytes())
	fmt.Println("NumTotalAccessUnits:", d.NumTotalAccessUnits())
	fmt.Println("NumBadAccessUnits:", d.NumBadAccessUnits())
	fmt.Println("SampleBits:", d.SampleBits())
	fmt.Println("PCM:", result.Len())

	// Output:
	// SampleRate: 44100
	// FrameSize: 1024
	// NumChannels: 2
	// AacSampleRate: 44100
	// Profile: 1
	// AudioObjectType: 2
	// ChannelConfig: 2
	// Bitrate: 33764
	// AacSamplesPerFrame: 1024
	// AacNumChannels: 2
	// ExtensionAudioObjectType: 0
	// ExtensionSamplingRate: 0
	// NumLostAccessUnits: 0
	// NumTotalBytes: 98
	// NumBadBytes: 0
	// NumTotalAccessUnits: 1
	// NumBadAccessUnits: 0
	// SampleBits: 16
	// PCM: 4096
}

func ExampleAacDecoder_ADTS_Stream_NotEnoughBits() {
	var err error
	d := fdkaac.NewAacDecoder()

	// @remark the partial stream only support by ADTS.
	if err := d.InitAdts(); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	// Fill then decode, aac stream.
	// AAC Frame #0, part 0
	result := new(bytes.Buffer)
	if err = d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54}, result); err != nil {
		fmt.Println("fill decoder failed, err is", err)
		return
	}

	if result.Len() != 0 {
		fmt.Println("not enough bits")
		return
	}

	// AAC Frame #0, part 1
	if err = d.Decode([]byte{0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}, result); err != nil {
		fmt.Println("fill decoder failed, err is", err)
		return
	}

	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("FrameSize:", d.FrameSize())
	fmt.Println("NumChannels:", d.NumChannels())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
	fmt.Println("Profile:", d.Profile())
	fmt.Println("AudioObjectType:", d.AudioObjectType())
	fmt.Println("ChannelConfig:", d.ChannelConfig())
	fmt.Println("Bitrate:", d.Bitrate())
	fmt.Println("AacSamplesPerFrame:", d.AacSamplesPerFrame())
	fmt.Println("AacNumChannels:", d.AacNumChannels())
	fmt.Println("ExtensionAudioObjectType:", d.ExtensionAudioObjectType())
	fmt.Println("ExtensionSamplingRate:", d.ExtensionSamplingRate())
	fmt.Println("NumLostAccessUnits:", d.NumLostAccessUnits())
	fmt.Println("NumTotalBytes:", d.NumTotalBytes())
	fmt.Println("NumBadBytes:", d.NumBadBytes())
	fmt.Println("NumTotalAccessUnits:", d.NumTotalAccessUnits())
	fmt.Println("NumBadAccessUnits:", d.NumBadAccessUnits())
	fmt.Println("SampleBits:", d.SampleBits())
	fmt.Println("PCM:", result.Len())

	// Output:
	// SampleRate: 44100
	// FrameSize: 1024
	// NumChannels: 2
	// AacSampleRate: 44100
	// Profile: 1
	// AudioObjectType: 2
	// ChannelConfig: 2
	// Bitrate: 33764
	// AacSamplesPerFrame: 1024
	// AacNumChannels: 2
	// ExtensionAudioObjectType: 0
	// ExtensionSamplingRate: 0
	// NumLostAccessUnits: 0
	// NumTotalBytes: 98
	// NumBadBytes: 0
	// NumTotalAccessUnits: 1
	// NumBadAccessUnits: 0
	// SampleBits: 16
	// PCM: 4096
}

func ExampleAacDecoder_ADTS_Frames() {
	var err error
	d := fdkaac.NewAacDecoder()

	if err := d.InitAdts(); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	// AAC FRAME #0
	result1 := new(bytes.Buffer)
	if err = d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}, result1); err != nil {
		fmt.Println("decode failed, err is", err)
		return
	}

	// AAC FRAME #1
	result2 := new(bytes.Buffer)
	if err = d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0b, 0xe0, 0xfc,
		0x21, 0x17, 0x55, 0x55, 0x19, 0x1a, 0x2a, 0x2d, 0x54, 0xce, 0x00, 0x58, 0x1a, 0x1e, 0x42, 0x0e,
		0x1f, 0xd2, 0xd4, 0x9c, 0x15, 0x77, 0xf4, 0x07, 0x38, 0x3d, 0xc5, 0x04, 0x19, 0x64, 0x39, 0x98,
		0x01, 0xae, 0x2e, 0xb1, 0xd0, 0x87, 0xca, 0x33, 0x17, 0xfb, 0x05, 0x00, 0x7a, 0x60, 0x47, 0x79,
		0x6b, 0x9b, 0xdf, 0x2d, 0xfd, 0x32, 0xc6, 0x9f, 0x1f, 0x21, 0x4b, 0x04, 0x9b, 0xe2, 0x4d, 0x62,
		0xc8, 0x01, 0xe0, 0x98, 0x0a, 0x37, 0x48, 0x44, 0x42, 0x02, 0x00, 0xd0, 0x7d, 0xae, 0xb4, 0x32,
		0xf1, 0xcc, 0x76, 0x5f, 0x18, 0xac, 0xae, 0x0e}, result2); err != nil {
		fmt.Println("decode failed, err is", err)
		return
	}

	fmt.Println("SampleRate:", d.SampleRate())
	fmt.Println("FrameSize:", d.FrameSize())
	fmt.Println("NumChannels:", d.NumChannels())
	fmt.Println("AacSampleRate:", d.AacSampleRate())
	fmt.Println("Profile:", d.Profile())
	fmt.Println("AudioObjectType:", d.AudioObjectType())
	fmt.Println("ChannelConfig:", d.ChannelConfig())
	fmt.Println("Bitrate:", d.Bitrate())
	fmt.Println("AacSamplesPerFrame:", d.AacSamplesPerFrame())
	fmt.Println("AacNumChannels:", d.AacNumChannels())
	fmt.Println("ExtensionAudioObjectType:", d.ExtensionAudioObjectType())
	fmt.Println("ExtensionSamplingRate:", d.ExtensionSamplingRate())
	fmt.Println("NumLostAccessUnits:", d.NumLostAccessUnits())
	fmt.Println("NumTotalBytes:", d.NumTotalBytes())
	fmt.Println("NumBadBytes:", d.NumBadBytes())
	fmt.Println("NumTotalAccessUnits:", d.NumTotalAccessUnits())
	fmt.Println("NumBadAccessUnits:", d.NumBadAccessUnits())
	fmt.Println("SampleBits:", d.SampleBits())
	fmt.Println("PCM0:", result1.Len())
	fmt.Println("PCM1:", result2.Len())

	// Output:
	// SampleRate: 44100
	// FrameSize: 1024
	// NumChannels: 2
	// AacSampleRate: 44100
	// Profile: 1
	// AudioObjectType: 2
	// ChannelConfig: 2
	// Bitrate: 32730
	// AacSamplesPerFrame: 1024
	// AacNumChannels: 2
	// ExtensionAudioObjectType: 0
	// ExtensionSamplingRate: 0
	// NumLostAccessUnits: 0
	// NumTotalBytes: 193
	// NumBadBytes: 0
	// NumTotalAccessUnits: 2
	// NumBadAccessUnits: 0
	// SampleBits: 16
	// PCM0: 4096
	// PCM1: 4096
}
