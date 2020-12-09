package api_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/model/admin"
	"github.com/olongfen/demo/app/service/admin"
)

// CtrlAdmin
type CtrlAdmin struct{}

// AddOne add Admin one record
// @tags Admin
// @Summary add Admin one record
// @Description add Admin one record
// @Accept json
// @Produce json
// @Param {} body srv_admin.AddAdminReqForm true "添加Admin表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/admin/add [post]
func (ct *CtrlAdmin) AddOne(c *gin.Context) {
	var (
		req  = &srv_admin.AddAdminReqForm{}
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
	if data, err = srv_admin.AddAdminOne(req); err != nil {
		return
	}

}

// AddList add Admin list record
// @tags Admin
// @Summary add Admin list record
// @Description add Admin list record
// @Accept json
// @Produce json
// @Param  {} body srv_admin.AdminBatchForm true "添加Admin表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/admin/addList [post]
func (ct *CtrlAdmin) AddList(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  srv_admin.AdminBatchForm
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

	if data, err = srv_admin.AddAdminBatch(req); err != nil {
		return
	}
}

// Edit edit Admin one record
// @tags Admin
// @Summary edit Admin one record
// @Description edit Admin one record
// @Accept json
// @Produce json
// @Param  {} body srv_admin.EditAdminReqForm true "编辑Admin表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/admin/edit [put]
func (ct *CtrlAdmin) Edit(c *gin.Context) {
	var (
		data interface{}
		req  = new(srv_admin.EditAdminReqForm)
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
	if data, err = srv_admin.EditAdminOne(req); err != nil {
		return
	}
}

// GetOne get Admin one record
// @tags Admin
// @Summary get Admin one record
// @Description get Admin one record
// @Accept json
// @Produce json
// @Param id query string true "Admin ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/admin/get  [get]
func (ct *CtrlAdmin) GetOne(c *gin.Context) {
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
	if data, err = srv_admin.GetAdminOne(id); err != nil {
		return
	}
}

// GetList get Admin list record
// @tags Admin
// @Summary get Admin list record
// @Description get Admin list record
// @Accept json
// @Produce json
// @Param {} body model_admin.QueryAdminForm true "获取Admin列表form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/admin/list  [get]
func (ct *CtrlAdmin) GetList(c *gin.Context) {
	var (
		data interface{}
		req  = new(model_admin.QueryAdminForm)
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
	if data, err = srv_admin.GetAdminPage(req); err != nil {
		return
	}
}

// DeleteOne delete Admin one record
// @tags Admin
// @Summary delete Admin one record
// @Description delete Admin one record
// @Accept json
// @Produce json
// @Param id body string true "Admin ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/admin/delete [delete]
func (ct *CtrlAdmin) DeleteOne(c *gin.Context) {
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
	if err = srv_admin.DeleteAdminOne(id); err != nil {
		return
	}
}

// DeleteList delete Admin list record
// @tags Admin
// @Summary delete Admin list record
// @Description delete Admin list record
// @Accept json
// @Produce json
// @Param ids body []string true "Admin ID list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/admin/deleteList [delete]
func (ct *CtrlAdmin) DeleteList(c *gin.Context) {
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
	if err = srv_admin.DeleteAdminBatch(ids); err != nil {
		return
	}
}
