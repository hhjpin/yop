package yop

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/hhjpin/goutils/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Post(uri string, request *Request, response interface{}) error {
	request.HttpMethod = "POST"
	serverUrl := c.richRequest(uri, request)
	err := c.signRsaParameter(uri, request)
	if err != nil {
		logger.Error(err)
		return err
	}
	resp := Response{}
	resp.body, err = c.doPostForm(serverUrl, request)
	if err != nil {
		logger.Error(err)
		return err
	}
	if err := resp.JsonUnmarshal(response); err != nil {
		logger.Error(err)
		return err
	}
	if err := resp.checkError(); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (c *Client) convertHeader(header http.Header) map[string]string {
	newHeaders := map[string]string{}
	for key := range header {
		newHeaders[key] = header.Get(key)
	}
	return newHeaders
}

func (c *Client) richRequest(uri string, request *Request) string {
	request.Method = uri
	serverUrl := request.ServerRoot + uri
	reg := regexp.MustCompile("@/rest/v([^/]+)/@i")
	version := reg.FindString(uri)
	if version != "" {
		request.Version = version
	}
	return serverUrl
}

func (c *Client) signRsaParameter(uri string, request *Request) error {
	headers := http.Header{}
	headers.Set("x-yop-appkey", request.AppKey)
	headers.Set("x-yop-request-id", request.RequestId)
	headers.Set("x-yop-sdk-langs", "go")
	headers.Set("x-yop-sdk-version", "3.1.12")

	timestamp := time.Now().Format("2006-01-02T15:04:05+0800") //ISO8601时间格式
	expiredSecond := "1800"
	protocolVersion := "yop-auth-v2"
	authString := fmt.Sprintf("%s/%s/%s/%s", protocolVersion, request.AppKey, timestamp, expiredSecond)

	headersToSignSet := []string{"x-yop-request-id"}
	headersToSign := c.getHeadersToSign(headers, headersToSignSet)
	signedHeaders := strings.ToLower(strings.Join(headersToSignSet, ";"))
	canonicalHeader := c.getCanonicalHeaders(headersToSign)
	canonicalUri := strings.Replace(url.QueryEscape(uri), "%2F", "/", -1)
	canonicalQueryString := c.getCanonicalQueryString(request)
	canonicalRequest := authString + "\n" + request.HttpMethod + "\n" + canonicalUri + "\n" + canonicalQueryString + "\n" + canonicalHeader

	//签名和编码
	sign, err := c.sign(request.SecretKey, canonicalRequest)
	if err != nil {
		logger.Error(err)
		return err
	}
	signToBase64 := base64.StdEncoding.EncodeToString(sign)
	signToBase64 = strings.Replace(signToBase64, "+", "-", -1)
	signToBase64 = strings.Replace(signToBase64, "/", "_", -1)
	//signToBase64 = strings.TrimRight(signToBase64, "=")
	signToBase64 += "$SHA256"

	headers.Set("Authorization", fmt.Sprintf("YOP-RSA2048-SHA256 %s/%s/%s/%s/%s/%s", protocolVersion, request.AppKey, timestamp, expiredSecond, signedHeaders, signToBase64))
	request.Headers = headers

	if request.Config.Debug {
		logger.Debug("authString=" + authString)
		logger.Debug("canonicalURI=" + canonicalUri)
		logger.Debug("canonicalQueryString=" + canonicalQueryString)
		logger.Debug("canonicalHeader=" + canonicalHeader)
		logger.Debug("canonicalRequest=" + canonicalRequest)
		logger.Debug("signToBase64=" + signToBase64)
	}
	return nil
}

func (c *Client) sign(privateKey string, data string) ([]byte, error) {
	fullKey := ""
	start := 0
	length := len(privateKey)
	for start <= length {
		if length >= start+64 {
			fullKey += privateKey[start:start+64] + "\n"
		} else {
			fullKey += privateKey[start:] + "\n"
		}
		start += 64
	}
	fullKey = "-----BEGIN RSA PRIVATE KEY-----\n" + fullKey + "-----END RSA PRIVATE KEY-----"

	//pem解码
	block, _ := pem.Decode([]byte(fullKey))
	//X509解码
	key, err1 := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err1 != nil {
		logger.Info(err1)
		return nil, err1
	}
	hash := sha256.New()
	hash.Write([]byte(data))
	newBytes := hash.Sum(nil)
	sign, err2 := rsa.SignPKCS1v15(rand.Reader, (key).(*rsa.PrivateKey), crypto.SHA256, newBytes)
	if err2 != nil {
		logger.Info(err2)
		return nil, err2
	}
	return sign, nil
}

func (c *Client) getCanonicalHeaders(headers http.Header) string {
	var headerStrings []string
	for key := range headers {
		key = url.QueryEscape(strings.ToLower(strings.Trim(key, " ")))
		val := url.QueryEscape(strings.Trim(headers.Get(key), " "))
		headerStrings = append(headerStrings, key+":"+val)
	}
	sort.Strings(headerStrings)
	return strings.Join(headerStrings, "\n")
}

func (c *Client) getHeadersToSign(headers http.Header, headersToSign []string) http.Header {
	var set []string
	for _, val := range headersToSign {
		set = append(set, strings.ToLower(strings.Trim(val, " ")))
	}
	isSetEmpty := len(set) == 0
	ret := http.Header{}
	for key, val := range headers {
		if (isSetEmpty && c.isDefaultHeaderToSign(key)) || (!isSetEmpty && strings.ToLower(key) != "authorization" && InStringArray(set, strings.ToLower(key))) {
			ret[key] = val
		}
	}
	return ret
}

func (c *Client) isDefaultHeaderToSign(header string) bool {
	header = strings.ToLower(strings.Trim(header, " "))
	return strings.Index(header, "x-yop-") == 0 || header == "host" || header == "content-type"
}

func (c *Client) getCanonicalQueryString(request *Request) string {
	var params []string
	if len(request.JsonParam) != 0 {
		return ""
	}
	for key, val := range request.ParamMap {
		if strings.ToLower(key) == "authorization" {
			continue
		}
		params = append(params, key+"="+url.QueryEscape(val))
	}
	sort.Strings(params)
	return strings.Join(params, "&")
}

func (c *Client) escapeParamMap(serverUrl string, request *Request) {
	for key, val := range request.ParamMap {
		if key != "_file" {
			request.ParamMap[key] = url.QueryEscape(val)
		}
	}
}

/*func (c *Client) doPost(uri string, request *Request) (string, error) {
	resp, _ := golang_utils.Requests.PostForm(uri, request.ParamMap, c.convertHeader(request.Headers))
	text := resp.Text()
	logger.Debug(text)
	return text, nil
}*/

func (c *Client) writeFileToPostBody(bodyWriter *multipart.Writer, fileField string, fileName string) error {
	fileWriter, err := bodyWriter.CreateFormFile(fileField, path.Base(fileName))
	if err != nil {
		logger.Error(err)
		return err
	}
	localFile, err := os.Open(fileName)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer func() {
		err := localFile.Close()
		if err != nil {
			logger.Error(err)
		}
	}()
	_, err = io.Copy(fileWriter, localFile)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (c *Client) doPostForm(uri string, request *Request) ([]byte, error) {
	//构造body
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	for key, value := range request.ParamMap {
		err := bodyWriter.WriteField(key, value)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}
	for fileField, fileName := range request.FileMap {
		err := c.writeFileToPostBody(bodyWriter, fileField, fileName)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
	}
	contentType := bodyWriter.FormDataContentType()
	err := bodyWriter.Close()
	if err != nil {
		logger.Error(err)
	}
	//http客户端请求
	httpClient := http.Client{
		Timeout: 12 * time.Second, //12s超时
	}
	req, err := http.NewRequest("POST", uri, bodyBuffer)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for key, val := range c.convertHeader(request.Headers) {
		req.Header.Set(key, val)
	}
	response, err := httpClient.Do(req)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error("Couldn't parse response body. %+v", err)
	}
	logger.Debug(string(body))
	return body, nil
}
