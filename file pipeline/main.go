// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: JOY ADIKWU
// Squad: THE DEPLOYABLES

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: THE DEPLOYABLES
// ───────────────────────────────────────────
// Input line types:
//   - Normal report lines
//   - Lines in ALL CAPS
//   - Lines in all lowercase
//   - Lines starting with TODO:
//   - Lines starting with CLASSIFIED:
//   - Lines that are only dashes or blanks
//   - Lines with extra leading/trailing spaces
//   - Lines containing numbers and symbols
//
// Transformation rules (in order):
//   1. Trim all leading and trailing whitespace
//   2. Remove lines that are only dashes or blanks
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Add a line number prefix to every remaining line (001., 002., ...)
//
// Output format:
//   Header: YES — "SENTINEL FIELD REPORT — PROCESSED"
//   Line numbering format: 3-digit (001., 002., ...)
//   Summary block in file: YES — includes:
//     - Total lines processed
//     - Total lines removed
//
// Terminal summary fields:
//   ✦ Lines read
//   ✦ Lines written
//   ✦ Lines removed
//   ✦ Rules applied
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Stats structure
type Stats struct {
	read    int
	written int
	removed int
}

func main() {
	// Validate arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Check same file
	if inputFile == outputFile {
		fmt.Println("Input and output cannot be the same file.")
		return
	}

	// Check input exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", inputFile)
		return
	}

	// Check output is not a directory
	if info, err := os.Stat(outputFile); err == nil && info.IsDir() {
		fmt.Println("Cannot write to output: path is a directory, not a file.")
		return
	}

	lines, err := readLines(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	stats := Stats{read: len(lines)}

	// Handle empty file
	if len(lines) == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
		writeOutput(outputFile, []string{}, stats)
		printSummary(stats)
		return
	}

	processed, removed := processLines(lines)
	stats.removed = removed
	stats.written = len(processed)

	err = writeOutput(outputFile, processed, stats)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	printSummary(stats)
}

// Read file line by line
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Process pipeline
func processLines(lines []string) ([]string, int) {
	var processed []string
	removed := 0

	for _, line := range lines {
		line = trimWhitespace(line)

		if isRemovable(line) {
			removed++
			continue
		}

		line = replaceTODO(line)
		line = replaceClassified(line)

		processed = append(processed, line)
	}

	processed = addLineNumbers(processed)

	return processed, removed
}

// Rule 1: Trim whitespace
func trimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// Rule 2: Remove blank or dash lines
func isRemovable(s string) bool {
	if s == "" {
		return true
	}

	for _, ch := range s {
		if ch != '-' {
			return false
		}
	}
	return true
}

// Rule 3: Replace TODO
func replaceTODO(s string) string {
	return strings.ReplaceAll(s, "TODO:", "ACTION:")
}

// Rule 4: Replace CLASSIFIED
func replaceClassified(s string) string {
	return strings.ReplaceAll(s, "CLASSIFIED:", "[REDACTED]:")
}

// Rule 5: Add numbering
func addLineNumbers(lines []string) []string {
	var result []string

	for i, line := range lines {
		numbered := fmt.Sprintf("%03d. %s", i+1, line)
		result = append(result, numbered)
	}

	return result
}

// Write output file
func writeOutput(filename string, lines []string, stats Stats) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Header
	file.WriteString("SENTINEL FIELD REPORT — PROCESSED\n\n")

	// Lines
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	// Summary block in file
	file.WriteString("\n--- SUMMARY ---\n")
	file.WriteString(fmt.Sprintf("Lines processed: %d\n", stats.read))
	file.WriteString(fmt.Sprintf("Lines removed: %d\n", stats.removed))

	return nil
}

// Terminal summary
func printSummary(stats Stats) {
	fmt.Println("\n--- SUMMARY ---")
	fmt.Println("✦ Lines read    :", stats.read)
	fmt.Println("✦ Lines written :", stats.written)
	fmt.Println("✦ Lines removed :", stats.removed)
	fmt.Println("✦ Rules applied : Trim, Remove blanks, TODO replace, CLASSIFIED replace, Numbering")
}
