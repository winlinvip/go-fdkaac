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
)

func ExampleAacDecoder_ADTS() {
	d := dec.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitAdts(asc); err != nil {
		return
	}

	fmt.Println("\nInitAdts:")
	fmt.Println(fmt.Sprintf("AacSampleRate=%v, Profile=%v, AudioObjectType=%v, ChannelConfig=%v, SamplePerFrame=%v",
		d.AacSampleRate(), d.Profile(), d.AudioObjectType(), d.ChannelConfig(), d.SamplesPerFrame()))

	fmt.Println("\nFinal:")
	fmt.Println(fmt.Sprintf("SampleRate=%v, FrameSize=%v, NumChannels=%v, AacSampleRate=%v,",
		d.SampleRate(), d.FrameSize(), d.NumChannels(), d.AacSampleRate()))
	fmt.Println(fmt.Sprintf("Profile=%v, AudioObjectType=%v, ChannelConfig=%v, Bitrate=%v, SamplesPerFrame=%v,",
		d.Profile(), d.AudioObjectType(), d.ChannelConfig(), d.Bitrate(), d.SamplesPerFrame()))
	fmt.Println(fmt.Sprintf("AacNumChannels=%v, ExtensionAudioObjectType=%v, ExtensionSamplingRate=%v,",
		d.AacNumChannels(), d.ExtensionAudioObjectType(), d.ExtensionSamplingRate()))
	fmt.Println(fmt.Sprintf("NumLostAccessUnits=%v, NumTotalBytes=%v, NumBadBytes=%v, NumTotalAccessUnits=%v,",
		d.NumLostAccessUnits(), d.NumTotalBytes(), d.NumBadBytes(), d.NumTotalAccessUnits()))
	fmt.Println(fmt.Sprintf("NumBadAccessUnits=%v", d.NumBadAccessUnits()))

	// Output:
	//
	// InitAdts:
	// AacSampleRate=44100, Profile=1, AudioObjectType=2, ChannelConfig=2, SamplePerFrame=1024
	//
	// Final:
	// SampleRate=0, FrameSize=0, NumChannels=0, AacSampleRate=44100,
	// Profile=1, AudioObjectType=2, ChannelConfig=2, Bitrate=0, SamplesPerFrame=1024,
	// AacNumChannels=0, ExtensionAudioObjectType=0, ExtensionSamplingRate=0,
	// NumLostAccessUnits=0, NumTotalBytes=0, NumBadBytes=0, NumTotalAccessUnits=0,
	// NumBadAccessUnits=0
}