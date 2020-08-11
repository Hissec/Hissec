package Hissec

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntToBytes(n int) (pre []byte, err error) {
	bytesBuffer := new(bytes.Buffer)
	err = binary.Write(bytesBuffer, binary.BigEndian, int32(n))
	pre = bytesBuffer.Bytes()
	return
}
func BytesToInt(b []byte) (n int, err error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err = binary.Read(bytesBuffer, binary.BigEndian, &x)
	n = int(x)
	return
}
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
func Out(context []byte) []byte {
	pre, err := IntToBytes(len(context))
	if err != nil {
		fmt.Println(err)
	}
	return BytesCombine(pre, context)
}
