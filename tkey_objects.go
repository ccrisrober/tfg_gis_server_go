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

package main

// KEY OBJECTS
type TKeyObject struct {
	Id    int
	PosX  float64
	PosY  float64
	Color string
}

func CreateKeyObject(id int, posx float64, posy float64, color string) TKeyObject {
	var k TKeyObject

	k.Id = id
	k.PosX = posx
	k.PosY = posy
	k.Color = color

	return k
}

func (k TKeyObject) getId() int {
	return k.Id
}

func (k TKeyObject) getPosX() float64 {
	return k.PosX
}

func (k TKeyObject) getPosY() float64 {
	return k.PosY
}

func (k TKeyObject) getColor() string {
	return k.Color
}
