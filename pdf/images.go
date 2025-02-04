package pdf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ncruces/zenity"
)

func (p *Pdf) addImagesToPdf() error {
	// Get the images for Page 2
	imagePaths, err := p.selectImages()
	if err != nil {
		log.Printf("Error selecting files: %v", err)
	}

	numImages := len(imagePaths)
	if numImages == 0 {
		return nil // No images selected, return
	}

	// Define page layout properties
	pageWidth := 210.0  // A4 width in mm
	pageHeight := 297.0 // A4 height in mm
	margin := 10.0      // Page margin in mm
	usableWidth := pageWidth - (2 * margin)
	usableHeight := pageHeight - (2 * margin)

	// Determine rows and columns based on the number of images
	var rows, cols int
	switch {
	case numImages == 1:
		rows, cols = 1, 1 // Single image fills most of the page
	case numImages == 2 || numImages == 3:
		rows, cols = 2, 2 // Split into 2 rows, 2 columns
	case numImages <= 6:
		rows, cols = 3, 2 // Max 6 images in 3 rows, 2 cols
	case numImages <= 10:
		rows, cols = 5, 2 // Max 10 images in 5 rows, 2 cols
	default:
		rows, cols = 6, 3 // More images use 6 rows, 3 cols
	}

	// Dynamically calculate image size
	imageWidth := (usableWidth - (float64(cols-1) * margin)) / float64(cols)
	imageHeight := (usableHeight - (float64(rows-1) * margin)) / float64(rows)

	// Start positions
	xPosition := margin
	yPosition := margin
	imagesPlaced := 0

	// Loop over images and add to the PDF
	for _, imgPath := range imagePaths {
		// Add image
		p.Pdf.Image(imgPath, xPosition, yPosition, imageWidth, imageHeight, false, "", 0, "")

		imagesPlaced++
		if imagesPlaced%cols == 0 {
			// Move to the next row
			xPosition = margin
			yPosition += imageHeight + margin
		} else {
			// Move to the next column
			xPosition += imageWidth + margin
		}

		// Add a new page if we exceed the bottom of the page
		if yPosition+imageHeight+margin > pageHeight {
			p.Pdf.AddPage()
			xPosition = margin
			yPosition = margin
			imagesPlaced = 0 // Reset counter for new page
		}
	}
	return nil
}

func (p *Pdf) selectImages() ([]string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	downloadsDir := filepath.Join(dir, "Downloads")
	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		// Create the downloads directory if it doesn't exist
		if err := os.Mkdir(downloadsDir, 0755); err != nil {
			return nil, err
		}
	}

	selectedFiles, err := zenity.SelectFileMultiple(
		zenity.Filename(downloadsDir),
		zenity.FileFilters{
			{Name: "Image files", Patterns: []string{"*.png", "*.gif", "*.svg", "*.ico", "*.jpg", "*.webp", "*.jpeg"}, CaseFold: true},
		})

	if err != nil {
		log.Printf("Error selecting files: %v", err)
	}
	return selectedFiles, nil
}
