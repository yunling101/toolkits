package common

type H map[string]interface{}
type Request map[string]interface{}

type Response struct {
	Data  interface{}
	Error error
}

type IdStringName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Option interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	GetAll() map[string]interface{}
	List() []string
	Delete(key string)
}
