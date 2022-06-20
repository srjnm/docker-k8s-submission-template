# docker-k8s-submission-template

This is a base template for CWiCS BOOST 2.0 Docker and Kubernetes.

This is set up in a way to utilize Github Actions for [CI to build and push](.github/workflows/docker-image-ci.yml) container images from Dockerfiles to the Github Container Registry.

## Naming Conventions

The container repository base will be `ghcr.io/<username>/<reponame>`.

The tag conventions are as follows:

- [`Dockerfile`](Dockerfile) will be built and pushed with `latest` and `<branchname>` tags
- [`Dockerfile.service`](Dockerfile.service) will be built and pushed with `latest-<service>` and `<branchname>-<service>` tags
- [`module/Dockerfile`](module/Dockerfile) will be built and pushed with `<module>-latest` and `<module>-<branchname>` tags

Use this naming conventions while referring to images in the Kubernetes manifests.

Place all the Kubernetes manifests under the [`manifests/`](manifests) directory.
