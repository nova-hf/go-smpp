// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdutext

import (
	"bytes"
	"testing"
	"reflect"
)

func TestEncode(t *testing.T) {
	test := []struct {
		codec Codec
		want []byte
	}{
		{Latin1([]byte("áéíóú moço")), []byte("\xe1\xe9\xed\xf3\xfa mo\xe7o")},
		{UCS2([]byte("áéíóú moço")), []byte("\x00\xe1\x00\xe9\x00\xed\x00\xf3\x00\xfa\x00 \x00m\x00o\x00\xe7\x00o")},
		{ISO88595([]byte(iso88595UTF8Bytes)), []byte(iso88595Bytes)},
	}
	for _, tc := range test {
		have := tc.codec.Encode()
		if !bytes.Equal(tc.want, have) {
			t.Fatalf("unexpected text for %s:\nwant: %q\nhave: %q",
				reflect.TypeOf(tc.codec), tc.want, have)
		}
	}
}

func TestDecode(t *testing.T) {
	test := []struct {
		want []byte
		codec Codec
	}{
		{[]byte("áéíóú moço"), Latin1([]byte("\xe1\xe9\xed\xf3\xfa mo\xe7o"))},
		{[]byte("áéíóú moço"), UCS2([]byte("\x00\xe1\x00\xe9\x00\xed\x00\xf3\x00\xfa\x00 \x00m\x00o\x00\xe7\x00o"))},
		{[]byte(iso88595UTF8Bytes), ISO88595([]byte(iso88595Bytes))},
	}
	for _, tc := range test {
		have := tc.codec.Decode()
		if !bytes.Equal(tc.want, have) {
			t.Fatalf("unexpected text for %s:\nwant: %q\nhave: %q",
				reflect.TypeOf(tc.codec), tc.want, have)
		}
	}
}

func TestTypes(t *testing.T) {
	test := []struct {
		want DataCoding
		test DataCoding
	}{
		{0x00, DefaultType},
		{0x03, Latin1Type},
		{0x06, ISO88595Type},
		{0x08, UCS2Type},
		{0x10, DefaultType | FlashType},
		{0x18, UCS2Type | FlashType},
	}
	for _, tc := range test {
		if tc.want != tc.test {
			t.Fatalf("unexpected data coding:\nwant: %x\nhave: %x", tc.want, tc.test)
		}
	}
}
