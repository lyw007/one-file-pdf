// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 3FF045                          [utest/doc_author.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DocAuthor is the unit test for
func DocAuthor(t *testing.T) {

	// -------------------------------------------------------------------------
	// (pdf *PDF) DocAuthor() string
	//
	fmt.Println("utest.DocAuthor")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocAuthor(), "")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocAuthor(), "")
	}()

	// -------------------------------------------------------------------------
	// (pdf *PDF) SetDocAuthor(s string) *PDF
	//
	fmt.Println("utest.SetDocAuthor")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	//
	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocAuthor("'Author' metadata entry")
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>>stream
		endstream
		5 0 obj<</Type/Info/Author ('Author' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000182 00000 n
		0000000221 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		283
		%%EOF
        `
		pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
	}()

} //                                                                   DocAuthor

//end
