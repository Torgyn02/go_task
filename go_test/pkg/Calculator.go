package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SdekAPI struct {
	Token      string
	TestMode   bool
	APIAddress string
}

type Address struct {
	CountryCode string `json:"country_code"`
	Postcode    string `json:"postcode"`
	City        string `json:"city"`
	Street      string `json:"street"`
	House       string `json:"house"`
}

type Size struct {
	Weight float64
	Length float64
	Width  float64
	Height float64
}

type PriceSending struct {
	TariffCode        int    `json:"tariff_code"`
	TariffName        string `json:"tariff_name"`
	TariffDescription string `json:"tariff_description"`
	DeliveryMode      int    `json:"delivery_mode"`
	DeliveryAmount    float64
	PeriodMin         int `json:"period_min"`
	PeriodMax         int `json:"period_max"`
}

type SdekResponse struct {
	FareCodes []PriceSending `json:"fare_codes"`
}

func (s *SdekAPI) Calculate(addrFrom, addrTo Address, size Size) ([]PriceSending, error) {
	form := url.Values{}
	form.Add("version", "1.0")
	form.Add("json", "true")
	form.Add("secure", "1")
	form.Add("auth_login", s.Token)
	form.Add("sender_city_postcode", addrFrom.Postcode)
	form.Add("receiver_city_postcode", addrTo.Postcode)
	form.Add("tariff_list", "136,137,233") // Tariff codes for standard and business delivery
	form.Add("goods_weight", strconv.FormatFloat(size.Weight, 'f', 2, 64))
	form.Add("goods_length", strconv.FormatFloat(size.Length, 'f', 2, 64))
	form.Add("goods_width", strconv.FormatFloat(size.Width, 'f', 2, 64))
	form.Add("goods_height", strconv.FormatFloat(size.Height, 'f', 2, 64))

	apiUrl := s.APIAddress
	if s.TestMode {
		apiUrl = strings.Replace(apiUrl, "oauth", "edu", 1)
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sdekResp SdekResponse
	err = json.Unmarshal(body, &sdekResp)
	if err != nil {
		return nil, err
	}

	return sdekResp.FareCodes, nil
}
