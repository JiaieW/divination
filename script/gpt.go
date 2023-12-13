package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 定义一个结构体来匹配 API 的请求格式
type ChatGPTRequest struct {
	Prompt string `json:"prompt"`
}

// 定义一个结构体来接收 API 的响应数据
type ChatGPTResponse struct {
	Responses []string `json:"responses"`
}

func main() {
	// 创建一个新的请求体实例
	requestBody := ChatGPTRequest{
		Prompt: "Hello, world!",
	}

	// 将请求体实例转换为 JSON 格式
	jsonReq, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/gpt-3.5-turbo/completions", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加必要的头部信息
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将响应数据解析到结构体中
	var chatGPTResp ChatGPTResponse
	if err := json.Unmarshal(body, &chatGPTResp); err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应
	fmt.Println("Response from ChatGPT:", chatGPTResp.Responses)
}
