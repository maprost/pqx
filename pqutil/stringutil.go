package pqutil

func Concate(a string, b string, separator string) string {
	if len(a) > 0 && len(b) > 0 {
		return a + separator + b
	}
	if len(a) > 0 {
		return a
	}
	return b
}
