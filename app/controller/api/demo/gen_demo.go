package api_demo

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/model/demo"
	"github.com/olongfen/demo/app/service/demo"
)

// CtrlDemo
type CtrlDemo struct{}

// AddOne add Demo one record
// @tags Demo
// @Summary add Demo one record
// @Description add Demo one record
// @Accept json
// @Produce json
// @Param {} body srv_demo.AddDemoReqForm true "添加Demo表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/demo/add [post]
func (ct *CtrlDemo) AddOne(c *gin.Context) {
	var (
		req  = &srv_demo.AddDemoReqForm{}
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
	if data, err = srv_demo.AddDemoOne(req); err != nil {
		return
	}

}

// AddList add Demo list record
// @tags Demo
// @Summary add Demo list record
// @Description add Demo list record
// @Accept json
// @Produce json
// @Param  {} body srv_demo.DemoBatchForm true "添加Demo表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/demo/addList [post]
func (ct *CtrlDemo) AddList(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  srv_demo.DemoBatchForm
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

	if data, err = srv_demo.AddDemoBatch(req); err != nil {
		return
	}
}

// Edit edit Demo one record
// @tags Demo
// @Summary edit Demo one record
// @Description edit Demo one record
// @Accept json
// @Produce json
// @Param  {} body srv_demo.EditDemoReqForm true "编辑Demo表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/demo/edit [put]
func (ct *CtrlDemo) Edit(c *gin.Context) {
	var (
		data interface{}
		req  = new(srv_demo.EditDemoReqForm)
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
	if data, err = srv_demo.EditDemoOne(req); err != nil {
		return
	}
}

// GetOne get Demo one record
// @tags Demo
// @Summary get Demo one record
// @Description get Demo one record
// @Accept json
// @Produce json
// @Param id query string true "Demo ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/demo/get  [get]
func (ct *CtrlDemo) GetOne(c *gin.Context) {
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
	if data, err = srv_demo.GetDemoOne(id); err != nil {
		return
	}
}

// GetList get Demo list record
// @tags Demo
// @Summary get Demo list record
// @Description get Demo list record
// @Accept json
// @Produce json
// @Param {} query model_demo.QueryDemoForm true "获取Demo列表form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/demo/list  [get]
func (ct *CtrlDemo) GetList(c *gin.Context) {
	var (
		data interface{}
		req  = new(model_demo.QueryDemoForm)
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
	if data, err = srv_demo.GetDemoPage(req); err != nil {
		return
	}
}

// DeleteOne delete Demo one record
// @tags Demo
// @Summary delete Demo one record
// @Description delete Demo one record
// @Accept json
// @Produce json
// @Param id body string true "Demo ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/demo/delete [delete]
func (ct *CtrlDemo) DeleteOne(c *gin.Context) {
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
	if err = srv_demo.DeleteDemoOne(id); err != nil {
		return
	}
}

// DeleteList delete Demo list record
// @tags Demo
// @Summary delete Demo list record
// @Description delete Demo list record
// @Accept json
// @Produce json
// @Param ids body []string true "Demo ID list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/demo/deleteList [delete]
func (ct *CtrlDemo) DeleteList(c *gin.Context) {
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
	if err = srv_demo.DeleteDemoBatch(ids); err != nil {
		return
	}
}
