package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func StillCmd() *cobra.Command {
	var outputDir string

	cmd := &cobra.Command{
		Use:   "still <project-dir>",
		Short: "Export each card as PNG using Remotion still",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir := args[0]

			if _, err := os.Stat(filepath.Join(projectDir, "package.json")); err != nil {
				return fmt.Errorf("not a Remotion project: %s (missing package.json)", projectDir)
			}

			if outputDir == "" {
				outputDir = filepath.Join(projectDir, "out", "stills")
			}
			if err := os.MkdirAll(outputDir, 0o755); err != nil {
				return err
			}

			// Count cards from src/cards/index.ts
			count, err := countCards(projectDir)
			if err != nil {
				return fmt.Errorf("count cards: %w", err)
			}

			fmt.Printf("Exporting %d cards as PNG...\n", count)

			for i := 0; i < count; i++ {
				outFile := filepath.Join(outputDir, fmt.Sprintf("card-%02d.png", i+1))
				props := fmt.Sprintf(`{"cardIndex":%d}`, i)

				c := exec.Command("npx", "remotion", "still",
					"src/index.ts", "CardStill",
					"--props", props,
					"--output", outFile,
				)
				c.Dir = projectDir
				c.Stdout = os.Stdout
				c.Stderr = os.Stderr

				if err := c.Run(); err != nil {
					return fmt.Errorf("render card %d: %w", i+1, err)
				}
				fmt.Printf("  [%d/%d] %s\n", i+1, count, filepath.Base(outFile))
			}

			fmt.Printf("\nDone! PNGs saved to %s\n", outputDir)
			return nil
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (default: <project>/out/stills)")

	return cmd
}

func countCards(projectDir string) (int, error) {
	cardsDir := filepath.Join(projectDir, "src", "cards")
	entries, err := os.ReadDir(cardsDir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, e := range entries {
		name := e.Name()
		if len(name) > 4 && name[:4] == "Card" && filepath.Ext(name) == ".tsx" && name != "CardTemplate.tsx" {
			count++
		}
	}
	return count, nil
}
