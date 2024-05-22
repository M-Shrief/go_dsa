package sorting

type Dirction string

const (
	ASC  Dirction = "ASC"
	DESC Dirction = "DESC"
)

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
