package utils

import "crypto/sha1"

func Encrypt(str string) string {
	var result string
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)[:10]
	for _, b := range bs {
		temp := int16(b)
		index := temp % 63
		result = result + string(alphabet[index])
	}
	return result
}
