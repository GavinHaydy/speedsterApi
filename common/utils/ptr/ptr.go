package ptr

// Value 返回指针的值，如果为 nil 则返回零值。
func Value[T any](v *T) T {
	var zero T
	if v == nil {
		return zero
	}
	return *v
}
