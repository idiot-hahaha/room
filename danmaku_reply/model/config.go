package model

import (
	"encoding/json"
	"os"
)

type Config struct {
	Grpc            *GrpcConfig       `json:"grpc"`
	Http            *HttpConfig       `json:"http"`
	PgConf          *PgConfig         `json:"pg_conf"`
	EmbeddingConfig *GrpcClientConfig `json:"embedding_config"`
}
type GrpcClientConfig struct {
	Host string `json:"host"`
}

type GrpcConfig struct {
	Host string `json:"host"`
}

type HttpConfig struct {
	Host string `json:"host"`
}

type PgConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Init() (conf *Config) {
	f, err := os.Open("reply/cmd/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	conf = new(Config)
	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		panic(err)
	}
	return conf
}

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
