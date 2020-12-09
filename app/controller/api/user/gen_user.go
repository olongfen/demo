package api_user

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/model/user"
	"github.com/olongfen/demo/app/service/user"
)

// CtrlUser
type CtrlUser struct{}

// AddOne add User one record
// @tags User
// @Summary add User one record
// @Description add User one record
// @Accept json
// @Produce json
// @Param {} body srv_user.AddUserReqForm true "添加User表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/user/add [post]
func (ct *CtrlUser) AddOne(c *gin.Context) {
	var (
		req  = &srv_user.AddUserReqForm{}
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
	if data, err = srv_user.AddUserOne(req); err != nil {
		return
	}

}

// AddList add User list record
// @tags User
// @Summary add User list record
// @Description add User list record
// @Accept json
// @Produce json
// @Param  {} body srv_user.UserBatchForm true "添加User表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/user/addList [post]
func (ct *CtrlUser) AddList(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  srv_user.UserBatchForm
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

	if data, err = srv_user.AddUserBatch(req); err != nil {
		return
	}
}

// Edit edit User one record
// @tags User
// @Summary edit User one record
// @Description edit User one record
// @Accept json
// @Produce json
// @Param  {} body srv_user.EditUserReqForm true "编辑User表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/user/edit [put]
func (ct *CtrlUser) Edit(c *gin.Context) {
	var (
		data interface{}
		req  = new(srv_user.EditUserReqForm)
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
	if data, err = srv_user.EditUserOne(req); err != nil {
		return
	}
}

// GetOne get User one record
// @tags User
// @Summary get User one record
// @Description get User one record
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/user/get  [get]
func (ct *CtrlUser) GetOne(c *gin.Context) {
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
	if data, err = srv_user.GetUserOne(id); err != nil {
		return
	}
}

// GetList get User list record
// @tags User
// @Summary get User list record
// @Description get User list record
// @Accept json
// @Produce json
// @Param {} body model_user.QueryUserForm true "获取User列表form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/user/list  [get]
func (ct *CtrlUser) GetList(c *gin.Context) {
	var (
		data interface{}
		req  = new(model_user.QueryUserForm)
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
	if data, err = srv_user.GetUserPage(req); err != nil {
		return
	}
}

// DeleteOne delete User one record
// @tags User
// @Summary delete User one record
// @Description delete User one record
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/user/delete [delete]
func (ct *CtrlUser) DeleteOne(c *gin.Context) {
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
	if err = srv_user.DeleteUserOne(id); err != nil {
		return
	}
}

// DeleteList delete User list record
// @tags User
// @Summary delete User list record
// @Description delete User list record
// @Accept json
// @Produce json
// @Param ids body []string true "User ID list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/user/deleteList [delete]
func (ct *CtrlUser) DeleteList(c *gin.Context) {
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
	if err = srv_user.DeleteUserBatch(ids); err != nil {
		return
	}
}
