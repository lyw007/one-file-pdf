// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-13 01:54:23 00DFA4                           [utest/font_name.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_FontName_ is the unit test for
func Test_PDF_FontName_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) FontName() string
	//
	fmt.Println("Test PDF.FontName()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.FontName(), "Helvetica")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.FontName(), "Helvetica")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetFontName(name string) *PDF
	//
	fmt.Println("Test PDF.SetFontName()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetFontName("Courier").FontName(), "Courier")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetFontName("Courier").FontName(), "Courier")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(1, 1).
			SetFont("Helvetica", 10).
			SetFontName("TimesRoman").
			DrawText("Hello World!")

		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 95>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Hello World!) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Roman
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000373 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		476
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
	}()

} //                                                          Test_PDF_FontName_

//end
