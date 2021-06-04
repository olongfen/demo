package services

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/demo/app/form"
	"github.com/olongfen/demo/app/models"
)

// CreateProjectBusinessIndicator create one record
func CreateProjectBusinessIndicator(req *form.CreateProjectBusinessIndicatorForm) (res *models.ProjectBusinessIndicator, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = new(models.ProjectBusinessIndicator)
	)
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	// if needed todo add you business logic code

	if err = models.CreateProjectBusinessIndicator(data); err != nil {
		return
	}

	//
	res = data
	return
}

// CreateProjectBusinessIndicatorBatch create ProjectBusinessIndicator  batch record
func CreateProjectBusinessIndicatorBatch(req *form.CreateProjectBusinessIndicatorBatchForm) (res []*models.ProjectBusinessIndicator, err error) {
	var (
		data []*models.ProjectBusinessIndicator
	)
	if err = mapstructure.Decode(req, &data); err != nil {
		return
	}
	// if needed todo add you business logic code
	if err = models.CreateProjectBusinessIndicatorBatch(data); err != nil {
		return
	}
	//
	res = data
	return
}

// UpProjectBusinessIndicator edit ProjectBusinessIndicator one record
func UpProjectBusinessIndicator(id interface{}, req *form.UpProjectBusinessIndicatorForm) (err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = models.NewProjectBusinessIndicator()
	)
	// if needed todo add you business logic code code
	if err = mapstructure.Decode(req, data); err != nil {
		return
	}
	if err = models.UpProjectBusinessIndicator(id, data); err != nil {
		return
	}

	return
}

// GetProjectBusinessIndicatorBatch get ProjectBusinessIndicator list  data
func GetProjectBusinessIndicatorBatch(req *form.QueryProjectBusinessIndicatorForm) (res []*form.ProjectBusinessIndicatorRes, err error) {
	var (
		data []*form.ProjectBusinessIndicatorRes
	)
	// if needed todo add you business logic code code
	if err = models.GetProjectBusinessIndicatorBatch(req, &data); err != nil {
		return
	}

	//
	res = data
	return
}

// GetProjectBusinessIndicator get ProjectBusinessIndicator one record
func GetProjectBusinessIndicator(field string, value interface{}) (res *form.ProjectBusinessIndicatorRes, err error) {
	var (
		data = new(form.ProjectBusinessIndicatorRes)
	)
	switch field {
	case "id":
		if err = models.GetProjectBusinessIndicatorID(value, data); err != nil {
			return
		}
	case "project_code":
		if err = models.GetProjectBusinessIndicatorProjectCode(value, data); err != nil {
			return
		}
	default:
		err = fmt.Errorf("field: %s not support in this way", field)
	}
	res = data
	return
}

// DelProjectBusinessIndicator delete ProjectBusinessIndicator one record
func DelProjectBusinessIndicator(field string, value interface{}) (err error) {
	switch field {
	case "id":
		if err = models.DelProjectBusinessIndicatorID(value); err != nil {
			return
		}
	case "project_code":
		if err = models.DelProjectBusinessIndicatorProjectCode(value); err != nil {
			return
		}
	default:
		err = fmt.Errorf("field: %s not support in this way", field)
	}
	return
}

// DelProjectBusinessIndicatorBatch delete ProjectBusinessIndicator batch record
func DelProjectBusinessIndicatorBatch(ids []string) (err error) {
	// if needed todo add you business logic code
	return models.DelProjectBusinessIndicatorBatch(ids)
}
