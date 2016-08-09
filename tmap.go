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

import "strings"
import "strconv"

// TMAP
type TMap struct {
	Id         int
	MapFields  string
	Width      int
	Height     int
	KeyObjects map[string]TKeyObject
}

var RealObjects map[string]TKeyObject = make(map[string]TKeyObject)

/*
var RealObjects = struct{
    sync.RWMutex
    m map[int]TObjectUser
}{m: make(map[int]TKeyObject)}
RealObjects.RLock()
n := RealObjects.m[123]
RealObjects.RUnlock()
*/

func CreateTMap(i int, mf string, w int, h int, keys []TKeyObject) TMap {
	var m TMap

	m.Id = i
	m.MapFields = strings.Replace(mf, " ", "", -1)
	m.Width = w
	m.Height = h
	m.KeyObjects = make(map[string]TKeyObject)
	for _, k := range keys {
		m.KeyObjects[strconv.Itoa(k.Id)] = k
		RealObjects[strconv.Itoa(k.Id)] = k
	}

	return m
}

func (m TMap) getObjects() map[string]TKeyObject {
	return m.KeyObjects
}

func (m TMap) getID() int {
	return m.Id
}

func (m TMap) getMapFields() string {
	return m.MapFields
}

func (m TMap) getWidth() int {
	return m.Width
}

func (m TMap) getHeight() int {
	return m.Height
}

func (m TMap) removeKey(idx int) TKeyObject {
	id_str := strconv.Itoa(idx)
	delete(m.KeyObjects, id_str)
	return RealObjects[id_str]
}

func (m TMap) addKey(idx int, x float64, y float64) TKeyObject {
	id_str := strconv.Itoa(idx)
	obj := RealObjects[id_str]
	obj.PosX = x
	obj.PosY = y
	RealObjects[id_str] = obj
	m.KeyObjects[id_str] = obj
	return obj

}
