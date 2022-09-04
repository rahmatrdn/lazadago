package lazadago

import (
	"fmt"
	"testing"

	lazadaConfig "github.com/rahmatrdn/lazadago/config"
)

func TestLazada(t *testing.T) {
	api := NewApi(&lazadaConfig.Config{
		AppKey:      "your app key",
		AccessToken: "your Access Token", //刚开始可以为空字符串
		AppSecret:   "your app Secret",
		Country:     "ph",
	})
	fmt.Println(api.AuthorizationURL("callback"))
}
