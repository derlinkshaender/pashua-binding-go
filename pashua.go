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

// CompletionMode is a type to store the completion mode for a combobox
type CompletionMode string

const (
	NoCompletion    CompletionMode = "0"
	CaseSensitive                  = "1"
	CaseInsensitive                = "2"
)

// FontSize is a type to store the fontsize of a Pashua window
type FontSize string

const (
	Regular FontSize = "regular"
	Small            = "small"
	Mini             = "mini"
)

// PashuaButton is a structure that holds all information for a PashuaButton
type PashuaButton struct {
	Label    string
	X        int
	Y        int
	Disabled bool
	Tooltip  string
}

// PashuaCancelButton is a structure that holds all information for a PashuaCancelButton
type PashuaCancelButton struct {
	Label    string
	Disabled bool
	Tooltip  string
}

// PashuaCheckbox is a structure that holds all information for a PashuaCheckbox
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

// PashuaCombobox is a structure that holds all information for a PashuaCombobox
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

// PashuaDate is a structure that holds all information for a PashuaDate
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

// PashuaDefaultButton is a structure that holds all information for a PashuaDefaultButton
type PashuaDefaultButton struct {
	Label    string
	Disabled bool
	Tooltip  string
}

// PashuaImage is a structure that holds all information for a PashuaImage
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

// PashuaOpenBrowser is a structure that holds all information for a PashuaOpenBrowser
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

// PashuaPassword is a structure that holds all information for a PashuaPassword
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

// PashuaPopup is a structure that holds all information for a PashuaPopup
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

// PashuaRadioButton is a structure that holds all information for a PashuaRadioButton
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

// PashuaSaveBrowser is a structure that holds all information for a PashuaSaveBrowser
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

// PashuaText is a structure that holds all information for a PashuaText
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

// PashuaTextBox is a structure that holds all information for a PashuaTextBox
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

// PashuaTextField is a structure that holds all information for a PashuaTextField
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

// PashuaComponents is type for th elist of components contained in a Pashua window
type PashuaComponents map[string]interface{}

//  PashuaWindow is the top-most structure when defininng a dialog window for Pashua
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

// fileExists is a helper function that returns true
// if a specified file exists and is a file
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return (info != nil) && !info.IsDir()
}

// LocatePashua is one of the two main binding function
// and tries to find the Pashua.app and the executable contained
// within the app container by iterating over the standard
// file locations. Returns the location path and an error code
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
		// yeah, in Go there is no "insert" function,
		// so this is a little tricky at first
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

// RunPashua is one of the two main binding function.
// it sets up a pipe and executes Pashua as an external command,
// then converts the STdOut and StdErr to strings and parses the output
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

// RunPashuaWithStruct is a convenience function that saves you
// from having to convert a struct-based window definition to a string first
func RunPashuaWithStruct(pashuaWindow *PashuaWindow, pashuaPath string) (map[string]string, error) {
	configString := pashuaWindow.ToString()
	return RunPashua(configString, pashuaPath)
}

// parsePashuaOutput takes a list of lines and
// converts key=value pairs to a map and
// skips empty lines of lines without an "="
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

// getFieldValue converts various types to a string
// representation as Pashua needs the configuration as string
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

/*
the function below convert each possible component type
into a config string that Pashua con use.
the "ToString()" of the PashuaWindow iterates over all
defines component and combines them into a large config string
that can be provided to Pashua
*/

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

func (win *PashuaWindow) WindowToString() string {
	result := []string{}
	autoSaveKey := getFieldValue(win.AutoSaveKey)
	title := getFieldValue(win.Title)
	if title != "" {
		result = append(result, "*.title="+title)
	}
	if win.Transparency > 0 {
		// do not display invisible dialogs ;-)
		result = append(result, "*.transparency="+getFieldValue(win.Transparency))
	}
	if win.AutoCloseTime > 1 {
		result = append(result, "*.autoclosetime="+getFieldValue(win.AutoCloseTime))
	}
	if autoSaveKey != "" {
		result = append(result, "*.autosavekey="+autoSaveKey)
	} else {
		result = append(result, "*.x="+getFieldValue(win.X))
		result = append(result, "*.y="+getFieldValue(win.Y))
	}
	if win.Floating {
		result = append(result, "*.floating=1")
	}
	return strings.Join(result, "\n")
}

func (win *PashuaWindow) ToString() string {
	var result = []string{}
	result = append(result, win.WindowToString())
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

