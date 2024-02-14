package reflect

import (
	"testing"

	"gotest.tools/v3/assert"
)

const (
	cbool       = true
	cint        = 42
	cintbin     = 0b101010
	cintoct     = 0o52
	cinthex     = 0x2a
	cint8       = int8(42)
	cint16      = int16(42)
	cint32      = int32(42)
	crune       = '*'
	cint64      = int64(42)
	cuint       = uint(42)
	cuint8      = uint8(42)
	cuint16     = uint16(42)
	cuint32     = uint32(42)
	cuint64     = uint64(42)
	cuintptr    = uintptr(42)
	cfloat32    = float32(42)
	cfloat64    = 42.0
	cfloat64hex = 0x2ap0
	ccomplex64  = complex64(42i)
	ccomplex128 = 42i
	cstring     = "42"
	craw        = `42`
)

func TestAddr(t *testing.T) {
	vbool := true
	vint := 42
	vintbin := 0b101010
	vintoct := 0o52
	vinthex := 0x2a
	vint8 := int8(42)
	vint16 := int16(42)
	vint32 := int32(42)
	vrune := '*'
	vint64 := int64(42)
	vuint := uint(42)
	vuint8 := uint8(42)
	vuint16 := uint16(42)
	vuint32 := uint32(42)
	vuint64 := uint64(42)
	vuintptr := uintptr(42)
	vfloat32 := float32(42)
	vfloat64 := 42.0
	vfloat64hex := 0x2ap0
	vcomplex64 := complex64(42i)
	vcomplex128 := 42i
	vstring := "42"
	vraw := `42`

	tests := []struct {
		name     string
		literal  any
		constant any
		variable any
	}{
		{"bool", true, cbool, vbool},
		{"int", 42, cint, vint},
		{"intbin", 0b101010, cintbin, vintbin},
		{"intoct", 0o52, cintoct, vintoct},
		{"inthex", 0x2a, cinthex, vinthex},
		{"int8", int8(42), cint8, vint8},
		{"int16", int16(42), cint16, vint16},
		{"int32", int32(42), cint32, vint32},
		{"rune", '*', crune, vrune},
		{"int64", int64(42), cint64, vint64},
		{"uint", uint(42), cuint, vuint},
		{"uint8", uint8(42), cuint8, vuint8},
		{"uint16", uint16(42), cuint16, vuint16},
		{"uint32", uint32(42), cuint32, vuint32},
		{"uint64", uint64(42), cuint64, vuint64},
		{"uintptr", uintptr(42), cuintptr, vuintptr},
		{"float32", float32(42), cfloat32, vfloat32},
		{"float64", 42.0, cfloat64, vfloat64},
		{"float64hex", 0x2ap0, cfloat64hex, vfloat64hex},
		{"complex64", complex64(42i), ccomplex64, vcomplex64},
		{"complex128", 42i, ccomplex128, vcomplex128},
		{"string", "42", cstring, vstring},
		{"raw", `42`, craw, vraw},
	}

	t.Logf("Nr Name       Token Type       Value")
	for num, test := range tests {
		testAddr(t, num, test.name, "lit", test.literal)
		testAddr(t, num, test.name, "const", test.constant)
		testAddr(t, num, test.name, "var", test.variable)
		assert.Equal(t, test.literal, test.constant)
		assert.Equal(t, test.literal, test.variable)
		// constant and variable are also equal because of transitivity
	}

	const c = 42
	ca := Addr(c)
	assert.Equal(t, *ca, c)

	la := Addr(42)
	assert.Equal(t, *la, 42)
}

func testAddr(t *testing.T, num int, name string, token string, test any) {
	t.Logf("%2v %-10v %-5v %-10T %#v", num, name, token, test, test)
	want := &test
	got := Addr(test)
	assert.Equal(t, *got, *want)
}
