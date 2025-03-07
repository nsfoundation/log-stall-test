package main

import (
	"time"
	"os"
)


func main() {
		var buf []byte
		for j := 0; j <= 78; j++ {
			buf = append(buf, 'a')
		}

		var allBytes []byte
		for j := 0; j <= 255; j++ {
			allBytes = append(allBytes, byte(j))
		}

		for i := 0; i <= 100; i++ {
			time.Sleep(300*time.Millisecond)
			// long line
			os.Stdout.Write(buf)
			os.Stdout.Write(buf)
			os.Stdout.Write(buf)
			os.Stdout.Write(buf)
			os.Stdout.Write(buf)
			os.Stdout.WriteString("\n")
			// special line
			os.Stdout.Write(allBytes)
			os.Stdout.WriteString("\n")
			// short line
			os.Stdout.Write(buf)
			os.Stdout.WriteString("\n")
		}
}
