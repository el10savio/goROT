package rot

func Encrypt(sentence string) string {
	inputBytes := stringToBytes(sentence)
	encryptedBytes := ROT13(inputBytes)
	encryptedSentence := bytesToString(encryptedBytes)
	return encryptedSentence
}

func Decrypt(sentence string) string {
	return Encrypt(sentence)
}

func ROT13(inputBytes []byte) []byte {
	outputBytes := make([]byte, len(inputBytes))

	for index, char := range inputBytes {
		// Uppercase A-Z
		if char > 64 && char < 91 {
			outputBytes[index] = ((char - 65 + 13) % 26) + 65
			continue
		}

		// Lowercase a-z
		if char > 96 && char < 123 {
			outputBytes[index] = ((char - 97 + 13) % 26) + 97
			continue
		}

		outputBytes[index] = char
	}

	return outputBytes
}

func stringToBytes(sentence string) []byte {
	return []byte(sentence)
}

func bytesToString(bytes []byte) string {
	return string(bytes)
}
