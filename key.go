package paysdk

import (
	"strings"
	"bytes"
)

func parsePublicKey(privateKey string)[]byte{
	return parse(privateKey, "-----BEGIN PUBLIC KEY-----", "-----END PUBLIC KEY-----")
}
func parsePrivateKey(publicKey string)[]byte{
	return parse(publicKey, "-----BEGIN RSA PRIVATE KEY-----", "-----END RSA PRIVATE KEY-----")
}
func parse(keyStr, prefix, suffix string)[]byte{
	if keyStr == "" {
		return nil
	}
	newKeyStr :=replace(keyStr)

	constLen := 64
	keyLen:=len(newKeyStr)
	var count = keyLen / constLen 
	if  keyLen % constLen  > 0 {
		count = count + 1
	}
	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < count; i++ {
		var b = i * constLen
		var e = b + constLen
		if e > keyLen {
			buf.WriteString(newKeyStr[b:])
		} else {
			buf.WriteString(newKeyStr[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}
func replace(data string)string{
	rep := strings.NewReplacer("\n", "", "\t", "", "\r", ""," ","")
	return rep.Replace(data)
}