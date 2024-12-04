package dto

// CustomInsertReq 创建请求结构
type CustomInsertReq struct {
	Content interface{} `json:"content" binding:"required"` // 接收任意类型数据
}

// CustomGetResp 响应结构
type CustomGetResp struct {
	Id      uint        `json:"id"`
	Content interface{} `json:"content"`
}
