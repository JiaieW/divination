package routes

import (
	"divination/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.GET("/qigua", controllers.QiGua)
	r.GET("/guaxiang", func(c *gin.Context) {
		type Gua struct {
			BenGua      []string     `json:"bengua"`
			BianGua     []string     `json:"biangua"`
			BianIndexes map[int]bool `json:"bian_indexes"`
		}
		data := Gua{
			BenGua:      []string{"yin", "yin", "yang", "yang", "yang", "yin"},
			BianGua:     []string{"yin", "yin", "yang", "yin", "yin", "yang"},
			BianIndexes: make(map[int]bool),
		}
		for _, index := range []int{4, 5, 6} {
			data.BianIndexes[index-1] = true // 使用index-1因为Go中数组索引是从0开始的
		}
		// 使用HTML模板渲染数据
		fmt.Println(data)
		c.HTML(200, "guaxiang.html", data)
	})
	return r
}
