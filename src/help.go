package f2

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func helpText() string {
	description := fmt.Sprintf(
		"%s\n\t\t{{.Usage}}\n\n",
		pterm.Yellow("DESCRIPTION"),
	)
	usage := fmt.Sprintf(
		"%s\n\t\t{{.HelpName}} {{if .UsageText}}{{ .UsageText }}{{end}}\n\n",
		pterm.Yellow("USAGE"),
	)
	author := fmt.Sprintf(
		"{{if len .Authors}}%s\n\t\t{{range .Authors}}{{ . }}{{end}}{{end}}\n\n",
		pterm.Yellow("AUTHOR"),
	)

	version := fmt.Sprintf(
		"{{if .Version}}%s\n\t\t{{.Version}}{{end}}\n\n",
		pterm.Yellow("VERSION"),
	)
	flags := fmt.Sprintf(
		"{{if .VisibleFlags}}%s\n{{range .VisibleFlags}}{{ if (eq .Name `find` `undo` `replace` `csv`) }}\t\t{{if .Aliases}}-{{range $element := .Aliases}}%s,{{end}}{{end}} %s\n\t\t\t\t{{.Usage}}\n\n{{end}}{{end}}",
		pterm.Yellow("FLAGS"),
		pterm.Green("{{$element}}"),
		pterm.Green("--{{.Name}} {{.DefaultText}}"),
	)
	options := fmt.Sprintf(
		"%s\n{{range .VisibleFlags}}{{ if not (eq .Name `find` `undo` `replace` `csv`) }}\t\t{{if .Aliases}}-{{range $element := .Aliases}}%s,{{end}}{{end}} %s\n\t\t\t\t{{.Usage}}\n\n{{end}}{{end}}{{end}}",
		pterm.Yellow("OPTIONS"),
		pterm.Green("{{$element}}"),
		pterm.Green("--{{.Name}} {{.DefaultText}}"),
	)

	docs := fmt.Sprintf(
		"%s\n\t\t%s\n\n",
		pterm.Yellow("DOCUMENTATION"),
		"https://github.com/ayoisaiah/f2/wiki",
	)
	website := fmt.Sprintf(
		"%s\n\t\thttps://github.com/ayoisaiah/f2",
		pterm.Yellow("WEBSITE"),
	)

	return description + usage + author + version + flags + options + docs + website
}

func shortHelp(app *cli.App) string {
	heading := fmt.Sprintf(
		"F2 — Command-line bulk renaming tool [version %s]\n\n",
		app.Version,
	)

	usage := fmt.Sprintf("Usage: %s\n", app.UsageText)

	description := `
F2 helps you organise your filesystem through batch renaming.
The simplest usage is to do a basic find and replace:

$ f2 Screenshot Image
+--------------------+---------------+--------+
|       INPUT        |    OUTPUT     | STATUS |
+--------------------+---------------+--------+
| Screenshot (1).png | Image (1).png | ok     |
| Screenshot (2).png | Image (2).png | ok     |
| Screenshot (3).png | Image (3).png | ok     |
+--------------------+---------------+--------+

The first argument is the find string, while the second is the
replacement string. Any other arguments are interpreted as paths
to files or directories where the renaming operation should take 
place. The current directory is used by default.

F2 supports many command-line options. Use the --help flag to examine
the full list. For extensive usage examples, visit the project wiki:
https://github.com/ayoisaiah/f2/wiki`

	return heading + usage + description
}
