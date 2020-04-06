# yop

易宝支付基础包 - golang库



## usage

```sh
go get -u github.com/hhjpin/yop
```


```go
package main

import (
    "fmt"
    "os"

    "github.com/hhjpin/yop"
)

func main() {
    var (
        //收单子商户
        demoMerchantNo = os.Getenv("YopMerchantNo")
        //父商编
        demoParentMerchantNo = os.Getenv("YopParentMerchantNo")
        //子商户对称密钥,可调密钥获取接口获取,下单生成hmac使用
        demoHmacKey = os.Getenv("YopHmacKey")
        //父商编私钥
        demoPrivateKey = os.Getenv("YopPrivateKey")
        //易宝公钥
        demoYopPublicKey = os.Getenv("YopPublicKey")

        business = yop.NewBusiness(yop.NewConfig(demoParentMerchantNo, demoPrivateKey, demoHmacKey, demoYopPublicKey))
    )

    res, err := business.MerChantBalanceQuery(demoMerchantNo)
	fmt.Printf("++Response++: %+v\n", res)
	if err != nil {
		panic(err)
	}
}
```


更多例子请看 `business_test.go` 文件。



## license

MIT