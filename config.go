package yop

const (
	// 加密算法
	//algAes  = "AES"
	//algSha  = "SHA"
	//algSha1 = "SHA1"

	// 保护参数
	serverRoot    = "https://openapi.yeepay.com/yop-center" //"http://58.83.141.72/yop-center"//
	yosServerRoot = "https://yos.yeepay.com/yop-center"

	//encoding      = "UTF-8"
	//success       = "SUCCESS"
	//callback      = "callback"

	// 方法的默认参数名
	method = "method"

	// 会话id默认参数名
	sessionId = "sessionId"

	// 应用键的默认参数名
	appKey = "appKey"

	// 服务版本号的默认参数名
	version = "v"

	// 签名的默认参数名
	sign = "sign"

	// 商户编号
	customerNo = "customerNo"

	// 时间戳
	//timestamp = "ts"
)

type Config struct {
	ServerRoot     string
	YosServerRoot  string
	AppKey         string
	AesSecretKey   string //非对称加密密钥
	HmacSecretKey  string //对称加密密钥
	YopPublicKey   string //易宝公钥
	Debug          bool
	ConnectTimeout int
	ReadTimeout    int
	MaxUploadLimit int
	PublicEdKeys   []string
}

func NewDefaultConfig() *Config {
	return NewConfig(parentMerchantNo, privateKey, hmacKey, yopPublicKey)
}

func NewConfig(parentMerchantNo string, privateKey string, hmacKey string, yopPublicKey string) *Config {
	return &Config{
		ServerRoot:     serverRoot,
		YosServerRoot:  yosServerRoot,
		AppKey:         parentMerchantNo,
		AesSecretKey:   privateKey,
		HmacSecretKey:  hmacKey,
		YopPublicKey:   yopPublicKey,
		Debug:          false,
		ConnectTimeout: 30000,
		ReadTimeout:    60000,
		MaxUploadLimit: 4096000,
		PublicEdKeys:   []string{appKey, version, sign, method, sessionId, customerNo, ""},
	}
}

func (c *Config) IsPublicEdKey(key string) bool {
	for _, val := range c.PublicEdKeys {
		if val == key {
			return true
		}
	}
	return false
}
