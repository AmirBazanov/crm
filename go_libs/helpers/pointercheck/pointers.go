package pointercheck

func DerefOrDefault[T any](ptr *T, def T) T {
	if ptr == nil {
		return def
	}
	return *ptr
}

func ToPtr[T any](v T) *T {
	return &v
}
