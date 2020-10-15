package runner

import (
	"strings"
	"tix/creator"
	"tix/creator/dryrun"
	"tix/creator/github"
	"tix/logger"
	"tix/md"
	"tix/render"
	"tix/runner/env"
	"tix/runner/field"
	"tix/settings"
	"tix/ticket"
)

type GithubRunner struct {
	envMap   map[string]string
	settings *settings.Settings
}

func NewGithubRunner(envMap map[string]string, settings *settings.Settings) TixRunner {
	return &GithubRunner{
		envMap:   envMap,
		settings: settings,
	}
}

func (r GithubRunner) Run(markdownData []byte) error {
	if !r.settings.Github.Configured() {
		return nil
	}

	err := env.CheckGithubEnvironment(r.envMap)
	if err != nil {
		return err
	}

	tickets, err := r.parseMarkdown(markdownData)

	if err != nil {
		return err
	}

	r.githubCreator().CreateTickets(tickets)

	return nil
}

func (r GithubRunner) parseMarkdown(markdownData []byte) ([]*ticket.Ticket, error) {
	fieldState := field.GithubFieldState(*r.settings)
	markdownParser := md.NewParser(fieldState)
	return markdownParser.Parse(markdownData)
}

func (r GithubRunner) githubCreator() *github.Creator {
	api := r.createGithubApi()
	cache := github.NewCache(api)
	renderer := render.NewGithubBodyRenderer()
	options := github.NewOptions(renderer)
	issueCreator := github.NewIssueCreator(api, cache, options)
	projectCreator := github.NewProjectCreator(api, cache, options)

	if r.settings.Github.NoProjects {
		return github.NewCreatorWithoutProjects(issueCreator, projectCreator)
	} else {
		return github.NewCreatorWithProjects(issueCreator, projectCreator)
	}
}

func (r GithubRunner) createGithubApi() github.Api {
	return github.NewGithubApi(r.envMap[env.GithubApiToken], r.settings.Github.Owner, r.settings.Github.Repo)
}

func (r GithubRunner) DryRun(markdownData []byte) error {
	if !r.settings.Github.Configured() {
		return nil
	}

	err := env.CheckGithubEnvironment(r.envMap)
	if err != nil {
		return err
	}

	tickets, err := r.parseMarkdown(markdownData)

	if err != nil {
		return err
	}

	var builder strings.Builder
	r.dryRunCreator(&builder).CreateTickets(tickets)
	logger.Message(builder.String())

	return nil
}

func (r GithubRunner) dryRunCreator(builder *strings.Builder) creator.TicketCreator {
	var startingTicketLevel int
	if r.settings.Github.NoProjects {
		startingTicketLevel = 1
	} else {
		startingTicketLevel = 0
	}
	labels := []*dryrun.LevelLabel{
		dryrun.NewLevelLabel("project", "projects"),
		dryrun.NewLevelLabel("issue", "issues"),
	}
	renderer := render.NewGithubBodyRenderer()
	return dryrun.NewCreator(builder, labels, renderer, startingTicketLevel, "github")
}
