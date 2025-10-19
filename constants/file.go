package constants

const (
	MultipartFormFileName    = "file"
	MultipartFormContextName = "context"

	DefaultMaxImageHeight   = 300
	FirebaseStorageProvider = "firebase"
	LocalStorageProvider    = "local"
	UrlStorageProvider      = "url"
)

func JpegMimeType() []string {
	return []string{
		// jpg
		"image/jpeg",
		"image/pjpeg",
		"image/jpg",
		"image/x-citrix-jpeg",
		"image/x-citrix-jpg",
	}
}

func ImageMimeType() []string {
	mimeTypes := []string{
		// png
		"image/png",
		"image/x-png",
		"application/png",
	}
	mimeTypes = append(mimeTypes, JpegMimeType()...)

	return mimeTypes
}

func ValidFileMimeType() []string {
	mimeTypes := []string{
		// zip
		"application/zip",
		"application/x-zip-compressed",
		"multipart/x-zip",
		"application/x-zip",
		"application/x-zip-compressed",
		// pdf
		"application/pdf",
		"application/x-pdf",
		"application/x-bzpdf",
		"application/x-gzpdf",
		// ppt
		"application/vnd.ms-powerpoint",
		"application/mspowerpoint",
		"application/powerpoint",
		"application/x-mspowerpoint",
		"application/x-ms-powerpoint",
		// pptx
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"application/x-pptx",
		"application/vnd.ms-powerpoint.presentation.macroEnabled.12",
		// doc
		"application/msword",
		"application/doc",
		"application/msword",
		"application/x-msword",
		"application/vnd.msword",
		// docx
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/x-docx",
		"application/vnd.ms-word.document.macroEnabled.12",
		// xls
		"application/x-excel",
		"application/x-msexcel",
		"application/vnd.ms-office",
		"application/excel",
		// xlsx
		"application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-excel.sheet.macroEnabled.12",
		"application/vnd.ms-excel.sheet.binary.macroEnabled.12",
		"application/xlsx",
		// csv
		"text/csv",
		"application/csv",
		"application/x-csv",
		"text/x-csv",
		"application/vnd.ms-excel",
	}

	mimeTypes = append(mimeTypes, ImageMimeType()...)

	return mimeTypes
}
