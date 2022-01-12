package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"text/tabwriter"
)

var DirSizes sync.Map

func main() {

	paths := os.Args[1:]

	stdOut := os.Stdout

	if err := DiskUsage(paths, &DirSizes); err != nil {
		fmt.Printf("Error while executing command. Err: %v", err)
		os.Exit(1)
	}

	PrintResult(&DirSizes, stdOut)
}

func DiskUsage(paths []string, m *sync.Map) error {

	var wg sync.WaitGroup
	err := make(chan error)

	wg.Add(len(paths))

	for _, path := range paths {
		go GoToDir(path, m, &wg, err)

		select {
		case err := <-err:
			return err
		default:
		}
	}

	wg.Wait()

	return nil
}

func GoToDir(rootPath string, m *sync.Map, wg *sync.WaitGroup, errChan chan error) {

	defer wg.Done()

	if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, errOut error) error {
		if errOut != nil {
			return errOut
		}

		if d.IsDir() {
			m.Store(path, GetDirSize(path))
		}

		return nil
	}); err != nil {
		errChan <- err
	}
}

func GetDirSize(path string) int64 {
	var size int64

	dir, err := os.Open(path)
	if err != nil {
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return size
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		if file.IsDir() {
			size += GetDirSize(fmt.Sprintf("%s/%s", path, file.Name()))
		} else {
			size += file.Size()
		}
	}

	return size
}

func PrintResult(m *sync.Map, out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 8, 0, '\t', tabwriter.AlignRight)
	defer w.Flush()

	m.Range(func(k, v interface{}) bool {
		fmt.Fprintf(w, "%d\t%s\n", int64(v.(int64)/1024.0), k)
		return true
	})
}

func HumanReadable(d interface{}) string {
	return ""
}
