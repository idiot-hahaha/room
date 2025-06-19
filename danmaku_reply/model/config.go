package model

import (
	"encoding/json"
	"os"
)

type Config struct {
	//Grpc               *GrpcConfig               `json:"grpc"`
	Http            *HttpConfig               `json:"http"`             // http接口
	PgConf          *PgConfig                 `json:"pg_conf"`          // PostgreSQL配置信息
	MysqlConf       *MysqlConfig              `json:"mysql_conf"`       // Mysql配置信息
	EmbeddingConfig *GrpcClientConfig         `json:"embedding_config"` //
	DigitalConfig   *DigitalModelServerConfig `json:"digital_config"`   // 数字人服务配置
}
type GrpcClientConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
type DigitalModelServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type HttpConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

type PgConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

//func Init() (conf *Config) {
//	f, err := os.Open("reply/cmd/config.json")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	conf = new(Config)
//	if err := json.NewDecoder(f).Decode(&conf); err != nil {
//		panic(err)
//	}
//	return conf
//}

func NewConfig() (conf *Config) {
	f, err := os.Open("danmaku_reply/cmd/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	conf = new(Config)
	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		panic(err)
	}
	return
}
