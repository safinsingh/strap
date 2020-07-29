package main

func main() {
	repo, outputDir := parseFlags()
	cloneRepo(repo, outputDir)
}
