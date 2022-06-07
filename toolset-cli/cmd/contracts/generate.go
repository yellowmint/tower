package contracts

import (
	"fmt"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
)

const outputDirectory = "./contracts/gen/"
const protoBuilderImage = "artdecoction.registry.jetbrains.space/p/wt/tools/proto-builder:0.0.1"

func Generate(service string, skipLint, skipBreaking bool) {
	lint(skipLint, service)
	//runBufCommand("breaking --against .git#branch=main")

	//removeDirectory(outputDirectory)
	//runBufCommand("generate")
	//removeDirectory(outputDirectory)
}

func lint(skipLint bool, service string) {
	if skipLint {
		fmt.Println("Skipping linting. Please do not do this in your final proto definitions shape.")
		return
	}

	out, err := bufCommand("lint", service).CombinedOutput()
	if err != nil {
		fmt.Printf("Linting failed with err: %e\nCommand output:\n%s", err, out)
		os.Exit(1)
	}

	fmt.Println("Linting: OK")
}

func bufCommand(command, service string) *exec.Cmd {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Unable to create docker client")
		os.Exit(1)
	}

	//return exec.Command(
	//	"docker",
	//	"run",
	//	"--rm",
	//	`--volume /home/jack/art/wt/tower/toolset-cli:/workspace`,
	//	protoBuilderImage,
	//	command,
	//	service,
	//)
}

func removeDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Printf("Cannot remove direcotry %s, err: %e\n", dir, err)
		os.Exit(1)
	}
}
