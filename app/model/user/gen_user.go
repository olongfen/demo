package model_user

import (
	"errors"
	"github.com/olongfen/demo/app/model/common"
	"gorm.io/gorm"
)

// Error
var (
	ErrCreateUser = errors.New("create User failed")
	ErrDeleteUser = errors.New("delete User failed")
	ErrGetUser    = errors.New("get User failed")
	ErrUpdateUser = errors.New("update User failed")
)

// User
type User struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}

// NewUser new
func NewUser() *User {
	return new(User)
}

// Add add one record
func (t *User) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateUser
		return
	}
	return
}

// Delete delete record
func (t *User) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		err = ErrDeleteUser
		return
	}
	return
}

// Updates update record
func (t *User) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(t).Where("id = ?", t.ID).Updates(m).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrUpdateUser
		return
	}
	return
}

// GetUserAll get all record
func GetUserAll(db *gorm.DB) (ret []*User, err error) {
	if err = db.Find(&ret).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// GetUserCount get count
func GetUserCount(db *gorm.DB) (ret int64) {
	db.Model(&User{}).Count(&ret)
	return
}

// DeleteUserBatch delete User batch
func DeleteUserBatch(db *gorm.DB, ids []string) (err error) {
	if err = db.Model(&User{}).Delete("id in ?", ids).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}

// AddUserBatch add User batch
func AddUserBatch(db *gorm.DB, datas []*User) (err error) {
	if err = db.Model(&User{}).Create(datas).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateUser
		return
	}
	return
}

//  QueryUserForm query form
type QueryUserForm struct {
	CreatedAt *model_common.FieldData `json:"createdAt" form:"createdAt"` // if required, add binding:"required" to tag by self
	UpdatedAt *model_common.FieldData `json:"updatedAt" form:"updatedAt"` // if required, add binding:"required" to tag by self
	Name      *model_common.FieldData `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Age       *model_common.FieldData `json:"age" form:"age"`             // if required, add binding:"required" to tag by self
	Class     *model_common.FieldData `json:"class" form:"class"`         // if required, add binding:"required" to tag by self

	Order    []string `json:"order" form:"order"`
	PageNum  int      `json:"pageNum" form:"pageNum" binding:"required"`
	PageSize int      `json:"pageSize" form:"pageSize" binding:"required" `
}

// GetUserList get User list some field value or some condition
func GetUserList(db *gorm.DB, q *QueryUserForm) (ret []*User, err error) {
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
func (t *User) SetQueryByID(id uint) *User {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *User) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *User) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}
