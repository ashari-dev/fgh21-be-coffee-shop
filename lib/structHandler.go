package lib

type Respont struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"pageInfo,omitempty"`
	Result   any    `json:"result,omitempty"`
}

type PageInfo struct {
	TotalData int  `json:"totalData"`
	TotalPage int  `json:"totalPage"`
	Page      int  `json:"page"`
	Limit     int  `json:"limit"`
	Next      *int `json:"next,omitempty"`
	Prev      *int `json:"prev,omitempty"`
}
