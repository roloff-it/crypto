// decrypt package consists of all the algorithms
package decrypt

// decrypts by reducing the acii code by 3 for each character
func Nimbus(str string) string {
	decryptedStr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
