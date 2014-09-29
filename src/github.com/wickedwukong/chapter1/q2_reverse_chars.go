package chapter1

func Reverse(valuePointer *string) {
	var reversed []byte

	for i := len(*valuePointer) - 1; i >= 0; i-- {
		reversed = append(reversed, (*valuePointer)[i])
	}
	*valuePointer = string(reversed)
}