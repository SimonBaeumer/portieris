package _dev

import (
	"context"
	"fmt"
	dockertrust "github.com/docker/cli/cli/trust"
	"github.com/docker/docker/api/types"
	registrytypes "github.com/docker/docker/api/types/registry"
	"log"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	//index := &registrytypes.IndexInfo{
	//	Name: "docker.io",
	//	Official: true,
	//	Secure: true,
	//}

	img, err := dockertrust.GetImageReferencesAndAuth(ctx, nil, func (ctx context.Context, index *registrytypes.IndexInfo) types.AuthConfig {
		fmt.Println("Get Image Reference")
		return types.AuthConfig{
			Username: os.Getenv("REGISTRY_USERNAME"),
			Password: os.Getenv("REGISTRY_PASSWORD"),
			Email: os.Getenv("REGISTRY_EMAIL"),
			ServerAddress: "https://index.docker.io/v1/",
		}
	}, "docker.io/library/nginx:1.14.2")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := dockertrust.GetNotaryRepository(os.Stdout, os.Stdin,
		"testing",
		img.RepoInfo(),
		img.AuthConfig(),
		dockertrust.ActionsPullOnly...,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo.GetGUN())
	meta, err := repo.GetAllTargetMetadataByName("1.20")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", meta)
}
