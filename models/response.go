package models

type NetworkResponse struct {
	Status    int         `json:"status"`    // 状态码
	Message   string      `json:"message"`   // 消息
	Data      interface{} `json:"data"`      // 数据
	Timestamp int64       `json:"timestamp"` // 时间戳
}
