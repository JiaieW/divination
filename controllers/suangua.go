package controllers

import (
	"bytes"
	"divination/controllers/util"
	"divination/database"
	"divination/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const DAYAN = 49

type Orcale struct {
	BenGuaCode  []int  `json:"bengua_code"`
	BianGuaCode []int  `json:"biangua_code"`
	BianYaos    []int  `json:"bian_yaos"`
	BianNum     int    `json:"bian_num"`
	Master      string `json:"master"`
	Slave       string `json:"slave"`
	Bengua      models.Gua64
	Biangua     models.Gua64
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func QiGua(c *gin.Context) {
	db := database.DBConn
	benGuaCode, bianGuaCode := []int{}, []int{}
	bianYaos := []int{}
	yao, bian := 0, false
	for i := 1; i < 7; i++ {
		//经过三变得到代表爻的数字
		yaoNum := util.Change(DAYAN, 1)
		yaoNum = yaoNum / 4
		//0表示阴爻，1表示阳爻
		switch yaoNum {
		case 6: //老阴
			yao = 0
			bian = true
		case 7: //少阳
			yao = 1
		case 8: //少阴
			yao = 0
		case 9: //老阳
			yao = 1
			bian = true
		}
		//用6位二进制表示64卦
		benGuaCode = append(benGuaCode, yao)
		if bian {
			bianYaos = append(bianYaos, i) //老阴为变卦
			yao = yao ^ 1
		}
		bianGuaCode = append(bianGuaCode, yao)
	}
	benGuaCodeStr, bianGuaCodeStr := util.ArrayToString(benGuaCode), util.ArrayToString(bianGuaCode)
	bengua, biangua := &models.Gua64{}, &models.Gua64{}
	benguaDetail, bianguaDetail := &[]models.Yao386{}, &[]models.Yao386{}
	db.Where(models.Gua64{Code: benGuaCodeStr}).First(bengua)
	db.Where(models.Gua64{Code: bianGuaCodeStr}).First(biangua)

	db.Where(models.Yao386{GuaId: bengua.Id}).Find(benguaDetail)
	db.Where(models.Yao386{GuaId: biangua.Id}).Find(bianguaDetail)

	for _, yao := range *benguaDetail {
		bengua.Yaos = append(bengua.Yaos, yao)
	}
	for _, yao := range *bianguaDetail {
		biangua.Yaos = append(biangua.Yaos, yao)
	}
	o := Orcale{}
	o.BenGuaCode = benGuaCode
	o.BianGuaCode = bianGuaCode
	o.BianYaos = bianYaos
	o.BianNum = len(bianYaos)
	o.Bengua = *bengua
	o.Biangua = *biangua
	//根据变爻的数目，确定使用哪个爻辞判定凶吉
	switch o.BianNum {
	case 0: //本卦卦辞
		o.Master = bengua.Guaci
	case 1: //本卦变爻
		o.Master = bengua.Yaos[bianYaos[0]-1].YaoTrans
		fmt.Println("本卦变爻：", bengua.Yaos[bianYaos[0]-1])
	case 2: //如果卦里有两个爻发生变动，那就用本卦里这两个变爻的占辞来判断吉凶，并以位置靠上的那一个变爻的占辞为主
		o.Master = bengua.Yaos[bianYaos[1]-1].YaoTrans
		o.Slave = bengua.Yaos[bianYaos[0]-1].YaoTrans
		fmt.Println("本卦变爻（主）：", bengua.Yaos[bianYaos[1]-1])
		fmt.Println("本卦变爻：", bengua.Yaos[bianYaos[0]-1])
	case 3: //本卦变卦卦辞 ，以本卦卦辞为主
		o.Master = bengua.Guaci
		o.Slave = biangua.Guaci
		fmt.Println("本卦卦辞 (主)：", bengua.Guaci)
		fmt.Println("变卦卦辞：", biangua.Guaci)
	case 4: //变卦两个不变爻
		str := bytes.Buffer{}
		for i := 1; i < 7; i++ {
			res := util.IssetInSlice(i, bianYaos)
			if res != 0 {
				str.WriteString(biangua.Yaos[res-1].YaoTrans)
				fmt.Println("变卦不变爻：", biangua.Yaos[res-1])
			}
		}
		o.Master = str.String()
	case 5: //变卦的一个不变爻
		var n int
		for _, i := range bianYaos {
			n += i
		}
		o.Master = biangua.Yaos[21-n-1].YaoTrans
		fmt.Println("变卦不变爻：", biangua.Yaos[21-n-1])
	case 6: // 变卦卦辞
		o.Master = biangua.Guaci
		fmt.Println("变卦卦辞：", biangua.Guaci)
	}
	oJson, _ := json.Marshal(o)
	c.JSON(http.StatusOK, gin.H{"data": string(oJson)})
}
