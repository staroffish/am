package view

import "encoding/json"

// JSONRequest - 页面的请求
type JSONRequest struct {
	// 方法
	Method string `json:"method"`
	// 参数
	Params []interface{} `json:"params"`
}

// ParsePostData - 解析JSON
func ParsePostData(postData []byte) (*JSONRequest, error) {
	var jr JSONRequest
	err := json.Unmarshal(postData, &jr)
	if err != nil {
		return nil, err
	}
	return &jr, nil
}
