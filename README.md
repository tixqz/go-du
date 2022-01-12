# go-du
Naive implementation of du command tool written in go. It lacks most of features from original du. Currently, it can only walk trough passed directories and return size for 

### Usage

`./dugo`"path/to/dir1" "path/to/dir2" "path/to/dir/etc"

Program processing different directories via different goroutines: 1 directory -> 1 goroutine.

As the storage I used sync.Map to support RW operations from many goroutines in time.

  

To build du locally use `make` tool 

```
cd ~/path/to/go-du
make build
```

To run tests
`make test`

and for test coverage
`make test_coverage`

For cleanup use
`make clean`

### TODO

- [ ] add human-readable flag
- [ ] add tests
- [ ] encapsulate storage in struct (easier testing)
