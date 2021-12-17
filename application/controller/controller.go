package controller

type Controller interface {
	Handle() ([]byte, error)
	Bind(interface{}) error
}
