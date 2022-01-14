package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {

	// 生成 私钥
	// 注意保存私钥 s.D.Bytes()
	genKey, err := GenerateKey()
	if err != nil {
		log.Println(err)
		return
	}

	// 波场地址有两种， hex 和 base58 编码的，我们一般使用 base58 编码，不考虑 hex 编码
	// 生成 地址 ( Base58)
	tronAddr := PubkeyToAddress(genKey.PublicKey)

	log.Println(tronAddr.String()) // Base58地址 THEjT4G2VEJs25kJXQUwWbba5cqLL8kzs4
	log.Println(tronAddr.Hex())    // Hex地址    414fb8896d6240303a97a315475d9a747adfae3cea

	// 解析Base58地址, 如果不是合法的地址，会返回error
	reparse, err := Base58ToAddress(tronAddr.String())
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(reparse.String()) // THEjT4G2VEJs25kJXQUwWbba5cqLL8kzs4

	// 根据私钥生成公钥， 这也是一种验证方式
	puk := ecdsa.PublicKey{
		Curve: secp256k1.S256(),
	}
	puk.X, puk.Y = pk.Curve.ScalarBaseMult(s.D.Bytes())
	log.Println(PubkeyToAddress(puk).String())
}
