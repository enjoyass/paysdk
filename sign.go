package paysdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

/**
* @param pemKeyFile string类型 pem文件路径,例如：/home/private.pem
* @return *pem.Block类型，参考golang标准库encoding/pem
* @return error 类型
*
* type Block struct {
*    Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
*    Headers map[string]string // 可选的头项
*    Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
* }
 */
func ParsePemKeyFromFile(pemKeyFile string) (*pem.Block, error) {
	byteKey, err := ioutil.ReadFile(pemKeyFile)
	if err != nil {
		return nil, err
	}
	var block *pem.Block
	block, _ = pem.Decode(byteKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	return block, nil
}

/**
* @param byteKey []byet类型 privateKey的字节数组 ,例如：
*
*	var privateKey = `-----BEGIN RSA PRIVATE KEY-----
*	MIIEpAIBAAKCAQEAyseL5HsZRZvxHKf+7ksjKkqfeLEw3IrF8OvGuoI1+E49qawC
*	...(此处省略中间字符)
*	amt34/fIOgsPZJdVS/MxUd4zAwyDKJXWacs5Z5NbHks9vTAg8AAK4g==
*	-----END RSA PRIVATE KEY-----`
*	var byteKey = []byte(privateKey)
*
* @return *pem.Block类型，参考golang标准库encoding/pem
* @return error 类型
*
* type Block struct {
*    Type    string            // 得自前言的类型（如"RSA PRIVATE KEY"）
*    Headers map[string]string // 可选的头项
*    Bytes   []byte            // 内容解码后的数据，一般是DER编码的ASN.1结构
* }
 */
func ParsePemKeyFromByte(byteKey []byte) (*pem.Block, error) {
	var block *pem.Block
	block, _ = pem.Decode(byteKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	return block, nil
}

/**
* @param src []byte类型,生成签名的数据源
* @param hash crypto.Hash类型，生成消息摘要方法，例如crypto.SHA256
* @param privateKeyFile string类型,privateKey文件路径
* @retrun []byte 签名后的字节数组
* @return error
**/
func SignPKCS1v15ByPemFile(src []byte, hash crypto.Hash, privateKeyFile string) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	block, err := ParsePemKeyFromFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, prikey, hash, hashed)
}

/**
* @param src []byte类型,生成签名的数据源
* @param hash crypto.Hash类型，生成消息摘要方法，例如crypto.SHA256
* @param privateKey []byte,privateKey的字节数组
* @retrun []byte 签名后的字节数组
* @return error
**/
func SignPKCS1v15ByPemByte(src, privateKey []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	block, err := ParsePemKeyFromByte(privateKey)
	if err != nil {
		return nil, err
	}

	prikey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, prikey, hash, hashed)
}

/**
* @param src []byte类型,生成签名的数据源
* @param hash crypto.Hash类型，生成消息摘要方法，例如crypto.SHA256
* @param publicKeyFile string类型,publicKeyFile文件路径
* @param sign []byte 签名
* @return error
**/
func VerifyPKCS1v15ByPemFile(src, sign []byte, publicKeyFile string, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	block, err := ParsePemKeyFromFile(publicKeyFile)
	if err != nil {
		return err
	}
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}

/**
* @param src []byte类型,生成签名的数据源
* @param hash crypto.Hash类型，生成消息摘要方法，例如crypto.SHA256
* @param publickey []byte,publickey的字节数组
* @param sign []byte 签名
* @return error
**/
func VerifyPKCS1v15ByPemByte(src, sign, publickey []byte, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	block, err := ParsePemKeyFromByte(publickey)
	if err != nil {
		return err
	}
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}
