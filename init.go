package console

func Init(short bool) {
	if short {
		logfunc = shortlog
		errorfunc = shorterror
	} else {
		logfunc = longlog
		errorfunc = longerror
	}
}
