package pdf

import (
	"os/exec"
)

func CreateCover(pdfPath string) (coverPath string, err error) {
	_, err = exec.Command("convert", pdfPath+"[0]", pdfPath+".jpg").CombinedOutput()
	if err != nil {
		return coverPath, err
	}
	return pdfPath + ".jpg", err
}
