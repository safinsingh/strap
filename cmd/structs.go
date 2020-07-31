package cmd

type ProjectConfig struct {
	Version  string          `json:"version"`
	Commands ProjectCommands `json:"commands"`
}

type ProjectCommands struct {
	Default DefaultCommand `json:"default"`
}

type DefaultCommand struct {
	Steps      []string `json:"steps"`
	Entrypoint []string `json:"entrypoint"`
}
