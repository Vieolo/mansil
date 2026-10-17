package generator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

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

// A Unicode code point, written in the JSON source as a hex string, such as "0x1F300"
type CodePoint uint32

func (c *CodePoint) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	hex, ok := strings.CutPrefix(s, "0x")
	if !ok {
		return fmt.Errorf("code point %q must start with 0x", s)
	}
	v, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return fmt.Errorf("invalid code point %q: %w", s, err)
	}
	*c = CodePoint(v)
	return nil
}

// An inclusive range of code points that take the same number of terminal columns
type CharRange struct {
	From CodePoint `json:"from"`
	To   CodePoint `json:"to"`
	Name string    `json:"name"`
}

// The characters that do not take exactly one terminal column
type Widths struct {
	Zero []CharRange `json:"zero"`
	Wide []CharRange `json:"wide"`
}

type Source struct {
	Styles   []Style   `json:"styles"`
	Colors   []Color   `json:"colors"`
	Controls []Control `json:"controls"`
	Widths   Widths    `json:"widths"`
}
