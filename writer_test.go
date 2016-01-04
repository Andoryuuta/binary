package binary

import (
	"testing"
	_"fmt"
	"bytes"
	"encoding/binary"
)

func TestUint8(t *testing.T){
	input		:= uint8(5)
	expected	:= []byte{0x05}

	
	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Uint8(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestUint16(t *testing.T){
	input		:= uint16(30000)
	expected	:= []byte{0x30, 0x75}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Uint16(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestUint32(t *testing.T){
	input		:= uint32(1000000)
	expected	:= []byte{0x40, 0x42, 0x0F, 0x00}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Uint32(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestUint64(t *testing.T){
	input		:= uint64(3000000000)
	expected	:= []byte{0x00, 0x5E, 0xD0, 0xB2, 0x00, 0x00, 0x00, 0x00}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Uint64(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestInt8(t *testing.T){
	input		:= int8(-5)
	expected	:= []byte{0xFB}

	
	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Int8(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestInt16(t *testing.T){
	input		:= int16(-30000)
	expected	:= []byte{0xD0, 0x8A}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Int16(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestInt32(t *testing.T){
	input		:= int32(-1000000)
	expected	:= []byte{0xC0, 0xBD, 0xF0, 0xFF}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Int32(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestInt64(t *testing.T){
	input		:= int64(-3000000000)
	expected	:= []byte{0x00, 0xA2, 0x2F, 0x4D, 0xFF, 0xFF, 0xFF, 0xFF}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Int64(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestFloat32(t *testing.T){
	input		:= float32(1.5523)
	expected	:= []byte{0xC4, 0xB1, 0xC6, 0x3F}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Float32(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}

func TestFloat64(t *testing.T){
	input		:= float64(552355.5)
	expected	:= []byte{0x00, 0x00, 0x00, 0x00, 0x47, 0xDB, 0x20, 0x41}

	var buf bytes.Buffer
	binWtr := NewBinaryWriter(&buf, binary.LittleEndian)
	binWtr.Float64(input)
	
	actual := buf.Bytes()
	if !bytes.Equal(expected, actual){
		t.Errorf("Test Failed, Got:[% X], Expected:[% X]\n", actual, expected)
	}
}


