package binary

import (
	"encoding/binary"
	"io"
	"math"
)

type BinaryWriter struct{
	outWriter  io.Writer
	byteOrder binary.ByteOrder
}

func NewBinaryWriter(writer io.Writer, byteOrder binary.ByteOrder) *BinaryWriter{
	return &BinaryWriter{writer, byteOrder}
}

func (b *BinaryWriter) SetLittleEndian(){
	b.byteOrder = binary.LittleEndian
}
func (b *BinaryWriter) SetBigEndian(){
	b.byteOrder = binary.BigEndian
}

func (b *BinaryWriter) Uint8(tmp uint8){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Uint16(tmp uint16){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Uint32(tmp uint32){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Uint64(tmp uint64){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Int8(tmp int8){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Int16(tmp int16){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Int32(tmp int32){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Int64(tmp int64){
	err := binary.Write(b.outWriter, b.byteOrder, &tmp)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Float32(tmp float32){
	tmpUint32 := math.Float32bits(tmp)
	err := binary.Write(b.outWriter, b.byteOrder, &tmpUint32)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) Float64(tmp float64){
	tmpUint64 := math.Float64bits(tmp)
	err := binary.Write(b.outWriter, b.byteOrder, &tmpUint64)
	if err != nil{
		panic(err)
	}
}

func (b *BinaryWriter) UTF8Fixed(tmp string){
	panic("UTF8Fixed is not implemented.")
	
}
func (b *BinaryWriter) UTF8Null(tmp string){
	panic("UTF8Null is not implemented.")
}

func (b *BinaryWriter) UTF16Fixed(tmp string){
	panic("UTF16Fixed is not implemented.")
}

func (b *BinaryWriter) UTF16Null(tmp string){
	panic("UTF16Null is not implemented.")
}