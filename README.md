# github-org-iac

Repository to manage users and teams in the github organization with Pulumi and Golang. Everything are defined in the `config/org.yaml` file and automated with GitHub Actions. Follow the steps below to add a new user, team or repository.

## Adding a new user

1. Add the user to the map `members` list in `config/org.yaml`
2. Push the changes to GitHub
3. Create a Pull Request
4. Wait for the approval and merge

## Adding a new team

1. Add the team to the map `teams` list in `config/org.yaml`
2. Push the changes to GitHub
3. Create a Pull Request
4. Wait for the approval and merge

## Adding a new repository

1. Add the repository to the map `repositories` list in `config/org.yaml`
2. Push the changes to GitHub
3. Create a Pull Request
4. Wait for the approval and merge

## References

- [Pulumi GitHub Provider](https://www.pulumi.com/docs/intro/cloud-providers/github/setup/)
- [Pulumi GitHub Provider API Reference](https://www.pulumi.com/docs/reference/pkg/github/)
