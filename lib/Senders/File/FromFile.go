package fromfile

import (
	common "Schmottky/lib/Senders"
	"encoding/binary"
	"io"
	"log"
	"math"
	"os"
)

func byteToFloat64(b []byte) float64 {
	bits := binary.LittleEndian.Uint64(b)
	return math.Float64frombits(bits)
}

func byteToComplex128(b []byte) complex128 {
	r := byteToFloat64(b[:8])
	i := byteToFloat64(b[8:16])
	return complex(r, i)
}

func Read(fname string) (ret int) {
	// open file
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buff := make([]byte, 16)
	// loop over file
	for {
		n, err := f.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			close(common.PointChannel)
			break
		}
		ret += n
		common.PointChannel <- byteToComplex128(buff)
	}
	return
}
