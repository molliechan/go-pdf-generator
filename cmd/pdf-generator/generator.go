package main

import (
	"bytes"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/molliechan/go-pdf-generator/internal/template"
	"github.com/molliechan/go-pdf-generator/internal/user"
)

type PDFService struct{}

func NewPDFService() *PDFService {
	return &PDFService{}
}

func (p PDFService) generatePDF(templ string, data interface{}, pdfFileName string) ([]byte, error) {

	templBytes, err := template.ParseTemplate(templ, user.GetUser())
	if err != nil {
		return nil, err
	}

	// initalize a wkhtmltopdf generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(templBytes))

	// enable it if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to the generator
	pdfg.AddPage(page)

	// manipulate page attributes
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeB5)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// Create Pdf
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./" + pdfFileName)
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
