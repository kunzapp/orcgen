package rod

type Option func(*Config)

type Config struct {
	Preferences string
	WorkingDir  string
	UserDataDir string
}

func WithPreferences(preferences string) Option {
	return func(cfg *Config) {
		cfg.Preferences = preferences
	}
}

func WithWorkingDir(workingDir string) Option {
	return func(cfg *Config) {
		cfg.WorkingDir = workingDir
	}
}

func WithUserDataDir(userDataDir string) Option {
	return func(cfg *Config) {
		cfg.UserDataDir = userDataDir
	}
}
