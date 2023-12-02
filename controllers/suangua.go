package controllers

import (
	"bytes"
	"divination/controllers/util"
	"divination/database"
	"divination/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const DAYAN = 49

type Orcale struct {
	Master        string `json:"master"`
	MasterExplain string `json:"master_explain"`
	Slave         string `json:"slave"`
	SlaveExplain  string `json:"slave_explain"`
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

type GuaXiang struct {
	BenGua      []string     `json:"bengua"`
	BianGua     []string     `json:"biangua"`
	BianIndexes map[int]bool `json:"bian_indexes"`
	BenGuaInfo  models.Gua64 `json:"bengua_info"`
	BianGuaInfo models.Gua64 `json:"biangua_info"`
	Orcale      Orcale       `json:"orcale"`
}

func QiGua(c *gin.Context) {

	data := GuaXiang{
		BenGua:      make([]string, 7),
		BianGua:     make([]string, 7),
		BianIndexes: make(map[int]bool),
	}

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
	//
	for i := 0; i < 6; i++ {
		if benGuaCode[5-i] == 0 {
			data.BenGua[6-i] = "yin"
		} else {
			data.BenGua[6-i] = "yang"
		}
	}
	for i := 0; i < 6; i++ {
		if bianGuaCode[5-i] == 0 {
			data.BianGua[6-i] = "yin"
		} else {
			data.BianGua[6-i] = "yang"
		}
	}
	for _, index := range bianYaos {
		data.BianIndexes[index] = true
	}

	//c.HTML(200, "guaxiang.html", data)
	benGuaCodeStr, bianGuaCodeStr := util.ArrayToString(benGuaCode), util.ArrayToString(bianGuaCode)

	bengua, biangua := models.Gua64{}, models.Gua64{}
	benguaDetail, bianguaDetail := []models.Yao386{}, []models.Yao386{}
	db.Where(models.Gua64{Code: benGuaCodeStr}).First(&bengua)
	db.Where(models.Gua64{Code: bianGuaCodeStr}).First(&biangua)
	data.BenGuaInfo, data.BianGuaInfo = bengua, biangua
	db.Where(models.Yao386{GuaId: bengua.Id}).Find(&benguaDetail)
	db.Where(models.Yao386{GuaId: biangua.Id}).Find(&bianguaDetail)
	data.BenGuaInfo.Yaos = benguaDetail
	data.BianGuaInfo.Yaos = bianguaDetail

	// //根据变爻的数目，确定使用哪个爻辞判定凶吉
	switch len(bianYaos) {
	case 0: //本卦卦辞
		data.Orcale.Master = data.BenGuaInfo.Guaci
	case 1: //本卦变爻
		data.Orcale.Master = data.BenGuaInfo.Yaos[bianYaos[0]-1].YaoTrans

	case 2: //如果卦里有两个爻发生变动，那就用本卦里这两个变爻的占辞来判断吉凶，并以位置靠上的那一个变爻的占辞为主
		data.Orcale.Master = data.BenGuaInfo.Yaos[bianYaos[1]-1].YaoTrans
		data.Orcale.Slave = data.BenGuaInfo.Yaos[bianYaos[0]-1].YaoTrans
	case 3: //本卦变卦卦辞 ，以本卦卦辞为主
		data.Orcale.Master = data.BenGuaInfo.Guaci
		data.Orcale.Slave = data.BianGuaInfo.Guaci
	case 4: //变卦两个不变爻
		str := bytes.Buffer{}

		for i := 1; i < 7; i++ {
			res := util.IssetInSlice(i, bianYaos)
			fmt.Println("变卦不变爻---------------", res)
			if res != 0 {
				str.WriteString(data.BianGuaInfo.Yaos[res-1].YaoTrans) // todo
			}
		}
		data.Orcale.Master = str.String()
	case 5: //变卦的一个不变爻
		var n int
		for _, i := range bianYaos {
			n += i
		}
		data.Orcale.Master = data.BianGuaInfo.Yaos[21-n-1].YaoTrans
	case 6: // 变卦卦辞
		data.Orcale.Master = data.BianGuaInfo.Guaci
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
	//c.HTML(200, "guaxiang.html", data)
}
