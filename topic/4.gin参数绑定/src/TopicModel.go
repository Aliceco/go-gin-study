package src
type Topic struct {
	// 把数据当成json输入的时候会取json后面的id来当字段（映射出来的字段）
	TopicId int `json:"id"`
	TopicTitle string `json:"title"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

// 参数绑定
type TopicQuery struct {
	// form是get请求地址栏的参数映射
	// binding: "required" 必传参数
	UserName string `json:"userName" form:"userName"`
	Page int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"pageSize" form:"pageSize"`
}