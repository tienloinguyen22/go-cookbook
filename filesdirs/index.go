package filesdirs

func RunFilesdirs() {
	if err := operate(); err != nil {
		panic(err)
	}
	if err := tryCaplitalizer(); err != nil {
		panic(err)
	}
}