package generics

type List[T any] interface {
	Add(idx int, t T)
}
