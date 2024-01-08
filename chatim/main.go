package main

import (
	"divination/llm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态文件路由，用于提供 HTML 页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})

	// 加载 HTML 模板
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	// 处理前端发送的消息
	r.POST("/message", func(c *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		llmM := llm.NewLLM()
		anwser := ""
		for name, m := range llmM {
			m.Connect()
			m.SendRequest(req.Message)
			anwser += name + " : " + m.GetAnswer()
			// if anwser != "" {
			// 	break
			// }
		}
		if anwser == "" {
			anwser = "请稍后"
		}
		// 这里我们模拟 GPT 的回复
		reply := "模拟回复: " + anwser

		c.JSON(http.StatusOK, gin.H{"reply": reply})
	})

	r.Run(":8765") // 在 localhost:8080 上监听
}
