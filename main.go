package main

import (
	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new team with the default settings.
		_, err := github.NewTeam(ctx, "Collaborators", &github.TeamArgs{
			CreateDefaultMaintainer: pulumi.Bool(false),
			Description:             pulumi.String("Members of Cloud Native Rioja"),
			Name:                    pulumi.String("Collaborators"),
			Privacy:                 pulumi.String("closed"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		return nil
	})
}
