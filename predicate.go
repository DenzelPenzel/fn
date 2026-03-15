package fn

// Pred is a predicate function over type A
type Pred[A any] = func(A) bool

// PredAnd returns a predicate that is true when all given predicates are true
func PredAnd[A any](preds ...Pred[A]) Pred[A] {
	return func(a A) bool {
		for _, p := range preds {
			if !p(a) {
				return false
			}
		}
		return true
	}
}

// PredOr returns a predicate that is true when any given predicate is true
func PredOr[A any](preds ...Pred[A]) Pred[A] {
	return func(a A) bool {
		for _, p := range preds {
			if p(a) {
				return true
			}
		}
		return false
	}
}

// PredNot negates a predicate
func PredNot[A any](p Pred[A]) Pred[A] {
	return func(a A) bool {
		return !p(a)
	}
}
