package functions

func doSomething() (string, string) {
	result := "result"
	err := "err"
	return result, err
}

func plus(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func main() {
	_, err := doSomething()
	if err != nil {
		// エラー処理
	}
}
