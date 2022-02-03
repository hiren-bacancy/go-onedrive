package onedrive

// ShareLinkType the possible values for the scope property of SharingLink
type ShareLinkScope int

const (
	Anonymous ShareLinkScope = iota
	Organization
)

func (shareLinkScope ShareLinkScope) toString() string {
	return [...]string{"anonymous", "organization"}[shareLinkScope]
}
