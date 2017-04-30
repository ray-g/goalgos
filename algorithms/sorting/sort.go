package sorting

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
