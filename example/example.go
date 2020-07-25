package main

import (
	"fmt"
	pashua "github.com/derlinkshaender/pashua-binding-go"
)

func main() {
	// look for Pashua and bail out, if not found in the standard locations
	p, e := pashua.LocatePashua("")
	if e != nil {
		panic(e)
	}
	fmt.Println("Found Pashua at", p)

	// define a window to show some of the components
	cfg := pashua.PashuaWindow{
		Title:        "Dialog Box",
		AutoSaveKey: "hurga",
		Transparency: 1,
		Components: pashua.PashuaComponents{
			"tf": pashua.PashuaTextField{
				Label:   "Gib was ein",
				Default: "42",
				Width:   100,
				Y:       20,
			},
			"cb": pashua.PashuaCombobox{
				Label:          "My combobox label",
				Option:         []string{"Gromit", "Wallace", "Harold", "Maude"},
				Width:          220,
				Tooltip:        "Choose from the list",
				CompletionMode: pashua.CaseInsensitive,
				Y:              60,
			},
			"dt": pashua.PashuaDate{
				Label:   "TickTock",
				Tooltip: "A Date/Time control",
				UseDate: true,
				UseTime: false,
				Default: "2020-07-04",
				Textual: false,
				Y:       120,
			},
			"ok": pashua.PashuaDefaultButton{
				Label:   "OK",
				Tooltip: "Click here to submit dialog",
			},
			"cancel": pashua.PashuaCancelButton{
				Label:   "Cancel",
				Tooltip: "",
			},
		},
	}

	// take the struct definition of a Pashua window and
	// convert it into a string that Pashua can use
	// then run the tool
	//
	// this is functionally euqivalent to calling the
	// convenience function "RunPashuaWithStruct" like so:
	// res, err := RunPashuaWithStruct(&cfg, "")
	s := cfg.ToString()
	res, err := pashua.RunPashua(s, "")
	if err != nil {
		panic(err)
	}

	// display the resulting map
	fmt.Println("== Result ==")
	for key, value := range res {
		fmt.Println(key, "=", value)
	}
	fmt.Println("Done.")
}

