package pashua

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"strings"
)

type CompletionMode string

type FontSize string

const (
	NoCompletion    CompletionMode = "0"
	CaseSensitive                  = "1"
	CaseInsensitive                = "2"
)

const (
	Regular FontSize = "regular"
	Small            = "small"
	Mini             = "mini"
)

type PashuaButton struct {
	Label    string
	X        int
	Y        int
	Disabled bool
	Tooltip  string
}

type PashuaCancelButton struct {
	Label    string
	Disabled bool
	Tooltip  string
}

type PashuaCheckbox struct {
	Label    string
	Default  bool
	Disabled bool
	Tooltip  string
	X        int
	Y        int
	RelX     int
	RelY     int
}

type PashuaCombobox struct {
	Label          string
	Option         []string
	CompletionMode CompletionMode
	Mandatory      bool
	Rows           int
	Placeholder    string
	Disabled       bool
	Tooltip        string
	Width          int
	X              int
	Y              int
	RelX           int
	RelY           int
}

type PashuaDate struct {
	Label    string
	Textual  bool
	UseDate  bool
	UseTime  bool
	Default  string
	Disabled bool
	Tooltip  string
	X        int
	Y        int
}

type PashuaDefaultButton struct {
	Label    string
	Disabled bool
	Tooltip  string
}

type PashuaImage struct {
	Label     string
	Path      string
	Border    bool
	Width     int
	Height    int
	MaxWidth  int
	MaxHeight int
	UpScale   bool
	Tooltip   string
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaOpenBrowser struct {
	Label       string
	DefaultPath string
	Width       int
	Filetype    string
	Placeholder string
	Mandatory   bool
	X           int
	Y           int
	RelX        int
	RelY        int
}

type PashuaPassword struct {
	Label     string
	Default   bool
	Disabled  bool
	Mandatory bool
	Tooltip   string
	Width     int
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaPopup struct {
	Option    []string
	Default   string
	Label     string
	Disabled  bool
	Tooltip   string
	Mandatory bool
	Width     int
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaRadioButton struct {
	Option    []string
	Default   string
	Label     string
	Disabled  bool
	Tooltip   string
	Mandatory bool
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaSaveBrowser struct {
	Label       string
	DefaultPath string
	Width       int
	Filetype    string
	Placeholder string
	Mandatory   bool
	X           int
	Y           int
	RelX        int
	RelY        int
}

type PashuaText struct {
	Label   string
	Text    string
	Tooltip string
	Width   int
	X       int
	Y       int
	RelX    int
	RelY    int
}

type PashuaTextBox struct {
	Label     string
	Default   string
	Tooltip   string
	FixedFont bool
	FontSize  FontSize // regular small mini
	Mandatory bool
	Disabled  bool
	Width     int
	Height    int
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaTextField struct {
	Label     string
	Default   string
	Tooltip   string
	Mandatory bool
	Disabled  bool
	Width     int
	X         int
	Y         int
	RelX      int
	RelY      int
}

type PashuaComponents map[string]interface{}

type PashuaWindow struct {
	AutoCloseTime int
	AutoSaveKey   string
	Floating      bool
	Title         string
	Transparency  float64
	X             int
	Y             int
	Components    PashuaComponents
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return (info != nil) && !info.IsDir()
}

func LocatePashua(pashuaPath string) (string, error) {
	const bundlePath = "Pashua.app/Contents/MacOS/Pashua"
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	usrhome := usr.HomeDir
	pashuaPlaces := []string{
		path.Join(path.Dir(os.Args[0]), "Pashua"),
		path.Join(path.Dir(os.Args[0]), bundlePath),
		"./" + bundlePath,
		path.Join("/Applications", bundlePath),
		path.Join(usrhome, "Applications", bundlePath),
		path.Join("/usr/local/bin", bundlePath),
	}
	if pashuaPath != "" {
		// insert at index 0
		pashuaPlaces = append(pashuaPlaces, "")
		copy(pashuaPlaces[1:], pashuaPlaces[0:])
		pashuaPlaces[0] = pashuaPath

	}
	for _, p := range pashuaPlaces {
		if fileExists(p) {
			return p, nil
		}
	}
	return "", fmt.Errorf("Could not locate pashua")
}

func parsePashuaOutput(output string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		pos := strings.Index(line, "=")
		if line != "" && pos > 0 {
			key := line[:pos]
			value := line[pos+1:]
			result[key] = value
		}
	}
	return result
}

func RunPashua(configData string, pashuaPath string) (map[string]string, error) {
	var err error
	result := make(map[string]string)
	appPath, err := LocatePashua(pashuaPath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(appPath, "-")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = bytes.NewBuffer([]byte(configData))
	err = cmd.Run()
	if err != nil {
		return result, fmt.Errorf("Error: %w, pashua error text: %s", err, string(stderr.Bytes()))
	}
	outStr := string(stdout.Bytes())
	result = parsePashuaOutput(outStr)
	return result, err
}

func getFieldValue(field interface{}) string {
	result := ""
	switch field.(type) {
	case string:
		result = field.(string)
	case int:
		result = strconv.Itoa(field.(int))
	case int8:
		result = strconv.Itoa(field.(int))
	case int16:
		result = strconv.Itoa(field.(int))
	case int32:
		result = strconv.Itoa(field.(int))
	case int64:
		result = strconv.Itoa(field.(int))
	case float32:
		result = strconv.FormatFloat(field.(float64), 'f', 4, 32)
	case float64:
		result = strconv.FormatFloat(field.(float64), 'f', 4, 64)
	case bool:
		if field.(bool) {
			return "1"
		} else {
			return "0"
		}
	case CompletionMode:
		result = string(field.(CompletionMode))
	case FontSize:
		result = string(field.(FontSize))
	default:
		result = field.(string)
	}
	return result
}

func (btn *PashuaButton) ToString(key string) string {
	result := []string{key + ".type=button"}
	result = append(result, key+".label="+getFieldValue(btn.Label))
	result = append(result, key+".tooltip="+getFieldValue(btn.Tooltip))
	result = append(result, key+".disabled="+getFieldValue(btn.Disabled))
	result = append(result, key+".x="+getFieldValue(btn.X))
	result = append(result, key+".y="+getFieldValue(btn.Y))
	return strings.Join(result, "\n")
}

func (btn *PashuaDate) ToString(key string) string {
	result := []string{key + ".type=date"}
	result = append(result, key+".label="+getFieldValue(btn.Label))
	result = append(result, key+".tooltip="+getFieldValue(btn.Tooltip))
	result = append(result, key+".disabled="+getFieldValue(btn.Disabled))
	result = append(result, key+".default="+getFieldValue(btn.Default))
	result = append(result, key+".date="+getFieldValue(btn.UseDate))
	result = append(result, key+".time="+getFieldValue(btn.UseTime))
	result = append(result, key+".textual="+getFieldValue(btn.Textual))
	result = append(result, key+".x="+getFieldValue(btn.X))
	result = append(result, key+".y="+getFieldValue(btn.Y))
	return strings.Join(result, "\n")
}

func (btn *PashuaDefaultButton) ToString(key string) string {
	result := []string{key + ".type=defaultbutton"}
	result = append(result, key+".label="+getFieldValue(btn.Label))
	result = append(result, key+".tooltip="+getFieldValue(btn.Tooltip))
	result = append(result, key+".disabled="+getFieldValue(btn.Disabled))
	return strings.Join(result, "\n")
}

func (txt *PashuaCancelButton) ToString(key string) string {
	result := []string{key + ".type=cancelbutton"}
	result = append(result, key+".label="+getFieldValue(txt.Label))
	result = append(result, key+".tooltip="+getFieldValue(txt.Tooltip))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	return strings.Join(result, "\n")
}

func (txt *PashuaCheckbox) ToString(key string) string {
	result := []string{key + ".type=checkbox"}
	result = append(result, key+".label="+getFieldValue(txt.Label))
	result = append(result, key+".default="+getFieldValue(txt.Default))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".tooltip="+getFieldValue(txt.Tooltip))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaCombobox) ToString(key string) string {
	result := []string{key + ".type=combobox"}
	result = append(result, key+".label="+getFieldValue(txt.Label))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".tooltip="+getFieldValue(txt.Tooltip))
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".rows="+getFieldValue(txt.Rows))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	result = append(result, key+".placeholder="+getFieldValue(txt.Placeholder))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".completion="+getFieldValue(txt.CompletionMode))
	for _, k := range txt.Option {
		result = append(result, key+".option="+getFieldValue(k))
	}
	return strings.Join(result, "\n")
}

func (txt *PashuaImage) ToString(key string) string {
	result := []string{key + ".type=image"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".path="+txt.Path)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".height="+getFieldValue(txt.Height))
	result = append(result, key+".maxwidth="+getFieldValue(txt.MaxWidth))
	result = append(result, key+".maxheight="+getFieldValue(txt.MaxHeight))
	result = append(result, key+".upscale="+getFieldValue(txt.UpScale))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaOpenBrowser) ToString(key string) string {
	result := []string{key + ".type=openbrowser"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".default="+txt.DefaultPath)
	result = append(result, key+".filetype="+txt.Filetype)
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".placeholder="+getFieldValue(txt.Placeholder))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaSaveBrowser) ToString(key string) string {
	result := []string{key + ".type=savebrowser"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".default="+txt.DefaultPath)
	result = append(result, key+".filetype="+txt.Filetype)
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".placeholder="+getFieldValue(txt.Placeholder))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaPassword) ToString(key string) string {
	result := []string{key + ".type=password"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".default="+getFieldValue(txt.Default))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaPopup) ToString(key string) string {
	result := []string{key + ".type=popup"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".width="+getFieldValue(txt.Width))
	result = append(result, key+".default="+getFieldValue(txt.Default))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	for _, k := range txt.Option {
		result = append(result, key+".option="+getFieldValue(k))
	}
	return strings.Join(result, "\n")
}

func (txt *PashuaRadioButton) ToString(key string) string {
	result := []string{key + ".type=radiobutton"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".default="+getFieldValue(txt.Default))
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	for _, k := range txt.Option {
		result = append(result, key+".option="+getFieldValue(k))
	}
	return strings.Join(result, "\n")
}

func (txt *PashuaText) ToString(key string) string {
	result := []string{key + ".type=text"}
	result = append(result, key+".label="+txt.Label)
	s := getFieldValue(txt.Text)
	s = strings.Replace(s, "\n", "[return]", -1)
	result = append(result, key+".text="+s)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaTextBox) ToString(key string) string {
	result := []string{key + ".type=textbox"}
	result = append(result, key+".label="+txt.Label)
	s := getFieldValue(txt.Default)
	s = strings.Replace(s, "\n", "[return]", -1)
	result = append(result, key+".default="+s)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	s = getFieldValue(txt.FixedFont)
	if s == "1" {
		s = "fixed"
	} else {
		s = ""
	}
	result = append(result, key+".fonttype="+s)
	result = append(result, key+".fontsize="+getFieldValue(txt.FontSize))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (txt *PashuaTextField) ToString(key string) string {
	result := []string{key + ".type=textfield"}
	result = append(result, key+".label="+txt.Label)
	result = append(result, key+".default="+txt.Default)
	result = append(result, key+".tooltip="+txt.Tooltip)
	result = append(result, key+".disabled="+getFieldValue(txt.Disabled))
	result = append(result, key+".mandatory="+getFieldValue(txt.Mandatory))
	result = append(result, key+".x="+getFieldValue(txt.X))
	result = append(result, key+".y="+getFieldValue(txt.Y))
	result = append(result, key+".relx="+getFieldValue(txt.RelX))
	result = append(result, key+".rely="+getFieldValue(txt.RelY))
	return strings.Join(result, "\n")
}

func (win *PashuaWindow) ToString() string {
	var result = []string{}
	result = append(result, "*.title="+getFieldValue(win.Title))
	result = append(result, "*.transparency="+getFieldValue(win.Transparency))
	result = append(result, "*.autoclosetime="+getFieldValue(win.AutoCloseTime))
	result = append(result, "*.autosavekey="+getFieldValue(win.AutoSaveKey))
	result = append(result, "*.floating="+getFieldValue(win.Floating))
	result = append(result, "*.x="+getFieldValue(win.X))
	result = append(result, "*.y="+getFieldValue(win.Y))
	for key, comp := range win.Components {
		switch comp.(type) {
		case PashuaButton:
			btn := comp.(PashuaButton)
			result = append(result, btn.ToString(key))
		case PashuaCancelButton:
			txt := comp.(PashuaCancelButton)
			result = append(result, txt.ToString(key))
		case PashuaCheckbox:
			txt := comp.(PashuaCheckbox)
			result = append(result, txt.ToString(key))
		case PashuaCombobox:
			txt := comp.(PashuaCombobox)
			result = append(result, txt.ToString(key))
		case PashuaDate:
			txt := comp.(PashuaDate)
			result = append(result, txt.ToString(key))
		case PashuaDefaultButton:
			txt := comp.(PashuaDefaultButton)
			result = append(result, txt.ToString(key))
		case PashuaImage:
			txt := comp.(PashuaImage)
			result = append(result, txt.ToString(key))
		case PashuaOpenBrowser:
			txt := comp.(PashuaOpenBrowser)
			result = append(result, txt.ToString(key))
		case PashuaPassword:
			txt := comp.(PashuaPassword)
			result = append(result, txt.ToString(key))
		case PashuaPopup:
			txt := comp.(PashuaPopup)
			result = append(result, txt.ToString(key))
		case PashuaRadioButton:
			txt := comp.(PashuaRadioButton)
			result = append(result, txt.ToString(key))
		case PashuaSaveBrowser:
			txt := comp.(PashuaSaveBrowser)
			result = append(result, txt.ToString(key))
		case PashuaText:
			txt := comp.(PashuaText)
			result = append(result, txt.ToString(key))
		case PashuaTextBox:
			txt := comp.(PashuaTextBox)
			result = append(result, txt.ToString(key))
		case PashuaTextField:
			txt := comp.(PashuaTextField)
			result = append(result, txt.ToString(key))
		}
	}
	return strings.Join(result, "\n")
}

func main() {
	p, e := LocatePashua("")
	fmt.Println(p, e)

	cfg := PashuaWindow{
		Title:        "Dialog Box",
		Transparency: 0.75,
		Components: PashuaComponents{
			"tf": PashuaTextField{
				Label:   "Gib was ein",
				Default: "42",
				Width:   100,
				Y:       20,
			},
			"cb": PashuaCombobox{
				Label:          "My combobox label",
				Option:         []string{"Gromit", "Wallace", "Harold", "Maude"},
				Width:          220,
				Tooltip:        "Choose from the list",
				CompletionMode: CaseInsensitive,
				Y:              60,
			},
			"dt": PashuaDate{
				Label:   "TickTock",
				Tooltip: "A Date/Time control",
				UseDate: true,
				UseTime: true,
				Default: "2020-07-23",
				Textual: false,
				Y:       120,
			},
			"ok": PashuaDefaultButton{
				Label:   "OK",
				Tooltip: "Hier klicken zum Ausführen",
			},
		},
	}
	s := cfg.ToString()
	fmt.Println(s)
	res, err := RunPashua(s, "")
	fmt.Println(res, err)
	fmt.Println("Done.")

}
