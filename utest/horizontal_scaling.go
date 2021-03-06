// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-13 01:54:23 935F4E                  [utest/horizontal_scaling.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_HorizontalScaling_ is the unit test for PDF.HorizontalScaling()
func Test_PDF_HorizontalScaling_(t *testing.T) {
	fmt.Println("Test PDF.HorizontalScaling()")
	//
	// Horizontal Scaling of new PDF must be 100
	func() {
		var doc pdf.PDF
		TEqual(t, doc.HorizontalScaling(), 100)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.HorizontalScaling(), 100)
	}()
	//
	// SetHorizontalScaling() has effect on the property?
	func() {
		var doc pdf.PDF
		doc.SetHorizontalScaling(149)
		TEqual(t, doc.HorizontalScaling(), 149)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetHorizontalScaling(149)
		TEqual(t, doc.HorizontalScaling(), 149)
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		var doc = pdf.NewPDF("A4-L")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Times-Bold", 20).
			SetXY(1, 1).
			DrawText("Horizontal Scaling Property")

		for i, hscaling := range []int{50, 100, 150, 200, 250} {
			var y = 2.5 + float64(i)*2.5
			doc.SetXY(1, y).
				SetFont("Helvetica", 10).
				SetHorizontalScaling(100).
				DrawText(fmt.Sprintf("Horizontal Scaling = %d", hscaling)).
				SetXY(1, y+0.7).
				SetHorizontalScaling(uint16(hscaling)).
				SetFontSize(20).
				DrawText("Five hexing wizard bots jump quickly")
		}

		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 841 595]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R>> >> >>
		endobj
		4 0 obj <</Length 899>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 566 Td (Horizontal Scaling Property) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 524 Td (Horizontal Scaling = 50) Tj ET
		BT /FNT2 20 Tf ET
		BT 50 Tz ET
		BT 28 504 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 453 Td (Horizontal Scaling = 100) Tj ET
		BT /FNT2 20 Tf ET
		BT 28 433 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 382 Td (Horizontal Scaling = 150) Tj ET
		BT /FNT2 20 Tf ET
		BT 150 Tz ET
		BT 28 362 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 311 Td (Horizontal Scaling = 200) Tj ET
		BT /FNT2 20 Tf ET
		BT 200 Tz ET
		BT 28 291 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 240 Td (Horizontal Scaling = 250) Tj ET
		BT /FNT2 20 Tf ET
		BT 250 Tz ET
		BT 28 221 Td (Five hexing wizard bots jump quickly) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		6 0 obj <</Type/Font/Subtype/Type1/Name/FNT2
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 7
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000241 00000 n
		0000001191 00000 n
		0000001293 00000 n
		trailer
		<</Size 7/Root 1 0 R>>
		startxref
		1394
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
		doc.SaveFile("~~test_horizontal_scaling.pdf")
	}()

} //                                                 Test_PDF_HorizontalScaling_

//end
