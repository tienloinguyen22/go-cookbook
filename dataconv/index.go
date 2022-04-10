package dataconv

func RunDataconv() {
	showConv()
	
	if err := myStrConv(); err != nil {
		panic(err)
	}

	myInterfaces()
}