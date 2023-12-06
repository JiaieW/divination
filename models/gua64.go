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
	//Yong         string `json:"yong"`
	//Zonglun      string `json:"zonglun"`

	Yaos []*Yao386 `gorm:"-"`
}
