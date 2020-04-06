package yop

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

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
)

var (
	//business = NewDefaultBusiness() //实际业务的例子
	business = NewBusiness(NewConfig(demoParentMerchantNo, demoPrivateKey, demoHmacKey, demoYopPublicKey)) //demo例子
)

func TestMerChantBalanceQuery(t *testing.T) {
	res, err := business.MerChantBalanceQuery(demoMerchantNo)
	fmt.Printf("++Response++: %+v\n", res)
	if err != nil {
		t.FailNow()
	}
}

func TestFileUpload(t *testing.T) {
	buf := "Hello, World"
	file, err := ioutil.TempFile("", "tmp_file")
	if err != nil {
		t.FailNow()
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte(buf)); err != nil {
		t.FailNow()
	}
	fmt.Println("local file:", file.Name())

	res, err := business.FileUpload(file.Name())
	fmt.Printf("++Response++: %+v\n", res)
	if err != nil {
		t.FailNow()
	}
}
