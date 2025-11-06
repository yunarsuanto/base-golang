package objects

import "mime/multipart"

type FileUploadMultipartRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Context    string
}

type FileUploadBase64Request struct {
	File    string
	Context string
}

type FileUpload struct {
	Url      string
	Path     string
	MimeType string
	Size     int64
}

type FileGetUrlRequest struct {
	Path string
}

type FileGetUrl struct {
	Url string
}
