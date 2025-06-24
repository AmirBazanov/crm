package _map

func Map[TSource, TTarget any](source []TSource, trans func(TSource) TTarget) []TTarget {
	target := make([]TTarget, 0, len(source))
	for _, s := range source {
		target = append(target, trans(s))
	}
	return target
}
