package handler

type Handler struct {
	ServiceHandler interface{}
}

func NewHandler(
	Sh interface{},
) *Handler {
	return &Handler{
		ServiceHandler: Sh,
	}
}
