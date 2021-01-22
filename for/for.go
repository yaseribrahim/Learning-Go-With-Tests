package iteration

const repeatCount = 5

//Repeat takes a character and returns a string of repeated characters
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
