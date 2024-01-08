package xinghuo

import (
	"crypto/hmac"
	"crypto/sha256"
	"divination/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	ws "github.com/gorilla/websocket"
)

type XingHuo struct {
	conn *ws.Conn
}

// 请求结构体
type Header struct {
	AppID string `json:"app_id"`
}

type Chat struct {
	Domain      string  `json:"domain"`
	Temperature float64 `json:"temperature"`
	TopK        int64   `json:"top_k"`
	MaxTokens   int64   `json:"max_tokens"`
	Auditing    string  `json:"auditing"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	MessageData struct {
		Text []Message `json:"text"`
	} `json:"message"`
}

type RequestParams struct {
	Header    Header `json:"header"`
	Parameter struct {
		Chat Chat `json:"chat"`
	} `json:"parameter"`
	Payload Payload `json:"payload"`
}

// 返回结构体
type ResponseHeader struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	SID     string `json:"sid"`
	Status  int    `json:"status"`
}

type TextContent struct {
	Content string `json:"content"`
	Role    string `json:"role"`
	Index   int    `json:"index"`
}

type Choices struct {
	Status int           `json:"status"`
	Seq    int           `json:"seq"`
	Text   []TextContent `json:"text"`
}

type UsageText struct {
	QuestionTokens   int `json:"question_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Usage struct {
	Text *UsageText `json:"text,omitempty"` // 使用指针类型表示该字段可能为空
}

type ResPayload struct {
	Choices Choices `json:"choices"`
	Usage   *Usage  `json:"usage,omitempty"` // 使用指针类型表示该字段可能为空
}

type APIResponse struct {
	Header  ResponseHeader `json:"header"`
	Payload ResPayload     `json:"payload"`
}

var (
	hostUrl   string
	appid     string
	apiSecret string
	apiKey    string
)

func init() {
	conf := config.C
	hostUrl = conf.Xinghuo.HostUrl
	appid = conf.Xinghuo.Appid
	apiSecret = conf.Xinghuo.ApiSecret
	apiKey = conf.Xinghuo.ApiKey

}

func NewXingHuoModel() *XingHuo {
	return &XingHuo{}
}

func (x *XingHuo) Connect() error {
	d := ws.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthUrl(hostUrl, apiKey, apiSecret), nil)
	if err != nil {
		panic(readResp(resp) + err.Error())
		return err
	} else if resp.StatusCode != 101 {
		panic(readResp(resp) + err.Error())
		return err
	}
	x.conn = conn
	return nil
}
func (x *XingHuo) SendRequest(talk string) {
	go func() {
		data := genParams(appid, talk)
		x.conn.WriteJSON(data)
	}()
}

func (x *XingHuo) GetAnswer() string {

	var answer string
	// 获取返回的数据
	for {
		_, msg, err := x.conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}

		var apiResp APIResponse
		err = json.Unmarshal(msg, &apiResp)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return ""
		}
		fmt.Println(string(msg))

		// 解析数据
		if apiResp.Header.Code != 0 {
			fmt.Println("Error code from server:", apiResp.Header.Code)
			return ""
		}

		if len(apiResp.Payload.Choices.Text) > 0 {
			content := apiResp.Payload.Choices.Text[0].Content
			status := apiResp.Payload.Choices.Status

			if status != 2 {
				answer += content
			} else {
				fmt.Println("收到最终结果")
				answer += content
				if apiResp.Payload.Usage != nil && apiResp.Payload.Usage.Text != nil {
					totalTokens := apiResp.Payload.Usage.Text.TotalTokens
					fmt.Println("total_tokens:", totalTokens)
				}
				x.conn.Close()
				break
			}
		}
	}
	return answer
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

// 生成参数
func genParams(appid, question string) RequestParams { // 根据实际情况修改返回的数据结构和字段名
	messages := []Message{
		{Role: "user", Content: question},
	}

	return RequestParams{
		Header: Header{
			AppID: appid,
		},
		Parameter: struct {
			Chat Chat `json:"chat"`
		}{
			Chat: Chat{
				Domain:      "general",
				Temperature: 0.8,
				TopK:        6,
				MaxTokens:   2048,
				Auditing:    "default",
			},
		},
		Payload: Payload{
			MessageData: struct {
				Text []Message `json:"text"`
			}{
				Text: messages,
			},
		},
	}
}
