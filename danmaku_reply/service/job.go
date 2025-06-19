package service

import (
	"log"
	"strconv"
)

// StartJob 单个服务器时使用
func (s *Service) StartJob() (err error) {
	_, err = s.cron.AddFunc("*/10 * * * * *", s.UpdateRoomStates)
	if err != nil {
		return
	}
	s.cron.Start()
	return
}

func (s *Service) UpdateRoomStates() {
	rooms, err := s.dao.SelectAllRooms()
	if err != nil {
		log.Printf("dao.SelectAllRooms err: %v\n", err)
	}
	for _, room := range rooms {
		// 目前只有抖音
		if _, err := strconv.Atoi(room); err != nil {
			log.Printf("room id is not number: %s\n", room)
			continue
		}
		living, err := s.GetDouyinLiveStatus(room)
		if err != nil {
			log.Printf("s.GetDouyinLiveStatus err: %v\n", err)
			return
		}
		err = s.dao.UpdateLiveStatusByRoomID(room, living)
		if err != nil {
			log.Printf("UpdateLiveStatusByRoomID err: %v\n", room)
			return
		}
	}
}
