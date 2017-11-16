package Figo

import (
	"crypto"
	"errors"
	"crypto/rsa"
	"encoding/pem"
	"crypto/x509"
	"github.com/quexer/utee"
	"crypto/rand"
	"fmt"
	"math/big"
)

type RsaHelp struct {
	pubKey *rsa.PublicKey
	priKey *rsa.PrivateKey
	hash   map[crypto.Hash][]byte
}

func NewRsaHelp(pubKey, priKey string) RsaHelp {
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		utee.Chk(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	utee.Chk(err)

	block, _ = pem.Decode([]byte(pubKey))
	if block == nil {
		utee.Chk(errors.New("public key error!"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	utee.Chk(err)
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	return RsaHelp{
		pubKey: pub,
		priKey: priv,
		hash: map[crypto.Hash][]byte{
			crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
			crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
			crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
			crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
			crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
			crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
			crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
			crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
		},
	}
}

func (p *RsaHelp) PriEnc(data []byte) ([]byte, error) {
	signData, err := rsa.SignPKCS1v15(nil, p.priKey, crypto.Hash(0), data)
	if err != nil {
		return nil, err
	}
	return signData, nil
}

func (p *RsaHelp) PriDec(data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, p.priKey, data)
}

func (p *RsaHelp) PubEnc(data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, p.pubKey, data)
}

func (p *RsaHelp) PubDec(sig []byte) (out []byte, err error) {
	hashLen, prefix, err := p.pkcs1v15HashInfo(crypto.Hash(0), len(p.hash[crypto.Hash(0)]))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (p.pubKey.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)

	encrypt := func(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
		e := big.NewInt(int64(pub.E))
		c.Exp(m, e, pub.N)
		return c
	}

	m := encrypt(new(big.Int), p.pubKey, c)
	em := p.leftPad(m.Bytes(), k)
	out = p.unLeftPad(em)

	err = nil
	return
}

// copy from crypt/rsa/pkcs1v5.go
func (p *RsaHelp) pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := p.hash[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v5.go
func (p *RsaHelp) leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func (p *RsaHelp) unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}
