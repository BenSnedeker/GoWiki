package gowiki

// Choose either raw HTML or Text
type ReturnType int

const (
	Html ReturnType = iota
	Text
)

// Choose which sections to return
type ReturnContent int

const (
	All ReturnContent = iota
	Limited
	Summary
	References
)

// Choose either colorful fancy text or a clean plain string
type ReturnStyle int

const (
	Fancy ReturnStyle = iota
	Clean
)

// A flag struct to contain all categories of flags.
// Each flag category can only have 1 state selected.
type Flags struct {
	Name    string
	Type    ReturnType
	Content ReturnContent
	Style   ReturnStyle
}

func NewFlags() Flags {
	return Flags{
		Name:    "",
		Type:    ReturnType(Text),
		Content: ReturnContent(All),
		Style:   ReturnStyle(Fancy),
	}
}
