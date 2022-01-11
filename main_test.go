package main

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

var MockFS = fstest.MapFS{
	"/root-dir/":         {Mode: fs.ModeDir},
	"/root-dir/test.txt": {Data: []byte("this is test file")},
}

func TestWalk(t *testing.T) {

}

func TestGoToDir(t *testing.T) {

}

func TestGetDirSize(t *testing.T) {

}

func TestPrintResult(t *testing.T) {

}
