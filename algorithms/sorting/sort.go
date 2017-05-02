package sorting

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Changeable interface {
	Len() int
	Get(i int) interface{}
	Change(i int, v interface{})
}
