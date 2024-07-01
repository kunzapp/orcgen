package director

type Config struct {
	Bin         string
	Preferences string
	WorkingDir  string
	UserDataDir string

	InLambda bool
}
