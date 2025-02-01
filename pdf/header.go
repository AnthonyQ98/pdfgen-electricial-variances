package pdf

import "time"

func (p *Pdf) generateHeader() error {
	// Logo Properties
	logoPath := "media/logo.png" // Ensure this path is correct
	logoX := 10.0                // X Position (Left Margin)
	logoY := 10.0                // Y Position (Top Margin)
	logoWidth := 30.0            // Logo Width
	logoHeight := 30.0           // Logo Height

	// Add Logo
	p.Pdf.Image(logoPath, logoX, logoY, logoWidth, logoHeight, false, "", 0, "")

	// Move cursor below the logo for text
	textX := logoX + logoWidth + 10 // Move right of logo
	textY := logoY + 5              // Align with logo

	p.Pdf.SetXY(textX, textY)

	// Company Name
	p.Pdf.SetFont("Arial", "B", 16)
	p.Pdf.Cell(100, 10, "My Company Limited")
	p.Pdf.Ln(8)

	p.Pdf.SetFont("Arial", "", 12)
	p.Pdf.SetX(textX)
	p.Pdf.Cell(100, 10, "Email: info@company.ie")
	p.Pdf.Ln(8)

	// Document title
	p.Pdf.SetFont("Arial", "", 12)
	p.Pdf.SetX(textX)
	p.Pdf.Cell(100, 10, "Electrical Variance Report")
	p.Pdf.Ln(8)

	// Move cursor to avoid overlap with logo
	p.Pdf.Ln(15)

	// Project & Client Information
	p.Pdf.SetFont("Arial", "", 12)
	p.Pdf.Cell(190, 10, "Project: New Electrical Installation - Site A")
	p.Pdf.Ln(6)
	p.Pdf.Cell(190, 10, "Client: John Doe - ABC Constructions")
	p.Pdf.Ln(6)

	// Date & Reference Number
	currentTime := time.Now().Format("January 2, 2006")
	p.Pdf.Cell(190, 10, "Date: "+currentTime)
	p.Pdf.Ln(6)
	p.Pdf.Cell(190, 10, "Reference No: EVR-2024-001")
	p.Pdf.Ln(10)

	// Add a Divider Line
	p.Pdf.SetLineWidth(0.5)
	p.Pdf.Line(10, p.Pdf.GetY(), 200, p.Pdf.GetY())
	p.Pdf.Ln(10)
	return nil
}
