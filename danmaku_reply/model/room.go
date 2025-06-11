package model

type Platform string

const (
	Douyin Platform = "douyin"
)

type Room struct {
	Platform Platform `json:"platform"`
	RoomID   int64    `json:"room_id"`
	Status   bool     `json:"status"`
}

type User struct {
	UserID int64 `json:"user_id"`
}

type SubRoomParam struct {
	UserID   int64  `json:"user_id"`
	RoomID   int64  `json:"room_id"`
	Platform string `json:"platform"`
}

type DeleteRoomParam struct {
	UserID   int64  `json:"user_id"`
	RoomID   int64  `json:"room_id"`
	Platform string `json:"platform"`
}
