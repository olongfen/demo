package model_region

import (
	"errors"
	"github.com/olongfen/demo/app/model/common"
	"gorm.io/gorm"
)

// Error
var (
	ErrCreateRegion = errors.New("create Region failed")
	ErrDeleteRegion = errors.New("delete Region failed")
	ErrGetRegion    = errors.New("get Region failed")
	ErrUpdateRegion = errors.New("update Region failed")
)

// Region
type Region struct {
	ID        int64  `json:"id" gorm:"primarykey"`
	Name      string `json:"name" gorm:"type:varchar(64)"`            //
	Cname     string `json:"cname" gorm:"type:varchar(64)"`           //
	LowerName string `json:"lowerName" gorm:"type:varchar(64)"`       //
	Longitude string `json:"longitude" gorm:"type:varchar(12);index"` // 经度
	Latitude  string `json:"latitude" gorm:"type:varchar(12);index"`  // 纬度
}

func init() {
	model_common.Tables = append(model_common.Tables, &Region{})
}

// NewRegion new
func NewRegion() *Region {
	return new(Region)
}

// Add add one record
func (t *Region) Add(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Create(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateRegion
		return
	}
	return
}

// Delete delete record
func (t *Region) Delete(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		err = ErrDeleteRegion
		return
	}
	return
}

// Updates update record
func (t *Region) Updates(m map[string]interface{}, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(t).Updates(m).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrUpdateRegion
		return
	}
	return
}

// GetRegionAll get all record
func GetRegionAll(dbs ...*gorm.DB) (ret []*Region, err error) {
	if err = model_common.GetDB(dbs...).Find(&ret).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetRegion
		return
	}
	return
}

// GetRegionCount get count
func GetRegionCount(dbs ...*gorm.DB) (ret int64) {
	model_common.GetDB(dbs...).Model(&Region{}).Count(&ret)
	return
}

// DeleteRegionBatch delete Region batch
func DeleteRegionBatch(ids []string, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&Region{}).Delete("id in ?", ids).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteRegion
		return
	}
	return
}

// AddRegionBatch add Region batch
func AddRegionBatch(datas []*Region, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&Region{}).Create(datas).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateRegion
		return
	}
	return
}

//  QueryRegionForm query form
type QueryRegionForm struct {
	Name      *string `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Cname     *string `json:"cname" form:"cname"`         // if required, add binding:"required" to tag by self
	LowerName *string `json:"lowerName" form:"lowerName"` // if required, add binding:"required" to tag by self
	Longitude *string `json:"longitude" form:"longitude"` // if required, add binding:"required" to tag by self
	Latitude  *string `json:"latitude" form:"latitude"`   // if required, add binding:"required" to tag by self

	Order    []string `json:"order" form:"order"`
	PageNum  int      `json:"pageNum" form:"pageNum" binding:"required"`
	PageSize int      `json:"pageSize" form:"pageSize" binding:"required" `
}

// GetRegionList get Region list some field value or some condition
func GetRegionList(q *QueryRegionForm, dbs ...*gorm.DB) (ret []*Region, err error) {
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
	if q.Name != nil {
		db = db.Where("name = ?", *q.Name)
	}
	if q.Cname != nil {
		db = db.Where("cname = ?", *q.Cname)
	}
	if q.LowerName != nil {
		db = db.Where("lower_name = ?", *q.LowerName)
	}
	if q.Longitude != nil {
		db = db.Where("longitude = ?", *q.Longitude)
	}
	if q.Latitude != nil {
		db = db.Where("latitude = ?", *q.Latitude)
	}

	if err = db.Find(&ret).Error; err != nil {
		return
	}
	return
}

// QueryByID query cond by ID
func (t *Region) SetQueryByID(id int64) *Region {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *Region) GetByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).First(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetRegion
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *Region) DeleteByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteRegion
		return
	}
	return
}
