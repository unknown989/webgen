/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

type TemplateType struct {
	url        string
	folder     string
	to_replace string
}

var templates map[string]TemplateType = map[string]TemplateType{
	"default": {
		url:        "https://github.com/unknown989/webgen",
		folder:     "templates/default",
		to_replace: "index.html",
	},
	"react": {
		url:        "https://github.com/unknown989/webgen",
		folder:     "templates/react",
		to_replace: "src/App.jsx,package.json",
	},
	"react-ts": {},
}

func checkGit() error {
	_, err := exec.Command("git", "--help").Output()
	if err != nil {
		return err
	}
	return nil
}

func cloneRepoAndFolder(url string, folder string, output_name string) bool {
	_, err := exec.Command("git", "clone", url, output_name).Output()

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("git", "sparse-checkout", "set", "--no-cone", folder)
	cmd.Dir = output_name
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "checkout")
	cmd.Dir = output_name
	_, err = cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	var paths = []string{}

	err = filepath.Walk(folder, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, path)
		return nil
	})

	for p := range paths {
		fmt.Println("Type of p %T", p)
		// d := strings.Replace(p, "", "", 0)
	}

	if err != nil {
		log.Fatal(err)
	}

	return true

}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen - Generates the web app",
	Long: `the 'gen' commands handle the generation according to what template
	you specified which is usually one of these 'default', 'react', 'react-ts'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("False input, use it like this 'webgen gen TEMPLATE APP_NAME")
			fmt.Println("TEMPLATE can be one of these(more coming soon):")
			fmt.Println("	x default - normal legacy web app (HTML/CSS/JS)")
			fmt.Println("	x react - a JavaScript react web app")
			fmt.Println("	x react-ts - a TypeScript react web app")
			return
		}

		template_asked := args[0]
		app_name := args[1]
		if template, ok := templates[template_asked]; ok {
			err := checkGit()
			if err != nil {
				log.Fatal(err)
			}

			cloneRepoAndFolder(template.url, template.folder, app_name)

		} else {
			fmt.Println("Template does not exist, please use on of these:")
			fmt.Println("	x default - normal legacy web app (HTML/CSS/JS)")
			fmt.Println("	x react - a JavaScript react web app")
			fmt.Println("	x react-ts - a TypeScript react web app")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
