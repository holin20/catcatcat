package ezgo

import (
	"bytes"
	"os"
)

func TailFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	const readSize = 1024
	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	var lastLine []byte
	buf := make([]byte, readSize)
	fileSize := stat.Size()
	offset := fileSize

	for {
		if offset == 0 {
			break
		}

		if offset < readSize {
			buf = make([]byte, offset)
			offset = 0
		} else {
			offset -= readSize
		}

		_, err := file.ReadAt(buf, offset)
		if err != nil {
			return "", err
		}

		if idx := bytes.LastIndexByte(buf, '\n'); idx != -1 {
			lastLine = buf[idx+1:]
			break
		}
	}

	return string(lastLine), nil
}
