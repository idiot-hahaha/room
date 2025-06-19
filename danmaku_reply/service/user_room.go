package service

import (
	"context"
	"danmaku/danmaku_reply/model"
	"log"
	"strconv"
	"sync"
)

func (s *Service) AddQAByGroupID(ctx context.Context, msg string) (err error) {

	return
}

func (s *Service) GetQAByGroupID(ctx context.Context, msg string) (err error) {
	return
}

func (s *Service) GetRoomsByID(ctx context.Context, msg string) (err error) {
	return
}

func (s *Service) InsertUserRoom(ctx context.Context, msg string) (err error) {
	return
}

func (s *Service) SubRoom(ctx context.Context, param *model.SubRoomParam) (err error) {
	err = s.dao.InsertUserRoom(ctx, param.UserID, param.RoomID, param.Platform)
	if err != nil {
		return
	}
	return
}

func (s *Service) DeleteRoom(ctx context.Context, param *model.DeleteRoomParam) (err error) {
	err = s.dao.DeleteUserRoom(ctx, param.UserID, param.RoomID, param.Platform)
	if err != nil {
		return
	}
	return
}

func (s *Service) RoomsByUserID(ctx context.Context, userID int64) (rooms []model.Room, err error) {
	rooms, err = s.dao.SelectUserRoom(ctx, userID)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	for i, _ := range rooms {
		wg.Add(1)
		go func(idx int) {
			switch rooms[idx].Platform {
			
			case model.Douyin:
				liveID := strconv.Itoa(int(rooms[idx].RoomID))
				rooms[idx].Status, _ = s.GetDouyinLiveStatus(liveID)

			default:
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}

func (s *Service) GetDouyinLiveStatus(liveId string) (ok bool, err error) {
	d, err := NewDouyinLive(s, liveId)
	if err != nil {
		return false, err
	}
	result, err := d.getUrl(d.liveurl + d.liveid)
	if err != nil {
		return false, err
	}
	str := extractMatch(isLiveRegexp, result, 2)
	log.Printf("直播状态: %v\n", str)
	return str == "2", nil
}
