package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/response"
	"github.com/olongfen/demo/app/form"
	"github.com/olongfen/demo/app/services"
)

// CtlProjectBusinessIndicator ctrl
type CtlProjectBusinessIndicator struct{}

// Create create ProjectBusinessIndicator one record
// @tags ProjectBusinessIndicator
// @Summary create ProjectBusinessIndicator one record
// @Description create ProjectBusinessIndicator one record
// @Accept json
// @Produce json
// @Param {} body form.CreateProjectBusinessIndicatorForm true "添加ProjectBusinessIndicator表单"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/project_business_indicators [post]
func (ct *CtlProjectBusinessIndicator) Create(c *gin.Context) {
	var (
		req  = new(form.CreateProjectBusinessIndicatorForm)
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
	if data, err = services.CreateProjectBusinessIndicator(req); err != nil {
		return
	}

}

// AddBatch create ProjectBusinessIndicator batch record
// @tags ProjectBusinessIndicator
// @Summary create ProjectBusinessIndicator batch record
// @Description create ProjectBusinessIndicator batch record
// @Accept json
// @Produce json
// @Param  {} body form.CreateProjectBusinessIndicatorBatchForm true "添加ProjectBusinessIndicator表单列表"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/project_business_indicators/batch [post]
func (ct *CtlProjectBusinessIndicator) AddBatch(c *gin.Context) {
	var (
		data interface{}
		code = response.CodeFail
		req  = new(form.CreateProjectBusinessIndicatorBatchForm)
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

	if data, err = services.CreateProjectBusinessIndicatorBatch(req); err != nil {
		return
	}
}

// Update update ProjectBusinessIndicator one record
// @tags ProjectBusinessIndicator
// @Summary edit ProjectBusinessIndicator one record
// @Description edit ProjectBusinessIndicator one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param {} body form.UpProjectBusinessIndicatorForm true "update ProjectBusinessIndicator form"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/project_business_indicators/:id [put]
func (ct *CtlProjectBusinessIndicator) Update(c *gin.Context) {
	var (
		req  = new(form.UpProjectBusinessIndicatorForm)
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
	if err = services.UpProjectBusinessIndicator(key, req); err != nil {
		return
	}
}

// Get get ProjectBusinessIndicator one record
// @tags ProjectBusinessIndicator
// @Summary get ProjectBusinessIndicator one record
// @Description get ProjectBusinessIndicator one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/project_business_indicators/:id  [get]
func (ct *CtlProjectBusinessIndicator) Get(c *gin.Context) {
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
	if data, err = services.GetProjectBusinessIndicator("id", key); err != nil {
		return
	}
}

// GetBatch get ProjectBusinessIndicator batch record
// @tags ProjectBusinessIndicator
// @Summary get ProjectBusinessIndicator batch record
// @Description get ProjectBusinessIndicator batch record
// @Accept json
// @Produce json
// @Param	facilities query string  false "Facilities"
// @Param	qualitativeDescription query string  false "QualitativeDescription"
// @Param	issue query string  false "Issue"
// @Param	quantitativeAssessment query string  false "QuantitativeAssessment"
// @Param	change query string  false "Change"
// @Param	ageMap query string  false  "example: ageMap[>]=some value&ageMap[<]=some value; key must be >,>=,<,<=,!=,=,gt,ge,lt,le,ne,eq"
// @Param orderMap query string false "example: orderMap[column]=desc"
// @Param pageSize query int false "page size"
// @Param pageNum query int false "page num"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router /api/v1/project_business_indicators  [get]
func (ct *CtlProjectBusinessIndicator) GetBatch(c *gin.Context) {
	var (
		data interface{}
		req  = new(form.QueryProjectBusinessIndicatorForm)
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

	req.AgeMap = c.QueryMap("ageMap")
	if data, err = services.GetProjectBusinessIndicatorBatch(req); err != nil {
		return
	}
}

// Delete delete ProjectBusinessIndicator one record
// @tags ProjectBusinessIndicator
// @Summary delete ProjectBusinessIndicator one record
// @Description delete ProjectBusinessIndicator one record
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/project_business_indicators/:id [delete]
func (ct *CtlProjectBusinessIndicator) Delete(c *gin.Context) {
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
	if err = services.DelProjectBusinessIndicator("id", key); err != nil {
		return
	}
}

// DeleteBatch delete ProjectBusinessIndicator list record
// @tags ProjectBusinessIndicator
// @Summary delete ProjectBusinessIndicator list record
// @Description delete ProjectBusinessIndicator list record
// @Accept json
// @Produce json
// @Param ids body []string true "ProjectBusinessIndicator id list"
// @Success 200  {object} response.Response
// @Failure 500  {object} response.Response
// @router  /api/v1/project_business_indicators/batch [delete]
func (ct *CtlProjectBusinessIndicator) DeleteBatch(c *gin.Context) {
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
	if err = services.DelProjectBusinessIndicatorBatch(ids); err != nil {
		return
	}
}
