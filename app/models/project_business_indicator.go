package models

import (
	"errors"

	v1 "github.com/jinzhu/gorm"
	"github.com/olongfen/demo/app/form"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Error
var (
	ErrCreateProjectBusinessIndicator = errors.New("create ProjectBusinessIndicator failed")
	ErrDeleteProjectBusinessIndicator = errors.New("delete ProjectBusinessIndicator failed")
	ErrGetProjectBusinessIndicator    = errors.New("get ProjectBusinessIndicator failed")
	ErrUpdateProjectBusinessIndicator = errors.New("update ProjectBusinessIndicator failed")
)

// ProjectBusinessIndicator table
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

func init() {
	Tables = append(Tables, &ProjectBusinessIndicator{})
}

// NewProjectBusinessIndicator new
func NewProjectBusinessIndicator() *ProjectBusinessIndicator {
	return new(ProjectBusinessIndicator)
}

// ProjectBusinessIndicatorTableName TableName
func ProjectBusinessIndicatorTableName() string {
	return "project_business_indicators"
}

// CreateProjectBusinessIndicator create one record
func CreateProjectBusinessIndicator(data *ProjectBusinessIndicator, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Create(data).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrCreateProjectBusinessIndicator
		return
	}
	return
}

// DelProjectBusinessIndicator delete record
func DelProjectBusinessIndicator(id interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Delete("id = ?", id).Error; err != nil {
		err = ErrDeleteProjectBusinessIndicator
		return
	}
	return
}

// UpProjectBusinessIndicator update record
func UpProjectBusinessIndicator(id, m interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Where("id = ?", id).Updates(m).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrUpdateProjectBusinessIndicator
		return
	}
	return
}

// GetAllProjectBusinessIndicator get all record
func GetAllProjectBusinessIndicator(dbs ...*gorm.DB) (res []*ProjectBusinessIndicator, err error) {
	if err = GetDB(dbs...).Find(&res).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrGetProjectBusinessIndicator
		return
	}
	return
}

// CountProjectBusinessIndicator get count
func CountProjectBusinessIndicator(dbs ...*gorm.DB) (res int64) {
	GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Count(&res)
	return
}

// DelProjectBusinessIndicatorBatch delete ProjectBusinessIndicator batch
func DelProjectBusinessIndicatorBatch(ids []string, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Delete("id in ?", ids).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrDeleteProjectBusinessIndicator
		return
	}
	return
}

// CreateProjectBusinessIndicatorBatch create ProjectBusinessIndicator batch
func CreateProjectBusinessIndicatorBatch(data []*ProjectBusinessIndicator, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Create(data).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrCreateProjectBusinessIndicator
		return
	}
	return
}

// GetProjectBusinessIndicatorBatch get ProjectBusinessIndicator list some field value or some condition
func GetProjectBusinessIndicatorBatch(q *form.QueryProjectBusinessIndicatorForm, res interface{}, dbs ...*gorm.DB) (err error) {
	var (
		db = GetDB(dbs...).Model(&ProjectBusinessIndicator{})
	)

	if q.Facilities != nil {
		db = db.Where("facilities = ?", q.Facilities)
	}
	if q.QualitativeDescription != nil {
		db = db.Where("qualitative_description = ?", q.QualitativeDescription)
	}
	if q.Issue != nil {
		db = db.Where("issue = ?", q.Issue)
	}
	if q.QuantitativeAssessment != nil {
		db = db.Where("quantitative_assessment = ?", q.QuantitativeAssessment)
	}
	if q.Change != nil {
		db = db.Where("change = ?", q.Change)
	}
	for k, v := range q.AgeMap {
		if k, err = ValidFieldSymbol("age", k); err != nil {
			return
		}
		db = db.Where("age "+k+" ?", v)
	}

	for k, v := range q.OrderMap {
		var (
			desc bool
		)
		if v == "desc" {
			desc = true
		} else {
			desc = false
		}
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v1.ToColumnName(k)}, Desc: desc})
	}
	if q.PageSize != 0 {
		db = db.Limit(q.PageSize)
	}
	if q.PageNum != 0 {
		q.PageNum = (q.PageNum - 1) * q.PageSize
		db = db.Offset(q.PageNum)
	}
	if err = db.Find(res).Error; err != nil {
		return
	}
	return
}

// GetProjectBusinessIndicatorID get one record by ID
func GetProjectBusinessIndicatorID(id interface{}, res interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).First(res, "id = ?", id).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrGetProjectBusinessIndicator
		return
	}
	return
}

// DelProjectBusinessIndicatorID delete record by ID
func DelProjectBusinessIndicatorID(id interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Delete("id = ?", id).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrDeleteProjectBusinessIndicator
		return
	}
	return
}

// GetProjectBusinessIndicatorProjectCode get one record by ProjectCode
func GetProjectBusinessIndicatorProjectCode(projectCode interface{}, res interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).First(res, "project_code = ?", projectCode).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrGetProjectBusinessIndicator
		return
	}
	return
}

// DelProjectBusinessIndicatorProjectCode delete record by ProjectCode
func DelProjectBusinessIndicatorProjectCode(projectCode interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&ProjectBusinessIndicator{}).Delete("project_code = ?", projectCode).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrDeleteProjectBusinessIndicator
		return
	}
	return
}
