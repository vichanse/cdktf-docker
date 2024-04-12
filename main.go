package main

import (
	"cdk.tf/go/stack/generated/kreuzwerker/docker/container"
	"cdk.tf/go/stack/generated/kreuzwerker/docker/image"
	"cdk.tf/go/stack/generated/kreuzwerker/docker/network"
	docker "cdk.tf/go/stack/generated/kreuzwerker/docker/provider"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	backendPort  = 8080
	frontendPort = 8000
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Initialize the Docker provider
	docker.NewDockerProvider(stack, jsii.String("docker"), &docker.DockerProviderConfig{})

	// Pull the Backend image
	backendImage := image.NewImage(stack, jsii.String("backendImage"), &image.ImageConfig{
		Name:        jsii.String("scraly/backend-docker:1.0.0"),
		KeepLocally: jsii.Bool(false),
	})

	// Pull the Frontend Watcher image
	frontendImage := image.NewImage(stack, jsii.String("frontendImage"), &image.ImageConfig{
		Name:        jsii.String("scraly/frontend-docker:1.0.1"),
		KeepLocally: jsii.Bool(false),
	})

	// Create a Docker network to allows our containers to comunicate to each others
	gophersNetwork := network.NewNetwork(stack, jsii.String("my_network"), &network.NetworkConfig{
		Name: jsii.String("my_network"),
	})

	// Create the backend container
	container.NewContainer(stack, jsii.String("backendContainer"), &container.ContainerConfig{
		Image: backendImage.Name(),
		Name:  jsii.String("backend"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(backendPort), External: jsii.Number(backendPort),
		}},
		NetworksAdvanced: &[]*container.ContainerNetworksAdvanced{{
			Name:    gophersNetwork.Name(),
			Aliases: jsii.Strings(*jsii.String("my_network")),
		}},
	})

	// Create the frontend container
	container.NewContainer(stack, jsii.String("frontendContainer"), &container.ContainerConfig{
		Image: frontendImage.Name(),
		Name:  jsii.String("frontend"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(frontendPort), External: jsii.Number(frontendPort),
		}},
		NetworksAdvanced: &[]*container.ContainerNetworksAdvanced{{
			Name:    gophersNetwork.Name(),
			Aliases: jsii.Strings(*jsii.String("my_network")),
		}},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-docker")

	app.Synth()
}