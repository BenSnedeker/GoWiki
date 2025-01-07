package gowiki

type Flags struct {
	Html       bool
	Text       bool
	All        bool
	Limited    bool
	Summary    bool
	References bool
	Wikilinks  bool
	Clean      bool
	Fancy      bool
}

func NewFlags() Flags {
	return Flags{
		// Raw HTML or tagless string
		Html: false,
		Text: true,
		// Which sections to return
		All:        false,
		Limited:    true,
		Summary:    false,
		References: false,
		Wikilinks:  false,
		// Raw string or colored string
		Clean: false,
		Fancy: true,
	}
}
