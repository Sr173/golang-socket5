package encrypt

type Cipher struct {
	encode [256]byte
	decode [256]byte
}

func NewCipher(password []byte) *Cipher {
	var temp_cipher Cipher
	for i, v := range password {
		temp_cipher.decode[v] = byte(i)
	}
	for i := 0; i < 256; i++ {
		temp_cipher.encode[i] = password[i]
	}
	return &temp_cipher
}

func (cipher *Cipher) EncodeData(data []byte) {
	for i := range data {
		data[i] = cipher.encode[data[i]] ^ cipher.encode[i%0x100]
	}
}

func (cipher *Cipher) DecodeData(data []byte) {
	for i := range data {
		data[i] = cipher.decode[data[i]^cipher.encode[i%0x100]]
	}
}
