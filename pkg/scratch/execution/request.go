package execution

type Request interface {
	Accept(visitor *RequestHandler) error
}

type RequestGenerator interface {
	NextRequest() Request
}
