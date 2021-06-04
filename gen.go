package main

//go:generate gengo  -i ./gen.go  -m github.com/olongfen/demo -g gen-model
type ProjectBusinessIndicator struct {
	ID                     uint   `json:"id" gorm:"primaryKey"`
	ProjectCode            string `json:"projectCode" gorm:"type:varchar(36);uniqueIndex"`
	Facilities             string `json:"facilities" gorm:"type:varchar(128)" comment:"主要设备"`
	QualitativeDescription string `json:"qualitativeDescription" gorm:"type:varchar(256)" comment:"定性描述"`
	Issue                  string `json:"issue" gorm:"40" comment:"期号"`
	QuantitativeAssessment string `json:"quantitativeAssessment" gorm:"type:varchar(256)" comment:"定量评估"`
	Change                 string `json:"change" gorm:"type:varchar(256)" comment:"与上期报告相比的变化"`
	Age                    int    `json:"age"`
}

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Uid  string `json:"uid" gorm:"type:varchar(36)"`
	Name string `json:"name"`
}
