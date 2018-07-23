#!/bin/sh
# TODO(bentheelder): this is just a test script for installing from upstream
# sources... remove this

set -o errexit
set -x

# install kubeadm / kubelet / kubectl from upstream
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y --no-install-recommends kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl

echo "KUBELET_EXTRA_ARGS=--fail-swap-on=false" >> /etc/default/kubelet

systemctl show kubelet.service
echo "aaa"
cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
echo "blah"
cat /etc/systemd/system/kubelet.service.d/20-kubeadm.conf
echo "baz"
cat /etc/default/kubelet

# boot the master node
# --ignore-preflight-errors=Swap
kubeadm init --ignore-preflight-errors=all
