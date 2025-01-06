package ezgo

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func TailFile(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}
	fileSize := fileInfo.Size()
	if fileSize == 0 {
		return "", fmt.Errorf("file is empty")
	}

	// Define buffer size for chunks
	const bufferSize = 4096
	var remaining []byte

	// Start reading from the end
	for offset := fileSize; offset > 0; {
		// Calculate the size of the chunk to read
		readSize := bufferSize
		if int64(readSize) > offset {
			readSize = int(offset)
		}
		offset -= int64(readSize)

		// Seek and read the chunk
		buf := make([]byte, readSize)
		if _, err := file.ReadAt(buf, offset); err != nil {
			return "", fmt.Errorf("failed to read file: %w", err)
		}

		// Prepend the chunk to the remaining buffer
		remaining = append(buf, remaining...)

		// Split the buffer into lines
		lines := bytes.Split(remaining, []byte("\n"))

		// Process lines in reverse
		for i := len(lines) - 1; i >= 0; i-- {
			line := strings.TrimSpace(string(lines[i]))
			if line != "" {
				return line, nil
			}
		}

		// Keep only the first line fragment for the next iteration
		remaining = lines[0]
	}

	return "", fmt.Errorf("file contains no non-empty lines")
}
