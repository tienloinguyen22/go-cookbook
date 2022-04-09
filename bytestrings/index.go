package bytestrings

func RunBytestrings() {
	err := workWithBuffer()
	if err != nil {
		panic(err)
	}

	searchString()
	modifyString()
	myStrReader()
}