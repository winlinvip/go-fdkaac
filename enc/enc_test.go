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

// @remark we must use another packet for utest, because the cgo will dup symbols.
package enc_test

import (
	"testing"
	"github.com/winlinvip/go-fdkaac/enc"
)

func TestLcInit(t *testing.T) {
	e := enc.NewAacEncoder()
	if err := e.InitLc(2, 44100, 48000); err != nil {
		t.Error("Init encoder failed, err is", err)
		return
	}
	defer e.Close()

	if 1024 != e.FrameSize() {
		t.Error("FrameSize not 1024, actual is", e.FrameSize())
		return
	}
}