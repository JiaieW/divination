package qianwen

import (
	"bytes"
	"divination/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// qianwen 结构体定义
type QianWen struct {
	apiURL      string
	apiKey      string
	requestBody GenerationRequest
	response    []byte
}

// GenerationRequest 结构体
type GenerationRequest struct {
	Model      string   `json:"model"`
	Input      Input    `json:"input"`
	Parameters struct{} `json:"parameters"`
}

// APIResponse 用于解析 API 响应
type APIResponse struct {
	Output struct {
		Text string `json:"text"`
	} `json:"output"`
}

// Input 结构体
type Input struct {
	Messages []Message `json:"messages"`
}

// Message 结构体
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

var ApiKey string
var apiURL string

func init() {
	conf := config.C
	ApiKey = conf.Qianwen.ApiKey
	apiURL = conf.Qianwen.ApiUrl
}

// Connect 方法实现
func (q *QianWen) Connect() error {
	// 在这里可以实现任何初始化或连接逻辑
	// 例如，可以在这里设置 API URL 和 API Key
	q.apiURL = apiURL
	q.apiKey = ApiKey // 替换为您的 API Key
	return nil
}

// SendRequest 方法实现
func (q *QianWen) SendRequest(talk string) {
	// 构建请求体
	q.requestBody = GenerationRequest{
		Model: "qwen-max",
		Input: Input{
			Messages: []Message{
				{Role: "system", Content: "你的身份是一个精通周易命理，擅长通过周易卦辞爻辞占卜推理吉凶，并给出人生经验的大师，其他无关的问题简短的回复不知道即可"},
				{Role: "user", Content: talk},
			},
		},
	}

	// 将请求体编码为 JSON
	jsonData, err := json.Marshal(q.requestBody)
	if err != nil {
		fmt.Println("Error encoding request data:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", q.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+q.apiKey)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("X-DashScope-SSE", "enable")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	q.response = body
}

// GetAnswer 方法实现
func (q *QianWen) GetAnswer() string {
	var resp APIResponse
	fmt.Println(string(q.response))
	err := json.Unmarshal(q.response, &resp)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return ""
	}
	return resp.Output.Text
}

func NewQianWenModel() *QianWen {
	return &QianWen{}
}
