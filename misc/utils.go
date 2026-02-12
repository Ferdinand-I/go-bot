package misc

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
