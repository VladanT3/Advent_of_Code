package shared

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}
