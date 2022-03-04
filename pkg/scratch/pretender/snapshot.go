package pretender

type NodeSnapshot struct {
	NodeCount int
}

func newEmptyNodeSnapshot(state *nodeState) *NodeSnapshot {
	return &NodeSnapshot{
		NodeCount: 0,
	}
}

func makeNodeSnapshot(state *nodeState) *NodeSnapshot {
	snapshot := newEmptyNodeSnapshot(state)
	for _, traits := range state.pods {
		for _, trait := range traits {
			trait.Apply(snapshot)
		}
	}
	return snapshot
}

type StateSnapshot map[string]*NodeSnapshot
