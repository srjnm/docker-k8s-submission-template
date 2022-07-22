# docker-k8s-submission-template

This is a base template for CWiCS BOOST 2.0 Docker and Kubernetes.

This is set up in a way to utilize Github Actions for [CI to build and push](.github/workflows/docker-image-ci.yml) container images from Dockerfiles to the Github Container Registry.

## Initial Steps

- [Fork this repository](https://github.com/adyanth/docker-k8s-submission-template/fork)
- Go to the Actions tab and enable Actions by clicking on I understand my workflows
- Clone the repo or open the online editor and start commiting code!

## Naming Conventions

The container repository base will be `ghcr.io/<username>/<reponame>`. The tags will be the branch name. The main branch will also get the `latest` tag.

Use this naming convention while referring to images in the Kubernetes manifests.

Place all the Kubernetes manifests under the [`manifests/`](manifests) directory. Utilize your github username for namespaces.
