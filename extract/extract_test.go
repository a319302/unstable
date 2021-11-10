package extract

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExtractTextFromPdf(t *testing.T) {
	f, err := os.Open("f:/wang.pdf")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromPdf(b)
	fmt.Println(text)
}

func TestExtractTextFromDocx(t *testing.T) {
	f, err := os.Open("f:/ji.docx")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromDocx(b)
	fmt.Println(text)
}

func TestExtractTextFromDoc(t *testing.T) {
	f, err := os.Open("f:/ji.doc")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromDoc(b)
	fmt.Println(text)
}

func TestExtractTextFromXlsx(t *testing.T) {
	f, err := os.Open("f:/q.xlsx")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromXLSX(b)
	fmt.Println(text)
}

func TestExtractTextFromXls(t *testing.T) {
	f, err := os.Open("f:/q.xls")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromXLS(b)
	fmt.Println(text)
}

func TestExtractTextFromZIP(t *testing.T) {
	f, err := os.Open("f:/f.zip")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	text, _ := ExtractTextFromZIP(b)
	fmt.Println(text)
}
