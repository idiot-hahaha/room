package service

type DanmakuFetcher interface {
	Start()
	Close() error
}
