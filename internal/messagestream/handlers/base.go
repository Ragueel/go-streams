package handlers

type IBaseHandler interface {
	Handle(msg interface{}) error
}
