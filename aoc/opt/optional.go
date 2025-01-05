package opt

type Some[T any] struct {
	t *T
}

func None[T any]() Some[T] {
	return Some[T]{}
}

func From[T any](t T) Some[T] {
	return Some[T]{t: &t}
}

func OrNone[T any](t *T) Some[T] {
	if t == nil {
		return None[T]()
	}

	return From[T](*t)
}

func (v Some[T]) Or(o T) T {
	if v.t == nil {
		return o
	}
	return *v.t
}

func (v *Some[T]) Set(t T) {
	v.t = &t
}

func (v Some[T]) Ok() bool {
	return v.t != nil
}

func (v Some[T]) IfOk(fn func(t T)) {
	if v.t != nil {
		fn(*v.t)
	}
}

func (v Some[T]) Is(fn func(t T) bool) bool {
	return v.Ok() && fn(*v.t)
}

func (v Some[T]) Get() T {
	if v.t == nil {
		panic("Some.Get on empty")
	}

	return *v.t
}

func Map[In any, Out any](opt Some[In], mapFn func(in In) Out) Some[Out] {
	if !opt.Ok() {
		return None[Out]()
	}
	out := mapFn(*opt.t)
	return From[Out](out)
}
