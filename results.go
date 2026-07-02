package t

type Result[R any] struct {
	Result R
	Error error
}