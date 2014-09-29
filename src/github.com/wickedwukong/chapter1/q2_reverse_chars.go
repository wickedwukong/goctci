package chapter1

func Reverse(valuePointer *string) {
	value := *valuePointer
	reversed := make([]byte, 0)

	for i := len(value) - 1; i >= 0; i-- {
		reversed = append(reversed, value[i])
	}
	*valuePointer = string(reversed)
}