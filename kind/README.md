<!--TODO(bentheelder): fill this in much more thoroughly-->
# kind - **K**ubernetes **IN** **D**ocker

`kind` is a toolset for running local Kubernetes clusters using Docker container "nodes".

It consists of:
 - Go [packages](./pkg) implementing cluster creation, image build, etc.
 - A command line interface ([`kind`](./cmd/kind)) built on these packages
 - [`kubetest`](https://github.com/kubernetes/test-infra/tree/master/kubetest) integration also built on these packages (WIP)

Kind bootstraps each "node" with [Kubeadm](https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/).

## Usage

`kind create` will create a cluster

`kind delete` will delete a cluster

## Advanced

`kind build image` will build the node image

## Related work
