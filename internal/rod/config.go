package rod

type Option func(*Config)

type Config struct {
	Bin         string
	Preferences string
	WorkingDir  string
	UserDataDir string
}

func WithBin(bin string) Option {
	return func(cfg *Config) {
		cfg.Bin = bin
	}
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
