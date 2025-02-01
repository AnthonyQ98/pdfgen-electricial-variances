package pdf

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

type Pdf struct {
	Pdf  *fpdf.Fpdf
	name string
}

func NewPdf(filename string) *Pdf {
	return &Pdf{Pdf: fpdf.New("P", "mm", "A4", ""), name: filename}
}

func (p *Pdf) save() error {
	err := p.Pdf.OutputFileAndClose(p.name)
	if err != nil {
		return fmt.Errorf("error saving %s header: %v", p.name, err)
	}
	return nil
}

func (p *Pdf) Generate() error {
	// Create the 1st page
	p.Pdf.AddPage()

	p.Pdf.SetFont("Arial", "B", 16)
	err := p.generateHeader()
	if err != nil {
		return fmt.Errorf("error generating header: %v", err)
	}

	err = p.generateBody()
	if err != nil {
		return fmt.Errorf("error generating body: %v", err)
	}

	// Create the 2nd page
	p.Pdf.AddPage()

	// Add images to the second page
	err = p.addImagesToPdf()
	if err != nil {
		return fmt.Errorf("error adding images: %v", err)
	}

	err = p.save()
	if err != nil {
		return fmt.Errorf("error saving %s: %v", p.name, err)
	}

	return nil
}
