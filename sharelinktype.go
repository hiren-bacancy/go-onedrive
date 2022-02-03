package onedrive

// ShareLinkType the possible values for the type property of SharingLink
type ShareLinkType int

const (
	View ShareLinkType = iota
	Edit
	Embed
)

func (shareLinkType ShareLinkType) toString() string {
	return [...]string{"view", "edit", "embed"}[shareLinkType]
}
