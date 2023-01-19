package errors

import "fmt"

type FilesystemError Error

type filesystemError error

var (
	ErrFileNotFound = FilesystemError{
		Error: filesystemError(fmt.Errorf("file not found")),
		Code:  "FS01",
	}
	ErrInvalidFile = FilesystemError{
		Error: filesystemError(fmt.Errorf("invalid file")),
		Code:  "FS02",
	}
)
