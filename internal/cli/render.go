package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func RenderCmd() *cobra.Command {
	var (
		output      string
		composition string
		concurrency int
	)

	cmd := &cobra.Command{
		Use:   "render <project-dir>",
		Short: "Render card news video as MP4 using Remotion",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir := args[0]

			if _, err := os.Stat(filepath.Join(projectDir, "package.json")); err != nil {
				return fmt.Errorf("not a Remotion project: %s", projectDir)
			}

			if output == "" {
				output = filepath.Join(projectDir, "out", "output.mp4")
			}
			if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
				return err
			}

			fmt.Printf("Rendering video: %s\n", output)

			cmdArgs := []string{"remotion", "render",
				"src/index.ts", composition,
				"--output", output,
				fmt.Sprintf("--concurrency=%d", concurrency),
			}

			c := exec.Command("npx", cmdArgs...)
			c.Dir = projectDir
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if err := c.Run(); err != nil {
				return fmt.Errorf("remotion render: %w", err)
			}

			fmt.Printf("\nDone! Video saved to %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file (default: <project>/out/output.mp4)")
	cmd.Flags().StringVarP(&composition, "composition", "c", "CardNewsVideoClean", "Remotion composition ID")
	cmd.Flags().IntVar(&concurrency, "concurrency", 4, "Render concurrency")

	return cmd
}
