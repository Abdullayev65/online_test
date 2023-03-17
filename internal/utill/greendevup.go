package utill

func Map[I any, O any](in []I, mapper func(i I) O) []O {
	os := make([]O, 0, len(in))
	for _, i := range in {
		o := mapper(i)
		os = append(os, o)
	}
	return os
}
