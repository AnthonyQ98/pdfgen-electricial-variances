package dialog

import (
	"errors"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	datepicker "github.com/sdassow/fyne-datepicker"
	"golang.org/x/exp/rand"
)

type Dialog struct {
	ClientName       string
	Name             string
	TimeTaken        string
	Description      string
	DateTimeReported string
	DateTimeEvent    string
	IncludeImgs      bool
}

func init() {
	rand.Seed(uint64(time.Now().UnixNano())) // Ensure randomness only once
}

func setEventDateTime() string {
	now := time.Now()
	randomHour := 18 + rand.Intn(3) // Random hour between 18, 19, or 20
	randomMinute := rand.Intn(60)   // Random minute between 0 and 59

	eventTime := time.Date(now.Year(), now.Month(), now.Day(), randomHour, randomMinute, 0, 0, now.Location())
	return eventTime.Format("2006/01/02 15:04")
}

func (d *Dialog) OpenInputDialog() error {
	a := app.New()
	w := a.NewWindow("Form Example")
	w.Resize(fyne.NewSize(300, 250))

	// Text Inputs
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter employee name...")
	nameEntry.Text = "Thomas Quinn"

	clientName := widget.NewEntry()
	clientName.SetPlaceHolder("Enter client name...")
	clientName.Text = "Rostek & Tritech Group"

	timeTaken := widget.NewEntry()
	timeTaken.SetPlaceHolder("Time taken to fix...")

	descriptionEntry := widget.NewMultiLineEntry()
	descriptionEntry.SetPlaceHolder("Enter description...")

	// Date & Time of Event
	eventDateTimeInput := widget.NewEntry()
	eventDateTimeInput.SetPlaceHolder("YYYY/MM/DD HH:MM")
	now := time.Now()

	eventDateTimeInput.Text = setEventDateTime()

	eventDateTimeInput.ActionItem = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		showDatePicker(w, eventDateTimeInput)
	})

	// Reported Time
	reportedTimeInput := widget.NewEntry()
	reportedTimeInput.SetPlaceHolder("YYYY/MM/DD HH:MM")
	fixedTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location()) // set the time to 9am
	reportedTimeInput.Text = fixedTime.Format("2006/01/02 15:04")

	reportedTimeInput.ActionItem = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		showDatePicker(w, reportedTimeInput)
	})
	// Checkbox (Checked by Default)
	imageCheckbox := widget.NewCheck("Do you want to include images?", nil)
	imageCheckbox.SetChecked(true)

	// Submit Button
	submitBtn := widget.NewButtonWithIcon("Submit", theme.ConfirmIcon(), func() {
		// Validate input fields
		if nameEntry.Text == "" {
			dialog.ShowError(errors.New("name cannot be empty"), w)
			return
		}
		if descriptionEntry.Text == "" {
			dialog.ShowError(errors.New("description cannot be empty"), w)
			return
		}
		if reportedTimeInput.Text == "" {
			dialog.ShowError(errors.New("date and time cannot be empty"), w)
			return
		}

		// Validate date format
		_, err := time.Parse("2006/01/02 15:04", reportedTimeInput.Text)
		if err != nil {
			dialog.ShowError(errors.New("invalid date format. Use YYYY/MM/DD HH:MM"), w)
			return
		}

		if eventDateTimeInput.Text == "" {
			dialog.ShowError(errors.New("date and time cannot be empty"), w)
			return
		}

		// Validate date format
		_, err = time.Parse("2006/01/02 15:04", eventDateTimeInput.Text)
		if err != nil {
			dialog.ShowError(errors.New("invalid date format. Use YYYY/MM/DD HH:MM"), w)
			return
		}

		// Store user inputs in struct
		d.ClientName = clientName.Text
		d.TimeTaken = timeTaken.Text
		d.Name = nameEntry.Text
		d.Description = descriptionEntry.Text
		d.DateTimeEvent = eventDateTimeInput.Text
		d.DateTimeReported = reportedTimeInput.Text
		d.IncludeImgs = imageCheckbox.Checked

		// Log successful input
		log.Println("Form submitted successfully:")
		log.Println("Name:", d.Name)
		log.Println("Description:", d.Description)
		log.Println("Event Time:", d.DateTimeEvent)
		log.Println("Reported Time:", d.DateTimeReported)
		log.Println("Include Images:", d.IncludeImgs)

		// Close window
		w.Close()
	})

	// Cancel Button
	cancelBtn := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
		w.Close()
		log.Printf("Exiting! you hit cancel.")
		os.Exit(0)
	})

	// Form Layout
	form := container.NewVBox(
		widget.NewLabel("Employee Name:"),
		nameEntry,
		widget.NewLabel("Client Name:"),
		clientName,
		widget.NewLabel("Description:"),
		descriptionEntry,
		widget.NewLabel("Logged Time (on-site):"),
		reportedTimeInput,
		widget.NewLabel("Reported Time (now):"),
		eventDateTimeInput,
		widget.NewLabel("Time Taken:"),
		timeTaken,
		imageCheckbox,
		container.NewHBox(submitBtn, cancelBtn),
	)

	// Show Window
	w.SetContent(form)
	w.ShowAndRun()

	return nil

}

// showDatePicker opens a date-time picker and updates the given input field
func showDatePicker(w fyne.Window, input *widget.Entry) {
	// Parse existing date or use current time
	// Parse existing date or use current time
	when := time.Now()
	if input.Text != "" {
		t, err := time.Parse("2006/01/02 15:04", input.Text)
		if err == nil {
			when = t
		}
	}

	// Create a new date-time picker
	picker := datepicker.NewDateTimePicker(when, time.Monday, func(selected time.Time, confirm bool) {
		// Callback function: update input field with chosen date only if confirmed
		if confirm {
			input.SetText(selected.Format("2006/01/02 15:04"))
		}
	})

	// Display the picker inside a dialog
	dialog.ShowCustomConfirm(
		"Choose date and time",
		"OK",
		"Cancel",
		picker,
		picker.OnActioned,
		w,
	)
}
