package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	writer io.Writer
	err    error
}

func (sw safeWriter) writeLn(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.writer, s)
}

func proverbs(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	w := safeWriter{writer: file}
	w.writeLn("Errors are values.")
	w.writeLn("Don't just check errors, handle them gracefully.")
	w.writeLn("Don't panic")

	return w.err
}

func main() {
	filename := "proverbs.txt"
	err := proverbs(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e := os.Remove(filename)
	if e != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
