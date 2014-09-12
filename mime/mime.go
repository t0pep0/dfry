package mime

import (
	"os/exec"
	"strings"
)

func Identify(filePath string) (mimeType string, err error) {
	out, err := exec.Command("file", "--mime-type", filePath).CombinedOutput()
	if err != nil {
		return mimeType, err
	}
	mimeType = strings.Split(string(out), ": ")[1]
	return mimeType, err
}
