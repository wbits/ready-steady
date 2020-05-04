package main

import "fmt"

func main() {
	cipherText := "CS OIT EUI WUI ZNS ROCNKFD"
	keyword := "GOLANG"

	decrypted := cipher(cipherText, keyword, false)
	encrypted := cipher(decrypted, keyword, true)

	fmt.Println(decrypted)
	fmt.Println(encrypted)
	fmt.Printf("AssertEncrypted: %v", cipherText == encrypted)
}

func cipher(txt, keyword string, encrypt bool) string {
	ciphered := ""
	keywordIndex := 0
	for i := 0; i < len(txt); i++ {
		character := txt[i]
		if character == ' ' {
			ciphered += string(character)
			continue
		}

		c := character - 'A'
		k := keyword[keywordIndex] - 'A'

		if encrypt {
			c = (c+k)%26 + 'A'
		} else {
			c = (c-k+26)%26 + 'A'
		}
		ciphered += string(c)
		keywordIndex++
		keywordIndex %= len(keyword)
	}

	return ciphered
}
