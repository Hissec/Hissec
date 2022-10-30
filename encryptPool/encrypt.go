package encryptPool

type Encrypt interface {
	Encode([]byte) []byte
	Decode([]byte) ([]byte, error)
}

