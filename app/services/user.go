package services

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/demo/app/form"
	"github.com/olongfen/demo/app/models"
)

// CreateUser create one record
func CreateUser(req *form.CreateUserForm) (res *models.User, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = new(models.User)
	)
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	// if needed todo add you business logic code

	if err = models.CreateUser(data); err != nil {
		return
	}

	//
	res = data
	return
}

// CreateUserBatch create User  batch record
func CreateUserBatch(req *form.CreateUserBatchForm) (res []*models.User, err error) {
	var (
		data []*models.User
	)
	if err = mapstructure.Decode(req, &data); err != nil {
		return
	}
	// if needed todo add you business logic code
	if err = models.CreateUserBatch(data); err != nil {
		return
	}
	//
	res = data
	return
}

// UpUser edit User one record
func UpUser(id interface{}, req *form.UpUserForm) (err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = models.NewUser()
	)
	// if needed todo add you business logic code code
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	if err = models.UpUser(id, data); err != nil {
		return
	}

	return
}

// GetUserBatch get User list  data
func GetUserBatch(req *form.QueryUserForm) (res []*form.UserRes, err error) {
	var (
		data []*form.UserRes
	)
	// if needed todo add you business logic code code
	if err = models.GetUserBatch(req, &data); err != nil {
		return
	}

	//
	res = data
	return
}

// GetUser get User one record
func GetUser(field string, value interface{}) (res *form.UserRes, err error) {
	var (
		data = new(form.UserRes)
	)
	switch field {
	case "id":
		if err = models.GetUserID(value, data); err != nil {
			return
		}
	default:
		err = fmt.Errorf("field: %s not support in this way", field)
	}
	res = data
	return
}

// DelUser delete User one record
func DelUser(field string, value interface{}) (err error) {
	switch field {
	case "id":
		if err = models.DelUserID(value); err != nil {
			return
		}
	default:
		err = fmt.Errorf("field: %s not support in this way", field)
	}
	return
}

// DelUserBatch delete User batch record
func DelUserBatch(ids []string) (err error) {
	// if needed todo add you business logic code
	return models.DelUserBatch(ids)
}
