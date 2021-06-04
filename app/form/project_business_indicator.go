package form

// QueryProjectBusinessIndicatorForm query ProjectBusinessIndicator  form ;  if some field is required, create binding:"required" to tag by self
type QueryProjectBusinessIndicatorForm struct {
	Facilities             *string           `json:"facilities" form:"facilities"`                         // cond Facilities
	QualitativeDescription *string           `json:"qualitativeDescription" form:"qualitativeDescription"` // cond QualitativeDescription
	Issue                  *string           `json:"issue" form:"issue"`                                   // cond Issue
	QuantitativeAssessment *string           `json:"quantitativeAssessment" form:"quantitativeAssessment"` // cond QuantitativeAssessment
	Change                 *string           `json:"change" form:"change"`                                 // cond Change
	AgeMap                 map[string]string `json:"ageMap" form:"ageMap"`                                 // example: AgeMap[>]=some value&AgeMap[<]=some value; key must be ">,>=,<,<=,="
	OrderMap               map[string]string `json:"orderMap" form:"orderMap"`                             // example: orderMap[column]=desc
	PageNum                int               `json:"pageNum" form:"pageNum"`                               // get all without uploading
	PageSize               int               `json:"pageSize" form:"pageSize"`                             // get all without uploading
}

// CreateProjectBusinessIndicatorForm create ProjectBusinessIndicator form
type CreateProjectBusinessIndicatorForm struct {
	ProjectCode            string `json:"projectCode" form:"projectCode" binding:"required"`    // projectCode
	Facilities             string `json:"facilities" form:"facilities"`                         // facilities
	QualitativeDescription string `json:"qualitativeDescription" form:"qualitativeDescription"` // qualitativeDescription
	Issue                  string `json:"issue" form:"issue"`                                   // issue
	QuantitativeAssessment string `json:"quantitativeAssessment" form:"quantitativeAssessment"` // quantitativeAssessment
	Change                 string `json:"change" form:"change"`                                 // change
	Age                    int    `json:"age" form:"age"`                                       // age
}

// Valid create ProjectBusinessIndicator  form verify
func (a *CreateProjectBusinessIndicatorForm) Valid() (err error) {
	return
}

type CreateProjectBusinessIndicatorBatchForm []*CreateProjectBusinessIndicatorForm

// UpProjectBusinessIndicatorForm  edit ProjectBusinessIndicator form
type UpProjectBusinessIndicatorForm struct {
	ProjectCode            string `json:"projectCode" form:"projectCode"`                       // projectCode
	Facilities             string `json:"facilities" form:"facilities"`                         // facilities
	QualitativeDescription string `json:"qualitativeDescription" form:"qualitativeDescription"` // qualitativeDescription
	Issue                  string `json:"issue" form:"issue"`                                   // issue
	QuantitativeAssessment string `json:"quantitativeAssessment" form:"quantitativeAssessment"` // quantitativeAssessment
	Change                 string `json:"change" form:"change"`                                 // change
	Age                    int    `json:"age" form:"age"`                                       // age
}

// Valid  edit ProjectBusinessIndicator form verify
func (a *UpProjectBusinessIndicatorForm) Valid() (err error) {
	return
}

type ProjectBusinessIndicatorRes struct {
	Facilities             string `json:"facilities,omitempty"`
	QualitativeDescription string `json:"qualitativeDescription,omitempty"`
	Issue                  string `json:"issue,omitempty"`
	QuantitativeAssessment string `json:"quantitativeAssessment,omitempty"`
	Change                 string `json:"change,omitempty"`
	Age                    int    `json:"age,omitempty"`
}
