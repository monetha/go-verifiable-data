package files

import (
	"errors"
	"io"
	"os"
)

var (
	// ErrNotDirectory returned by NextFile() when it's called not for directory
	ErrNotDirectory = errors.New("couldn't call NextFile(), this isn't a directory")
	// ErrNotReader returned when Read() or Close() called on a directory
	ErrNotReader = errors.New("this file is a directory, can't use Reader functions")
)

// File is an interface that provides functionality for handling
// files/directories as values that can be supplied to commands. For
// directories, child files are accessed serially by calling `NextFile()`.
type File interface {
	// Files implement ReadCloser, but can only be read from or closed if
	// they are not directories
	io.ReadCloser

	// FileName returns a filename associated with this file
	FileName() string

	// FullPath returns the full path used when adding with this file
	FullPath() string

	// IsDirectory returns true if the File is a directory (and therefore
	// supports calling `NextFile`) and false if the File is a normal file
	// (and therefor supports calling `Read` and `Close`)
	IsDirectory() bool

	// NextFile returns the next child file available (if the File is a
	// directory). It will return (nil, io.EOF) if no more files are
	// available. If the file is a regular file (not a directory), NextFile
	// will return a non-nil error.
	NextFile() (File, error)
}

// StatFile provides os.FileInfo for file
type StatFile interface {
	File

	Stat() os.FileInfo
}

// PeekFile returns file length and provides Peek method
type PeekFile interface {
	SizeFile

	Peek(n int) File
	Length() int
}

// SizeFile is a file that has size
type SizeFile interface {
	File

	Size() (int64, error)
}

// FileInfo provides os.FileInfo and absolute path of file
type FileInfo interface {
	AbsPath() string
	Stat() os.FileInfo
}
