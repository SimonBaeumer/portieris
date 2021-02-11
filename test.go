package main

import (
	"fmt"
	"github.com/SimonBaeumer/portieris/helpers/oauth"
	notaryclient "github.com/SimonBaeumer/portieris/pkg/notary"
	//"github.com/SimonBaeumer/portieris/pkg/verifier/trust"
	"github.com/golang/glog"
	"log"
	"os"
)

func main() {
	pass := os.Getenv("REGISTRY_PASSWORD")
	user := os.Getenv("REGISTRY_USERNAME")
	writeAccess := false
	service := "notary"
	hostname := "notary.docker.io"
	repo := "docker.io/library/nginx"

	resp, err := oauth.Request(pass, repo, user, writeAccess, service, hostname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", resp)

	//var location string
	//location = "~/.kube/config"
	//kubeClientConfig := kube.GetKubeClientConfig(&location)
	//kubeClientset := kube.GetKubeClient(kubeClientConfig)
	//kubeWrapper := kubernetes.NewKubeClientsetWrapper(kubeClientset)

	trust, err := notaryclient.NewClient(".trust", []byte{})
	if err != nil {
		glog.Fatal("Could not get trust client", err)
	}


	//cr := registryclient.NewClient()
	//_ = notaryverifier.NewVerifier(kubeWrapper, trust, cr)
	//_ = notaryverifier.NewVerifier(kubeWrapper, nv, cr)
	notaryRepo, err := trust.GetNotaryRepo("https://index.docker.io", "docker.io/library/nginx", resp.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("GUN:", notaryRepo.GetGUN())
	fmt.Println("Metadata")

	fmt.Println(notaryRepo.GetAllTargetMetadataByName("1.14.2"))
	//notaryRepo.GetAllTargetMetadataByName()

}
