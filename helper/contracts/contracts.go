package contracts

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const DefinitionsDirectory = "../contracts/"
const GoContractsDestination = "../be/services/contracts/"

func GenerateContracts(serviceName string) {
	filesToProcess := findFilesToProcess(serviceName)

	clearOutputDirectory(GoContractsDestination + serviceName)

	cmd := exec.Command(
		"protoc",
		"--proto_path="+DefinitionsDirectory,
		"--go_out="+GoContractsDestination,
		"--go_opt=paths=source_relative",
		"--go-grpc_out="+GoContractsDestination,
		"--go-grpc_opt=paths=source_relative",
		strings.Join(filesToProcess, " "),
	)

	out, err := cmd.CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated contracts definitions for " + serviceName + " service")
}

func findFilesToProcess(serviceName string) (contractFiles []string) {
	err := filepath.Walk(DefinitionsDirectory+serviceName, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		contractFiles = append(contractFiles, path)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return contractFiles
}

func clearOutputDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
}
