package yop

type MerchantPersonRegInfoAddParam struct {
	LegalName        string `json:"LegalName"`
	LegalIdCard      string `json:"LegalIdCard"`
	MerLegalPhone    string `json:"MerLegalPhone"`
	MerLegalEmail    string `json:"MerLegalEmail"`
	MerLevel1No      string `json:"MerLevel1No"`
	MerLevel2No      string `json:"MerLevel2No"`
	MerProvince      string `json:"MerProvince"`
	MerCity          string `json:"MerCity"`
	MerDistrict      string `json:"MerDistrict"`
	MerAddress       string `json:"MerAddress"`
	MerScope         string `json:"MerScope"`
	CardNo           string `json:"CardNo"`
	HeadBankCode     string `json:"HeadBankCode"`
	BankCode         string `json:"BankCode"`
	BankProvince     string `json:"BankProvince"`
	BankCity         string `json:"BankCity"`
	ProductInfo      string `json:"ProductInfo"`
	FileInfo         string `json:"FileInfo"`
	BusinessFunction string `json:"BusinessFunction"`
	NotifyUrl        string `json:"NotifyUrl"`
	MerAuthorizeType string `json:"MerAuthorizeType"`
}

type MerchantIndividualRegInfoAddParam struct {
	MerFullName      string `json:"merFullName"`
	MerShortName     string `json:"merShortName"`
	MerCertNo        string `json:"merCertNo"`
	LegalName        string `json:"legalName"`
	LegalIdCard      string `json:"legalIdCard"`
	MerLegalPhone    string `json:"merLegalPhone"`
	MerLegalEmail    string `json:"merLegalEmail"`
	MerLevel1No      string `json:"merLevel1No"`
	MerLevel2No      string `json:"merLevel2No"`
	MerProvince      string `json:"merProvince"`
	MerCity          string `json:"merCity"`
	MerDistrict      string `json:"merDistrict"`
	MerAddress       string `json:"merAddress"`
	MerScope         string `json:"merScope"`
	CardNo           string `json:"cardNo"`
	HeadBankCode     string `json:"headBankCode"`
	BankCode         string `json:"bankCode"`
	BankProvince     string `json:"bankProvince"`
	BankCity         string `json:"bankCity"`
	ProductInfo      string `json:"productInfo"`
	FileInfo         string `json:"fileInfo"`
	BusinessFunction string `json:"businessFunction"`
	NotifyUrl        string `json:"notifyUrl"`
	MerAuthorizeType string `json:"merAuthorizeType"`
}

type MerchantEnterpriseRegInfoAddParam struct {
	MerFullName      string `json:"merFullName"`
	MerShortName     string `json:"merShortName"`
	MerCertNo        string `json:"merCertNo"`
	LegalName        string `json:"legalName"`
	LegalIdCard      string `json:"legalIdCard"`
	MerLegalPhone    string `json:"merLegalPhone"`
	MerLegalEmail    string `json:"merLegalEmail"`
	MerLevel1No      string `json:"merLevel1No"`
	MerLevel2No      string `json:"merLevel2No"`
	MerProvince      string `json:"merProvince"`
	MerCity          string `json:"merCity"`
	MerDistrict      string `json:"merDistrict"`
	MerAddress       string `json:"merAddress"`
	MerContactName   string `json:"merContactName"`
	MerContactPhone  string `json:"merContactPhone"`
	MerContactEmail  string `json:"merContactEmail"`
	TaxRegistCert    string `json:"taxRegistCert"`
	AccountLicense   string `json:"accountLicense"`
	OrgCode          string `json:"orgCode"`
	IsOrgCodeLong    string `json:"isOrgCodeLong"`
	OrgCodeExpiry    string `json:"orgCodeExpiry"`
	CardNo           string `json:"cardNo"`
	HeadBankCode     string `json:"headBankCode"`
	BankCode         string `json:"bankCode"`
	BankProvince     string `json:"bankProvince"`
	BankCity         string `json:"bankCity"`
	ProductInfo      string `json:"productInfo"`
	FileInfo         string `json:"fileInfo"`
	BusinessFunction string `json:"businessFunction"`
	NotifyUrl        string `json:"notifyUrl"`
	MerAuthorizeType string `json:"merAuthorizeType"`
}

type MerchantServiceMerSettleInfoUpdateParam struct {
	MerchantNo      string `json:"merchantNo"`
	MerAuthorizeNum string `json:"merAuthorizeNum"`
	BankcardNo      string `json:"bankcardNo"`
	HeadBankName    string `json:"headBankName"`
	BankName        string `json:"bankName"`
	BankProvince    string `json:"bankProvince"`
	BankCity        string `json:"bankCity"`
	CallbackUrl     string `json:"callbackurl"`
}

type TradeOrderParam struct {
	MerchantNo         string `json:"merchantNo"`
	OrderId            string `json:"orderId"`
	OrderAmount        string `json:"orderAmount"`
	TimeoutExpress     string `json:"timeoutExpress"`
	TimeoutExpressType string `json:"timeoutExpressType"`
	RequestDate        string `json:"requestDate"`
	RedirectUrl        string `json:"redirectUrl"`
	NotifyUrl          string `json:"notifyUrl"`
	AssureType         string `json:"assureType"`
	AssurePeriod       string `json:"assurePeriod"`
	GoodsParamExt      string `json:"goodsParamExt"`
	PaymentParamExt    string `json:"paymentParamExt"`
	IndustryParamExt   string `json:"industryParamExt"`
	RiskParamExt       string `json:"riskParamExt"`
	Memo               string `json:"memo"`
	FundProcessType    string `json:"fundProcessType"`
	Hmac               string `json:"hmac"`
	DivideDetail       string `json:"divideDetail"`
	CsUrl              string `json:"csUrl"`
	DivideNotifyUrl    string `json:"divideNotifyUrl"`
	TimeoutNotifyUrl   string `json:"timeoutNotifyUrl"`
}

type TradeOrderCloseParam struct {
	MerchantNo    string `json:"merchantNo"`
	OrderId       string `json:"orderId"`
	UniqueOrderNo string `json:"uniqueOrderNo"`
	Description   string `json:"description"`
	Hmac          string `json:"hmac"`
}

type TradeOrderQueryParam struct {
	MerchantNo    string `json:"merchantNo"`
	OrderId       string `json:"orderId"`
	UniqueOrderNo string `json:"uniqueOrderNo"`
	Hmac          string `json:"hmac"`
}

type TradeRefundParam struct {
	MerchantNo      string `json:"merchantNo"`
	OrderId         string `json:"orderId"`
	RefundRequestId string `json:"refundRequestId"`
	UniqueOrderNo   string `json:"uniqueOrderNo"`
	RefundAmount    string `json:"refundAmount"`
	AccountDivided  string `json:"accountDivided"`
	Description     string `json:"description"`
	Memo            string `json:"memo"`
	NotifyUrl       string `json:"notifyUrl"`
	Hmac            string `json:"hmac"`
}

type TradeRefundQueryParam struct {
	MerchantNo      string `json:"merchantNo"`
	OrderId         string `json:"orderId"`
	RefundRequestId string `json:"refundRequestId"`
	UniqueRefundNo  string `json:"uniqueRefundNo"`
	Hmac            string `json:"hmac"`
}

type TradeDivideParam struct {
	MerchantNo               string `json:"merchantNo"`
	OrderId                  string `json:"orderId"`
	UniqueOrderNo            string `json:"uniqueOrderNo"`
	DivideRequestId          string `json:"divideRequestId"`
	ContractNo               string `json:"contractNo"`
	DivideDetail             string `json:"divideDetail"`
	IsUnfreezeResidualAmount string `json:"isUnfreezeResidualAmount"`
	Hmac                     string `json:"hmac"`
}

type TradeDivideQueryParam struct {
	MerchantNo      string `json:"merchantNo"`
	OrderId         string `json:"orderId"`
	UniqueOrderNo   string `json:"uniqueOrderNo"`
	DivideRequestId string `json:"divideRequestId"`
	Hmac            string `json:"hmac"`
}

type TradeFullSettleParam struct {
	MerchantNo    string `json:"merchantNo"`
	OrderId       string `json:"orderId"`
	UniqueOrderNo string `json:"uniqueOrderNo"`
	Hmac          string `json:"hmac"`
}

type TradeSettlementsQueryParam struct {
	MerchantNo      string `json:"merchantNo"`
	StartSettleDate string `json:"startSettleDate"`
	EndSettleDate   string `json:"endSettleDate"`
	Hmac            string `json:"hmac"`
}

type NcCashierApiPayParam struct {
	PayTool            string `json:"payTool"`
	PayType            string `json:"payType"`
	Token              string `json:"token"`
	AppId              string `json:"appId"`
	OpenId             string `json:"openId"`
	Version            string `json:"version"`
	PayEmpowerNo       string `json:"payEmpowerNo"`
	MerchantTerminalId string `json:"merchantTerminalId"`
	MerchantStoreNo    string `json:"merchantStoreNo"`
	UserIp             string `json:"userIp"`
	ExtParamMap        string `json:"extParamMap"`
	UserNo             string `json:"userNo"`
	UserType           string `json:"userType"`
}
