// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type ClassObject struct {
	Value Map
}

func (this ClassObject) Type() TYPE {
	return CLASSOBJECT
}

func (this ClassObject) StringValue() string {
	return fmt.Sprint(this.Value.StringValue())
}

func (this ClassObject) Init() ClassObject {
	this.Value = Map{}.Init()
	return this
}

func (this ClassObject) Get(key String) Object {
	return this.Value.Get(key)
}

func (this *ClassObject) Set(key string, value Object) {
	this.Value.Set(String{Value: key}, value)
}
