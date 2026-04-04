package main

import (
	"fmt"
	"os"

	"github.com/pureugong/mmk-cardnews-v2/internal/cli"
	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	root := &cobra.Command{
		Use:     "mmk-cn",
		Short:   "Card news generator — Remotion-based, AI-powered",
		Version: version,
	}

	root.AddCommand(cli.NewCmd())
	root.AddCommand(cli.ImagegenCmd())
	root.AddCommand(cli.StillCmd())
	root.AddCommand(cli.RenderCmd())
	root.AddCommand(cli.PdfCmd())
	root.AddCommand(cli.PreviewCmd())

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
