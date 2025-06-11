package dao

import (
	"context"
	"danmaku/danmaku_reply/model"
	"log"
)

const (
	// 用户房间订阅关系
	createUserRoomTable    = `CREATE TABLE user_room (id SERIAL PRIMARY KEY, user_id INTEGER NOT NULL, room_id INTEGER NOT NULL, platform TEXT NOT NULL, created_at TIMESTAMP DEFAULT now(), UNIQUE (user_id, room_id, platform));`
	insertUserRoom         = `INSERT INTO user_room (user_id, room_id, platform) VALUES ($1, $2, $3);`
	selectUserRoomByUserID = "Select room_id, platform from user_room where user_id = $1;"
	deleteUserRoomByID     = `DELETE FROM user_room where user_id = $1 and room_id = $2 and platform = $3;`
	selectAllUserRoom      = "Select user_id, room_id, platform from user_room;"
)

func (d *Dao) InsertUserRoom(ctx context.Context, userID int64, roomID int64, platform string) (err error) {
	if _, err = d.pgClient.Exec(insertUserRoom, userID, roomID, platform); err != nil {
		log.Println("insert user room err:", err)
		return
	}
	return
}

func (d *Dao) SelectUserRoom(ctx context.Context, userID int64) (rooms []model.Room, err error) {
	rooms = make([]model.Room, 0)
	row, err := d.pgClient.Query(selectUserRoomByUserID, userID)
	if err != nil {
		log.Println("select user room err:", err)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var roomID int64
		var s string
		var platform model.Platform
		if err = row.Scan(&roomID, &s); err != nil {
			log.Println("scan user room err:", err)
			return nil, err
		}
		if s == "douyin" {
			platform = model.Douyin
		} else {
			log.Println("invalid platform:", s)
			continue
		}
		rooms = append(rooms, model.Room{
			Platform: platform,
			RoomID:   roomID,
		})
	}
	return
}

func (d *Dao) DeleteUserRoom(ctx context.Context, userID int64, roomID int64, platform string) (err error) {
	if _, err = d.pgClient.Exec(deleteUserRoomByID, userID, roomID, platform); err != nil {
		log.Println("delete user room err:", err)
		return
	}
	return
}
