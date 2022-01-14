# 波场地址使用示例

## golang

生成、验证等方法参见 golang/main.go

```go
    // 生成
    genKey, _ := GenerateKey()
    // 波场地址有两种， hex 和 base58 编码的，我们一般使用 base58 编码，不考虑 hex 编码
    // 生成 地址 ( Base58)
    tronAddr := PubkeyToAddress(genKey.PublicKey)
    log.Println(tronAddr.String()) // Base58地址 THEjT4G2VEJs25kJXQUwWbba5cqLL8kzs4
    log.Println(tronAddr.Hex())    // Hex地址    414fb8896d6240303a97a315475d9a747adfae3cea

    // 解析Base58地址, 如果不是合法的地址，会返回error
    reparse, err := Base58ToAddress(tronAddr.String())
    if err != nil {
        panic(err)
    }
    log.Println(reparse.String()) // Base58地址 THEjT4G2VEJs25kJXQUwWbba5cqLL8kzs4
```

## js

js 只需验证地址即可

> 参考 https://github.com/tronprotocol/tronweb/blob/5fe79d5677/src/utils/crypto.js#L136

可以提取出验证方法 utils.crypto.isAddressValid

```js
import TronWeb from 'tronweb';
console.log(TronWeb.utils.crypto.isAddressValid("TH6NL3WcsAzY1wPCuoeuhbmC3gUCDvvBA"));
console.log(TronWeb.utils.crypto.isAddressValid("TH6NL3WcsAzY1wPCuoeuhbmC3gUCDvvBA8"));

```
