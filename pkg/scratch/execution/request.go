package execution

type Request interface {
	Accept(handler *RequestHandler) error
}

type RequestGenerator interface {
	NextRequest() Request
}

type StaticRequestGenerator struct {
	// FIXME may be better implemented using channels
	requests []Request
	i        int
}

func (g *StaticRequestGenerator) NextRequest() Request {
	if g.i >= len(g.requests) {
		return nil
	}
	req := g.requests[g.i]
	g.i++
	return req
}

func NewStaticRequestGenerator(requests []Request) *StaticRequestGenerator {
	return &StaticRequestGenerator{requests: requests, i: 0}
}
