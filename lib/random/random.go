package random

import "math/rand"

func NewRandomString(length int) string {
	var letters []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"} //scalable
	var res string = ""                                                                         // result string
	for i := 0; i < length; i++ {                                                               //for each lenght
		res += letters[rand.Intn(len(letters))] //append random letter to result string
	}
	return res
}
