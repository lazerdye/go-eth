package maker

func AsciiToByte(ascii string) [32]byte {
	var retBytes [32]byte
	copy(retBytes[:], []byte(ascii))
	return retBytes
}
