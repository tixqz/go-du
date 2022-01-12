package main

import (
	"sync"
	"testing"
)

func TestDiskUsageNoError(t *testing.T) {

	pathOne, pathTwo := t.TempDir(), t.TempDir()

	var (
		TestMap sync.Map
	)

	paths := []string{
		pathOne,
		pathTwo,
	}

	if err := DiskUsage(paths, &TestMap); err != nil {
		t.Errorf("")
	}

}

func TestsDiskUsageWithError(t *testing.T) {

}

func TestGoToDirNoSubDirs(t *testing.T) {
	var (
		TestMap sync.Map
		wg      sync.WaitGroup
	)

	errChan := make(chan error)
	path := t.TempDir()

	wg.Add(1)
	GoToDir(path, &TestMap, &wg, errChan)

}

func TestGoToDirWithSubDirs(t *testing.T) {

}

func TestGetDirSize(t *testing.T) {

}

func TestPrintResult(t *testing.T) {

}

func TestHumanReadable(t *testing.T) {

}

func checkMapForPaths(m *sync.Map, paths []string) {

}

func checkMapForSizes(m *sync.Map, dirSizes map[string]int) {

}
