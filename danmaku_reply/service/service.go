package service

import (
	"context"
	"danmaku/danmaku_reply/api"
	"danmaku/danmaku_reply/dao"
	"danmaku/danmaku_reply/model"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"sync"
)

func GenFetcherKey(userId string, roomId string, platform string) (s string) {
	return fmt.Sprintf("%s_%s_%s", userId, roomId, platform)
}

type Service struct {
	mu              sync.Mutex
	danmakuFetcher  map[string]DanmakuFetcher
	wsClient        map[string]*websocket.Conn
	config          model.Config
	dao             *dao.Dao
	embeddingClient api.EmbeddingServiceClient
}

func NewService(c *model.Config) (s *Service, err error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	embeddingClient := api.NewEmbeddingServiceClient(conn)
	_, err = embeddingClient.Ping(context.Background(), &api.Empty{})
	if err != nil {
		return nil, err
	}
	d, err := dao.NewDao(c)
	if err != nil {
		return nil, err
	}
	s = &Service{
		danmakuFetcher:  make(map[string]DanmakuFetcher),
		embeddingClient: embeddingClient,
		dao:             d,
		wsClient:        make(map[string]*websocket.Conn),
	}
	return
}

func (s *Service) Close() (err error) {
	for _, f := range s.danmakuFetcher {
		f.Close()
	}
	return
}

func (s *Service) FetchRoomDanmaku(ctx context.Context, userId string, roomId string, platform string, ws *websocket.Conn) (err error) {
	switch platform {
	case "douyin":
		s.mu.Lock()
		defer s.mu.Unlock()
		key := GenFetcherKey(userId, roomId, platform)
		if _, ok := s.danmakuFetcher[key]; ok {
			log.Printf("danmakuFetcher.FetchRoomDanmaku: key %s already exists", key)
			return
		}

		f, err := NewDouyinLive(s, roomId, ws)
		if err != nil {
			log.Printf("DanmakuFetcher.NewDouyinLive err %v\n", err)
			return err
		}
		s.danmakuFetcher[key] = f
		go f.Start()
	}
	return
}

func (s *Service) StopFetchRoomDanmaku(ctx context.Context, userId string, roomId string, platform string) (err error) {
	key := GenFetcherKey(userId, roomId, platform)
	log.Println("StopFetchRoomDanmaku")
	s.mu.Lock()
	defer s.mu.Unlock()
	f, ok := s.danmakuFetcher[key]
	if !ok {
		log.Printf("danmakuFetcher.StopFetchRoomDanmaku: key %s not found", key)
		return
	}
	f.Close()
	delete(s.danmakuFetcher, roomId)
	return
}

func (s *Service) Ping(ctx context.Context, req *api.PingReq) (res *api.PingRes, err error) {
	fmt.Println("DanmakuReplyService.Ping called")
	return &api.PingRes{}, nil
}

func (s *Service) ReplyByGroupID(ctx context.Context, req *api.ReplyByGroupIDReq) (res *api.ReplyByGroupIDRes, err error) {
	getEmbeddingReq := &api.EmbeddingRequest{Text: req.Question}
	embeddingRes, err := s.embeddingClient.GetEmbedding(ctx, getEmbeddingReq)
	if err != nil {
		log.Println(err)
		return
	}
	questionID, question, distance, err := s.dao.SelectMostSimilarQuestionByGroup(req.GroupID, embeddingRes.Embedding)
	if err != nil {
		return
	}
	if distance == 0 { // todo:设定一个阈值
		return
	}
	answers, err := s.dao.SelectAnswersByQuestionID(questionID)
	if err != nil {
		log.Printf("SelectAnswersByQuestionID err:%v\n", err)
		return
	}
	if len(answers) == 0 {
		return
	}
	randomAnswer := answers[rand.Intn(len(answers))]
	fmt.Println(question, randomAnswer)
	return &api.ReplyByGroupIDRes{
		Reply: randomAnswer,
	}, nil
}

func (s *Service) ReplyWithGroupID(ctx context.Context, groupID int64, question string) (matchQuestion, reply string, err error) {
	getEmbeddingReq := &api.EmbeddingRequest{Text: question}
	embeddingRes, err := s.embeddingClient.GetEmbedding(ctx, getEmbeddingReq)
	if err != nil {
		log.Println(err)
		return
	}
	questionID, matchQuestion, distance, err := s.dao.SelectMostSimilarQuestionByGroup(groupID, embeddingRes.Embedding)
	if err != nil {
		return
	}
	if distance == 0 { // todo:设定一个阈值
		return
	}
	answers, err := s.dao.SelectAnswersByQuestionID(questionID)
	if err != nil {
		log.Printf("SelectAnswersByQuestionID err:%v\n", err)
		return
	}
	if len(answers) == 0 {
		return
	}
	reply = answers[rand.Intn(len(answers))]
	return
}

func (s *Service) CreateQAGroup(ctx context.Context, args *model.CreateGroupArgs) (res *model.CreateGroupResp, err error) {
	id, err := s.dao.CreateQAGroup(ctx, args.GroupName)
	if err != nil {
		return
	}
	res = &model.CreateGroupResp{
		GroupID:   id,
		GroupName: args.GroupName,
	}
	return
}

func (s *Service) SendReplyToTTS(ctx context.Context, msg string) (err error) {

	return
}

func (s *Service) AddWsConn(ctx context.Context, userID, platform, roomID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := GenFetcherKey(userID, roomID, platform)
	s.wsClient[key] = conn
	return
}
