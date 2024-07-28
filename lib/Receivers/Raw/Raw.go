package raw

import (
	common "Schmottky/lib/Senders"
	"encoding/binary"
	"log"
	"math"
	"os"
)

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func complex128ToByte(c complex128) []byte {
	buffReal := float64ToByte(real(c))
	buffImag := float64ToByte(imag(c))

	return append(buffReal, buffImag...)
}

func ToFile(fname string) (ret int) {
	// create file
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write bytes to the file
	for {
		p, ok := <-common.PointChannel
		if !ok {
			break
		}
		bytes := complex128ToByte(p)
		n, err := f.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
		ret += n
	}
	return
}
