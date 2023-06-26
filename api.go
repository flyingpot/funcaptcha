package funcaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GetTokenOptions struct {
	PKey     string            `json:"pkey"`
	SURL     string            `json:"surl,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
	Site     string            `json:"site,omitempty"`
	Location string            `json:"location,omitempty"`
	Proxy    string            `json:"proxy,omitempty"`
}

type GetTokenResult struct {
	ChallengeURL          string `json:"challenge_url"`
	ChallengeURLCDN       string `json:"challenge_url_cdn"`
	ChallengeURLCDNSRI    string `json:"challenge_url_cdn_sri"`
	DisableDefaultStyling bool   `json:"disable_default_styling"`
	IFrameHeight          int    `json:"iframe_height"`
	IFrameWidth           int    `json:"iframe_width"`
	KBio                  bool   `json:"kbio"`
	MBio                  bool   `json:"mbio"`
	NoScript              string `json:"noscript"`
	TBio                  bool   `json:"tbio"`
	Token                 string `json:"token"`
}

func GetToken(options *GetTokenOptions) (GetTokenResult, error) {
	if options.SURL == "" {
		options.SURL = "https://client-api.arkoselabs.com"
	}
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}
	if _, ok := options.Headers["User-Agent"]; !ok {
		options.Headers["User-Agent"] = DEFAULT_USER_AGENT
	}

	options.Headers["Accept-Language"] = "en-US,en;q=0.9"
	options.Headers["Sec-Fetch-Site"] = "same-origin"
	options.Headers["Accept"] = "*/*"
	options.Headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	options.Headers["sec-fetch-mode"] = "cors"

	if options.Site != "" {
		options.Headers["Origin"] = options.SURL
		options.Headers["Referer"] = fmt.Sprintf("%s/v2/%s/1.4.3/enforcement.%s.html", options.SURL, options.PKey, Random())
	}

	ua := options.Headers["User-Agent"]
	formData := url.Values{
		"bda":         {GetBda(ua, options.Headers["Referer"], options.Location)},
		"public_key":  {options.PKey},
		"site":        {options.Site},
		"userbrowser": {ua},
		"rnd":         {strconv.FormatFloat(rand.Float64(), 'f', -1, 64)},
	}

	if options.Site == "" {
		formData.Del("site")
	}

	for key, value := range options.Data {
		formData.Add("data["+key+"]", value)
	}

	form := strings.ReplaceAll(formData.Encode(), "+", "%20")
	form = strings.ReplaceAll(form, "%28", "(")
	form = strings.ReplaceAll(form, "%29", ")")
	req, err := http.NewRequest("POST", options.SURL+"/fc/gt2/public_key/"+options.PKey, bytes.NewBufferString(form))
	if err != nil {
		return GetTokenResult{}, err
	}

	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GetTokenResult{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetTokenResult{}, err
	}

	var result GetTokenResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return GetTokenResult{}, err
	}

	return result, nil
}
