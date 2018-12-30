// +build windows
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

func rev(b []byte) {
	for i := len(b)/2 - 1; i >= 0; i-- {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
}

func decodeByte(buf []byte) byte {
	const chars = "BCDFGHJKMPQRTVWXY2346789"
	acc := 0
	for j := 14; j >= 0; j-- {
		acc *= 256
		acc += int(buf[j])
		buf[j] = byte((acc / len(chars)) & 0xFF)
		acc %= len(chars)
	}
	return chars[acc]
}

func binaryKeyToASCII(buf []byte) string {
	var out bytes.Buffer
	for i := 28; i >= 0; i-- {
		if (29-i)%6 == 0 {
			out.WriteByte('-')
			i--
		}
		out.WriteByte(decodeByte(buf))
	}
	outBytes := out.Bytes()
	rev(outBytes)
	return string(outBytes)
}

var (
	regKeyPath       = `SOFTWARE\Microsoft\Windows NT\CurrentVersion`
	regKeyValueName  = `DigitalProductId4`
	productKeyOffset = 52
)

func init() {
	flag.StringVar(&regKeyPath, "p", regKeyPath, "registry path")
	flag.StringVar(&regKeyValueName, "n", regKeyValueName, "registry key value name")
	flag.IntVar(&productKeyOffset, "o", productKeyOffset, "product key offset in bytes")
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
}

func main() {
	r, err := registry.OpenKey(registry.LOCAL_MACHINE, regKeyPath, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	digitalProductID, _, err := r.GetBinaryValue(regKeyValueName)
	if err != nil {
		log.Fatal(err)
	}

	binaryKey := digitalProductID[productKeyOffset:]
	fmt.Println(binaryKeyToASCII(binaryKey))
}
