package model

type FetchDanmakuParam struct {
	UserID   string `json:"user_id" binding:"required"`
	RoomID   string `json:"room_id" binding:"required"`
	Platform string `json:"platform" binding:"required"`
}

type SelectArgs struct {
	GroupID  string `json:"group_id" binding:"required"`
	Question string `json:"question"`
}

type SelectResp struct {
	Question []Question `json:"question"`
}

type Question struct {
	QuestionID int64    `json:"question_id"`
	Content    string   `json:"content"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	AnswerID int64  `json:"answer_id"`
	Content  string `json:"content"`
}

type CreateGroupArgs struct {
	GroupName string `json:"group_name" binding:"required"`
}

type CreateGroupResp struct {
	GroupName string `json:"group_name"`
	GroupID   int64  `json:"group_id"`
}
type CreateQAArgs struct {
	GroupID    int64    `json:"group_id" binding:"required"`
	Question   string   `json:"question" binding:"required"`
	QuestionID int64    `json:"question_id"`
	Answers    []string `json:"answers"`
}

type SubscribeRoomArgs struct {
	UserID   string `json:"user_id" binding:"required"` // 后续应该把userID转到token中，不应该在请求参数中使用
	RoomID   string `json:"room_id" binding:"required"`
	Platform string `json:"platform" binding:"required"`
}

type SubscribeRoomResp struct {
}
