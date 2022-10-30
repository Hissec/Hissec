package encryptPool

import "fmt"

type encryptManager struct {
	allEncrypt map[string]Encrypt
}

func (m *encryptManager) regist(name string, encry Encrypt) {
	m.allEncrypt[name] = encry
}

var enMager = encryptManager{
	allEncrypt: make(map[string]Encrypt),
}

func regist(name string, encry Encrypt) {
	enMager.regist(name, encry)
}

func GetEncrypt(name string) (encrypt Encrypt, err error) {
	encrypt, ok := enMager.allEncrypt[name]
	if !ok {
		err = fmt.Errorf("not found encrypt type of %s", name)
	}
	return
}

func GetEncryptList() (encryptList []string) {
	for k, _ := range enMager.allEncrypt {
		encryptList = append(encryptList, k)
	}
	return
}
