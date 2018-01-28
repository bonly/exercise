// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reverse implements an Android in Go
package bonly

import (
	"Java/android/databinding/DataBindingUtil"
	"Java/android/os"
	"Java/android/support/v7/app"
	rlayout "Java/go/bonly/R/layout"
	"Java/go/bonly/databinding/ActivityMainBinding"
	// "fmt"
)

type MainActivity struct {
	app.AppCompatActivity;
}

func (a *MainActivity) OnCreate1(this app.AppCompatActivity, b os.Bundle) {
	this.Super().OnCreate1(b);
	db := DataBindingUtil.SetContentView2(this, rlayout.Activity_main);
	mainBind := ActivityMainBinding.Cast(db);
	mainBind.SetAct(this);
	// fmt.Println(ActivityMainBinding.Wv);
	MvBinding.SetText("ok");
}

//app:imageUrl="http://f.tqn.com/y/inventors/1/L/z/V/macintosh.jpg"