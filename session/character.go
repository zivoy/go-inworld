package session

type Assets struct {
	AvatarImage         string
	AvatarImageOriginal string

	RpmModelUri         string
	RpmImageUriPortrait string
	RpmImageUriPosture  string
}

type Character struct {
	Id           string
	ResourceName string
	DisplayName  string
	Assets       *Assets
}
