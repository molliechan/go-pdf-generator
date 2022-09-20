package main

import (
	"log"

	"github.com/molliechan/go-pdf-generator/internal/user"
)

func main() {
	pdfService := NewPDFService()

	_, err := pdfService.generatePDF(
		"../../template/sample.gohtml",
		user.GetUser(),
		"sample.pdf",
	)

	if err != nil {
		log.Println(err)
	}
}
