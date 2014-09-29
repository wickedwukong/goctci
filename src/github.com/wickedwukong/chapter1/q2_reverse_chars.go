package chapter1

func Reverse(valuePointer *string) {
	reversed := make([]byte, 0)

	for i := len(*valuePointer) - 1; i >= 0; i-- {
		reversed = append(reversed, (*valuePointer)[i])
	}
	*valuePointer = string(reversed)
}