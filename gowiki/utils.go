package gowiki

type Flags struct {
	html       bool
	text       bool
	all        bool
	limited    bool
	summary    bool
	references bool
	wikilinks  bool
	clean      bool
	fancy      bool
}

func NewFlags() Flags {
	return Flags{
		// Raw HTML or tagless string
		html: false,
		text: true,
		// Which sections to return
		all:        false,
		limited:    true,
		summary:    false,
		references: false,
		wikilinks:  false,
		// Raw string or colored string
		clean: false,
		fancy: true,
	}
}
