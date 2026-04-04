package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pureugong/mmk-cardnews-v2/internal/scaffold"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var slides int
	var outputDir string

	cmd := &cobra.Command{
		Use:   "new <title>",
		Short: "Create a new card news project with Remotion boilerplate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			title := args[0]

			if outputDir == "" {
				slug := slugify(title)
				outputDir = filepath.Join("output", slug)
			}

			if _, err := os.Stat(outputDir); err == nil {
				return fmt.Errorf("directory already exists: %s", outputDir)
			}

			cfg := scaffold.DefaultConfig()
			cfg.Title = title
			cfg.Slides = slides
			cfg.OutputDir = outputDir

			fmt.Printf("Creating card news project: %s (%d slides)\n", title, slides)
			fmt.Printf("Output: %s/\n", outputDir)

			if err := scaffold.Create(cfg); err != nil {
				return fmt.Errorf("scaffold: %w", err)
			}

			fmt.Println("\nDone! Next steps:")
			fmt.Printf("  1. cd %s && npm install\n", outputDir)
			fmt.Printf("  2. Edit src/cards/Card01.tsx ~ Card%02d.tsx\n", slides)
			fmt.Printf("  3. mmk-cn preview %s     # Remotion Studio\n", outputDir)
			fmt.Printf("  4. mmk-cn still %s       # Export PNGs\n", outputDir)
			fmt.Printf("  5. mmk-cn render %s      # Export MP4\n", outputDir)

			return nil
		},
	}

	cmd.Flags().IntVarP(&slides, "slides", "s", 7, "Number of slides")
	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (default: output/<slug>)")

	return cmd
}

var nonAlphaNum = regexp.MustCompile(`[^a-z0-9가-힣]+`)

func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = nonAlphaNum.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if len(s) > 50 {
		s = s[:50]
	}
	return s
}
