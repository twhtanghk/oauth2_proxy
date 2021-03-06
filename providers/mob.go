package providers

import (
	"log"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bitly/oauth2_proxy/api"
)

type MobProvider struct {
	*ProviderData
}

func NewMobProvider(p *ProviderData) *MobProvider {
	p.ProviderName = "mob"
	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   "mob.myvnc.com",
			Path:   "/org/oauth2/authorize/",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   "mob.myvnc.com",
			Path:   "/org/oauth2/token/",
		}
	}
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   "mob.myvnc.com",
			Path:   "/org/oauth2/verify/",
		}
	}
	if p.Scope == "" {
		p.Scope = "https://mob.myvnc.com/org/users"
	}
	return &MobProvider{ProviderData: p}
}

func (p *MobProvider) GetEmailAddress(s *SessionState) (string, error) {

	req, err := http.NewRequest("GET", p.ValidateURL.String(), nil)
        req.Header = make(http.Header)
        req.Header.Set("Accept", "application/json")
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("user").Get("email").String()
}
