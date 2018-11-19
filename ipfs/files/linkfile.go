package files

import (
	"io"
	"os"
	"strings"
)

// Symlink represents symbolic link
type Symlink struct {
	name   string
	path   string
	Target string
	stat   os.FileInfo

	reader io.Reader
}

// NewLinkFile returns an instance of Symlink
func NewLinkFile(name, path, target string, stat os.FileInfo) File {
	return &Symlink{
		name:   name,
		path:   path,
		Target: target,
		stat:   stat,
		reader: strings.NewReader(target),
	}
}

// IsDirectory returns true if the File is a directory (and therefore
// supports calling `NextFile`) and false if the File is a normal file
// (and therefor supports calling `Read` and `Close`)
func (lf *Symlink) IsDirectory() bool {
	return false
}

// NextFile returns the next child file available (if the File is a
// directory). It will return (nil, io.EOF) if no more files are
// available. If the file is a regular file (not a directory), NextFile
// will return a non-nil error.
func (lf *Symlink) NextFile() (File, error) {
	return nil, io.EOF
}

// FileName returns a filename associated with this file
func (lf *Symlink) FileName() string {
	return lf.name
}

// Close implements io.Closer
func (lf *Symlink) Close() error {
	return nil
}

// FullPath returns the full path used when adding with this file
func (lf *Symlink) FullPath() string {
	return lf.path
}

// Read implements io.Reader
func (lf *Symlink) Read(b []byte) (int, error) {
	return lf.reader.Read(b)
}
