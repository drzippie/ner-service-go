package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"ner-service-go/internal/config"
	"ner-service-go/internal/ner"
)

var (
	modelPath string
	inputFile string
	outputJSON bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ner-cli",
		Short: "Named Entity Recognition CLI for Spanish text",
		Long:  "A CLI tool to perform Named Entity Recognition on Spanish text using MITIE",
		Run:   runNER,
	}

	rootCmd.Flags().StringVarP(&modelPath, "model", "m", "", "Path to MITIE model file (default: models/ner_model.dat)")
	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file path (if not provided, reads from stdin)")
	rootCmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "Output in JSON format")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runNER(cmd *cobra.Command, args []string) {
	cfg := config.Load()
	if modelPath != "" {
		cfg.ModelPath = modelPath
	}

	nerService, err := ner.NewService(cfg.ModelPath)
	if err != nil {
		log.Fatalf("Failed to initialize NER service: %v", err)
	}
	defer nerService.Close()

	var text string
	if inputFile != "" {
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		text = string(data)
	} else if len(args) > 0 {
		text = args[0]
	} else {
		fmt.Println("Please provide text as argument or use --file flag")
		os.Exit(1)
	}

	entities, err := nerService.ExtractEntities(text)
	if err != nil {
		log.Fatalf("Error extracting entities: %v", err)
	}

	if outputJSON {
		jsonOutput, err := json.MarshalIndent(entities, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %v", err)
		}
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Printf("Found %d entities:\n\n", len(entities))
		for i, entity := range entities {
			fmt.Printf("%d. %s (%s) - Score: %s\n", i+1, entity.Label, entity.Tag, entity.Score)
		}
	}
}