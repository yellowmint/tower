package contracts

import (
	"fmt"
	"github.com/otiai10/copy"
	"os"
	"os/exec"
	"strings"
)

const contractsDirectory = "contracts/"
const outputDirectory = contractsDirectory + "gen/"

const protoBuilderImage = "artdecoction.registry.jetbrains.space/p/wt/tools/proto-builder:0.0.1"

func Generate(service string, skipLint, skipBreaking bool) {
	if service == "all" {
		service = ""
	}

	lint(skipLint, service)
	breaking(skipBreaking)

	removeDirectory(outputDirectory)
	compile(service)
	changeGeneratedFilesOwnership()
	copyGeneratedFilesToServices(service)
}

func copyGeneratedFilesToServices(service string) {
	source := getProjectRootDir() + outputDirectory + "go/"
	destination := getProjectRootDir() + "be/services/contracts/"

	if service != "" {
		source += service + "/"
		destination += service + "/"
	}

	removeDirectory(destination)

	err := copy.Copy(source, destination)
	if err != nil {
		fmt.Printf("Copying generated files to services failed with: %s\n", err.Error())
		os.Exit(1)
	}
}

func lint(skipLint bool, service string) {
	if skipLint {
		fmt.Println("Skipping linting. Please do not do this in your final proto definitions shape.")
		return
	}

	out, err := bufCommand("lint", service).CombinedOutput()
	if err != nil {
		fmt.Printf("Linting failed with: %s\nCommand output:\n%s", err.Error(), out)
		os.Exit(1)
	}

	fmt.Println("Linting: OK")
}

func breaking(skipBreaking bool) {
	if skipBreaking {
		fmt.Println("Skipping breaking check. Please do not do this in your final proto definitions shape.")
		return
	}

	out, err := bufCommandBreaking().CombinedOutput()
	if err != nil {
		fmt.Printf("Breaking check failed with: %s\nCommand output:\n%s", err.Error(), out)
		os.Exit(1)
	}

	fmt.Println("Breaking: OK")
}

func compile(service string) {
	out, err := bufCommand("generate", service).CombinedOutput()
	if err != nil {
		fmt.Printf("Compiling failed with: %s\nCommand output:\n%s", err.Error(), out)
		os.Exit(1)
	}

	fmt.Println("Compiling: OK")
}

func changeGeneratedFilesOwnership() {
	out, err := bufCommandChangeGeneratedFilesOwnership().CombinedOutput()
	if err != nil {
		fmt.Printf("Changing generated files ownership failed with: %s\nCommand output:\n%s", err.Error(), out)
		os.Exit(1)
	}

	fmt.Println("Changing generated files ownership: OK")
}

func bufCommand(command, service string) *exec.Cmd {
	args := []string{
		"run",
		"--rm",
		"--volume",
		getProjectRootDir() + contractsDirectory + ":/workspace",
		protoBuilderImage,
		command,
	}

	if service != "" {
		args = append(args, service)
	}

	return exec.Command("docker", args...)
}

func bufCommandBreaking() *exec.Cmd {
	return exec.Command(
		"docker",
		"run",
		"--rm",
		"--volume",
		getProjectRootDir()+":/workspace",
		protoBuilderImage,
		"breaking",
		"./contracts/",
		"--against",
		".git#branch=main,subdir=contracts",
	)
}

func bufCommandChangeGeneratedFilesOwnership() *exec.Cmd {
	return exec.Command(
		"docker",
		"run",
		"--rm",
		"--volume",
		getProjectRootDir()+contractsDirectory+":/workspace",
		"--entrypoint",
		"chown",
		protoBuilderImage,
		"-R",
		fmt.Sprintf("%d:%d", os.Getuid(), os.Getgid()),
		"./gen/",
	)
}

func getProjectRootDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get current working directory")
		os.Exit(1)
	}

	return strings.ReplaceAll(workingDir, "/toolset-cli", "") + "/"
}

func removeDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Printf("Cannot remove direcotry %s, err: %e\n", dir, err)
		os.Exit(1)
	}
}
