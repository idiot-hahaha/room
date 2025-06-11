package service

import (
	"bytes"
	"context"
	"danmaku/danmaku_reply/api"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	c := NewModelClient("gemma3")
	reply, err := c.GenReply("你好")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(reply)
}

type GenerateResponse struct {
	Response string `json:"response"`
}

func TestReply(t *testing.T) {

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// 构建请求内容
	reqBody := GenerateRequest{
		Model: "gemma3",
		Prompt: `你是一个带货主播，请用直播话术来回答用户问题，只需要回答问题。

商品信息：
名称：蜂蜜柚子茶
特点：无添加糖、纯天然果粒、维C丰富、适合夏日解暑

用户问题：和冰红茶比哪个好喝？`,
		Stream: false,
	}

	// JSON 编码请求体
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	// 构造请求
	req, err := http.NewRequest("POST", "http://172.27.96.1:11434/api/generate", bytes.NewReader(bodyBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 解析响应体
	var result GenerateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("原始响应：", string(body)) // 打印原始数据方便调试
		panic(err)
	}

	// 打印返回的内容
	fmt.Println("回答：", result.Response)
}

func TestEmbedding(t *testing.T) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewEmbeddingServiceClient(conn)
	req := &api.EmbeddingRequest{
		Text: "Hello, this is a test sentence.",
	}

	// 调用 RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.GetEmbedding(ctx, req)
	if err != nil {
		log.Fatalf("❌ 调用失败: %v", err)
	}

	// 打印前 5 个 embedding
	fmt.Println("✅ 接收到 embedding（前5维）:")
	for i, v := range res.Embedding {
		if i >= 5 {
			break
		}
		fmt.Printf("  [%d] %.4f\n", i, v)
	}
}

func TestSDP(t *testing.T) {
	fmt.Println(GenSDP())
}

func TestSessionID(t *testing.T) {
	payload := bytes.NewBufferString(fmt.Sprintf(`{"type":"offer","sdp":%s}`, GenSDP()))
	resp, err := http.Post("https://192.168.10.204:/offer", "application/json", payload)
	if err != nil {
		fmt.Println("Error posting data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

func TestQuickSort(t *testing.T) {

}

func Cal(arr []int) (res int) {
	sum, res := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if sum < 0 {
			sum = 0
		}
		res = max(res, sum+arr[i])
		sum += arr[i]
	}
	return
}

var wg sync.WaitGroup

var data int = 0
var flag bool = false

func write() {
	data = 42   // 写数据
	flag = true // 发信号
}

func read() {
	if flag && data == 0 {
		panic(123)
	}
}

func TestTest(t *testing.T) {
	for i := 0; i < 10_000_000; i++ {
		data = 0
		flag = false

		wg.Add(2)

		go func() {
			defer wg.Done()
			write()
		}()

		go func() {
			defer wg.Done()
			read()
		}()

		wg.Wait()
	}
}
