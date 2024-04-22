package auth

import (
	"fmt"
	"testing"

	lazadaConfig "github.com/zhouzongyan/lazadago/config"
)

func TestAuth(t *testing.T) {
	cfg := &lazadaConfig.Config{
		AppKey:      "113865",
		AccessToken: "yXVeaOlP2VY04gwOZzx1ig6BxfEUrZUk",
		Country:     "ph",
	}
	a := Auth{
		Config: cfg,
	}
	fmt.Println(a.AuthorizationURL("https://www.xingtool.cn/lazada_callback"))
}
