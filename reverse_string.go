package reverse_string

func ReverseString(input string) (output string) {
	temp := []rune(input)
	n := len(temp)
	for i := 0; i < n/2; i++ {
		temp[i], temp[n-1-i] = temp[n-1-i], temp[i]
	}
	output = string(temp)
	return output
}
