package api_region

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/model/region"
	"github.com/olongfen/demo/app/service/region"
)

// CtrlRegion
type CtrlRegion struct{}

// AddOne add Region one record
// @tags Region
// @Summary add Region one record
// @Description add Region one record
// @Accept json
// @Produce json
// @Param {} body srv_region.AddRegionReqForm true "添加Region表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/region/add [post]
func (ct *CtrlRegion) AddOne(c *gin.Context) {
	var (
		req  = &srv_region.AddRegionReqForm{}
		data interface{}
		code = response.CodeFail
		err  error
	)

	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()

	if err = c.ShouldBind(req); err != nil {
		return
	}
	if data, err = srv_region.AddRegionOne(req); err != nil {
		return
	}

}

// AddList add Region list record
// @tags Region
// @Summary add Region list record
// @Description add Region list record
// @Accept json
// @Produce json
// @Param  {} body srv_region.RegionBatchForm true "添加Region表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/region/addList [post]
func (ct *CtrlRegion) AddList(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  srv_region.RegionBatchForm
		err  error
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		return
	}

	if data, err = srv_region.AddRegionBatch(req); err != nil {
		return
	}
}

// Edit edit Region one record
// @tags Region
// @Summary edit Region one record
// @Description edit Region one record
// @Accept json
// @Produce json
// @Param  {} body srv_region.EditRegionReqForm true "编辑Region表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/region/edit [put]
func (ct *CtrlRegion) Edit(c *gin.Context) {
	var (
		data interface{}
		req  = new(srv_region.EditRegionReqForm)
		err  error
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		return
	}
	if data, err = srv_region.EditRegionOne(req); err != nil {
		return
	}
}

// GetOne get Region one record
// @tags Region
// @Summary get Region one record
// @Description get Region one record
// @Accept json
// @Produce json
// @Param id query string true "Region ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/region/get  [get]
func (ct *CtrlRegion) GetOne(c *gin.Context) {
	var (
		data interface{}
		id   string
		err  error
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	id = c.Query("id")
	if data, err = srv_region.GetRegionOne(id); err != nil {
		return
	}
}

// GetList get Region list record
// @tags Region
// @Summary get Region list record
// @Description get Region list record
// @Accept json
// @Produce json
// @Param {} query model_region.QueryRegionForm true "获取Region列表form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/region/list  [get]
func (ct *CtrlRegion) GetList(c *gin.Context) {
	var (
		data interface{}
		req  = new(model_region.QueryRegionForm)
		err  error
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBindQuery(req); err != nil {
		return
	}
	if data, err = srv_region.GetRegionPage(req); err != nil {
		return
	}
}

// DeleteOne delete Region one record
// @tags Region
// @Summary delete Region one record
// @Description delete Region one record
// @Accept json
// @Produce json
// @Param id body string true "Region ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/region/delete [delete]
func (ct *CtrlRegion) DeleteOne(c *gin.Context) {
	var (
		data interface{}
		err  error
		id   string
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBind(&id); err != nil {
		return
	}
	if err = srv_region.DeleteRegionOne(id); err != nil {
		return
	}
}

// DeleteList delete Region list record
// @tags Region
// @Summary delete Region list record
// @Description delete Region list record
// @Accept json
// @Produce json
// @Param ids body []string true "Region ID list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/region/deleteList [delete]
func (ct *CtrlRegion) DeleteList(c *gin.Context) {
	var (
		data interface{}
		ids  []string
		err  error
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBind(&ids); err != nil {
		return
	}
	if err = srv_region.DeleteRegionBatch(ids); err != nil {
		return
	}
}
