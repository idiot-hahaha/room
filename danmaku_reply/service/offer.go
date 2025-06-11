package service

import (
	"log"

	"github.com/pion/webrtc/v3"
)

func GenSDP() string {
	mediaEngine := &webrtc.MediaEngine{}
	if err := mediaEngine.RegisterDefaultCodecs(); err != nil {
		log.Fatal(err)
	}

	api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

	// 创建 PeerConnection
	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个数据通道（非必须）
	_, err = peerConnection.CreateDataChannel("temp", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置本地描述（必须设置，ICE candidate 才能启动）
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		log.Fatal(err)
	}

	// 等待 Gathering 完成
	<-webrtc.GatheringCompletePromise(peerConnection)

	// 输出 SDP 字符串
	return peerConnection.LocalDescription().SDP
}
