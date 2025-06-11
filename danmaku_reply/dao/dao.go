package dao

import (
	"danmaku/danmaku_reply/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Dao struct {
	pgClient *sql.DB
}

func NewDao(conf *model.Config) (dao *Dao, err error) {
	dao = &Dao{}
	dao.pgClient, err = newPgClient(conf.PgConf)
	if err != nil {
		return nil, err
	}
	return dao, nil
}

func newPgClient(conf *model.PgConfig) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.User, conf.Password, conf.Database))
}

func (dao *Dao) Close() {
	dao.pgClient.Close()
}
