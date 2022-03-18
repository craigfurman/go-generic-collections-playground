package collections

////////////////////////////////////////
// Pipes
////////////////////////////////////////

// Is code generation the only way to sort-of have variadic type parameters
// here?
func Pipe4[S, T, U, V any](s S, fS func(S) T, fT func(T) U, fU func(U) V) V {
	return fU(fT(fS(s)))
}

func PipeMap[S, T any](f func(S) T) func([]S) []T {
	return func(l []S) []T {
		var ret []T
		for _, e := range l {
			ret = append(ret, f(e))
		}
		return ret
	}
}

func PipeFilter[S any](pred func(S) bool) func([]S) []S {
	return func(l []S) []S {
		var ret []S
		for _, e := range l {
			if pred(e) {
				ret = append(ret, e)
			}
		}
		return ret
	}
}

func PipeReduce[S, T any](initial T, f func(T, S) T) func([]S) T {
	acc := initial
	return func(l []S) T {
		for _, e := range l {
			acc = f(acc, e)
		}
		return acc
	}
}
