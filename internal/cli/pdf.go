package cli

import (
	"fmt"
	"path/filepath"

	"github.com/pureugong/mmk-cardnews-v2/internal/pdf"
	"github.com/spf13/cobra"
)

func PdfCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "pdf <png-dir>",
		Short: "Stitch PNG card images into a PDF carousel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pngDir := args[0]

			if output == "" {
				output = filepath.Join(filepath.Dir(pngDir), "carousel.pdf")
			}

			fmt.Printf("Generating PDF from %s\n", pngDir)

			if err := pdf.GenerateFromImages(pngDir, output); err != nil {
				return fmt.Errorf("generate pdf: %w", err)
			}

			fmt.Printf("Done! PDF saved to %s\n", output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output PDF path (default: carousel.pdf)")

	return cmd
}
