package form

// QueryUserForm query User  form ;  if some field is required, create binding:"required" to tag by self
type QueryUserForm struct {
	Uid      *string           `json:"uid" form:"uid"`           // cond Uid
	Name     *string           `json:"name" form:"name"`         // cond Name
	OrderMap map[string]string `json:"orderMap" form:"orderMap"` // example: orderMap[column]=desc
	PageNum  int               `json:"pageNum" form:"pageNum"`   // get all without uploading
	PageSize int               `json:"pageSize" form:"pageSize"` // get all without uploading
}

// CreateUserForm create User form
type CreateUserForm struct {
	Uid  string `json:"uid" form:"uid"`   // uid
	Name string `json:"name" form:"name"` // name
}

// Valid create User  form verify
func (a *CreateUserForm) Valid() (err error) {
	return
}

type CreateUserBatchForm []*CreateUserForm

// UpUserForm  edit User form
type UpUserForm struct {
	Uid  string `json:"uid" form:"uid"`   // uid
	Name string `json:"name" form:"name"` // name
}

// Valid  edit User form verify
func (a *UpUserForm) Valid() (err error) {
	return
}

type UserRes struct {
	Uid  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}
