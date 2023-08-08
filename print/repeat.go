package print

func Repeat(str string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += str
	}
	return result
}
