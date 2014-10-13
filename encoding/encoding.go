package encoding

type ToUnicode struct {
	convTable [256]uint16
}

func (to *ToUnicode) Convert(input string) (output string) {
	output = ""
	for _, char := range []byte(input) {
		output += string(to.convTable[char])
	}
	return output
}

func New(convTable [256]uint16) *ToUnicode {
	to := new(ToUnicode)
	to.convTable = convTable
	return to
}
