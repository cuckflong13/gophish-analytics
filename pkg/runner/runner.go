package runner

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/avct/uasurfer"
	"github.com/cuckflong/gophish-analytics/pkg/asset"
)

type Runner struct {
	inputFile string
	inputDir  string
	mode      int
	emails    []*asset.Email
	f         *excelize.File
}

func NewRunner(inputFile string, mode int) *Runner {
	r := &Runner{
		inputFile: inputFile,
		mode:      mode,
		emails:    []*asset.Email{},
	}
	r.f = excelize.NewFile()
	r.parseFile()
	return r
}

func (r *Runner) RunMode() {
	r.runTimestamp()
	r.runPayload()
	r.runBrowser()
	r.f.DeleteSheet("Sheet1")
	if err := r.f.SaveAs("out.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func (r *Runner) runTimestamp() {
	titles := map[string]string{"A1": "Email", "B1": "Email Sent Time", "C1": "Time Taken To Open Email", "D1": "Time Taken To Click Link", "E1": "Time Taken To Submit", "F1": "Time Taken To Report Email"}
	sheet := "Timestamp"
	r.f.NewSheet(sheet)
	for k, v := range titles {
		r.f.SetCellStr(sheet, k, v)
	}
	row := 2
	for _, e := range r.emails {
		rowString := strconv.Itoa(row)
		r.f.SetCellStr(sheet, "A"+rowString, e.GetEmail())
		r.f.SetCellStr(sheet, "B"+rowString, e.TimeSentString())
		r.f.SetCellStr(sheet, "C"+rowString, e.TimeSpentOpen())
		r.f.SetCellStr(sheet, "D"+rowString, e.TimeSpentClick())
		r.f.SetCellStr(sheet, "E"+rowString, e.TimeSpentSubmit())
		row++
	}
}

func (r *Runner) runPayload() {
	titles := map[string]string{"A1": "Email", "B1": "Time", "C1": "Message", "D1": "rid", "E1": "Login", "F1": "Password", "G1": "Address", "H1": "User-Agent"}
	sheet := "Payload"
	r.f.NewSheet(sheet)
	for k, v := range titles {
		r.f.SetCellStr(sheet, k, v)
	}
	row := 2
	for _, e := range r.emails {
		for _, event := range e.GetClickEvents() {
			rowString := strconv.Itoa(row)
			r.f.SetCellStr(sheet, "A"+rowString, e.GetEmail())
			r.f.SetCellStr(sheet, "B"+rowString, e.TimeClickedString())
			r.f.SetCellStr(sheet, "C"+rowString, event.GetMessage())
			r.f.SetCellStr(sheet, "D"+rowString, event.GetRid())
			r.f.SetCellStr(sheet, "E"+rowString, event.GetLogin())
			r.f.SetCellStr(sheet, "F"+rowString, event.GetPasswd())
			r.f.SetCellStr(sheet, "G"+rowString, event.GetAddr())
			r.f.SetCellStr(sheet, "H"+rowString, event.GetAgent())
			row++
		}
		for _, event := range e.GetSubmitEvents() {
			rowString := strconv.Itoa(row)
			r.f.SetCellStr(sheet, "A"+rowString, e.GetEmail())
			r.f.SetCellStr(sheet, "B"+rowString, e.TimeClickedString())
			r.f.SetCellStr(sheet, "C"+rowString, event.GetMessage())
			r.f.SetCellStr(sheet, "D"+rowString, event.GetRid())
			r.f.SetCellStr(sheet, "E"+rowString, event.GetLogin())
			r.f.SetCellStr(sheet, "F"+rowString, event.GetPasswd())
			r.f.SetCellStr(sheet, "G"+rowString, event.GetAddr())
			r.f.SetCellStr(sheet, "H"+rowString, event.GetAgent())
			row++
		}
	}
}

func (r *Runner) runBrowser() {
	sheet := "Browser"
	r.f.NewSheet(sheet)
	list := make(map[string]int)
	for _, e := range r.emails {
		for _, event := range e.GetClickEvents() {
			ua := uasurfer.Parse(event.GetAgent())
			list[ua.Browser.Name.String()]++
		}
	}
	r.f.SetCellStr(sheet, "A1", "Browser")
	r.f.SetCellStr(sheet, "B1", "Number")
	row := 2
	for k, v := range list {
		rowString := strconv.Itoa(row)
		r.f.SetCellStr(sheet, "A"+rowString, k)
		r.f.SetCellInt(sheet, "B"+rowString, v)
		row++
	}
}
