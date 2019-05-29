package interpreter

var hex_ascii = []byte{ 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46}

/**
 * Convierte de informacion binaria a ASCII hex.
 */
func Interpret(data []byte, offset int) []byte {
	b:=make([]byte, len(data) * 2)

	for i:=0; i<len(data); i++ {
		b[offset + i * 2] = hex_ascii[ (data[i] & 0XF0) >> 4]
		b[offset + i * 2 + 1] = hex_ascii[data[i] & 0x0F]
	}
	return b
}