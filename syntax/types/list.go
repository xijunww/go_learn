package types

type List interface {
	Add(ind int, val any) error
	Append(val any)
	Delete(idx int) (any, error)
}
