package alipay

// SystemOauthToken https://docs.open.alipay.com/284/106001/
func (this *AliPay) SystemOauthToken(param SystemOauthToken) (results *SystemOauthTokenResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
