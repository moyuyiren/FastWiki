package Model

import (
	"github.com/gin-gonic/gin"
)

// ========================================后端数据
type DocMassage struct {
	Uniposcod     string `json:"uniposcod" xorm:"uniposcod"`
	TitleName     string `json:"titleName" xorm:"titlename"`
	TitleImage    string `json:"titleImage" xorm:"titleimage"`
	TitleDocument string `json:"titleDocument" xorm:"titledocument"`
	Matter        string `json:"matter" xorm:"document"`
}

// ========================================前端数据
type Massage struct {
	Uniposcod     string              `json:"uniposcod"`
	TitleName     string              `json:"titleName"`
	TitleImage    string              `json:"titleImage"`
	TitleDocument string              `json:"titleDocument"`
	MassageInfo   []map[string]string `json:"massageInfo"`
}

type Table struct {
	TableNameFirst      string `mapstructure:"tableNameFirst"`
	TableImagesFirst    string `mapstructure:"tableImagesFirst"`
	TableDocumentFirst  string `mapstructure:"tableDocumentFirst"`
	TableNameSecond     string `mapstructure:"tableNameSecond"`
	TableImagesSecond   string `mapstructure:"tableImagesSecond"`
	TableDocumentSecond string `mapstructure:"tableDocumentSecond"`
	TableNameThird      string `mapstructure:"tableNameThird"`
	TableImagesThird    string `mapstructure:"tableImagesThird"`
	TableDocumentThird  string `mapstructure:"tableDocumentThird"`
	TableNameFourth     string `mapstructure:"tableNameFourth"`
	TableImagesFourth   string `mapstructure:"tableImagesFourth"`
	TableDocumentFourth string `mapstructure:"tableDocumentFourth"`
}

func GetDocMassage(c gin.Context) (err error, Docc DocMassage) {

}
