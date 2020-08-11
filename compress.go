package Hissec

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

type Compress interface {
	Compre(data []byte) (result []byte)
	Uncompre(data []byte) (result []byte)
}

type Hgzip struct{}

func (Hgzip) Compre(data []byte) (result []byte) {
	var res bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&res, gzip.BestCompression)
	_, _ = gz.Write(data)
	_ = gz.Flush()
	defer func() { gz.Close() }()
	result = res.Bytes()
	return
}

func (Hgzip) Uncompre(data []byte) (result []byte) {

	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err == nil {
		defer func() { gz.Close() }()
		result, _ = ioutil.ReadAll(gz)
	}
	return
}
