package pdf

func (p *Pdf) addImages(imagePaths []string) error {
	// Image properties
	imageWidth := 60.0   // Width of each image
	imageHeight := 60.0  // Height of each image
	margin := 10.0       // Margin between images
	xPosition := 10.0    // Start x position
	yPosition := 10.0    // Start y position
	maxImagesPerRow := 2 // Maximum number of images per row

	// Loop over the image paths and add them to the page
	for i, imgPath := range imagePaths {
		// Add image
		p.Pdf.Image(imgPath, xPosition, yPosition, imageWidth, imageHeight, false, "", 0, "")

		// Move to next position
		if (i+1)%maxImagesPerRow == 0 {
			// If we reach max images per row, move to the next line
			xPosition = 10.0
			yPosition += imageHeight + margin
		} else {
			// Otherwise, move horizontally
			xPosition += imageWidth + margin
		}

		// Add a new page if we exceed the bottom of the page
		if yPosition+imageHeight+margin > 270 {
			p.Pdf.AddPage()
			xPosition = 10.0
			yPosition = 10.0
		}
	}
	return nil
}
