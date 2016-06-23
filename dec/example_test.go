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

package dec_test

import (
	"fmt"
	"github.com/winlinvip/go-fdkaac/dec"
	"bytes"
)

func ExampleAacDecoder_ADTS() {
	d := dec.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitAdts(asc); err != nil {
		fmt.Println("init decoder failed, err is", err)
		return
	}
	defer d.Close()

	buf := &bytes.Buffer{}
	buf.Write([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38})
	if err := d.Fill(buf); err != nil {
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

	// Output:
	// SampleRate: 0
	// FrameSize: 0
	// NumChannels: 0
	// AacSampleRate: 44100
	// Profile: 1
	// AudioObjectType: 2
	// ChannelConfig: 2
	// Bitrate: 0
	// AacSamplesPerFrame: 1024
	// AacNumChannels: 0
	// ExtensionAudioObjectType: 0
	// ExtensionSamplingRate: 0
	// NumLostAccessUnits: 0
	// NumTotalBytes: 0
	// NumBadBytes: 0
	// NumTotalAccessUnits: 0
	// NumBadAccessUnits: 0
}