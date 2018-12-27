package alipay

// see: https://docs.open.alipay.com/284/106001/
type SystemOauthToken struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (s SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (s SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	if s.Code != "" {
		m["grant_type"] = "authorization_code"
		m["code"] = s.Code
	} else {
		m["grant_type"] = "refresh_token"
		m["refresh_token"] = s.RefreshToken
	}

	return m
}

func (s SystemOauthToken) ExtJSONParamName() string {
	return ""
}

func (s SystemOauthToken) ExtJSONParamValue() string {
	return ""
}

type SystemOauthTokenResponse struct {
	AlipaySystemOauthTokenResponse struct {
		AccessToken  string `json:"access_token"`
		UserID       string `json:"user_id"`
		ExpiresIn    int    `json:"expires_in"`
		ReExpiresIn  int    `json:"re_expires_in"`
		RefreshToken string `json:"refresh_token"`
	} `json:"alipay_system_oauth_token_response"`
	Sign string `json:"sign"`
}
