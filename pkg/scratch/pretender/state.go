package pretender

type PodsSnapshot struct{}

type PodTraits struct{}

type nodeState struct {
	name        string
	cpuCapacity float64
	memCapacity float64
	pods        []*PodTraits
}

type State struct {
	nodes map[string]*nodeState
}

func (c State) GetSnapshot() PodsSnapshot {
	return PodsSnapshot{
		// TODO
	}
}

func NewState() State {
	return State{}
}
