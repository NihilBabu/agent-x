package file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

//ListFiles is for listing all files in a dir
func ListFiles() (matches []string) {
	files, err := filepath.Glob("/Users/nihil/Desktop/go/agent-x/log/*.log")
	if err != nil {
		log.Fatal(err)
		return []string{}
	}
	return files
}

//Tail a file
func Tail(filename string, out io.Writer) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	oldSize := info.Size()
	for {
		for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			if prefix {
				fmt.Fprint(out, string(line))
			} else {
				fmt.Fprintln(out, string(line))
			}
		}
		pos, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			panic(err)
		}
		for {
			time.Sleep(time.Second)
			newinfo, err := f.Stat()
			if err != nil {
				panic(err)
			}
			newSize := newinfo.Size()
			if newSize != oldSize {
				if newSize < oldSize {
					f.Seek(0, 0)
				} else {
					f.Seek(pos, io.SeekStart)
				}
				r = bufio.NewReader(f)
				oldSize = newSize
				break
			}
		}
	}
}
