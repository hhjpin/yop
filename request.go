package yop

import (
	"github.com/satori/go.uuid"
	"net/http"
	"reflect"
)

type Request struct {
	Config *Config

	HttpMethod string
	Method     string
	Version    string
	SignAlg    string

	// 商户编号，易宝商户可不注册开放应用(获取appKey)也可直接调用API
	//CustomerNo string

	Headers          http.Header
	ParamMap         map[string]string
	FileMap          map[string]string
	JsonParam        map[string]string
	IgnoreSignParams []string

	RequestId string

	// 连接超时时间
	ConnectTimeout int

	// 读取返回结果超时
	ReadTimeout int

	// 可支持不同请求使用不同的appKey及secretKey
	AppKey string

	// 可支持不同请求使用不同的appKey及secretKey,secretKey只用于本地签名，不会被提交
	SecretKey string

	// 可支持不同请求使用不同的appKey及secretKey、serverRoot,secretKey只用于本地签名，不会被提交
	YopPublicKey string

	// 可支持不同请求使用不同的appKey及secretKey、serverRoot,secretKey只用于本地签名，不会被提交
	ServerRoot string
}

func NewRequest(config *Config) *Request {
	requestId := uuid.NewV4()
	request := &Request{
		Config:           config,
		Method:           "zh_CN",
		Version:          "1.0",
		Headers:          http.Header{},
		ParamMap:         map[string]string{},
		FileMap:          map[string]string{},
		JsonParam:        map[string]string{},
		IgnoreSignParams: []string{"sign"},
		RequestId:        requestId.String(),
		ConnectTimeout:   30000,
		ReadTimeout:      60000,
		AppKey:           "OPR:" + config.AppKey,
		SecretKey:        config.AesSecretKey,
		YopPublicKey:     config.YopPublicKey,
		ServerRoot:       config.ServerRoot,
	}
	return request
}

func (r *Request) AddParam(key string, value string) {
	r.ParamMap[key] = value
}

func (r *Request) AddFile(key string, filePath string) {
	//r.IgnoreSignParams = append(r.IgnoreSignParams, key)
	r.FileMap[key] = filePath
}

func (r *Request) AddParamsWithStruct(s interface{}) {
	t := reflect.TypeOf(s).Elem()
	v := reflect.ValueOf(s).Elem()
	numField := t.NumField()
	for i := 0; i < numField; i++ {
		r.ParamMap[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	}
}
