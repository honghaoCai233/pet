package exp

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"os"
	"testing"
)

func TestEPDF(t *testing.T) {
	err := license.SetMeteredKey("7a590c1a7ef8a21754b8dc5d25509d4b836008dd3d5b1c69ea94773344b7503b")
	if err != nil {
		t.Log(err)
	}
	f, err := os.Open("C:\\Users\\Tangerg\\Desktop\\我是驴友-舟山攻略.pdf")
	if err != nil {
		t.Log(err)
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		t.Log(err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(numPages)
	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < 2; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			t.Log(err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			t.Log(err)
		}
		_, err = ex.ExtractText()
		if err != nil {
			t.Log(err)
		}
		fonts, err := ex.ExtractFonts(nil)
		for _, font := range fonts.Fonts {
			t.Log(font.FontName)
			t.Log(string(font.FontData))
		}
	}
}
