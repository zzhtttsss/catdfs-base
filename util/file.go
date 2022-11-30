package util

import "os"

// createFile creates a RDWR file whose path is #{des} and size is #{fileSize}
func createFile(des string, fileSize int64, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return nil, err
	}
	err = file.Truncate(fileSize)
	if err != nil {
		return nil, err
	}
	return file, nil
}
