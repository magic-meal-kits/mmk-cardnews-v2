package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func PreviewCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "preview <project-dir>",
		Short: "Open Remotion Studio for live preview",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectDir := args[0]

			if _, err := os.Stat(filepath.Join(projectDir, "package.json")); err != nil {
				return fmt.Errorf("not a Remotion project: %s", projectDir)
			}

			url := fmt.Sprintf("http://localhost:%d", port)
			fmt.Printf("Starting Remotion Studio at %s\n", url)

			c := exec.Command("npx", "remotion", "studio",
				"--port", fmt.Sprintf("%d", port),
			)
			c.Dir = projectDir
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			// Open browser after a short delay
			go func() {
				openBrowser(url)
			}()

			return c.Run()
		},
	}

	cmd.Flags().IntVar(&port, "port", 3000, "Studio port")

	return cmd
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	}
	if cmd != nil {
		_ = cmd.Run()
	}
}
