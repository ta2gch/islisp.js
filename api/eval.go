/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"fmt"
	"html"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/ta2gch/iris/runtime"
	"github.com/ta2gch/iris/runtime/ilos/instance"
)

const version = "7cd80c3"

func main() {
	print(`Welcome to Iris (` + version + `). Iris is an ISLisp implementation on Go.
This library works with gopherjs and has no methods to get input.
For more infomation, see https://islisp.js.org.

Copyright &copy; 2017 TANIGUCHI Masaya All Rights Reserved.`)
	js.Global.Set("islisp", map[string]interface{}{"eval": eval})
}

func eval(s string) string {
	r := strings.NewReader(s)
	w := new(bytes.Buffer)
	runtime.TopLevel.StandardInput = instance.NewStream(r, nil)
	runtime.TopLevel.StandardOutput = instance.NewStream(nil, w)
	runtime.TopLevel.ErrorOutput = instance.NewStream(nil, w)
	runtime.TopLevel.Function.Set(instance.NewSymbol("READ"), nil)
	runtime.TopLevel.Function.Set(instance.NewSymbol("READ-LINE"), nil)
	runtime.TopLevel.Function.Set(instance.NewSymbol("READ-CHAR"), nil)
	p, err := runtime.Read(runtime.TopLevel)
	if err != nil {
		fmt.Fprint(w, html.EscapeString(err.String()))
		return w.String()
	}
	e, err := runtime.Eval(runtime.TopLevel, p)
	if err != nil {
		fmt.Fprint(w, html.EscapeString(err.String()))
		return w.String()
	}
	fmt.Fprint(w, html.EscapeString(e.String()))
	return w.String()
}
