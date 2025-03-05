package main

import (
	"bytes"
	"fmt"
	"os"
)


func main() {
		var buf bytes.Buffer

		for i := 0; i <= 256; i++ {
			for j := 0; j <= 20; j++ {
				buf.WriteByte('a')
			}
			buf.WriteByte('\n')
		}

		_, err := fmt.Fprint(os.Stderr, buf.String())
		if err != nil {
			fmt.Printf("failed to flush: %v", err)
			return
		}

		fmt.Printf("flushed a %d byte string to stderr", buf.Len())

		_, err = fmt.Fprint(os.Stdout, buf.String())
		if err != nil {
			fmt.Printf("failed to flush: %v", err)
			return
		}

		fmt.Printf("flushed a %d byte string to stderr", buf.Len())

	        os.Exit(1)

}
