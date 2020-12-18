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
	Name  string `grom:"unique_index"`
	Age   int
	Class string
}

func init() {
	model_common.Tables = append(model_common.Tables, &User{})
}

// NewUser new
func NewUser() *User {
	return new(User)
}

// Add add one record
func (t *User) Add(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Create(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrCreateUser
		return
	}
	return
}

// Delete delete record
func (t *User) Delete(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		err = ErrDeleteUser
		return
	}
	return
}

// Updates update record
func (t *User) Updates(m map[string]interface{}, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(t).Updates(m).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrUpdateUser
		return
	}
	return
}

// GetUserAll get all record
func GetUserAll(dbs ...*gorm.DB) (ret []*User, err error) {
	if err = model_common.GetDB(dbs...).Find(&ret).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// GetUserCount get count
func GetUserCount(dbs ...*gorm.DB) (ret int64) {
	model_common.GetDB(dbs...).Model(&User{}).Count(&ret)
	return
}

// DeleteUserBatch delete User batch
func DeleteUserBatch(ids []string, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&User{}).Delete("id in ?", ids).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}

// AddUserBatch add User batch
func AddUserBatch(datas []*User, dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Model(&User{}).Create(datas).Error; err != nil {
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

	Age   *model_common.FieldData `json:"age" form:"age"`     // if required, add binding:"required" to tag by self
	Class *string                 `json:"class" form:"class"` // if required, add binding:"required" to tag by self

	Order    []string `json:"order" form:"order"`
	PageNum  int      `json:"pageNum" form:"pageNum" binding:"required"`
	PageSize int      `json:"pageSize" form:"pageSize" binding:"required" `
}

// GetUserList get User list some field value or some condition
func GetUserList(q *QueryUserForm, dbs ...*gorm.DB) (ret []*User, err error) {
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
func (t *User) SetQueryByID(id uint) *User {
	t.ID = id
	return t
}

// GetByID get one record by ID
func (t *User) GetByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).First(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *User) DeleteByID(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}

// QueryByName query cond by Name
func (t *User) SetQueryByName(name string) *User {
	t.Name = name
	return t
}

// GetByName get one record by Name
func (t *User) GetByName(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).First(t, "name = ?", t.Name).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *User) DeleteByName(dbs ...*gorm.DB) (err error) {
	if err = model_common.GetDB(dbs...).Delete(t, "name = ?", t.Name).Error; err != nil {
		model_common.ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}
