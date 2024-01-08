package llm

import (
	"divination/llm/qianwen"
)

type LLMMap map[string]LLM

type LLM interface {
	Connect() error
	SendRequest(talk string)
	GetAnswer() string
}

func NewLLM() LLMMap {
	llmMap := LLMMap{}
	//llmMap["XingHuo"] = xinghuo.NewXingHuoModel()
	llmMap["QianWen"] = qianwen.NewQianWenModel()
	return llmMap
}
