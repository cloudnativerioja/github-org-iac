package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v2"
)

// Parse the YAML file and return the Organization struct
type Team struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Context     string `yaml:"context"`
}
type Organization struct {
	Org   string `yaml:"org"`
	Teams []Team `yaml:"teams"`
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		//import accurate team info from yaml
		orgFilePath, err := filepath.Abs("./config/org.yaml")
		if err != nil {
			return err
		}
		yamlFile, err := ioutil.ReadFile(orgFilePath)
		if err != nil {
			return err
		}
		var org Organization
		err = yaml.Unmarshal(yamlFile, &org)
		if err != nil {
			return err
		}

		for _, team := range org.Teams {
			_, err := github.NewTeam(ctx, team.Context, &github.TeamArgs{
				Description: pulumi.String(team.Description),
				Name:        pulumi.String(team.Name),
				Privacy:     pulumi.String("closed"),
			}, pulumi.Protect(true))
			if err != nil {
				fmt.Println("encountered error creating new Pulumi GitHub team: ", team.Name)
				return err
			}
		}
		return nil
	})

}
