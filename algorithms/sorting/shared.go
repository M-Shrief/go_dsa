package sorting

type Direction string

const (
	ASC  Direction = "ASC"
	DESC Direction = "DESC"
)

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
