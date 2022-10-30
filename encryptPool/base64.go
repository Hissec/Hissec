package encryptPool

import (
	"Hissec"
	"encoding/base64"
)

type B64 struct {
}

func init() {
	regist("base64", &B64{})
}

func (*B64) Encode(data []byte) []byte {
	if len(data) < 32 {
		return data
	}
	rs := base64.StdEncoding.EncodeToString(data[32:])
	return Hissec.BytesCombine(data[:32], []byte(rs))
}

func (*B64) Decode(data []byte) ([]byte, error) {
	if len(data) < 32 {
		return data, nil
	}
	rs, err := base64.StdEncoding.DecodeString(string(data[32:]))
	return Hissec.BytesCombine(data[:32], rs), err
}
