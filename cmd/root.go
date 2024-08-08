package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gonew",
	Short: "Generate a simple go project",
	Long:  `Generate a simple go project: gonew project1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("project name is missing")
		}

		currentPath, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		projectName := args[0]
		fullPath := path.Join(currentPath, projectName)

		// create the folder name
		if err = createProjectFolder(fullPath); err != nil {
			log.Fatal(err)
		}

		// run go mod init
		if err = goMod(fullPath, projectName); err != nil {
			log.Fatal(err)
		}

		// run git init
		if err = gitInit(fullPath); err != nil {
			log.Fatal(err)
		}

		// touch README
		if err = readme(fullPath, projectName); err != nil {
			log.Fatal(err)
		}

		// touch CHANGELOG
		if err = changelog(fullPath, projectName); err != nil {
			log.Fatal(err)
		}

		// touch main.go
		if err = mainGo(fullPath, projectName); err != nil {
			log.Fatal(err)
		}
	},
}

func readme(fullPath string, projectName string) error {
	type Inventory struct {
		ProjectName string
	}

	sweaters := Inventory{projectName}
	tmpl, err := template.New("README").Parse("# {{.ProjectName}}")
	if err != nil {
		return err
	}

	// create a new file
	pathToReadme := path.Join(fullPath, "README.md")
	file, _ := os.Create(pathToReadme)
	defer file.Close()

	return tmpl.Execute(file, sweaters)
}

func changelog(fullPath string, projectName string) error {
	type Inventory struct {
		ProjectName string
	}

	sweaters := Inventory{projectName}
	tmpl, err := template.New("README").Parse("# CHANGELOG")
	if err != nil {
		return err
	}

	// create a new file
	pathToReadme := path.Join(fullPath, "CHANGELOG.md")
	file, _ := os.Create(pathToReadme)
	defer file.Close()

	return tmpl.Execute(file, sweaters)
}

func mainGo(fullPath string, projectName string) error {
	type Inventory struct {
		ProjectName string
	}

	sweaters := Inventory{projectName}
	tmpl, err := template.New("MAINGO").Parse("package main")
	if err != nil {
		return err
	}

	// create a new file
	pathToReadme := path.Join(fullPath, "main.go")
	file, _ := os.Create(pathToReadme)
	defer file.Close()

	return tmpl.Execute(file, sweaters)
}

func gitInit(fullPath string) error {
	command := exec.Command("git", "init")
	command.Dir = fullPath
	_, err := command.Output()
	return err
}

func createProjectFolder(fullPath string) error {
	return os.Mkdir(fullPath, os.ModePerm)
}

func goMod(projectFullpath string, projectName string) error {
	// go mod init github.com/<user>/<projectname>
	mod := os.Getenv("GONEW_GOMOD")

	if mod != "" {
		mod = fmt.Sprintf("%s/%s", mod, projectName)
	} else {
		mod = projectName
	}

	command := exec.Command("go", "mod", "init", mod)
	command.Dir = projectFullpath
	_, err := command.Output()
	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gonew.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
