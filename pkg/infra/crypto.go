package infra

import (
	//"github.com/zhigui-projects/gm-go/sm2"
	"github.com/zhigui-projects/gm-go/sm3"

	//"encoding/asn1"

	//"crypto/sha256"
	"crypto/x509"
	//"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"math/big"

	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/hyperledger/fabric/common/crypto"
	//"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"

	gm_x509 "github.com/zhigui-projects/gm-crypto/x509"
	"github.com/zhigui-projects/gm-plugins/gosm"
	// GM
	//"github.com/zhigui-projects/gm-go/sm2"
	//"github.com/zhigui-projects/gm-go/sm3"
	"github.com/zhigui-projects/gm-plugins/primitive"
)

type CryptoConfig struct {
	MSPID      string
	PrivKey    string
	SignCert   string
	TLSCACerts []string
}

type ECDSASignature struct {
	R, S *big.Int
}

type Crypto struct {
	Creator  []byte
	//PrivKey  *ecdsa.PrivateKey
	//SignCert *x509.Certificate
	PrivKey    *primitive.Sm2PrivateKey
	SignCert   *x509.Certificate
}

func (s *Crypto) Sign(message []byte) ([]byte, error) {
	//ri, si, err := ecdsa.Sign(rand.Reader, s.PrivKey, digest(message))
	//signature, err := s.PrivKey.Sign(rand.Reader, digest(message), nil)
	//signature, err := s.PrivKey.Sign(rand.Reader, digest(message), nil)
	//signature, err := primitive.Sm2Crypto.Sign(s.PrivKey, digest(message), nil)
	signature, err := new(gosm.GoSm2).Sign(s.PrivKey, digest(message), nil)
	if err != nil {
		return nil, err
	}

	//isTrue, _ := new(gosm.GoSm2).Verify(&s.PrivKey.Sm2PublicKey, signature, message, nil)
	//if !isTrue {
	//	fmt.Errorf("验证签名错误")
	//}

	//si, _, err = utils.ToLowS(&s.PrivKey.PublicKey, si)
	//if err != nil {
	//	return nil, err
	//}

	//return asn1.Marshal(ECDSASignature{ri, si})

	return signature, nil
}

func (s *Crypto) Serialize() ([]byte, error) {
	return s.Creator, nil
}

func (s *Crypto) NewSignatureHeader() (*common.SignatureHeader, error) {
	creator, err := s.Serialize()
	if err != nil {
		return nil, err
	}
	nonce, err := crypto.GetRandomNonce()
	if err != nil {
		return nil, err
	}

	return &common.SignatureHeader{
		Creator: creator,
		Nonce:   nonce,
	}, nil
}

func digest(in []byte) []byte {
	//h := sha256.New()
	//h.Write(in)
	//return h.Sum(nil)
	e := sm3.New()
	e.Write(in)
	//return e.Sum(nil)[:32]
	return e.Sum(nil)
}

func toPEM(in []byte) ([]byte, error) {
	d := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, err := base64.StdEncoding.Decode(d, in)
	if err != nil {
		return nil, err
	}
	return d[:n], nil
}

//func GetPrivateKey(f string) (*ecdsa.PrivateKey, error) {
//	in, err := ioutil.ReadFile(f)
//	if err != nil {
//		return nil, err
//	}
//
//	k, err := utils.PEMtoPrivateKey(in, []byte{})
//	if err != nil {
//		return nil, err
//	}
//
//	key, ok := k.(*ecdsa.PrivateKey)
//	if !ok {
//		return nil, errors.Errorf("expecting ecdsa key")
//	}
//
//	return key, nil
//}
func GetPrivateKey(f string) (*primitive.Sm2PrivateKey, error) {
		in, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		k, err := utils.PEMtoPrivateKey(in, []byte{})
		if err != nil {
			return nil, err
		}

		key, ok := k.(*primitive.Sm2PrivateKey)
		if !ok {
			return nil, errors.Errorf("expecting ecdsa key")
		}

		return key, nil
}

func GetCertificate(f string) (*x509.Certificate, []byte, error) {
	in, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, nil, err
	}

	block, _ := pem.Decode(in)

	//c, err := gm_x509.ParseCertificate(block.Bytes)
	gm_x509.InitX509("SM2")
	//c, err := gm_x509.GetX509SM2().ParseCertificate(block.Bytes)
	c, err := gm_x509.GetX509().ParseCertificate(block.Bytes)
	return c, in, err
}
