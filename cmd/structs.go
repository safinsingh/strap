package cmd

type ProjectConfig struct {
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
