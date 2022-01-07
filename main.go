package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

var DirSizes sync.Map

func main() {

	paths := os.Args[1:]

	stdOut := os.Stdout

	if err := Walk(paths); err != nil {
		fmt.Printf("Error while executing command. Err: %v", err)
		os.Exit(1)
	}

	PrintResult(&DirSizes, stdOut)
}

func Walk(paths []string) error {

	var wg sync.WaitGroup
	err := make(chan error)

	wg.Add(len(paths))

	for _, path := range paths {
		go GoToDir(path, &DirSizes, &wg)

		select {
		case err := <-err:
			return err
		default:
		}
	}

	wg.Wait()

	return nil
}

func GoToDir(rootPath string, m *sync.Map, wg *sync.WaitGroup) error {

	defer wg.Done()

	if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, errOut error) error {
		if errOut != nil {
			return errOut
		}

		currPath := fmt.Sprintf("%s/%s", rootPath, d.Name())

		info, err := d.Info()
		if err != nil {
			return err
		}
		m.Store(currPath, info.Size())

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func GetSize(path string) int64 {
	var size int64

	return size
}

func PrintResult(m *sync.Map, out io.Writer) {
	m.Range(func(k, v interface{}) bool {
		s := fmt.Sprintf("%d    %s\n", v, k)
		io.WriteString(out, s)
		return true
	})
}
