package console

func Init(short bool) {
	if short {
		log = shortlog
		error = shorterror
	} else {
		log = longlog
		error = longerror
	}
}
