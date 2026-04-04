package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pureugong/mmk-cardnews-v2/internal/imagegen"
	"github.com/spf13/cobra"
)

func ImagegenCmd() *cobra.Command {
	var (
		prompt      string
		output      string
		aspectRatio string
		imageSize   string
		envFile     string
		batchFile   string
		concurrency int
	)

	cmd := &cobra.Command{
		Use:   "imagegen",
		Short: "Generate images using Gemini API",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiKey, model := loadEnv(envFile)
			if apiKey == "" {
				return fmt.Errorf("GEMINI_API_KEY not set. Set in %s or environment", envFile)
			}

			client := imagegen.NewClient(apiKey, model)

			// Batch mode: read prompts from JSON file
			if batchFile != "" {
				return runBatch(client, batchFile, aspectRatio, imageSize, concurrency)
			}

			// Single mode
			if prompt == "" {
				return fmt.Errorf("--prompt or --batch required")
			}
			if output == "" {
				return fmt.Errorf("--output required")
			}

			fmt.Printf("Generating image: %s\n", output)
			result, err := client.Generate(imagegen.GenerateOpts{
				Prompt:      prompt,
				OutputPath:  output,
				AspectRatio: aspectRatio,
				ImageSize:   imageSize,
			})
			if err != nil {
				return err
			}

			out, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(out))
			return nil
		},
	}

	cmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Image generation prompt")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file path")
	cmd.Flags().StringVar(&aspectRatio, "aspect-ratio", "3:4", "Aspect ratio (1:1, 16:9, 9:16, 4:3, 3:4)")
	cmd.Flags().StringVar(&imageSize, "image-size", "2K", "Image size (512, 1K, 2K, 4K)")
	cmd.Flags().StringVar(&envFile, "env-file", ".env", "Path to .env file")
	cmd.Flags().StringVar(&batchFile, "batch", "", "JSON file with batch prompts [{prompt, output}]")
	cmd.Flags().IntVar(&concurrency, "concurrency", 3, "Max parallel image generations")

	return cmd
}

type batchItem struct {
	Prompt string `json:"prompt"`
	Output string `json:"output"`
}

func runBatch(client *imagegen.Client, batchFile, aspectRatio, imageSize string, concurrency int) error {
	data, err := os.ReadFile(batchFile)
	if err != nil {
		return fmt.Errorf("read batch file: %w", err)
	}

	var items []batchItem
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("parse batch file: %w", err)
	}

	fmt.Printf("Generating %d images (concurrency: %d)\n", len(items), concurrency)

	sem := make(chan struct{}, concurrency)
	var mu sync.Mutex
	var results []imagegen.Result
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{}

		go func(idx int, it batchItem) {
			defer wg.Done()
			defer func() { <-sem }()

			result, err := client.Generate(imagegen.GenerateOpts{
				Prompt:      it.Prompt,
				OutputPath:  it.Output,
				AspectRatio: aspectRatio,
				ImageSize:   imageSize,
			})

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				fmt.Printf("  [%d/%d] FAIL %s: %v\n", idx+1, len(items), filepath.Base(it.Output), err)
				results = append(results, imagegen.Result{Success: false, Error: err.Error()})
			} else {
				status := "OK"
				if !result.Success {
					status = "FAIL"
				}
				fmt.Printf("  [%d/%d] %s %s\n", idx+1, len(items), status, filepath.Base(it.Output))
				results = append(results, *result)
			}
		}(i, item)
	}

	wg.Wait()

	success := 0
	for _, r := range results {
		if r.Success {
			success++
		}
	}
	fmt.Printf("\nComplete: %d/%d succeeded\n", success, len(items))
	return nil
}

func loadEnv(envFile string) (apiKey, model string) {
	// Check environment first
	apiKey = os.Getenv("GEMINI_API_KEY")
	model = os.Getenv("GEMINI_MODEL")

	// Then try .env file
	data, err := os.ReadFile(envFile)
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		val = strings.Trim(val, `"'`)

		switch key {
		case "GEMINI_API_KEY":
			if apiKey == "" {
				apiKey = val
			}
		case "GEMINI_MODEL":
			if model == "" {
				model = val
			}
		}
	}

	return
}
