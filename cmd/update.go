package cmd

import (
	"bufio"
	"os"
	"strings"
	"text/template"

	"github.com/go-playground/log"
	"github.com/spf13/cobra"
)

var templateFile string
var outFile string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Replace the template file's place holders '{{.SAMPLE}}' with Environment Variables.",
	Long: `Replace the template file's place holders '{{.SAMPLE}}' with Environment Variables.
For example:

goenvtemplatemerge update -t sampleTemplate.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Working with template %+v", templateFile)
		log.Debug("loading environment variables")
		log.Infof("Loading %+v environment variables", len(os.Environ()))

		// Load the environment variables into a map
		mapOfLoadedEnvVariables := loadEnvVars()
		log.Info("Updating template with loaded environnent variables")

		// Updates the template with the loaded environment variables
		updateTemplate(mapOfLoadedEnvVariables)

		// Validates the template
		validateTemplate()
		log.Infof("Template %+v updated successfuly", templateFile)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVarP(&templateFile, "templatefile", "t", "", "Path to the template. (REQUIRED)")
	updateCmd.MarkPersistentFlagRequired("templatefile")
}

// loadEnvVars: loads the Environment Variables and returns them in a map.
func loadEnvVars() map[string]string {
	m := make(map[string]string)
	var totalLoadedVariables int
	log.Debug("Going through the list of env variables")
	log.Info("Loading environment variable: ")
	for i, loadedVariable := range os.Environ() {
		variablePair := strings.SplitN(loadedVariable, "=", 2)
		m[variablePair[0]] = variablePair[1]
		log.Infof("%+v. %+v", i+1, variablePair[0])
		log.Debugf("Writing env variable %+v with value %+v", variablePair[0], variablePair[1])
		totalLoadedVariables = i
		log.Debugf("Environment Variables loaded: %+v", totalLoadedVariables)
	}
	return m
}

// updateTemplate: Uses the map of environment variables to replace the placeholders inside the template.
func updateTemplate(mapOfLoadedEnvVariables map[string]string) {
	// create template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Error while parsing template %+v", err)
	}
	f, err := os.Create(templateFile)
	if err != nil {
		log.Fatalf("create file: %+v", err)
		return
	}
	err = tmpl.Execute(f, mapOfLoadedEnvVariables)
	if err != nil {
		log.Fatalf("Error while parsing template %+v", err)
	}
	if debug {
		tmpl.Execute(os.Stdout, mapOfLoadedEnvVariables)
	}
}

// validateTemplate: Verifies the updated template does not have <no value> (Environment Variable not found for placeholder).
func validateTemplate() {
	var lineNotUpdated []int
	fileToCheck, err := os.Open(templateFile)
	if err != nil {
		log.Fatalf("Error opening file: %+v\n%+v", templateFile, err)
	}
	defer fileToCheck.Close()
	scanner := bufio.NewScanner(fileToCheck)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<no value>") {
			lineNotUpdated = append(lineNotUpdated, lineNumber)
		}
		lineNumber++
	}
	if len(lineNotUpdated) > 0 {
		log.Fatalf("Missing environment variables in lines: \n%+v", lineNotUpdated)
	}
}
