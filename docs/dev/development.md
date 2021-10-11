# Development Documentation

- [Development Documentation](#development-documentation)
  - [Commit Message Scope](#commit-message-scope)
  - [Workflow](#workflow)
    - [Naming Convention for Feature branches](#naming-convention-for-feature-branches)
  - [Makefile](#makefile)


## Commit Message Scope
Timer API Server adopts [Angular Commit style](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format).

Below are scopes used in Timer API Server for each commit type.

**chore**
- makefile: changes related to makefiles.
- lint: changes related to lint configs(e.g. `.golangci.yml`).
- migration: changes related to migration files.
- config: changes related to config files(e.g. `config/example.yml`).
- git: changes related to git(e.g. `.gitignore`).

**feat, fix, style, refactor, perf, test**
- resource: changes to `internal/resource`.
- rest: changes to `internal/restserver`.
- grpc: changes to `internal/grpcserver`.
- pkg: changes to `internal/pkg`.
- app: changes to `internal/app`.

**ci**
- circleci: changes related to circleci(e.g. `.circleci/config.yml`).
- docker: changes related to docker (e.g. `build/docker/Dockerfile`).

**docs**
- readme: changes related to readme files.
- swagger: changes related to swagger docs(e.g. `api/rest/swagger/swagger.yml`).
- comment: changes to comments.

**build**
- \<No specific scope>: changes to dependencies(e.g. `go.mod`)

## Workflow
Timer API server uses [Gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow).

To make changes, a feature branch should be checked out from branch `develop` as the branch to work on. A Pull Request that intends to merge the feature branch to branch `develop` should be created on Github and make sure the CircleCI check passes.

### Naming Convention for Feature branches
The scope of feature branches are encouraged to be small. The gist of [Angular Commit style](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-format) applies also to the naming of a feature branch. For example,

```
feat(app)/diable-reading-from-config-file-by-default
docs(readme)/add-overview-image
chore(makefile)/remove-mock-cleaning-in-phony-go-clean
```

## Makefile
[Makefile](../../Makefile) provides many handy phonies to automate tasks during development. For description of each phony, see the comment right above it. For a list of common top-level phonies, run `make help` at the root directory of the project.
