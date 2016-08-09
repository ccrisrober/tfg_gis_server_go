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

// OBJECT USER
type TObjectUser struct {
	Id      int
	RollDie int
	PosX    float64
	PosY    float64
	Map     int
	Objects Set
}

func CreateObjectUser(id int, px float64, py float64) TObjectUser {
	return CreateObjectUser2(id, px, py, 0, 0)
}

func CreateObjectUser2(id int, px float64, py float64, m int, roll int) TObjectUser {
	var o TObjectUser

	o.Id = id
	o.PosX = px
	o.PosY = py
	o.Map = m
	o.RollDie = roll
	o.Objects = NewSet()

	return o
}

func (o TObjectUser) GetId() int {
	return o.Id
}

func (o TObjectUser) GetPosX() float64 {
	return o.PosX
}

func (o TObjectUser) GetPosY() float64 {
	return o.PosY
}

func (o TObjectUser) GetMap() int {
	return o.Map
}

func (o TObjectUser) GetRollDie() int {
	return o.RollDie
}

func (o TObjectUser) SetId(v int) {
	o.Id = v
}

func (o TObjectUser) SetPosX(v float64) {
	o.PosX = v
}

func (o TObjectUser) SetPosY(v float64) {
	o.PosY = v
}

func (o TObjectUser) SetMap(v int) {
	o.Map = v
}

func (o TObjectUser) SetRollDie(v int) {
	o.RollDie = v
}

func (o TObjectUser) AddObject(idx int32) {
	o.Objects.Add(idx)
}

func (o TObjectUser) RemoveObject(idx int32) {
	o.Objects.Remove(idx)
}
