// Copyright (c) 2015, maldicion069 (Cristian Rodríguez) <ccrisrober@gmail.con>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.package com.example

// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: tmap.go
// DO NOT EDIT!

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *TMap) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	buf.Grow(256)
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *TMap) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	var scratch fflib.FormatBitsScratch
	_ = scratch
	_ = obj
	_ = err
	buf.WriteString(`{"Id":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.Id), 10, mj.Id < 0)
	buf.WriteString(`,"MapFields":`)
	fflib.WriteJsonString(buf, string(mj.MapFields))
	buf.WriteString(`,"Width":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.Width), 10, mj.Width < 0)
	buf.WriteString(`,"Height":`)
	fflib.FormatBits(&scratch, buf, uint64(mj.Height), 10, mj.Height < 0)
	buf.WriteString(`,"KeyObjects":`)
	/* Falling back. type=map[string]main.TKeyObject kind=map */
	obj, err = json.Marshal(mj.KeyObjects)
	if err != nil {
		return err
	}
	buf.Write(obj)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_TMapbase = iota
	ffj_t_TMapno_such_key

	ffj_t_TMap_Id

	ffj_t_TMap_MapFields

	ffj_t_TMap_Width

	ffj_t_TMap_Height

	ffj_t_TMap_KeyObjects
)

var ffj_key_TMap_Id = []byte("Id")

var ffj_key_TMap_MapFields = []byte("MapFields")

var ffj_key_TMap_Width = []byte("Width")

var ffj_key_TMap_Height = []byte("Height")

var ffj_key_TMap_KeyObjects = []byte("KeyObjects")

func (uj *TMap) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *TMap) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_TMapbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_TMapno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'H':

					if bytes.Equal(ffj_key_TMap_Height, kn) {
						currentKey = ffj_t_TMap_Height
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffj_key_TMap_Id, kn) {
						currentKey = ffj_t_TMap_Id
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'K':

					if bytes.Equal(ffj_key_TMap_KeyObjects, kn) {
						currentKey = ffj_t_TMap_KeyObjects
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'M':

					if bytes.Equal(ffj_key_TMap_MapFields, kn) {
						currentKey = ffj_t_TMap_MapFields
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'W':

					if bytes.Equal(ffj_key_TMap_Width, kn) {
						currentKey = ffj_t_TMap_Width
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_TMap_KeyObjects, kn) {
					currentKey = ffj_t_TMap_KeyObjects
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_TMap_Height, kn) {
					currentKey = ffj_t_TMap_Height
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_TMap_Width, kn) {
					currentKey = ffj_t_TMap_Width
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_TMap_MapFields, kn) {
					currentKey = ffj_t_TMap_MapFields
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_TMap_Id, kn) {
					currentKey = ffj_t_TMap_Id
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_TMapno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_TMap_Id:
					goto handle_Id

				case ffj_t_TMap_MapFields:
					goto handle_MapFields

				case ffj_t_TMap_Width:
					goto handle_Width

				case ffj_t_TMap_Height:
					goto handle_Height

				case ffj_t_TMap_KeyObjects:
					goto handle_KeyObjects

				case ffj_t_TMapno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Id:

	/* handler: uj.Id type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Id = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MapFields:

	/* handler: uj.MapFields type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.MapFields = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Width:

	/* handler: uj.Width type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Width = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Height:

	/* handler: uj.Height type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Height = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_KeyObjects:

	/* handler: uj.KeyObjects type=map[string]main.TKeyObject kind=map */

	{
		/* Falling back. type=map[string]main.TKeyObject kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.KeyObjects)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}