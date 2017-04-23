package set1

// TransposeBlocks takes a byte array, and transposes it into blocks of keysize bytes.
func TransposeBlocks(b []byte, keysize int) [][]byte {
	ret := [][]byte{}

	// initialize block counter
	blockCounter := 0
	// initialize bytes per block counter
	bytesPerBlockCounter := 0
	// loop through the bytes
	for _, bytes := range b {
		// append byte to ret[blockCounter]
		ret[blockCounter] = append(ret[blockCounter], bytes)
		// increment bytesPerBlockCounter
		bytesPerBlockCounter++

		// if bytes per block counter >= keysize
		if bytesPerBlockCounter >= keysize {
			// reinitialize bytes per block counter
			bytesPerBlockCounter = 0
			// increment block counter
			blockCounter++
		} // endif
	} // endloop

	return ret
}
