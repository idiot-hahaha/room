package test_test

import (
	"bytes"
	"danmaku/danmaku_reply/service"
	"encoding/json"
	"fmt"
	"github.com/pion/webrtc/v3"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"danmaku/danmaku_reply/generated/new_douyin"
	"google.golang.org/protobuf/proto"
)

func TestNewDouyinLive(t *testing.T) {
	d, _ := service.NewDouyinLive(nil, "646454278948", nil)
	d.Subscribe(func(eventData *new_douyin.Webcast_Im_Message) (err error) {
		// t.Logf("msg received ,type:%s", eventData.Method)
		switch eventData.Method {
		case service.WebcastChatMessage:
			msg := &new_douyin.Webcast_Im_ChatMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			t.Logf("聊天消息:user=%d %s %s", msg.User.Id, msg.User.Nickname, msg.Content)
		case service.WebcastGiftMessage:
			msg := &new_douyin.Webcast_Im_GiftMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			t.Logf("礼物消息:user=%d %s %s", msg.User.Id, msg.User.Nickname, msg.Gift.Name)
		case service.WebcastLikeMessage:
			msg := &new_douyin.Webcast_Im_LikeMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			t.Logf("点赞消息:user=%d %s", msg.User.Id, msg.User.Nickname)
		case service.WebcastMemberMessage:
			msg := &new_douyin.Webcast_Im_MemberMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			t.Logf("成员消息:user=%d %s", msg.User.Id, msg.User.Nickname)
		case service.WebcastSocialMessage:
			msg := &new_douyin.Webcast_Im_SocialMessage{}
			proto.Unmarshal(eventData.Payload, msg)
			t.Logf("社交消息:user=%d %s", msg.User.Id, msg.User.Nickname)
		default:
			t.Logf("其他消息:type:%s", eventData.Method)
		}

		// if eventData.Method == DanmakuFetcher.WebcastChatMessage {
		// 	msg := &douyin.ChatMessage{}
		// 	proto.Unmarshal(eventData.Payload, msg)
		// 	marshal, _ := protojson.Marshal(msg)
		// 	log.Println("聊天msg", msg.User.Id, msg.User.NickName, msg.Content, string(marshal))
		// }
		return
	})

	d.Start()

}

type ReqStruct struct {
	SDP  string `json:"sdp"`
	Type string `json:"type"`
}
type RespStruct struct {
	SDP       string `json:"sdp"`
	Type      string `json:"type"`
	SessionID int    `json:"sessionid"`
}

func TestSDP(t *testing.T) {
	m := webrtc.MediaEngine{}
	if err := m.RegisterDefaultCodecs(); err != nil {
		log.Fatal(err)
	}

	// 创建 API 对象
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&m))

	// 配置 STUN 服务器
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlan,
	}

	// 创建 PeerConnection
	pc, err := api.NewPeerConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	// 添加一个音频轨道
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: "audio/opus",
	}, "audio", "pion")
	if err != nil {
		log.Fatal(err)
	}
	_, err = pc.AddTrack(audioTrack)
	if err != nil {
		log.Fatal(err)
	}

	// 添加一个视频轨道
	videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: "video/vp8",
	}, "video", "pion")
	if err != nil {
		log.Fatal(err)
	}
	_, err = pc.AddTrack(videoTrack)
	if err != nil {
		log.Fatal(err)
	}

	// 添加数据通道（可选）
	_, err = pc.CreateDataChannel("data", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Offer
	offer, err := pc.CreateOffer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置本地描述
	if err = pc.SetLocalDescription(offer); err != nil {
		log.Fatal(err)
	}

	// 打印 SDP
	fmt.Println("SDP Offer:")
	fmt.Println(offer.SDP)

	// 构造请求体
	requestBody := ReqStruct{
		SDP:  offer.SDP,
		Type: "offer",
	}
	marshaledRequest, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	// 发送 HTTP POST 请求
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", "http://localhost:8010/offer", bytes.NewBuffer(marshaledRequest))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("HTTP Response: %+v\n", resp)
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respStruct := &RespStruct{}
	err = json.Unmarshal(responseBody, respStruct)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("HTTP Response: %+v\n", respStruct)
}
