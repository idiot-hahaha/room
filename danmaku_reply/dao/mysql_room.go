package dao

import (
	"danmaku/danmaku_reply/model"
	"log"
)

const (
	selectAllRoomsLiveID      = "select room_number from live_room"
	UpdateLiveStatusByRoomID  = "update live_room set live_status = ? where room_number = ?"
	selectDigitalInfoByRoomID = "select digital_human_port, model_status where room_number = ?"
)

func (d *Dao) SelectAllRooms() (liveIDs []string, err error) {
	rows, err := d.mysqlClient.Query(selectAllRoomsLiveID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var liveID string
		if err := rows.Scan(&liveID); err != nil {
			log.Fatal(err)
		}
		liveIDs = append(liveIDs, liveID)
	}
	return liveIDs, nil
}

func (d *Dao) UpdateLiveStatusByRoomID(liveId string, status bool) error {
	statusNum := 0
	if status {
		statusNum = 1
	}
	_, err := d.mysqlClient.Exec(UpdateLiveStatusByRoomID, statusNum, liveId)
	if err != nil {
		return err
	}
	return err
}

func (d *Dao) GetRoomInfo(platform model.Platform, liveID string) (RoomInfo *model.DigitalInfo, err error) {
	row := d.mysqlClient.QueryRow(selectDigitalInfoByRoomID, liveID)
	var port, modelStatus int
	err = row.Scan(&port, &modelStatus)
	if err != nil {
		return nil, err
	}
	return &model.DigitalInfo{
		RoomNumber:       liveID,
		ModelStatus:      modelStatus,
		DigitalHumanPort: port,
	}, nil
}
