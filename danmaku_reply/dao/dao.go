package dao

import (
	"danmaku/danmaku_reply/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Dao struct {
	pgClient    *sql.DB
	mysqlClient *sql.DB
}

func NewDao(conf *model.Config) (dao *Dao, err error) {
	dao = &Dao{}
	dao.pgClient, err = newPgClient(conf.PgConf)
	if err != nil {
		return nil, err
	}
	dao.mysqlClient, err = newMysqlClient(conf.MysqlConf)
	if err != nil {
		return nil, err
	}
	return dao, nil
}

func newPgClient(conf *model.PgConfig) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return
}

func newMysqlClient(config *model.MysqlConfig) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return
}

func (d *Dao) Close() {
	d.pgClient.Close()
}
