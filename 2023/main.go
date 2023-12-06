package y2023

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	Day01(false)
	Day02(false)
	Day03(false)
	Day04(false)
	Day05(false)
	Day06(false)
}
