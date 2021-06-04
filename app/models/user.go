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
	ErrCreateUser = errors.New("create User failed")
	ErrDeleteUser = errors.New("delete User failed")
	ErrGetUser    = errors.New("get User failed")
	ErrUpdateUser = errors.New("update User failed")
)

// User table
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Uid  string `json:"uid" gorm:"type:varchar(36)"`
	Name string `json:"name"`
}

func init() {
	Tables = append(Tables, &User{})
}

// NewUser new
func NewUser() *User {
	return new(User)
}

// UserTableName TableName
func UserTableName() string {
	return "users"
}

// CreateUser create one record
func CreateUser(data interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Table(UserTableName()).Create(data).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrCreateUser
		return
	}
	return
}

// DelUser delete record
func DelUser(id interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).Delete("id = ?", id).Error; err != nil {
		err = ErrDeleteUser
		return
	}
	return
}

// UpUser update record
func UpUser(id, m interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).Where("id = ?", id).Updates(m).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrUpdateUser
		return
	}
	return
}

// GetAllUser get all record
func GetAllUser(dbs ...*gorm.DB) (res []*User, err error) {
	if err = GetDB(dbs...).Find(&res).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// CountUser get count
func CountUser(dbs ...*gorm.DB) (res int64) {
	GetDB(dbs...).Model(&User{}).Count(&res)
	return
}

// DelUserBatch delete User batch
func DelUserBatch(ids []string, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).Delete("id in ?", ids).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}

// CreateUserBatch create User batch
func CreateUserBatch(data []*User, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).Create(data).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrCreateUser
		return
	}
	return
}

// GetUserBatch get User list some field value or some condition
func GetUserBatch(q *form.QueryUserForm, res interface{}, dbs ...*gorm.DB) (err error) {
	var (
		db = GetDB(dbs...).Model(&User{})
	)
	if q.Uid != nil {
		db = db.Where("uid = ?", q.Uid)
	}
	if q.Name != nil {
		db = db.Where("name = ?", q.Name)
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

// GetUserID get one record by ID
func GetUserID(id interface{}, res interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).First(res, "id = ?", id).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrGetUser
		return
	}
	return
}

// DelUserID delete record by ID
func DelUserID(id interface{}, dbs ...*gorm.DB) (err error) {
	if err = GetDB(dbs...).Model(&User{}).Delete("id = ?", id).Error; err != nil {
		ModelLog.Errorln(err)
		err = ErrDeleteUser
		return
	}
	return
}
