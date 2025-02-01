package main

import (
	"log"
	"pdfgen-electricial-variations/pdf"
)

const pdfName = "electrical-variance1.pdf"

func main() {
	log.Println("PDF Generator for Electricial Variances")

	pdf := pdf.NewPdf(pdfName)
	err := pdf.Generate()
	if err != nil {
		panic(err)
	}

	log.Println("PDF generated successfully")

}
