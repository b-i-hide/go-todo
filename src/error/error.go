package error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
