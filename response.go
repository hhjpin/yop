package yop

import (
	"encoding/json"
	"errors"
	"fmt"
)

type BaseResponse struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	SubCode    string `json:"subCode"`
	SubMessage string `json:"subMessage"`
	RequestId  string `json:"requestId"`
}

type BaseResult struct {
	Code             string `json:"code"`
	Message          string `json:"message"`
	ReturnMsg        string `json:"returnMsg"`
	ReturnCode       string `json:"returnCode"`
	ParentMerchantNo string `json:"ParentMerchantNo"`
	MerchantNo       string `json:"MerchantNo"`
	RequestNo        string `json:"RequestNo"`
}

//通用结果
type CommonResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//商户查询余额结果
type MerChantBalanceQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerBalance string `json:"merBalance"`
	} `json:"result"`
}

//商户产品协议查询结果
type MerChantAgreeInfoQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		ExternalId       string `json:"externalId"`
		AgreementContent string `json:"agreementContent"`
	} `json:"result"`
}

//授权短验重发结果
type MerchantSendAuthorizeNumResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//授权码确认结果
type MerchantReceiveAuthorizeNumResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//授权页面获取结果
type MerchantAuthorizeUrlQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerAuthorizeWebUrl string `json:"merAuthorizeWebUrl"`
		Phone              string `json:"phone"`
	} `json:"result"`
}

//银行支行信息获取结果
type MerchantBankBranchInfoResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		BranchBankInfo string `json:"branchBankInfo"`
	} `json:"result"`
}

//子账户密钥查询结果
type MerchantHmacKeyQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerHmacKey string `json:"merHmacKey"`
	} `json:"result"`
}

//入网信息查询结果
type MerchantRegStatusQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		ExternalId        string `json:"externalId"`
		MerNetInOutStatus string `json:"merNetInOutStatus"`
	} `json:"result"`
}

//商户入网个人注册结果
type MerchantPersonRegInfoAddResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		ExternalId       string `json:"externalId"`
		AgreementContent string `json:"agreementContent"`
	} `json:"result"`
}

//商户入网个体注册结果
type MerchantIndividualRegInfoAddResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		ExternalId       string `json:"externalId"`
		AgreementContent string `json:"agreementContent"`
	} `json:"result"`
}

//商户入网企业注册结果
type MerchantEnterpriseRegInfoAddResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		ExternalId       string `json:"externalId"`
		AgreementContent string `json:"agreementContent"`
	} `json:"result"`
}

//子商户结算银行卡修改结果
type MerchantServiceMerSettleInfoUpdateResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//子商户费率修改结果
type MerchantServiceMerProductFeeUpdateResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//子商户结算卡修改获取短验码结果
type MerchantServiceSendMerSmsNotifyResponse struct {
	BaseResponse
	Result struct {
		BaseResult
	} `json:"result"`
}

//创建订单结果
type TradeOrderResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		OrderId         string `json:"orderId"`
		UniqueOrderNo   string `json:"uniqueOrderNo"`
		GoodsParamExt   string `json:"goodsParamExt"`
		Memo            string `json:"memo"`
		Token           string `json:"token"`
		FundProcessType string `json:"fundProcessType"`
	} `json:"result"`
}

//关闭订单结果
type TradeOrderCloseResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		BizSystemNo         string `json:"bizSystemNo"`
		OrderId             string `json:"orderId"`
		OrderCloseRequestId string `json:"orderCloseRequestId"`
	} `json:"result"`
}

//查询订单结果
type TradeOrderQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		BizSystemNo          string `json:"bizSystemNo"`
		OrderId              string `json:"orderId"`
		UniqueOrderNo        string `json:"uniqueOrderNo"`
		Status               string `json:"status"`
		OrderAmount          string `json:"orderAmount"`
		PayAmount            string `json:"payAmount"`
		MerchantFee          string `json:"merchantFee"`
		CustomerFee          string `json:"customerFee"`
		RequestDate          string `json:"requestDate"`
		PaySuccessDate       string `json:"paySuccessDate"`
		GoodsParamExt        string `json:"goodsParamExt"`
		Memo                 string `json:"memo"`
		InstCompany          string `json:"instCompany"`
		InstNumber           string `json:"instNumber"`
		IndustryParamExt     string `json:"industryParamExt"`
		CashierType          string `json:"cashierType"`
		PaymentProduct       string `json:"paymentProduct"`
		Token                string `json:"token"`
		BankId               string `json:"bankId"`
		CardType             string `json:"cardType"`
		PlatformType         string `json:"platformType"`
		PayURL               string `json:"payURL"`
		OpenID               string `json:"openID"`
		UnionID              string `json:"unionID"`
		FundProcessType      string `json:"fundProcessType"`
		BizChannelId         string `json:"bizChannelId"`
		BankPaySuccessDate   string `json:"bankPaySuccessDate"`
		BankTrxId            string `json:"bankTrxId"`
		HaveAccounted        string `json:"haveAccounted"`
		AccountPayMerchantNo string `json:"accountPayMerchantNo"`
		ResidualDivideAmount string `json:"residualDivideAmount"`
		PreauthStatus        string `json:"preauthStatus"`
		PreauthType          string `json:"preauthType"`
		PreauthAmount        string `json:"preauthAmount"`
		PaymentSysNo         string `json:"paymentSysNo"`
		PaymentStatus        string `json:"paymentStatus"`
		DivideRequestId      string `json:"divideRequestId"`
		ParentMerchantName   string `json:"parentMerchantName"`
		MerchantName         string `json:"merchantName"`
		CombPaymentDTO       struct {
			SecondPayOrderNo     string `json:"secondPayOrderNo"`
			SecondBankOrderNo    string `json:"secondBankOrderNo"`
			SecondAmount         string `json:"secondAmount"`
			SecondPaySuccessDate string `json:"secondPaySuccessDate"`
		} `json:"combPaymentDTO"`
		CashFee                  string `json:"cashFee"`
		SettlementFee            string `json:"settlementFee"`
		BankOrderId              string `json:"bankOrderId"`
		BankCardNo               string `json:"bankCardNo"`
		CalFeeMerchantNo         string `json:"calFeeMerchantNo"`
		BankPromotionInfoDTOList []struct {
			Id                 string `json:"id"`
			PayInterface       string `json:"payInterface"`
			BankOrderNo        string `json:"bankOrderNo"`
			TradeType          string `json:"tradeType"`
			PromotionId        string `json:"promotionId"`
			PromotionName      string `json:"promotionName"`
			PromotionScope     string `json:"promotionScope"`
			PromotionType      string `json:"promotionType"`
			Amount             string `json:"amount"`
			AmountRefund       string `json:"amountRefund"`
			ActivityId         string `json:"activityId"`
			ChannelContribute  string `json:"channelContribute"`
			MerchantContribute string `json:"merchantContribute"`
			OtherContribute    string `json:"otherContribute"`
			Memo               string `json:"memo"`
			OrderTime          string `json:"orderTime"`
		} `json:"bankPromotionInfoDTOList"`
		MobilePhoneNo string `json:"mobilePhoneNo"`
	} `json:"result"`
}

//订单退款结果
type TradeRefundResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		BizSystemNo       string `json:"bizSystemNo"`
		OrderId           string `json:"orderId"`
		RefundRequestId   string `json:"refundRequestId"`
		UniqueRefundNo    string `json:"uniqueRefundNo"`
		Status            string `json:"status"`
		RefundAmount      string `json:"refundAmount"`
		ResidualAmount    string `json:"residualAmount"`
		Description       string `json:"description"`
		RefundRequestDate string `json:"refundRequestDate"`
		AccountDivided    string `json:"accountDivided"`
	} `json:"result"`
}

//订单退款结果
type TradeRefundQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		OrderId           string `json:"orderId"`
		RefundRequestId   string `json:"refundRequestId"`
		UniqueOrderNo     string `json:"uniqueOrderNo"`
		UniqueRefundNo    string `json:"uniqueRefundNo"`
		RefundAmount      string `json:"refundAmount"`
		ReturnMerchantFee string `json:"returnMerchantFee"`
		ReturnCustomerFee string `json:"returnCustomerFee"`
		Status            string `json:"status"`
		Description       string `json:"description"`
		RefundRequestDate string `json:"refundRequestDate"`
		RefundSuccessDate string `json:"refundSuccessDate"`
		RealDeductAmount  string `json:"realDeductAmount"`
		RealRefundAmount  string `json:"realRefundAmount"`
		AccountDivided    string `json:"accountDivided"`
	} `json:"result"`
}

//订单分账结果
type TradeDivideResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerchantNo       string `json:"merchant_no"`
		ParentMerchantNo string `json:"parent_merchant_no"`
		BizSystemNo      string `json:"biz_system_no"`
		OrderId          string `json:"order_id"`
		UniqueOrderNo    string `json:"unique_order_no"`
		DivideRequestId  string `json:"divide_request_id"`
		Status           string `json:"status"`
		DivideDetail     string `json:"divide_detail"`
	} `json:"result"`
}

//订单分账查询结果
type TradeDivideQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerchantNo       string `json:"merchant_no"`
		ParentMerchantNo string `json:"parent_merchant_no"`
		BizSystemNo      string `json:"biz_system_no"`
		OrderId          string `json:"order_id"`
		UniqueOrderNo    string `json:"unique_order_no"`
		Status           string `json:"status"`
		DivideDetail     string `json:"divide_detail"`
	} `json:"result"`
}

//资金全额确认结果
type TradeFullSettleResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		MerchantNo       string `json:"merchant_no"`
		ParentMerchantNo string `json:"parent_merchant_no"`
		BizSystemNo      string `json:"biz_system_no"`
		OrderId          string `json:"order_id"`
		UniqueOrderNo    string `json:"unique_order_no"`
		DivideRequestId  string `json:"divide_request_id"`
	} `json:"result"`
}

//结算查询结果
type TradeSettlementsQueryResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		StartSettleDate  string `json:"startSettleDate"`
		EndSettleDate    string `json:"endSettleDate"`
		SettleRecordList []struct {
			SettlementDate          string `json:"settlementDate"`
			BeginSettlementDate     string `json:"beginSettlementDate"`
			EndSettlementDate       string `json:"endSettlementDate"`
			SumRealSettlementAmount string `json:"sumRealSettlementAmount"`
			RemitFee                string `json:"remitFee"`
			SumNetAmount            string `json:"sumNetAmount"`
			SettlementType          string `json:"settlementType"`
			RemitStatus             string `json:"remitStatus"`
			FailMsg                 string `json:"failMsg"`
		} `json:"settleRecordList"`
	} `json:"result"`
}

//聚合API收银台结果
type NcCashierApiPayResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		PayTool     string `json:"payTool"`
		PayType     string `json:"payType"`
		Token       string `json:"token"`
		ResultType  string `json:"resultType"`
		ResultData  string `json:"resultData"`
		ExtParamMap string `json:"extParamMap"`
	} `json:"result"`
}

//上传文件结果
type FileUploadResponse struct {
	BaseResponse
	Result struct {
		BaseResult
		TotalCount   int `json:"total_count"`
		SuccessCount int `json:"success_count"`
		Files        []struct {
			CreatedDate  string `json:"createdDate"`
			AppKey       string `json:"appKey"`
			FileId       string `json:"fileId"`
			FileName     string `json:"fileName"`
			FileType     string `json:"fileType"`
			FileLocation string `json:"fileLocation"`
			FileSize     int    `json:"fileSize"`
			Md5          string `json:"md5"`
			TempStorage  bool   `json:"tempStorage"`
		} `json:"files"`
	} `json:"result"`
}

//直接结果
type Response struct {
	body []byte
}

func (r *Response) JsonUnmarshal(response interface{}) error {
	return json.Unmarshal(r.body, response)
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) checkError() error {
	resp := CommonResponse{}
	err := json.Unmarshal(r.body, &resp)
	if err != nil {
		return err
	}
	if resp.Code != "" || (resp.Result.ReturnCode != "REG00000" && resp.Result.ReturnCode != "") {
		return errors.New(fmt.Sprintf("YOP:%s:%s:%s", resp.Message, resp.SubMessage, resp.Result.ReturnMsg))
	}
	return nil
}
