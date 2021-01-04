package model_demo

import (
	"errors"
	"github.com/olongfen/demo/app/model/common"
	"gorm.io/gorm"
)

// Error
var (
	ErrCreateDemo = errors.New("create Demo failed")
	ErrDeleteDemo = errors.New("delete Demo failed")
	ErrGetDemo    = errors.New("get Demo failed")
	ErrUpdateDemo = errors.New("update Demo failed")
)

// Demo
type Demo struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}

func init() {
	model_common.Tables = append(model_common.Tables, &Demo{})
}

// NewDemo new
func NewDemo() *Demo {
	return new(Demo)
}

// Add add one record
func (t *Demo) Add(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Create(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateDemo
		return
	}
	return
}

// Delete delete record
func (t *Demo) Delete(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		err = ErrDeleteDemo
		return
	}
	return
}

// Updates update record
func (t *Demo) Updates(m map[string]interface{}, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(t).Updates(m).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrUpdateDemo
		return
	}
	return
}

// GetDemoAll get all record
func GetDemoAll(dbs ...*gorm.DB) (ret []*Demo, err error) {
	if err = model_common.GetDB(dbs...).Find(&ret).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetDemo
		return
	}
	return
}

// GetDemoCount get count
func GetDemoCount(dbs ...*gorm.DB) (ret int64) {
	model_common.GetDB(dbs...).Model(&Demo{}).Count(&ret)
	return
}

// DeleteDemoBatch delete Demo batch
func DeleteDemoBatch(ids []string, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&Demo{}).Delete("id in ?", ids).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteDemo
		return
	}
	return
}

// AddDemoBatch add Demo batch
func AddDemoBatch(datas []*Demo, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&Demo{}).Create(datas).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateDemo
		return
	}
	return
}

//  QueryDemoForm query form
type QueryDemoForm struct {
	CreatedAt *model_common.FieldData `json:"createdAt" form:"createdAt"` // if required, add binding:"required" to tag by self
	UpdatedAt *model_common.FieldData `json:"updatedAt" form:"updatedAt"` // if required, add binding:"required" to tag by self
	Name      *string                 `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Age       *model_common.FieldData `json:"age" form:"age"`             // if required, add binding:"required" to tag by self
	Class     *string                 `json:"class" form:"class"`         // if required, add binding:"required" to tag by self

	Order    []string `json:"order" form:"order"`
	PageNum  int      `json:"pageNum" form:"pageNum" binding:"required"`
	PageSize int      `json:"pageSize" form:"pageSize" binding:"required" `
}

// GetDemoList get Demo list some field value or some condition
func GetDemoList(q *QueryDemoForm, dbs ...*gorm.DB) (ret []*Demo, err error) {
	var (
		db = model_common.GetDB(dbs...)
	)
	// order
	if len(q.Order) > 0 {
		for _, v := range q.Order {
			db = db.Order(v)
		}
	}
	// pageSize
	if q.PageSize != 0 {
		db = db.Limit(q.PageSize)
	}
	// pageNum
	if q.PageNum != 0 {
		q.PageNum = (q.PageNum - 1) * q.PageSize
		db = db.Offset(q.PageNum)
	}

	if q.CreatedAt != nil {
		db = db.Where("created_at"+q.CreatedAt.Symbol+"?", q.CreatedAt.Value)
	}

	if q.UpdatedAt != nil {
		db = db.Where("updated_at"+q.UpdatedAt.Symbol+"?", q.UpdatedAt.Value)
	}
	if q.Name != nil {
		db = db.Where("name = ?", *q.Name)
	}

	if q.Age != nil {
		db = db.Where("age"+q.Age.Symbol+"?", q.Age.Value)
	}
	if q.Class != nil {
		db = db.Where("class = ?", *q.Class)
	}

	if err = db.Find(&ret).Error; err != nil {
		return
	}
	return
}

// QueryByID query cond by ID
func (t *Demo) SetQueryByID(id uint) *Demo {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *Demo) GetByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).First(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetDemo
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *Demo) DeleteByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteDemo
		return
	}
	return
}
