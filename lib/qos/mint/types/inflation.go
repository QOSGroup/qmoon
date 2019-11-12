package types

import "time"

// 通胀阶段
type InflationPhrase struct {
	EndTime       time.Time `json:"end_time"`       // 结束时间
	TotalAmount   string    `json:"total_amount"`   // 通胀总量
	AppliedAmount string    `json:"applied_amount"` // 发行总量
}
