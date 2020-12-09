package model_admin

import (
	"errors"
	"github.com/olongfen/demo/app/model/common"
	"gorm.io/gorm"
)

// Error
var (
	ErrCreateAdmin = errors.New("create Admin failed")
	ErrDeleteAdmin = errors.New("delete Admin failed")
	ErrGetAdmin    = errors.New("get Admin failed")
	ErrUpdateAdmin = errors.New("update Admin failed")
)

// Admin
type Admin struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}

// NewAdmin new
func NewAdmin() *Admin {
	return new(Admin)
}

// Add add one record
func (t *Admin) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateAdmin
		return
	}
	return
}

// Delete delete record
func (t *Admin) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		err = ErrDeleteAdmin
		return
	}
	return
}

// Updates update record
func (t *Admin) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(t).Where("id = ?", t.ID).Updates(m).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrUpdateAdmin
		return
	}
	return
}

// GetAdminAll get all record
func GetAdminAll(db *gorm.DB) (ret []*Admin, err error) {
	if err = db.Find(&ret).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetAdmin
		return
	}
	return
}

// GetAdminCount get count
func GetAdminCount(db *gorm.DB) (ret int64) {
	db.Model(&Admin{}).Count(&ret)
	return
}

// DeleteAdminBatch delete Admin batch
func DeleteAdminBatch(db *gorm.DB, ids []string) (err error) {
	if err = db.Model(&Admin{}).Delete("id in ?", ids).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteAdmin
		return
	}
	return
}

// AddAdminBatch add Admin batch
func AddAdminBatch(db *gorm.DB, datas []*Admin) (err error) {
	if err = db.Model(&Admin{}).Create(datas).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateAdmin
		return
	}
	return
}

//  QueryAdminForm query form
type QueryAdminForm struct {
	CreatedAt *model_common.FieldData `json:"createdAt" form:"createdAt"` // if required, add binding:"required" to tag by self
	UpdatedAt *model_common.FieldData `json:"updatedAt" form:"updatedAt"` // if required, add binding:"required" to tag by self
	Name      *model_common.FieldData `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Age       *model_common.FieldData `json:"age" form:"age"`             // if required, add binding:"required" to tag by self
	Class     *model_common.FieldData `json:"class" form:"class"`         // if required, add binding:"required" to tag by self

	Order    []string `json:"order" form:"order"`
	PageNum  int      `json:"pageNum" form:"pageNum" binding:"required"`
	PageSize int      `json:"pageSize" form:"pageSize" binding:"required" `
}

// GetAdminList get Admin list some field value or some condition
func GetAdminList(db *gorm.DB, q *QueryAdminForm) (ret []*Admin, err error) {
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
		db = db.Where("name"+q.Name.Symbol+"?", q.Name.Value)
	}
	if q.Age != nil {
		db = db.Where("age"+q.Age.Symbol+"?", q.Age.Value)
	}
	if q.Class != nil {
		db = db.Where("class"+q.Class.Symbol+"?", q.Class.Value)
	}
	if err = db.Find(&ret).Error; err != nil {
		return
	}
	return
}

// QueryByID query cond by ID
func (t *Admin) SetQueryByID(id uint) *Admin {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *Admin) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetAdmin
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *Admin) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteAdmin
		return
	}
	return
}
