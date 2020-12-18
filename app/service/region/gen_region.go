package srv_region

import (
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/demo/app/model/region"
	"strconv"
)

// AddRegionReqForm
type AddRegionReqForm struct {
	Name      string `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Cname     string `json:"cname" form:"cname"`         // if required, add binding:"required" to tag by self
	LowerName string `json:"lowerName" form:"lowerName"` // if required, add binding:"required" to tag by self
	Longitude string `json:"longitude" form:"longitude"` // if required, add binding:"required" to tag by self
	Latitude  string `json:"latitude" form:"latitude"`   // if required, add binding:"required" to tag by self
}

func (a *AddRegionReqForm) Valid() (err error) {
	return
}

// EditRegionReqForm
type EditRegionReqForm struct {
	ID        int64   `json:"id" form:"id" binding:"required"`
	Name      *string `json:"name" form:"name"`           // if required, add binding:"required" to tag by self
	Cname     *string `json:"cname" form:"cname"`         // if required, add binding:"required" to tag by self
	LowerName *string `json:"lowerName" form:"lowerName"` // if required, add binding:"required" to tag by self
	Longitude *string `json:"longitude" form:"longitude"` // if required, add binding:"required" to tag by self
	Latitude  *string `json:"latitude" form:"latitude"`   // if required, add binding:"required" to tag by self
}

func (a *EditRegionReqForm) Valid() (err error) {
	return
}

func (a *EditRegionReqForm) ToMAP() (ret map[string]interface{}) {
	ret = make(map[string]interface{}, 0)
	if a.Name != nil {
		ret["name"] = *a.Name
	}
	if a.Cname != nil {
		ret["cname"] = *a.Cname
	}
	if a.LowerName != nil {
		ret["lower_name"] = *a.LowerName
	}
	if a.Longitude != nil {
		ret["longitude"] = *a.Longitude
	}
	if a.Latitude != nil {
		ret["latitude"] = *a.Latitude
	}
	return
}

// AddRegionOne add
func AddRegionOne(req *AddRegionReqForm) (ret *model_region.Region, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = new(model_region.Region)
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

type RegionBatchForm []*AddRegionReqForm

// AddRegionBatch add Region
func AddRegionBatch(req RegionBatchForm) (ret []*model_region.Region, err error) {
	var (
		datas []*model_region.Region
	)
	if err = mapstructure.Decode(req, &datas); err != nil {
		return
	}
	// if needed todo add you business logic code
	if err = model_region.AddRegionBatch(datas); err != nil {
		return
	}
	//
	ret = datas
	return
}

// EditRegionOne edit
func EditRegionOne(req *EditRegionReqForm) (ret *model_region.Region, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = model_region.NewRegion()
	)
	// if needed todo add you business logic code code

	if err = data.SetQueryByID(int64(req.ID)).Updates(req.ToMAP()); err != nil {
		return
	}

	//
	ret = data
	return
}

// GetRegionPage get page Region data
func GetRegionPage(req *model_region.QueryRegionForm) (ret []*model_region.Region, err error) {
	var (
		datas []*model_region.Region
	)
	// if needed todo add you business logic code code

	if datas, err = model_region.GetRegionList(req); err != nil {
		return
	}

	//
	ret = datas
	return
}

// GetRegionOne get Region
func GetRegionOne(in string) (ret *model_region.Region, err error) {
	var (
		id, _ = strconv.Atoi(in)
		d     = model_region.NewRegion().SetQueryByID(int64(id))
	)
	if err = d.GetByID(); err != nil {
		return
	}

	ret = d
	return
}

// DeleteRegionOne delete Region
func DeleteRegionOne(in string) (err error) {

	var (
		id, _ = strconv.Atoi(in)
		d     = model_region.NewRegion().SetQueryByID(int64(id))
	)
	// if needed todo add you business logic code
	return d.DeleteByID()
}

// DeleteRegionBatch delete Region
func DeleteRegionBatch(ids []string) (err error) {
	// if needed todo add you business logic code
	return model_region.DeleteRegionBatch(ids)
}
