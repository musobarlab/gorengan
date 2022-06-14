package shared

// Output struct
type Output[R any] struct {
	Result R
	Err    error
}
