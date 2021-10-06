package main

import (
	"docker/internal/api"
	"docker/internal/process"
	"docker/internal/types"
	"docker/internal/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
)

const (
	TempDirPrefix   = "docker-*"
	RegistryAddress = "https://registry.hub.docker.com"
)

func checkArgs(args []string) {
	if len(args) < 4 {
		log.Fatal("Not enough arguments passed")
	}

	if args[1] != "run" {
		log.Fatal("The first argument must be \"run\"")
	}
}

func main() {
	checkArgs(os.Args)

	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	imageName, imageTag := util.ImageAndTag(os.Args[2])

	client := &http.Client{}

	auth, err := api.Authenticate(&types.RegistryAuthConfig{
		ServerAddress: "https://auth.docker.io",
		Service:       "registry.docker.io",
		Scope:         fmt.Sprintf("repository:library/%s:pull", imageName),
	}, client)
	util.CheckError(err)

	manifest, err := api.PullManifest(&types.ManifestConfig{
		RegistryAddress: RegistryAddress,
		ImageName:       imageName,
		ImageReference:  imageTag,
	}, auth, client)
	util.CheckError(err)

	tempDir, err := ioutil.TempDir("", TempDirPrefix)
	util.CheckError(err)

	for _, layer := range manifest.FsLayers {
		layerResponse, err := api.PullLayer(&types.LayerConfig{
			RegistryAddress: RegistryAddress,
			ImageName:       imageName,
			Digest:          layer.BlobSum,
		}, auth, client)
		util.CheckError(err)

		layerFileName := filepath.Join(os.TempDir(), fmt.Sprintf("%s.tar.gz", layer.BlobSum))
		err = ioutil.WriteFile(layerFileName, layerResponse.Content, 0700)
		util.CheckError(err)

		err = util.Untargz(layerFileName, tempDir)
		util.CheckError(err)
	}

	err = syscall.Chroot(tempDir)
	util.CheckError(err)

	err = syscall.Chdir("/")
	util.CheckError(err)

	cmd, err := process.Exec(command, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cmd.ProcessState.ExitCode())
	}
}
