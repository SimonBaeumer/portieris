package notary

import (
	"context"
	"fmt"
	dockertrust "github.com/docker/cli/cli/trust"
	"github.com/docker/docker/api/types"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/golang/glog"
	notaryclient "github.com/theupdateframework/notary/client"
	"log"
	"os"
	"time"
)

func getNotaryRepo(user, pass, image string) (notaryclient.Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	os.Setenv("DOCKER_CONFIG", ".trust")
	glog.Infof("DOCKER IMAGE: %s", image)
	img, err := dockertrust.GetImageReferencesAndAuth(ctx, nil, func (ctx context.Context, index *registrytypes.IndexInfo) types.AuthConfig {
		fmt.Println("Get Image Reference")
		return types.AuthConfig{
			Username: user,
			Password: pass,
			Email: os.Getenv("REGISTRY_EMAIL"),
			ServerAddress: "https://index.docker.io/v1/",
		}
	}, image)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := dockertrust.GetNotaryRepository(os.Stdout, os.Stdin,
		"portieris",
		img.RepoInfo(),
		img.AuthConfig(),
		dockertrust.ActionsPullOnly...,
	)
	if err != nil {
		return nil, err
	}

	meta, err := repo.GetAllTargetMetadataByName(img.Tag())
	if err != nil {
		return nil, err
	}
	fmt.Printf("METADATA: %+v", meta)
	return repo, nil
}
