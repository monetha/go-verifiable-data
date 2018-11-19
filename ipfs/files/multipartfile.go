package files

import (
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/url"
)

const (
	multipartFormdataType = "multipart/form-data"
	multipartMixedType    = "multipart/mixed"

	applicationDirectory = "application/x-directory"
	applicationSymlink   = "application/symlink"
	applicationFile      = "application/octet-stream"

	contentTypeHeader = "Content-Type"
)

// MultipartFile implements File, and is created from a `multipart.Part`.
// It can be either a directory or file (checked by calling `IsDirectory()`).
type MultipartFile struct {
	File

	Part      *multipart.Part
	Reader    *multipart.Reader
	Mediatype string
}

// NewFileFromPart transforms part of multipart to File
func NewFileFromPart(part *multipart.Part) (File, error) {
	f := &MultipartFile{
		Part: part,
	}

	contentType := part.Header.Get(contentTypeHeader)
	switch contentType {
	case applicationSymlink:
		out, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, err
		}

		return &Symlink{
			Target: string(out),
			name:   f.FileName(),
		}, nil
	case "": // default to application/octet-stream
		fallthrough
	case applicationFile:
		return &ReaderFile{
			reader:   part,
			filename: f.FileName(),
			abspath:  part.Header.Get("abspath"),
			fullpath: f.FullPath(),
		}, nil
	}

	var err error
	f.Mediatype, _, err = mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// IsDirectory returns true if the File is a directory (and therefore
// supports calling `NextFile`) and false if the File is a normal file
// (and therefor supports calling `Read` and `Close`)
func (f *MultipartFile) IsDirectory() bool {
	return f.Mediatype == multipartFormdataType || f.Mediatype == applicationDirectory
}

// NextFile returns the next child file available (if the File is a
// directory). It will return (nil, io.EOF) if no more files are
// available. If the file is a regular file (not a directory), NextFile
// will return a non-nil error.
func (f *MultipartFile) NextFile() (File, error) {
	if !f.IsDirectory() {
		return nil, ErrNotDirectory
	}
	if f.Reader != nil {
		part, err := f.Reader.NextPart()
		if err != nil {
			return nil, err
		}

		return NewFileFromPart(part)
	}

	return nil, io.EOF
}

// FileName returns a filename associated with this file
func (f *MultipartFile) FileName() string {
	if f == nil || f.Part == nil {
		return ""
	}

	filename, err := url.QueryUnescape(f.Part.FileName())
	if err != nil {
		// if there is a unescape error, just treat the name as unescaped
		return f.Part.FileName()
	}
	return filename
}

// FullPath returns the full path used when adding with this file
func (f *MultipartFile) FullPath() string {
	return f.FileName()
}

func (f *MultipartFile) Read(p []byte) (int, error) {
	if f.IsDirectory() {
		return 0, ErrNotReader
	}
	return f.Part.Read(p)
}

// Close implements io.Closer
func (f *MultipartFile) Close() error {
	if f.IsDirectory() {
		return ErrNotReader
	}
	return f.Part.Close()
}
