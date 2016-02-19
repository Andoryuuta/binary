package binary

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

type BinaryReader struct {
	inReader  io.Reader
	byteOrder binary.ByteOrder
}

func NewBinaryReader(reader io.Reader, byteOrder binary.ByteOrder) *BinaryReader {
	return &BinaryReader{reader, byteOrder}
}

func NewBinaryReaderFromBytes(inputBytes []byte, byteOrder binary.ByteOrder) *BinaryReader {
	inReader := bytes.NewReader(inputBytes)
	return NewBinaryReader(inReader, byteOrder)
}

func (b *BinaryReader) SetLittleEndian() {
	b.byteOrder = binary.LittleEndian
}
func (b *BinaryReader) SetBigEndian() {
	b.byteOrder = binary.BigEndian
}

func (b *BinaryReader) Uint8() uint8 {
	var tmp uint8
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Uint16() uint16 {
	var tmp uint16
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Uint32() uint32 {
	var tmp uint32
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Uint64() uint64 {
	var tmp uint64
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Int8() int8 {
	var tmp int8
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Int16() int16 {
	var tmp int16
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Int32() int32 {
	var tmp int32
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Int64() int64 {
	var tmp int64
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return tmp
}

func (b *BinaryReader) Float32() float32 {
	var tmp uint32
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return math.Float32frombits(tmp)
}

func (b *BinaryReader) Float64() float64 {
	var tmp uint64
	err := binary.Read(b.inReader, b.byteOrder, &tmp)
	if err != nil {
		panic(err)
	}
	return math.Float64frombits(tmp)
}

//Reads in a UTF8 fixed length string.
func (b *BinaryReader) UTF8Fixed(size int) string {
	out := make([]byte, size)
	
	n, err := io.ReadFull(b.inReader, out)
	if err != nil{
		panic(err)
	} else if n != size {
		panic("UTF8Fixed: io.ReadFull couldn't read the full size given.")
	}
	
	return string(out)
	
	return "UTF8Fixed is not implemented."
	
}

//Reads in a UTF8 null terminated string.
func (b *BinaryReader) UTF8Null() string {
	//Create the temp buf to hold our char
	tmpBuf := make([]byte, 1)
	
	//Create a slice of bytes to hold our output
	var out []byte

	//Loop while tmpChar is not NULL
	for {
		n, err := io.ReadFull(b.inReader, tmpBuf)
		if err != nil {
			panic(err)
		} else if n < 1 {
			panic("UTF8Null: io.ReadFull didn't read any bytes!")
		}
		
		//Break on NULL byte
		if tmpBuf[0] == 0{
			break
		}
		
		//Append the byte
		out = append(out, tmpBuf[0])
	}
	return string(out)
}

func (b *BinaryReader) UTF16Fixed(size int) string{
	return "UTF16Fixed is not implemented."
}

func (b *BinaryReader) UTF16Null() string {
	return "UTF16Null is not implemented."
}