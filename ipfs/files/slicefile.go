package files

import (
	"errors"
	"io"
)

// SliceFile implements File, and provides simple directory handling.
// It contains children files, and is created from a `[]File`.
// SliceFiles are always directories, and can't be read from or closed.
type SliceFile struct {
	filename string
	path     string
	files    []File
	n        int
}

// NewSliceFile returns an instance of SliceFile
func NewSliceFile(filename, path string, files []File) *SliceFile {
	return &SliceFile{filename, path, files, 0}
}

// IsDirectory returns true if the File is a directory (and therefore
// supports calling `NextFile`) and false if the File is a normal file
// (and therefor supports calling `Read` and `Close`)
func (f *SliceFile) IsDirectory() bool {
	return true
}

// NextFile returns the next child file available (if the File is a
// directory). It will return (nil, io.EOF) if no more files are
// available. If the file is a regular file (not a directory), NextFile
// will return a non-nil error.
func (f *SliceFile) NextFile() (File, error) {
	if f.n >= len(f.files) {
		return nil, io.EOF
	}
	file := f.files[f.n]
	f.n++
	return file, nil
}

// FileName returns a filename associated with this file
func (f *SliceFile) FileName() string {
	return f.filename
}

// FullPath returns the full path used when adding with this file
func (f *SliceFile) FullPath() string {
	return f.path
}

func (f *SliceFile) Read(p []byte) (int, error) {
	return 0, io.EOF
}

// Close implements io.Closer
func (f *SliceFile) Close() error {
	return ErrNotReader
}

// Peek implements PeekFile
func (f *SliceFile) Peek(n int) File {
	return f.files[n]
}

// Length implements PeekFile
func (f *SliceFile) Length() int {
	return len(f.files)
}

// Size implements PeekFile
func (f *SliceFile) Size() (int64, error) {
	var size int64

	for _, file := range f.files {
		sizeFile, ok := file.(SizeFile)
		if !ok {
			return 0, errors.New("could not get size of child file")
		}

		s, err := sizeFile.Size()
		if err != nil {
			return 0, err
		}
		size += s
	}

	return size, nil
}
