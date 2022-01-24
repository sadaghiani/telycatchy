package storage

import "time"

// Result standard result
type Result struct {
	Content ResultContent `json:"content,omitempty"`
	Error   string        `json:"error,omitempty"`
}

// Result standard result
type ResultContent struct {
	Data     interface{} `json:"data,omitempty"`
	DataTime int64       `json:"dataTime,omitempty"`
}

func NewResultContent(data interface{}) *Result {
	return &Result{
		Content: ResultContent{
			Data:     data,
			DataTime: time.Now().Unix(),
		},
		Error: "",
	}
}

func NewResultError(err string) *Result {
	return &Result{
		Content: ResultContent{},
		Error:   err,
	}
}
