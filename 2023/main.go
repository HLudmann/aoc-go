package y2023

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Day01(false)
	Day02(false)
}
