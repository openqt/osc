package cmd

import (
	"fmt"

	"github.com/openqt/osc/internal/color"
	"github.com/openqt/osc/internal/config"
	"github.com/openqt/osc/internal/ui"
	"github.com/spf13/cobra"
)

func infoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Print configuration info",
		Long:  "Print configuration information",
		Run: func(cmd *cobra.Command, args []string) {
			printInfo()
		},
	}
}

func printInfo() {
	const fmat = "%-25s %s\n"

	printLogo(color.Cyan)
	printTuple(fmat, "Configuration", config.K9sConfigFile, color.Cyan)
	printTuple(fmat, "Logs", config.K9sLogs, color.Cyan)
	printTuple(fmat, "Screen Dumps", config.K9sDumpDir, color.Cyan)
}

func printLogo(c color.Paint) {
	for _, l := range ui.LogoSmall {
		fmt.Println(color.Colorize(l, c))
	}
	fmt.Println()
}
