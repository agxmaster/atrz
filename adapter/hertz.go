package adapter

type HertzCtxCore interface {
	JSON(httpCode int, body interface{})
	BindJSON(body interface{}) error
	BindAndValidate(body interface{}) error
	Param(key string) string
	Method() []byte
	Path() []byte
}
