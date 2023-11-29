package models

type Gua64 struct {
	Id           int    `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Alias        string `json:"alias"`
	Location     string `json:"location"`
	Form         string `json:"form"`
	Code         string `json:"code"`
	Guaci        string `json:"guaci"`
	GuaciFanyi   string `json:"guaci_fanyi"`
	GuaciExplain string `json:"guaci_explain"`
	Yao_1        string `json:"yao_1"`
	Yao_2        string `json:"yao_2"`
	Yao_3        string `json:"yao_3"`
	Yao_4        string `json:"yao_4"`
	Yao_5        string `json:"yao_5"`
	Yao_6        string `json:"yao_6"`
	Yong         string `json:"yong"`
	Zonglun      string `json:"zonglun"`

	Yaos []Yao386 `gorm:"-"`
}
