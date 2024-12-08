package _data

func Ptr[T any](x T) *T {
	return &x
}
