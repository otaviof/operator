# Local Operator Development

## Prerequisites

* Ensure you have access to a Kubernetes cluster via `kubectl` with cluster admin permissions.
* Install Go version 1.15 or higher.

## Building locally

Run `make build` to compile the operator binary.
The resulting binary will be saved to `bin/operator`.

To build the container image for local use, run `make ko-publish IMAGE_REPO=ko.local`.
This will do the following:

1. Compile the application.
2. Create a container image.
3. Push the container image to your local Docker daemon, with the ref `ko.local/operator:<TAG>`.

The following make options can be set:

* `IMAGE_REPO` - defaults to `quay.io/shipwright`.
  The following special repositories can be used for testing:

  * `ko.local` - this pushes the image to the local Docker daemon.
  * `kind.local` - pushes to a local KinD cluster.

* `TAG` - defaults to `latest`.
* `IMAGE_PUSH` - if false, does not push the image. Defaults to true.

Refer to the [ko documentation](https://github.com/google/ko#local-publishing-options) for more information.

## Deploy to Kubernetes

To test the operator on a Kubernetes cluster, you first must have the following:

* Access to a Kubernetes cluster v1.19 or higher, with cluster admin permissions.
* Install Tekton v0.21 on the cluster.

```bash
$ export KUBECONFIG=/path/to/kubeconfig
$ kubectl apply -f https://github.com/tektoncd/pipeline/releases/download/v0.21.0/release.notags.yaml
```

If pushing to an external image registry, you may need to provide credentials to ko:

```bash
$ make ko
$ ko auth login <IMAGE_REGISTRY> -u <USERNAME> -p <PASSWORD>
```

Next, build the operator image as specified above, and push to a container registry that can be accessed by the cluster:

```bash
$ make build IMAGE_REPO="<IMAGE_REGISTRY>/<USERNAME>" TAG="<TAG>"
```

Finally, use the `make deploy` command with appropriate `IMAGE_REPO` and `TAG` arguments to deploy to the cluster.

```bash
$ make deploy IMAGE_REPO="<IMAGE_REGISTRY>/<USERNAME>" TAG="<TAG>"
```

Scripts in `hack` folder may require `sed` (GNU), therefore in platforms other than Linux you may have it with a different name. For instance, on macOS it's usually named `gsed`, in this case provide the `SED_BIN` make variable with the alternative name.

```bash
$ make build SED_BIN=gsed ...
```
