package main

import (
	"fmt"
	pashua "github.com/derlinkshaender/pashua-binding-go"
)

func main() {
	p, e := pashua.LocatePashua("")
	if e != nil {
		panic(e)
	}
	fmt.Println("Found Pashua at", p)

	cfg := pashua.PashuaWindow{
		Title:        "Dialog Box",
		Transparency: 0.75,
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
				UseTime: true,
				Default: "2020-07-23",
				Textual: false,
				Y:       120,
			},
			"ok": pashua.PashuaDefaultButton{
				Label:   "OK",
				Tooltip: "Hier klicken zum Ausführen",
			},
		},
	}
	s := cfg.ToString()
	fmt.Println(s)
	res, err := pashua.RunPashua(s, "")
	fmt.Println(res, err)
	fmt.Println("Done.")

}

