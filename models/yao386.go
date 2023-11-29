package models

type Yao386 struct {
	Id         int    `json:"id" gorm:"primary_key"`
	GuaId      int    `json:"gua_id"`
	YaoPos     int    `json:"yao_pos"`
	YaoTrans   string `json:"yao_trans"`
	YaoExplain string `json:"yao_explain"`
	YaoStruct  string `json:"yao_struct"`
}
