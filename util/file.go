package util

import (
	"fmt"
	"os"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true

		}
		return false
	}
	return true
}

func FileExistWithExtensionName(path string, extensionName ...string) (string, bool) {
	for _, extension := range extensionName {
		fp := fmt.Sprintf("%s.%s", path, extension)
		if FileExist(fp) {
			return fp, true
		}
	}
	return "", false
}
