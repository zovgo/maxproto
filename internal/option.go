package internal

//func HasTrueOption(x []bool) bool {
//	return MustFirstOption(x)
//}

func MustFirstOptionOr[T any, S ~[]T](x S, or T) T {
	if len(x) > 0 {
		return x[0]
	}
	return or
}

//func MustFirstOption[T any, S ~[]T](x S) (zero T) {
//	return MustFirstOptionOr(x, zero)
//}
