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

		var totalBuf []byte

		for i := 0; i <= 100; i++ {
			time.Sleep(300*time.Millisecond)
			// long line
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.WriteString("\n")
			totalBuf = append(totalBuf, '\n')
			// special line
			os.Stdout.Write(allBytes)
			totalBuf = append(totalBuf, allBytes...)
			os.Stdout.WriteString("\n")
			totalBuf = append(totalBuf, '\n')
			// short line
			os.Stdout.Write(buf)
			totalBuf = append(totalBuf, buf...)
			os.Stdout.WriteString("\n")
			totalBuf = append(totalBuf, '\n')

			// long line
			os.Stderr.Write(buf)
			os.Stderr.Write(buf)
			os.Stderr.Write(buf)
			os.Stderr.Write(buf)
			os.Stderr.Write(buf)
			os.Stderr.WriteString("\n")
			// special line
			os.Stderr.Write(allBytes)
			os.Stderr.WriteString("\n")
			// short line
			os.Stderr.Write(buf)
			os.Stderr.WriteString("\n")
		}


		// and all at once, twice :)
		os.Stdout.Write(totalBuf)
		os.Stderr.Write(totalBuf)
		os.Stdout.Write(totalBuf)
		os.Stderr.Write(totalBuf)
}
