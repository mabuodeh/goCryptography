package set2

import "github.com/goCryptography/set1/functions"

// DecryptCbc takes a file, and decrypts it using the string key, iv string, by aes-128-cbc
func DecryptCbc(byteData, byteKey, byteIv []byte) []byte {

	var finalByteMsg []byte

	// initialize xor variable to be iv
	xorVar := byteIv
	blockSize := len(byteIv)

	// loop over the byte data
	for i := 0; i < len(byteData); i += blockSize {
		// fmt.Println(i)
		// take an iv-sized block of data
		dataBlock := byteData[i : i+blockSize]
		// send the block to the aes ecb function with the key
		// fmt.Println(string(dataBlock))
		// fmt.Println(string(byteKey))
		var plaintext []byte
		plaintext = set1.DecryptEcb(dataBlock, byteKey)
		// take the decrypted text and xor it with the xor variable using the xor function
		plaintext = set1.XorByteArrays(plaintext, xorVar)
		// take the iv-sized block of data and assign it to the xor variable
		xorVar = dataBlock
		// take the resultant plaintext and append it to final decrypted message byte array
		finalByteMsg = append(finalByteMsg, plaintext...)

	} // endloop

	// return plaintext in bytes
	return finalByteMsg

}

// EncryptCbc takes a file, and decrypts its data using the string key and hex string iv, by aes-128-cbc
func EncryptCbc(byteData, byteKey, byteIv []byte) []byte {

	var finalCiphertext []byte

	// initialize xor variable to be iv
	xorVar := byteIv
	blockSize := len(byteIv)

	// loop over the byte data
	for i := 0; i < len(byteData); i += blockSize {
		// fmt.Println(i)
		// take an iv-sized block of data
		dataBlock := byteData[i : i+blockSize]
		// take the decrypted text and xor it with the xor variable using the xor function
		var ciphertext []byte
		ciphertext = set1.XorByteArrays(dataBlock, xorVar)
		// send the block to the aes ecb function with the key
		ciphertext = set1.EncryptEcb(ciphertext, byteKey)
		// take the iv-sized block of data and assign it to the xor variable
		xorVar = ciphertext
		// take the resultant plaintext and append it to final decrypted message byte array
		finalCiphertext = append(finalCiphertext, ciphertext...)

	} // endloop

	// return plaintext in bytes
	return finalCiphertext

}
