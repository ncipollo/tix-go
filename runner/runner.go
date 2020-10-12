package runner

type TixRunner interface {
	Run(markdownData []byte) error
	DryRun(markdownData []byte) error
}