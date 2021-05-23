package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type name struct {
	// name string
}

func (i name) Write(p []byte) (n int, err error) {

	fmt.Println("inside io.writer impl")
	return 1, nil
}

func main() {

	var n Writer
	n = name{}

	fmt.Fprintln(n, "")
	fmt.Fprintln(os.Stdout, "outing")
}

func multiWriter() {

	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)

	fmt.Fprintln(mw, "hola multi writer")
	fmt.Println("from buffer: ", buf)
}

func multiPipelReader() {
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hola")
	footer := strings.NewReader("</msg>")

	// for _, r := range []io.Reader{header, body, footer} {
	// 	_, err := io.Copy(os.Stdout, r)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	r := io.MultiReader(header, body, footer)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}

}

func pipeReader() {
	pr, pw := io.Pipe()

	go func() {

		defer pw.Close()
		_, err := fmt.Fprintln(pw, "hi")
		if err != nil {
			panic(err)
		}
	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err)

	}
}
