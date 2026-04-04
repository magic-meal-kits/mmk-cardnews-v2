package pdf

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-pdf/fpdf"
)

const (
	defaultWidthMM  = 108.0 // 1080px at 254 DPI ≈ 108mm
	defaultHeightMM = 135.0 // 1350px at 254 DPI ≈ 135mm
)

func GenerateFromImages(pngDir, outputPath string) error {
	pngs, err := findPNGs(pngDir)
	if err != nil {
		return fmt.Errorf("find PNGs: %w", err)
	}
	if len(pngs) == 0 {
		return fmt.Errorf("no PNG files found in %s", pngDir)
	}

	// Get dimensions from first image
	w, h, err := imageDimensions(pngs[0])
	if err != nil {
		return fmt.Errorf("read image dimensions: %w", err)
	}

	// Scale to mm (assuming 254 DPI for 1080px)
	widthMM := float64(w) / 10.0
	heightMM := float64(h) / 10.0

	pdf := fpdf.NewCustom(&fpdf.InitType{
		UnitStr: "mm",
		Size:    fpdf.SizeType{Wd: widthMM, Ht: heightMM},
	})
	pdf.SetMargins(0, 0, 0)
	pdf.SetAutoPageBreak(false, 0)

	for _, png := range pngs {
		pdf.AddPage()
		pdf.Image(png, 0, 0, widthMM, heightMM, false, "PNG", 0, "")
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}

	return pdf.OutputFileAndClose(outputPath)
}

func findPNGs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var pngs []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".png" {
			pngs = append(pngs, filepath.Join(dir, e.Name()))
		}
	}
	sort.Strings(pngs)
	return pngs, nil
}

func imageDimensions(path string) (int, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}
