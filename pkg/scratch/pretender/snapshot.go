package pretender

type NodeSnapshot struct {
	NodeCount         int
	MilliCPURequested int64
	MemoryRequested   int64
	CPULoad           float64
}

func newEmptyNodeSnapshot(state *nodeState) *NodeSnapshot {
	return &NodeSnapshot{
		NodeCount: 0,
	}
}

// FIXME is functional style ok?
func makeNodeSnapshot(state *nodeState, time int64) *NodeSnapshot {
	snapshot := newEmptyNodeSnapshot(state)
	for _, traits := range state.pods {
		for _, trait := range traits {
			trait.Apply(snapshot, time)
		}
	}
	return snapshot
}

type StateSnapshot map[string]*NodeSnapshot
