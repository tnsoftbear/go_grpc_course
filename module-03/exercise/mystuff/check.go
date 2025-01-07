package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "gopher.png"
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	buf := make([]byte, 5*1024)
	var n int
	for {
		n, err = fd.Read(buf)
		fmt.Println("service read n:", n)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		if n == 0 {
			break
		}
	}
}
