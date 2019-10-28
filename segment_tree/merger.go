package segment_tree

type Merger interface {
	merge(interface{}, interface{}) interface{}
}
