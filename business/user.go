package business

import (
	"github.com/go-resty/resty/v2"
)

func CheckExternalUser(username string, password string) bool {
	client := resty.New()
	resp, err := client.R().
		SetFormData(map[string]string{
			"username": username,
			"password": password,
		}).
		SetHeader("User-Agent", "ZCST XMT OA/1.0").
		Post("https://authserver.zcst.edu.cn/cas/v1/tickets")
	return err == nil && resp.IsSuccess()
}
