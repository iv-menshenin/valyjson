package race

//json:marshal
type TestStruct struct {
	HasCodes bool   `json:"has_codes,omitempty"`
	Referer  string `json:"referer,omitempty"`
	SiteID   string `json:"site_id,omitempty"`
	URL      string `json:"url,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

//json:marshal
type TestStructTyped struct {
	HasCodes bool   `json:"has_codes,omitempty"`
	Referer  string `json:"referer,omitempty"`
	SiteID   string `json:"site_id,omitempty"`
	URL      URL    `json:"url,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

type URL string
