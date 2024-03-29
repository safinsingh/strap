package cmd

type ProjectConfig struct {
	Name     string          `json:"name"`
	Version  string          `json:"version"`
	Commands ProjectCommands `json:"commands"`
}

func (p *ProjectConfig) BumpVersion(version string) {
	p.Version = version
}

type ProjectCommands struct {
	Default DefaultCommand `json:"default"`
}

type DefaultCommand struct {
	Steps      []string `json:"steps"`
	Entrypoint []string `json:"entrypoint"`
}

type GlobalConfig struct {
	Aliases map[string]string `json:"aliases"`
}
