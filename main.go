package main

import (
	"io/ioutil"
	"log"

	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v2"
)

type Member struct {
	Username string `yaml:"username"`
	Role     string `yaml:"role"`
}

type Team struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Privacy     string   `yaml:"privacy"`
	Members     []Member `yaml:"members"`
}

type Repository struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type GitHubEntities struct {
	Repositories []Repository `yaml:"repositories"`
	Teams        []Team       `yaml:"teams"`
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		content, err := ioutil.ReadFile("config/org.yaml")
		if err != nil {
			log.Fatalln(err)
		}

		var data GitHubEntities
		if err := yaml.Unmarshal(content, &data); err != nil {
			log.Fatalln(err)
		}

		// Create Github repositories
		for _, repo := range data.Repositories {
			_, err := github.NewRepository(ctx, repo.Name, &github.RepositoryArgs{
				Description: pulumi.String(repo.Description),
				Name:        pulumi.String(repo.Name),
			})
			if err != nil {
				return err
			}
		}

		// Create Github teams and members
		for _, team := range data.Teams {
			newTeam, err := github.NewTeam(ctx, team.Name, &github.TeamArgs{
				Name:        pulumi.String(team.Name),
				Description: pulumi.String(team.Description),
				Privacy:     pulumi.String(team.Privacy),
			})
			if err != nil {
				return err
			}

			for _, member := range team.Members {
				_, err := github.NewTeamMembership(ctx, member.Username, &github.TeamMembershipArgs{
					Username: pulumi.String(member.Username),
					TeamId:   newTeam.ID(),
					Role:     pulumi.String(member.Role),
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
