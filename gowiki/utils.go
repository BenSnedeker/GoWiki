package gowiki

// Choose either raw HTML or Text
type ReturnType int

const (
	Html ReturnType = iota
	Text
)

// Choose which sections to return
// Ignored by WikiSearch
type ReturnContent int

const (
	All ReturnContent = iota
	Limited
	Summary
	References
)

// Choose either colorful fancy text or a clean plain string
// Defaults to Clean on HTML type response
type ReturnStyle int

const (
	Fancy ReturnStyle = iota
	Clean
)

// A flag struct to contain all categories of flags.
// Each flag category can only have 1 state selected.
type Flags struct {
	Term     string
	Type     ReturnType
	Content  ReturnContent
	Style    ReturnStyle
	Language string
}

func NewFlags() Flags {
	return Flags{
		Term:     "",
		Type:     ReturnType(Text),
		Content:  ReturnContent(All),
		Style:    ReturnStyle(Fancy),
		Language: "en",
	}
}
