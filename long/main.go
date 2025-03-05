package main

import (
	"bytes"
	"fmt"
	"os"
)


func main() {
		var buf bytes.Buffer

		for i := 0; i <= 256_000; i++ {
			buf.WriteByte('a')
		}

		_, err := fmt.Fprint(os.Stderr, buf.String())
		if err != nil {
			fmt.Printf("failed to flush: %v", err)
			return
		}

		fmt.Printf("flushed a %d byte string to stderr", buf.Len())
}
