package protocol

type UserAgent struct {
	DeviceType   string `json:"deviceType"`
	Locale       string `json:"locale"`
	DeviceLocale string `json:"deviceLocale"`
	OS           string `json:"osVersion"`
	BrowserName  string `json:"deviceName"`
	Header       string `json:"headerUserAgent"`
	AppVersion   string `json:"appVersion"`
	Screen       string `json:"screen"`
	Timezone     string `json:"timezone"`
}

var DefaultUserAgent = &UserAgent{
	DeviceType:   "WEB",
	Locale:       "en",
	DeviceLocale: "en",
	OS:           "Windows",
	BrowserName:  "Firefox",
	Header:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:140.0) Gecko/20100101 Firefox/140.0",
	AppVersion:   "26.5.8",
	Screen:       "1067x1707 1.5x",
	Timezone:     "Europe/Moscow",
}
