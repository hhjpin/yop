package yop

import "fmt"

//具体接口解释看易宝支付的亿企通手册v1.13
//文档链接：https://open.yeepay.com/docs/retail000001

const (
	urlMerChantBalanceQuery         = "/rest/v1.0/sys/merchant/balancequery"
	urlMerChantAgreeInfoQuery       = "/rest/v1.0/sys/merchant/agreeinfoquery"
	urlMerchantSendAuthorizeNum     = "/rest/v1.0/sys/merchant/sendauthorizenum"
	urlMerchantReceiveAuthorizeNum  = "/rest/v1.0/sys/merchant/receiveauthorizenum"
	urlMerchantAuthorizeUrlQuery    = "/rest/v1.0/sys/merchant/authorizeurlquery"
	urlMerchantBankBranchInfo       = "/rest/v1.0/sys/merchant/bankbranchinfo"
	urlMerchantHmacKeyQuery         = "/rest/v1.0/sys/merchant/hmackeyquery"
	urlMerchantRegStatusQuery       = "/rest/v1.0/sys/merchant/regstatusquery"
	urlMerchantPersonRegInfoAdd     = "/rest/v1.0/sys/merchant/personreginfoadd"
	urlMerchantIndividualRegInfoAdd = "/rest/v1.0/sys/merchant/individualreginfoadd"
	urlMerchantEnterpriseRegInfoAdd = "/rest/v1.0/sys/merchant/enterprisereginfoadd"

	urlMerchantServiceMerSettleInfoUpdate = "/rest/v1.0/merchantService/mer-settle/mer-settle-info-update-for-o2o"
	urlMerchantServiceMerProductFeeUpdate = "/rest/v1.0/merchantService/mer-product-info/mer-product-fee-update-for-o2o"
	urlMerchantServiceSendMerSmsNotify    = "/rest/v1.0/merchantService/mer-out-invoke/send-mer-sms-notify"

	urlTradeOrder      = "/rest/v1.0/sys/trade/order"
	urlTradeOrderClose = "/rest/v1.0/sys/trade/orderclose"
	urlTradeOrderQuery = "/rest/v1.0/sys/trade/orderquery"
	//urlTradeMultiOrderQuery  = "/rest/v1.0/sys/trade/multiorderquery"
	urlTradeRefund           = "/rest/v1.0/sys/trade/refund"
	urlTradeRefundQuery      = "/rest/v1.0/sys/trade/refundquery"
	urlTradeDivide           = "/rest/v1.0/sys/trade/divide"
	urlTradeDivideQuery      = "/rest/v1.0/sys/trade/dividequery"
	urlTradeFullSettle       = "/rest/v1.0/sys/trade/fullsettle"
	urlTradeSettlementsQuery = "/rest/v1.0/sys/trade/settlementsquery"

	urlNcCashierApiPay = "/rest/v1.0/nccashierapi/api/pay"
	urlFileUpload      = "/rest/v1.0/file/upload"
)

//全局实例
var Instance *Business

func init() {
	fmt.Println("yos business init +++")
	if Instance == nil {
		Instance = NewDefaultBusiness()
	}
}

type Business struct {
	config *Config
	client *Client
}

func NewDefaultBusiness() *Business {
	return NewBusiness(NewDefaultConfig())
}

func NewBusiness(config *Config) *Business {
	return &Business{
		config: config,
		client: NewClient(),
	}
}

//商户余额查询
func (b *Business) MerChantBalanceQuery(merchantNo string) (*MerChantBalanceQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	response := &MerChantBalanceQueryResponse{}
	err := b.client.Post(urlMerChantBalanceQuery, request, response)
	return response, err
}

//商户产品协议获取
func (b *Business) MerChantAgreeInfoQuery(merchantNo string) (*MerChantAgreeInfoQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	response := &MerChantAgreeInfoQueryResponse{}
	err := b.client.Post(urlMerChantAgreeInfoQuery, request, response)
	return response, err
}

//授权短验重发
func (b *Business) MerchantSendAuthorizeNum(merchantNo string, phone string) (*MerchantSendAuthorizeNumResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	request.AddParam("phone", phone)
	response := &MerchantSendAuthorizeNumResponse{}
	err := b.client.Post(urlMerchantSendAuthorizeNum, request, response)
	return response, err
}

//授权码确认
func (b *Business) MerchantReceiveAuthorizeNum(merchantNo string, phone string, merAuthorizeNum string) (*MerchantReceiveAuthorizeNumResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	request.AddParam("phone", phone)
	request.AddParam("merAuthorizeNum", merAuthorizeNum)
	response := &MerchantReceiveAuthorizeNumResponse{}
	err := b.client.Post(urlMerchantReceiveAuthorizeNum, request, response)
	return response, err
}

//授权页面获取
func (b *Business) MerchantAuthorizeUrlQuery(merchantNo string) (*MerchantAuthorizeUrlQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	response := &MerchantAuthorizeUrlQueryResponse{}
	err := b.client.Post(urlMerchantAuthorizeUrlQuery, request, response)
	return response, err
}

//银行支行信息获取
func (b *Business) MerchantBankBranchInfo(headBankCode string, provinceCode string, cityCode string) (*MerchantBankBranchInfoResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("headBankCode", headBankCode)
	request.AddParam("provinceCode", provinceCode)
	request.AddParam("cityCode", cityCode)
	response := &MerchantBankBranchInfoResponse{}
	err := b.client.Post(urlMerchantBankBranchInfo, request, response)
	return response, err
}

//子账户密钥查询
func (b *Business) MerchantHmacKeyQuery(merchantNo string) (*MerchantHmacKeyQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	response := &MerchantHmacKeyQueryResponse{}
	err := b.client.Post(urlMerchantHmacKeyQuery, request, response)
	return response, err
}

//入网信息查询
func (b *Business) MerchantRegStatusQuery(merchantNo string) (*MerchantRegStatusQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	response := &MerchantRegStatusQueryResponse{}
	err := b.client.Post(urlMerchantRegStatusQuery, request, response)
	return response, err
}

//商户入网个人注册
func (b *Business) MerchantPersonRegInfoAdd(param *MerchantPersonRegInfoAddParam) (*MerchantPersonRegInfoAddResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("requestNo", request.RequestId)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &MerchantPersonRegInfoAddResponse{}
	err := b.client.Post(urlMerchantPersonRegInfoAdd, request, response)
	return response, err
}

//商户入网个体注册
func (b *Business) MerchantIndividualRegInfoAdd(param *MerchantIndividualRegInfoAddParam) (*MerchantIndividualRegInfoAddResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("requestNo", request.RequestId)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &MerchantIndividualRegInfoAddResponse{}
	err := b.client.Post(urlMerchantIndividualRegInfoAdd, request, response)
	return response, err
}

//商户入网企业注册
func (b *Business) MerchantEnterpriseRegInfoAdd(param *MerchantEnterpriseRegInfoAddParam) (*MerchantEnterpriseRegInfoAddResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("requestNo", request.RequestId)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &MerchantEnterpriseRegInfoAddResponse{}
	err := b.client.Post(urlMerchantEnterpriseRegInfoAdd, request, response)
	return response, err
}

//子商户结算银行卡修改
func (b *Business) MerchantServiceMerSettleInfoUpdate(param *MerchantServiceMerSettleInfoUpdateParam) (*MerchantServiceMerSettleInfoUpdateResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("requestNo", request.RequestId)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &MerchantServiceMerSettleInfoUpdateResponse{}
	err := b.client.Post(urlMerchantServiceMerSettleInfoUpdate, request, response)
	return response, err
}

//子商户费率修改
func (b *Business) MerchantServiceMerProductFeeUpdate(merchantNo string, payProductMap string) (*MerchantServiceMerProductFeeUpdateResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("requestNo", request.RequestId)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	request.AddParam("payProductMap", payProductMap)
	response := &MerchantServiceMerProductFeeUpdateResponse{}
	err := b.client.Post(urlMerchantServiceMerProductFeeUpdate, request, response)
	return response, err
}

//子商户结算卡修改获取短验码
func (b *Business) MerchantServiceSendMerSmsNotify(merchantNo string, phone string) (*MerchantServiceSendMerSmsNotifyResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParam("merchantNo", merchantNo)
	request.AddParam("phone", phone)
	request.AddParam("merchantNo", "NET_IN")
	response := &MerchantServiceSendMerSmsNotifyResponse{}
	err := b.client.Post(urlMerchantServiceSendMerSmsNotify, request, response)
	return response, err
}

//创建订单
func (b *Business) TradeOrder(param *TradeOrderParam) (*TradeOrderResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeOrderResponse{}
	err := b.client.Post(urlTradeOrder, request, response)
	return response, err
}

//关闭订单
func (b *Business) TradeOrderClose(param *TradeOrderCloseParam) (*TradeOrderCloseResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeOrderCloseResponse{}
	err := b.client.Post(urlTradeOrderClose, request, response)
	return response, err
}

//查询订单
func (b *Business) TradeOrderQuery(param *TradeOrderQueryParam) (*TradeOrderQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeOrderQueryResponse{}
	err := b.client.Post(urlTradeOrderQuery, request, response)
	return response, err
}

//订单退款
func (b *Business) TradeRefund(param *TradeRefundParam) (*TradeRefundResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeRefundResponse{}
	err := b.client.Post(urlTradeRefund, request, response)
	return response, err
}

//订单退款查询
func (b *Business) TradeRefundQuery(param *TradeRefundQueryParam) (*TradeRefundQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeRefundQueryResponse{}
	err := b.client.Post(urlTradeRefundQuery, request, response)
	return response, err
}

//订单分账
func (b *Business) TradeDivide(param *TradeDivideParam) (*TradeDivideResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeDivideResponse{}
	err := b.client.Post(urlTradeDivide, request, response)
	return response, err
}

//订单分账查询
func (b *Business) TradeDivideQuery(param *TradeDivideQueryParam) (*TradeDivideQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeDivideQueryResponse{}
	err := b.client.Post(urlTradeDivideQuery, request, response)
	return response, err
}

//资金全额确认
func (b *Business) TradeFullSettle(param *TradeFullSettleParam) (*TradeFullSettleResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeFullSettleResponse{}
	err := b.client.Post(urlTradeFullSettle, request, response)
	return response, err
}

//结算查询
func (b *Business) TradeSettlementsQuery(param *TradeSettlementsQueryParam) (*TradeSettlementsQueryResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &TradeSettlementsQueryResponse{}
	err := b.client.Post(urlTradeSettlementsQuery, request, response)
	return response, err
}

//聚合API收银台
func (b *Business) NcCashierApiPay(param *NcCashierApiPayParam) (*NcCashierApiPayResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("parentMerchantNo", b.config.AppKey)
	request.AddParamsWithStruct(param)
	response := &NcCashierApiPayResponse{}
	err := b.client.Post(urlNcCashierApiPay, request, response)
	return response, err
}

//文件上传
func (b *Business) FileUpload(filePath string) (*FileUploadResponse, error) {
	request := NewRequest(b.config)
	request.AddParam("fileType", "IMAGE")
	request.AddFile("_file", filePath)
	response := &FileUploadResponse{}
	err := b.client.Post(urlFileUpload, request, response)
	return response, err
}
