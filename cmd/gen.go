/*
Copyright © 2023 Omar Mouttaki <unknown989@proton.me>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

type TemplateType struct {
	url        string
	folder     string
	to_replace string
}

// The templates
var templates map[string]TemplateType = map[string]TemplateType{
	"default": {
		url:    "https://github.com/unknown989/webgen",
		folder: "templates/default",
		// Files to replace custom variables to their corresponding values
		// Like ${APP_NAME} is replaced to the app name you chosed
		to_replace: "index.html",
	},
	"react": {
		url:    "https://github.com/unknown989/webgen",
		folder: "templates/react",
		// Files to replace custom variables to their corresponding values
		// Like ${APP_NAME} is replaced to the app name you chosed
		to_replace: "src/App.jsx,package.json",
	},
	"react-ts": {
		url:    "https://github.com/unknown989/webgen",
		folder: "templates/react-ts",
		// Files to replace custom variables to their corresponding values
		// Like ${APP_NAME} is replaced to the app name you chosed
		to_replace: "src/App.tsx,package.json",
	},
}

// if err is an error, then kill the app, if not, do nothing
func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}

	return
}

// Checking if git exists
func checkGit() bool {
	// using --help because it works on linux and windows and it returns status code '0'
	_, err := exec.Command("git", "--help").Output()
	assert(err)

	return true
}

func cloneRepoAndFolder(url string, folder string, output_name string) bool {

	// Cloning the repo (webgen) to 'output_name'
	_, err := exec.Command("git", "clone", url, output_name).Output()

	assert(err)

	// Checking out only the required folders
	cmd := exec.Command("git", "sparse-checkout", "set", "--no-cone", folder)
	cmd.Dir = output_name
	_, err = cmd.Output()
	assert(err)
	// Removing the rest
	cmd = exec.Command("git", "checkout")
	cmd.Dir = output_name
	_, err = cmd.Output()
	assert(err)

	return true
}

func isDirectory(path string) bool {
	file_info, err := os.Stat(path)

	assert(err)
	return file_info.IsDir()
}

func handleFoldersLocation(folder string, output_name string) {
	var err error

	paths := make([]string, 0)

	// Default the backslash expected on the paths to unix-style
	var backslash byte = '/'
	// Changing the backslash to windows style '\'
	if runtime.GOOS == "windows" {
		backslash = '\\'
	}
	// Adding a backslash to the output_name if nonxistent
	var full_path string = output_name + folder
	if output_name[len(output_name)-1] != backslash {
		full_path = output_name + string(backslash) + folder
	}
	// Changing the windows path-ing (foo\bar\test.txt) system to unix(foo/bar/test.txt) (if applicable)
	full_path = strings.ReplaceAll(full_path, "\\", "/")

	// Walking the folder to find all the cloned files and their respective paths
	err = filepath.Walk(full_path, func(path string, _ os.FileInfo, err error) error {
		assert(err)
		paths = append(paths, strings.ReplaceAll(path, "\\", "/"))
		return nil
	})
	to_clean := []string{output_name + "/templates"}
	// Looping through the path to remove unnecessary files
	for i := range paths {
		curr_path := paths[i]

		new_path := strings.ReplaceAll(curr_path, folder+"/", "")
		// removing the parent folder from the list because 'walking the filesystem' adds it a walked path
		// for example walking '/foo/bar' return [ '/foo/bar', '/foo/bar/...' ]
		if curr_path == new_path {
			to_clean = append(to_clean, curr_path)
			continue
		}
		if isDirectory(curr_path) {
			assert(os.Mkdir(new_path, 0750))
			to_clean = append(to_clean, curr_path)
			continue
		}

		err = os.Rename(curr_path, new_path)
		assert(err)
	}
	for i := range to_clean {
		// reversing the order so that we go from the children to the parent
		// to solve the issue of non-empty folders
		var index int = len(to_clean) - 1 - i
		assert(os.Remove(to_clean[index]))
	}

	assert(err)
}

func replaceVarsToFile(app_name string, files []string, values map[string]string) {

	for i := range files {
		path := app_name + "/" + files[i]
		c, err := os.ReadFile(path)
		assert(err)
		content := string(c)
		for n, v := range values {
			content = strings.ReplaceAll(content, fmt.Sprintf("${%s}", strings.ToUpper(n)), v)
		}

		assert(os.WriteFile(path, []byte(content), 0750))
	}

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
			fmt.Println("Checking GIT...")
			checkGit()
			fmt.Println("Downloading the template...")
			cloneRepoAndFolder(template.url, template.folder, app_name)
			fmt.Println("Cleaning the template folder...")
			handleFoldersLocation(template.folder, app_name)

			fmt.Println("Replacing necessary content...")
			vars := map[string]string{
				"template": template_asked,
				"app_name": app_name,
			}
			replaceVarsToFile(app_name, strings.Split(template.to_replace, ","), vars)
			fmt.Println("Done!")
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
