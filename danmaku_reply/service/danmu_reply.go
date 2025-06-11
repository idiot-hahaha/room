package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var url = "http://172.27.96.1:11434/api/generate"

type ModelClient struct {
	httpClient  *http.Client
	model       string
	chatContext map[string]string
}

func NewModelClient(model string) *ModelClient {
	if model == "" {
		model = "gamma3"
	}
	return &ModelClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:       100,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: false,
			},
		},
		model:       model,
		chatContext: map[string]string{},
	}
}

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func (c *ModelClient) GetReplyHttp(question string) (data []byte, err error) {
	chat := make([]byte, 0)
	chat = append(chat, []byte("你是一个带货主播，商品特征如下：\n")...)
	for k, v := range c.chatContext {
		chat = append(chat, []byte(k)...)
		chat = append(chat, []byte(":")...)
		chat = append(chat, []byte(v)...)
		chat = append(chat, []byte("\n")...)
	}
	chat = append(chat, []byte("请回答用户的问题：")...)
	chat = append(chat, []byte(question)...)
	reqBody := GenerateRequest{
		Model:  c.model,
		Prompt: string(chat),
		Stream: false,
	}
	fmt.Println(reqBody.Prompt)
	data, err = json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return responseData, nil
}

func (c *ModelClient) GenReply(question string) (reply string, err error) {
	if question == "" {
		return "", errors.New("question is empty")
	}
	res, err := c.GetReplyHttp(question)
	if err != nil {
		return "", err
	}
	switch c.model {
	case "gemma3":
		return string(res), nil
	default:
	}
	return question, nil
}
