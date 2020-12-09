package srv_admin

import (
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/demo/app/model/admin"
	"github.com/olongfen/demo/app/model/common"
	"strconv"
)

// AddAdminReqForm
type AddAdminReqForm struct {
	Name  string `json:"name" form:"name"`   // if required, add binding:"required" to tag by self
	Age   int    `json:"age" form:"age"`     // if required, add binding:"required" to tag by self
	Class string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *AddAdminReqForm) Valid() (err error) {
	return
}

// EditAdminReqForm
type EditAdminReqForm struct {
	ID    int64   `json:"id" form:"id" binding:"required"`
	Name  *string `json:"name" form:"name"`   // if required, add binding:"required" to tag by self
	Age   *int    `json:"age" form:"age"`     // if required, add binding:"required" to tag by self
	Class *string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *EditAdminReqForm) Valid() (err error) {
	return
}

func (a *EditAdminReqForm) ToMAP() (ret map[string]interface{}) {
	ret = make(map[string]interface{}, 0)
	if a.Name != nil {
		ret["name"] = *a.Name
	}
	if a.Age != nil {
		ret["age"] = *a.Age
	}
	if a.Class != nil {
		ret["class"] = *a.Class
	}
	return
}

// AddAdminOne add
func AddAdminOne(req *AddAdminReqForm) (ret *model_admin.Admin, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = new(model_admin.Admin)
	)
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	// if needed todo add you business logic code

	if err = data.Add(model_common.DB); err != nil {
		return
	}

	//
	ret = data
	return
}

type AdminBatchForm []*AddAdminReqForm

// AddAdminBatch add Admin
func AddAdminBatch(req AdminBatchForm) (ret []*model_admin.Admin, err error) {
	var (
		datas []*model_admin.Admin
	)
	if err = mapstructure.Decode(req, &datas); err != nil {
		return
	}
	// if needed todo add you business logic code
	if err = model_admin.AddAdminBatch(model_common.DB, datas); err != nil {
		return
	}
	//
	ret = datas
	return
}

// EditAdminOne edit
func EditAdminOne(req *EditAdminReqForm) (ret *model_admin.Admin, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = model_admin.NewAdmin()
	)
	// if needed todo add you business logic code code

	if err = data.SetQueryByID(uint(req.ID)).Updates(model_common.DB, req.ToMAP()); err != nil {
		return
	}

	//
	ret = data
	return
}

// GetAdminPage get page Admin data
func GetAdminPage(req *model_admin.QueryAdminForm) (ret []*model_admin.Admin, err error) {
	var (
		datas []*model_admin.Admin
	)
	// if needed todo add you business logic code code

	if datas, err = model_admin.GetAdminList(model_common.DB, req); err != nil {
		return
	}

	//
	ret = datas
	return
}

// GetAdminOne get Admin
func GetAdminOne(in string) (ret *model_admin.Admin, err error) {
	var (
		id, _ = strconv.Atoi(in)
		d     = model_admin.NewAdmin().SetQueryByID(uint(id))
	)
	if err = d.GetByID(model_common.DB); err != nil {
		return
	}

	ret = d
	return
}

// DeleteAdminOne delete Admin
func DeleteAdminOne(in string) (err error) {

	var (
		id, _ = strconv.Atoi(in)
		d     = model_admin.NewAdmin().SetQueryByID(uint(id))
	)
	// if needed todo add you business logic code
	return d.DeleteByID(model_common.DB)
}

// DeleteAdminBatch delete Admin
func DeleteAdminBatch(ids []string) (err error) {
	// if needed todo add you business logic code
	return model_admin.DeleteAdminBatch(model_common.DB, ids)
}
