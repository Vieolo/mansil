package generator

type Style struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Color struct {
	Name     string `json:"name"`
	Fg       string `json:"fg"`
	Bg       string `json:"bg"`
	FgBright string `json:"fgBright"`
	BgBright string `json:"bgBright"`
}

type Control struct {
	Name string `json:"name"`
	Seq  string `json:"seq,omitempty"`
	Fmt  string `json:"fmt,omitempty"`
}

type Source struct {
	Styles   []Style   `json:"styles"`
	Colors   []Color   `json:"colors"`
	Controls []Control `json:"controls"`
}
