package cmd

import (
	"fmt"
	"github.com/short-d/app/fw/cli"
	"github.com/short-d/fwcli/git"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const serviceTemplateURL = "https://github.com/short-d/app-template.git"
const defaultAppName = "sampleapp"

func newNew(factory cli.CommandFactory) cli.Command {
	config := cli.CommandConfig{
		Usage:        "new [service_name]",
		ShortHelpMsg: "Generate a new project at current directory",
		OnExecute: func(cmd cli.Command, args []string) {
			if len(args) < 1 {
				err := cmd.Help()
				if err != nil {
					crash(err)
				}
				return
			}

			serviceName := args[0]
			cloneTemplate(serviceName)
			updateAppName(serviceName)
			initializeRepo(serviceName)
			installDeps()
		},
	}
	return factory.NewCommand(config)
}


func initializeRepo(serviceName string) {
	fmt.Println("Initializing repository")
	gitPath := filepath.Join(serviceName, ".git")
	err := os.RemoveAll(gitPath)
	if err != nil {
		crash(err)
		return
	}

	err = git.InitializeRepo(serviceName)
	if err != nil {
		crash(err)
		return
	}

	err = os.Chdir(serviceName)
	if err != nil {
		crash(err)
		return
	}

	err = git.StageAllChanges()
	if err != nil {
		crash(err)
		return
	}

	err = git.Commit("Initialize repository")
	if err != nil {
		crash(err)
		return
	}
}

func installDeps() {
	installFrontendDeps()
	installBackendDeps()
}

func installFrontendDeps() {
	fmt.Println("Installing frontend dependencies")
	err := os.Chdir("web")
	if err != nil {
		crash(err)
		return
	}

	command := exec.Command("yarn")
	err = command.Run()
	if err != nil {
		crash(err)
		return
	}

	err = os.Chdir("../")
	if err != nil {
		crash(err)
		return
	}
}

func installBackendDeps() {
	fmt.Println("Installing backend dependencies")
	err := os.Chdir("backend")
	if err != nil {
		crash(err)
		return
	}

	command := exec.Command("go", "mod", "download")
	err = command.Run()
	if err != nil {
		crash(err)
		return
	}

	err = os.Chdir("../")
	if err != nil {
		crash(err)
		return
	}
}

func updateAppName(serviceName string) {
	fmt.Println("Updating service name in .drone.yml")
	droneConfigPath := filepath.Join(serviceName, ".drone.yml")
	err := replaceString(droneConfigPath, defaultAppName, serviceName)
	if err != nil {
		crash(err)
		return
	}

	fmt.Println("Updating service name in Kubernetes configs")
	k8sConfigPath := filepath.Join(serviceName, "k8s")
	err = replaceString(k8sConfigPath, defaultAppName, serviceName)
	if err != nil {
		crash(err)
	}

	fmt.Println("Updating go module name")
	backendCodePath := filepath.Join(serviceName, "backend")
	err = replaceString(backendCodePath, "github.com/short-d/app-template/backend", serviceName)
	if err != nil {
		crash(err)
	}
}

func replaceString(dir string, original string, new string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		content := strings.ReplaceAll(string(buf), original, new)
		return ioutil.WriteFile(path, []byte(content), info.Mode())
	})
}

func cloneTemplate(serviceName string) {
	fmt.Printf("Cloning template to %s\n", serviceName)
	if fileExists(serviceName) {
		fmt.Printf(
			`%s already exists.
Please first remove it or choose a different service name.
`,
			serviceName)
		return
	}
	err := git.CloneRepo(serviceTemplateURL, serviceName)
	if err != nil {
		crash(err)
	}
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func crash(err error) {
	fmt.Println(err)
	os.Exit(1)
}
