package pdf

import (
	"fmt"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/sqweek/dialog"
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
	p.Pdf.AddPage()
	p.Pdf.SetFont("Arial", "B", 16)
	err := p.generateHeader()
	if err != nil {
		return err
	}

	err = p.generateBody()
	if err != nil {
		return err
	}

	// Open file selection dialog for images (select multiple)
	selectedFiles, err := dialog.File().Filter("Image files", "png", "jpg", "jpeg", "gif", "bmp").Title("Select Images").Load()
	if err != nil {
		fmt.Println("Error selecting files:", err)
		return err
	}

	// Split selected files into an array (if multiple)
	imagePaths := strings.Split(selectedFiles, "\n")

	p.Pdf.AddPage()
	err = p.addImages(imagePaths)
	if err != nil {
		return err
	}

	err = p.save()
	if err != nil {
		return err
	}

	return nil
}
