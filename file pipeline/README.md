```markdown
 # The file pipeline
 It a tool that input and give you an output in the output.txt

 ## feautures
func main
func readLines
func processLines
func trimWhitespace
func isRemovable
func replaceTODO
func replaceClassified
func addLineNumbers
func writeOutput
func printSummary
## about
 Reads a text file, cleans and transforms its lines, and writes a standardized output with a summary.

Key Steps:

Validate Input/Output – checks arguments, file existence, avoids overwriting input.
Read Lines – reads the input file line by line.
Process Lines – applies 5 rules:
Trim whitespace
Remove blank or dash-only lines
Replace TODO: → ACTION:
Replace CLASSIFIED: → [REDACTED]:
Add line numbers (001. …)
Write Output – creates output.txt with header, processed lines, and file summary.
Print Terminal Summary – shows lines read, written, removed, and applied rules.

Edge Cases Handled: missing file, empty file, output is directory, input equals output, blank/dash lines, etc.
## usage

To run the code use: go run . input.txt output.txt 


       // CodeCrafters — Operation Gopher Protocol
       // Module: The pipline file
       // Author: [Joy Adikwu]
       // Squad:  [The Deployables]
