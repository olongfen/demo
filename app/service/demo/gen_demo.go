package srv_demo

import (
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/demo/app/model/demo"
	"strconv"
)

// AddDemoReqForm
type AddDemoReqForm struct {
	Name  string `json:"name" form:"name"`   // if required, add binding:"required" to tag by self
	Age   int    `json:"age" form:"age"`     // if required, add binding:"required" to tag by self
	Class string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *AddDemoReqForm) Valid() (err error) {
	return
}

// EditDemoReqForm
type EditDemoReqForm struct {
	ID    int64   `json:"id" form:"id" binding:"required"`
	Name  *string `json:"name" form:"name"`   // if required, add binding:"required" to tag by self
	Age   *int    `json:"age" form:"age"`     // if required, add binding:"required" to tag by self
	Class *string `json:"class" form:"class"` // if required, add binding:"required" to tag by self
}

func (a *EditDemoReqForm) Valid() (err error) {
	return
}

func (a *EditDemoReqForm) ToMAP() (ret map[string]interface{}) {
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

// AddDemoOne add
func AddDemoOne(req *AddDemoReqForm) (ret *model_demo.Demo, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = new(model_demo.Demo)
	)
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	// if needed todo add you business logic code

	if err = data.Add(); err != nil {
		return
	}

	//
	ret = data
	return
}

type DemoBatchForm []*AddDemoReqForm

// AddDemoBatch add Demo
func AddDemoBatch(req DemoBatchForm) (ret []*model_demo.Demo, err error) {
	var (
		datas []*model_demo.Demo
	)
	if err = mapstructure.Decode(req, &datas); err != nil {
		return
	}
	// if needed todo add you business logic code
	if err = model_demo.AddDemoBatch(datas); err != nil {
		return
	}
	//
	ret = datas
	return
}

// EditDemoOne edit
func EditDemoOne(req *EditDemoReqForm) (ret *model_demo.Demo, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = model_demo.NewDemo()
	)
	// if needed todo add you business logic code code

	if err = data.SetQueryByID(uint(req.ID)).Updates(req.ToMAP()); err != nil {
		return
	}

	//
	ret = data
	return
}

// GetDemoPage get page Demo data
func GetDemoPage(req *model_demo.QueryDemoForm) (ret []*model_demo.Demo, err error) {
	var (
		datas []*model_demo.Demo
	)
	// if needed todo add you business logic code code

	if datas, err = model_demo.GetDemoList(req); err != nil {
		return
	}

	//
	ret = datas
	return
}

// GetDemoOne get Demo
func GetDemoOne(in string) (ret *model_demo.Demo, err error) {
	var (
		id int64
	)
	if id, err = strconv.ParseInt(in, 10, 64); err != nil {
		return
	}
	var (
		d = model_demo.NewDemo().SetQueryByID(uint(id))
	)
	if err = d.GetByID(); err != nil {
		return
	}

	ret = d
	return
}

// DeleteDemoOne delete Demo
func DeleteDemoOne(in string) (err error) {
	var (
		id int64
	)
	if id, err = strconv.ParseInt(in, 10, 64); err != nil {
		return
	}
	var (
		d = model_demo.NewDemo().SetQueryByID(uint(id))
	)
	// if needed todo add you business logic code
	return d.DeleteByID()
}

// DeleteDemoBatch delete Demo
func DeleteDemoBatch(ids []string) (err error) {
	// if needed todo add you business logic code
	return model_demo.DeleteDemoBatch(ids)
}
