package alipay

// see: https://docs.open.alipay.com/284/106001/

// 第三步: 使用auth_code换取接口access_token及用户userId

type SystemOauthToken struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (this SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (this SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	if this.Code != "" {
		m["grant_type"] = "authorization_code"
		m["code"] = this.Code
	} else {
		m["grant_type"] = "refresh_token"
		m["refresh_token"] = this.RefreshToken
	}

	return m
}

func (this SystemOauthToken) ExtJSONParamName() string {
	return ""
}

func (this SystemOauthToken) ExtJSONParamValue() string {
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

// 第四步: 调用接口获取用户信息
// 如果scope=auth_base，在第三步就可以获取到用户的userId，无需走第四步。
// 如果scope=auth_user，才需要走第四步，通过access_token调用用户信息共享接口获取用户信息。

type UserInfoShare struct {
	Code      string `json:"code"`
	AuthToken string `json:"auth_token"`
}

func (this UserInfoShare) APIName() string {
	return "alipay.user.info.share"
}

func (this UserInfoShare) Params() map[string]string {
	var m = make(map[string]string)

	m["grant_type"] = "authorization_code"
	m["code"] = this.Code
	m["auth_token"] = this.AuthToken

	return m
}

func (this UserInfoShare) ExtJSONParamName() string {
	return ""
}

func (this UserInfoShare) ExtJSONParamValue() string {
	return ""
}

type UserInfoShareResponse struct {
	AlipayUserInfoShareResponse struct {
		Code               string `json:"code"`
		Msg                string `json:"msg"`
		UserID             string `json:"user_id"`
		Avatar             string `json:"avatar"`
		UserType           string `json:"user_type"`
		UserStatus         string `json:"user_status"`
		IsCertified        string `json:"is_certified"`
		Province           string `json:"province"`
		City               string `json:"city"`
		NickName           string `json:"nick_name"`
		IsStudentCertified string `json:"is_student_certified"`
		Gender             string `json:"gender"`
	} `json:"alipay_user_info_share_response"`
	Sign string `json:"sign"`
}
