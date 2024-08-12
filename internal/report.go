package internal

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"text/template"
	"time"

	gh "github.com/google/go-github/v63/github"
	"github.com/ryotaro612/dpcli/internal/github"
)

type deps struct {
	github.Client
}

type Option struct {
	awsProfile   *string
	templateFile *string
	verbose      bool
}

type Reporting struct {
	github    *github.Client
	generator *generator
	logger    *slog.Logger
}

func (r *Reporting) Report(ctx context.Context) error {
	offset, err := calcOffset(time.Now())
	if err != nil {
		return err
	}
	return r.ReportOffset(ctx, offset)
}

func (r *Reporting) ReportOffset(ctx context.Context, offset time.Time) error {
	pullRequests, err := r.github.ReadPullRequests(ctx, offset)
	if err != nil {
		return err
	}
	converted := convertPullRequests(pullRequests)
	err = r.generator.generate(&converted)
	return err
}

// NewReporting creates a new Reporting object with the given options.
func NewReporting(ctx context.Context, awsProfile string, verbose bool, template string) (Reporting, error) {
	logger := NewLogger(verbose)
	var r Reporting
	s, err := ReadSecret(ctx, logger, awsProfile)
	if err != nil {
		return r, err
	}
	g := github.NewClient(logger, s.GithubToken)
	gen, err := newGenerator(template)
	if err != nil {
		return r, err
	}
	return Reporting{github: &g, generator: &gen, logger: logger}, nil
}

func calcOffset(current time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return current, err
	}
	year, month, day := current.In(loc).Date()
	meeting := time.Date(year, month, day, 11, 0, 0, 0, loc)
	if meeting.After(current) {
		if meeting.Weekday() == time.Monday {
			return meeting.AddDate(0, 0, -3), nil
		}
		return meeting.AddDate(0, 0, -1), nil

	}
	return meeting, nil
}

type generator struct {
	template string
}

func (g *generator) generate(pullRequests *[]TemplatePullRequest) error {
	tmpl, err := template.New("report").Parse(g.template)

	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, pullRequests)
	if err != nil {
		return err
	}
	a := buf.String()
	fmt.Printf(a)

	// 	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// if err != nil { panic(err) }

	// if err != nil { panic(err) }
	return nil
}

func newGenerator(file string) (generator, error) {

	return generator{
		template: `
- ■ やったこと
{{range .}}{{- if eq .Login "ryotaro612" }}{{ printf "  - %s\n" .Title }}{{ end }}{{end}}
- ■ やること
- ■ 困っていること/ひとこと
`,
	}, nil

}

type TemplatePullRequest struct {
	Title string
	Login string
	URL   string
}

func convertPullRequests(pullRequests []*gh.PullRequest) []TemplatePullRequest {
	var tprs []TemplatePullRequest
	for _, pr := range pullRequests {
		tprs = append(tprs, TemplatePullRequest{
			Title: *pr.Title,
			Login: *pr.User.Login,
			URL:   *pr.URL,
		})
	}
	return tprs
}
