package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/form"
	"github.com/olongfen/demo/app/services"
)

// CtlUser ctrl
type CtlUser struct{}

// Create create User one record
// @tags User
// @Summary create User one record
// @Description create User one record
// @Accept json
// @Produce json
// @Param {} body form.CreateUserForm true "添加User表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/users [post]
func (ct *CtlUser) Create(c *gin.Context) {
	var (
		req  = new(form.CreateUserForm)
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
	if data, err = services.CreateUser(req); err != nil {
		return
	}

}

// AddBatch create User batch record
// @tags User
// @Summary create User batch record
// @Description create User batch record
// @Accept json
// @Produce json
// @Param  {} body form.CreateUserBatchForm true "添加User表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/users/batch [post]
func (ct *CtlUser) AddBatch(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  = new(form.CreateUserBatchForm)
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

	if data, err = services.CreateUserBatch(req); err != nil {
		return
	}
}

// Update update User one record
// @tags User
// @Summary edit User one record
// @Description edit User one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param {} body form.UpUserForm true "update User form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/users/:id [put]
func (ct *CtlUser) Update(c *gin.Context) {
	var (
		req  = new(form.UpUserForm)
		key  string
		err  error
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(nil).Response()
		}
	}()
	if key = c.Param("id"); len(key) == 0 {
		err = fmt.Errorf("%s must be send", "id")
		return
	}
	if err = c.ShouldBind(&req); err != nil {
		return
	}
	if err = services.UpUser(key, req); err != nil {
		return
	}
}

// Get get User one record
// @tags User
// @Summary get User one record
// @Description get User one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/users/:id  [get]
func (ct *CtlUser) Get(c *gin.Context) {
	var (
		data interface{}
		key  string
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
	if key = c.Param("id"); len(key) == 0 {
		err = fmt.Errorf("%s must be send", "id")
		return
	}
	if data, err = services.GetUser("id", key); err != nil {
		return
	}
}

// GetBatch get User batch record
// @tags User
// @Summary get User batch record
// @Description get User batch record
// @Accept json
// @Produce json
// @Param	uid query string  false "Uid"
// @Param	name query string  false "Name"// @Param orderMap query string false "example: orderMap[column]=desc"
// @Param pageSize query int false "page size"
// @Param pageNum query int false "page num"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/users  [get]
func (ct *CtlUser) GetBatch(c *gin.Context) {
	var (
		data interface{}
		req  = new(form.QueryUserForm)
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

	if data, err = services.GetUserBatch(req); err != nil {
		return
	}
}

// Delete delete User one record
// @tags User
// @Summary delete User one record
// @Description delete User one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/users/:id [delete]
func (ct *CtlUser) Delete(c *gin.Context) {
	var (
		data interface{}
		err  error
		key  string
		code = response.CodeFail
	)
	defer func() {
		if err != nil {
			response.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			response.NewGinResponse(c).Success(data).Response()
		}
	}()
	if key = c.Param("id"); len(key) == 0 {
		err = fmt.Errorf("%s must be send", "id")
		return
	}
	if err = services.DelUser("id", key); err != nil {
		return
	}
}

// DeleteBatch delete User list record
// @tags User
// @Summary delete User list record
// @Description delete User list record
// @Accept json
// @Produce json
// @Param ids body []string true "User id list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/users/batch [delete]
func (ct *CtlUser) DeleteBatch(c *gin.Context) {
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
	if err = services.DelUserBatch(ids); err != nil {
		return
	}
}
