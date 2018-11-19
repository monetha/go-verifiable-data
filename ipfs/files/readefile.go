package files

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// ReaderFile is a implementation of File created from an `io.Reader`.
// ReaderFiles are never directories, and can be read from and closed.
type ReaderFile struct {
	filename string
	fullpath string
	abspath  string
	reader   io.ReadCloser
	stat     os.FileInfo
}

// NewReaderFile creates new instance of ReaderFile
func NewReaderFile(filename, path string, reader io.ReadCloser, stat os.FileInfo) *ReaderFile {
	return &ReaderFile{filename, path, path, reader, stat}
}

// NewReaderPathFile retrieves an absolute representation of path and creates new instance of ReaderFile
func NewReaderPathFile(filename, path string, reader io.ReadCloser, stat os.FileInfo) (*ReaderFile, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return &ReaderFile{filename, path, abspath, reader, stat}, nil
}

// IsDirectory returns true if the File is a directory (and therefore
// supports calling `NextFile`) and false if the File is a normal file
// (and therefor supports calling `Read` and `Close`)
func (f *ReaderFile) IsDirectory() bool {
	return false
}

// NextFile returns the next child file available (if the File is a
// directory). It will return (nil, io.EOF) if no more files are
// available. If the file is a regular file (not a directory), NextFile
// will return a non-nil error.
func (f *ReaderFile) NextFile() (File, error) {
	return nil, ErrNotDirectory
}

// FileName returns a filename associated with this file
func (f *ReaderFile) FileName() string {
	return f.filename
}

// FullPath returns the full path used when adding with this file
func (f *ReaderFile) FullPath() string {
	return f.fullpath
}

// AbsPath implements FileInfo
func (f *ReaderFile) AbsPath() string {
	return f.abspath
}

// Read implements io.Reader
func (f *ReaderFile) Read(p []byte) (int, error) {
	return f.reader.Read(p)
}

// Close implements io.Closer
func (f *ReaderFile) Close() error {
	return f.reader.Close()
}

// Stat implements FileInfo
func (f *ReaderFile) Stat() os.FileInfo {
	return f.stat
}

// Size implements SizeFile
func (f *ReaderFile) Size() (int64, error) {
	if f.stat == nil {
		return 0, errors.New("File size unknown")
	}
	return f.stat.Size(), nil
}
