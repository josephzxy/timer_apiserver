package util

// PanicIfErr panics only if the given error is non-nil
func PanicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}
