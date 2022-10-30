package encryptPool

type empty struct {
}

func init() {
	regist("empty", &empty{})
}

func (*empty) Encode(data []byte) []byte {
	return data
}

func (*empty) Decode(data []byte) ([]byte, error) {
	return data, nil
}
