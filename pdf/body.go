package pdf

func (p *Pdf) generateBody() error {
	// Set font for body
	p.Pdf.SetFont("Arial", "", 12)

	// Date Reported
	p.Pdf.Cell(50, 10, "Date Identified:")
	p.Pdf.Cell(100, 10, p.Dialog.DateTimeReported) // Example data
	p.Pdf.Ln(8)

	// Date Identified
	p.Pdf.Cell(50, 10, "Date Reported:")
	p.Pdf.Cell(100, 10, p.Dialog.DateTimeEvent)
	p.Pdf.Ln(8)

	// Description (Multiline Text)
	p.Pdf.Cell(50, 10, "Description:")
	p.Pdf.Ln(6)
	p.Pdf.MultiCell(190, 8, p.Dialog.Description, "", "L", false)
	p.Pdf.Ln(10)

	// Description (Multiline Text)
	p.Pdf.Cell(50, 10, "Time Taken:")
	p.Pdf.Cell(100, 10, p.Dialog.TimeTaken)
	p.Pdf.Ln(10)

	// Employee Responsible
	p.Pdf.Cell(50, 10, "Employee Responsible:")
	p.Pdf.Cell(100, 10, p.Dialog.Name)
	p.Pdf.Ln(20) // Extra space before footer

	// Footer - "Pictures Next Page"
	p.Pdf.SetFont("Arial", "B", 12)
	p.Pdf.SetTextColor(255, 0, 0) // Red text
	p.Pdf.Cell(190, 10, "Pictures Next Page")
	p.Pdf.Ln(10)
	return nil
}
