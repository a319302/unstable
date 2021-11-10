package extract

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"time"

	"github.com/a319302/unstable/fileconversion"
	"github.com/ledongthuc/pdf"
	"github.com/lu4p/cat"
)

func ExtractTextFromPdf(content []byte) (string, error) {
	r := bytes.NewReader(content)
	//readAt := NewUnbufferedReaderAt(r)
	readAt := NewBufReaderAt(r, len(content))

	reader, err := pdf.NewReader(readAt, int64(len(content)))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	b, err := reader.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ExtractTextFromDocx(content []byte) (string, error) {
	return cat.FromBytes(content)
}

func ExtractTextFromDoc(content []byte) (string, error) {
	r := bytes.NewReader(content)

	reader, err := fileconversion.DOC2Text(r)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ExtractTextFromXLSX(content []byte) (string, error) {
	r := bytes.NewReader(content)
	readAt := NewBufReaderAt(r, len(content))
	buffer := new(bytes.Buffer)
	_, err := fileconversion.XLSX2Text(readAt, int64(len(content)), buffer, 10*1024*1024, 10000)
	if err != nil {
		return "", err
	}

	return string(buffer.String()), nil
}

func ExtractTextFromXLS(content []byte) (string, error) {
	r := bytes.NewReader(content)
	buffer := new(bytes.Buffer)
	_, err := fileconversion.XLS2Text(r, buffer, int64(len(content)))
	if err != nil {
		return "", err
	}

	return string(buffer.String()), nil
}

// func ExtractTextFromZIP(content []byte) (string, error) {
// 	r := bytes.NewReader(content)
// 	buffer := new(bytes.Buffer)

// 	zipReader, err := zip.NewReader(r, int64(len(content)))
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, file := range zipReader.File {
// 		if file.FileInfo().IsDir() {
// 			continue
// 		}
// 		rc, err := file.Open()
// 		if err != nil {
// 			return "", err
// 		}

// 		mimeType := MimeTypeByExtension(file.Name)

// 		b, err := ioutil.ReadAll(rc)
// 		if err != nil {
// 			return "", err
// 		}

// 		text, err := ExtractText(b, mimeType)
// 		if err == nil {
// 			buffer.WriteString(text)
// 		}

// 		rc.Close()
// 	}

// 	return string(buffer.String()), nil
// }

func ExtractTextFromZIP(content []byte) (string, error) {
	buffer := new(bytes.Buffer)
	fileconversion.ContainerExtractFiles(content, func(name string, size int64, date time.Time, data []byte) {
		mimeType := MimeTypeByExtension(name)
		text, err := ExtractText(data, mimeType)
		if err == nil {
			buffer.WriteString(text)
		}
	})
	return string(buffer.String()), nil
}

func MimeTypeByExtension(filename string) string {
	switch strings.ToLower(path.Ext(filename)) {
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".odt":
		return "application/vnd.oasis.opendocument.text"
	case ".pages":
		return "application/vnd.apple.pages"
	case ".pdf":
		return "application/pdf"
	case ".pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case ".rtf":
		return "application/rtf"
	case ".xml":
		return "text/xml"
	case ".xhtml", ".html", ".htm":
		return "text/html"
	case ".jpg", ".jpeg", ".jpe", ".jfif", ".jfif-tbnl":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".tif":
		return "image/tif"
	case ".tiff":
		return "image/tiff"
	case ".txt":
		return "text/plain"
	}
	return "unknown"
}

// Convert a file to plain text.
func ExtractText(content []byte, mimeType string) (string, error) {

	switch mimeType {
	case "application/msword", "application/vnd.ms-word":
		return ExtractTextFromDoc(content)

	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return ExtractTextFromDocx(content)

	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return "", fmt.Errorf("ExtractTextFromPPTX not implemented")

	case "application/vnd.oasis.opendocument.text":
		return "", fmt.Errorf("ExtractTextFromODT not implemented")

	case "application/vnd.apple.pages", "application/x-iwork-pages-sffpages":
		return "", fmt.Errorf("ExtractTextFromPages not implemented")

	case "application/pdf":
		return ExtractTextFromPdf(content)

	case "application/rtf", "application/x-rtf", "text/rtf", "text/richtext":
		return "", fmt.Errorf("ExtractTextFromRTF not implemented")

	case "text/html":
		return string(content), nil

	case "text/xml", "application/xml":
		return string(content), nil

	case "image/jpeg", "image/png", "image/tif", "image/tiff":
		return "", fmt.Errorf("ExtractTextFromImage not implemented")

	case "text/plain":
		return string(content), nil

	case "application/vnd.ms-excel":
		return ExtractTextFromXLS(content)
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return ExtractTextFromXLSX(content)
	case "application/zip", "application/vnd.rar", "application/x-tar", "application/x-7z-compressed":
		return ExtractTextFromZIP(content)
	}

	return "", fmt.Errorf("ExtractText not implemented for mimeType %s", mimeType)
}
