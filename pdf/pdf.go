package pdf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pdfgen-electricial-variations/dialog"
	"time"

	"github.com/go-pdf/fpdf"
)

type Pdf struct {
	Pdf                   *fpdf.Fpdf
	name, ReferenceNumber string
	dialog.Dialog
}

func NewPdf(filename string) *Pdf {
	return &Pdf{Pdf: fpdf.New("P", "mm", "A4", ""), name: filename}
}

func (p *Pdf) save() error {
	if p.ReferenceNumber == "" {
		return fmt.Errorf("reference number is empty, generateFileName() must be called first")
	}

	// Get today's folder path
	now := time.Now()
	dateFolder := now.Format("2006-01-02")
	outputDir := filepath.Join("output", "pdf", dateFolder)
	filePath := filepath.Join(outputDir, p.ReferenceNumber+".pdf") // Append ".pdf"

	// Save the PDF
	err := p.Pdf.OutputFileAndClose(filePath)
	if err != nil {
		return fmt.Errorf("error saving %s.pdf: %v", p.ReferenceNumber, err)
	}

	fmt.Printf("Saved PDF: %s\n", filePath)
	return nil
}

// getNextFileNumber finds the next available number for the filename pattern
func getNextFileNumber(folderPath, dateSuffix string) (int, error) {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return 1, nil // Folder empty, start with 1
	}

	maxNum := 0
	for _, file := range files {
		name := file.Name()
		var num int
		_, err := fmt.Sscanf(name, "%d-EV-"+dateSuffix+".pdf", &num)
		if err == nil && num > maxNum {
			maxNum = num
		}
	}

	return maxNum + 1, nil
}

func (p *Pdf) generateFileName() (string, error) {
	// Get current date
	now := time.Now()
	dateFolder := now.Format("2006-01-02") // YYYY-MM-DD

	// Create output folder if it doesn't exist
	outputDir := filepath.Join("output", "pdf", dateFolder)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create output directory: %v", err)
	}

	// Find the next available file number
	nextNumber, err := getNextFileNumber(outputDir, dateFolder)
	if err != nil {
		return "", fmt.Errorf("error finding next file number: %v", err)
	}

	// Construct filename without ".pdf": "1-EV-YY_MM_DD"
	fileName := fmt.Sprintf("%d-EV-%s", nextNumber, dateFolder)
	return fileName, nil
}

func (p *Pdf) Generate() error {
	fileName, err := p.generateFileName()
	if err != nil {
		return fmt.Errorf("error generating file name: %v", err)
	}
	p.ReferenceNumber = fileName

	// Get the input for the PDF
	p.Dialog.OpenInputDialog()

	log.Printf("Generating PDF: %s", p.Dialog.Name)
	// Create the 1st page
	p.Pdf.AddPage()

	p.Pdf.SetFont("Arial", "B", 16)
	err = p.generateHeader()
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
