#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

set -o xtrace

# Disable IPv6, because got sporadic dns errors involving IPv6 addresses
echo "net.ipv6.conf.all.disable_ipv6 = 1
      net.ipv6.conf.default.disable_ipv6 = 1
      net.ipv6.conf.lo.disable_ipv6 = 1
      net.ipv6.conf.eth0.disable_ipv6 = 1" >> /etc/sysctl.conf
sysctl -p


# Install golang an other required tools
apt-get install -y golang-1.19 git make

# Set PATH for go
grep "go-1.19" /etc/bash.bashrc || echo 'export PATH=/usr/lib/go-1.19/bin/:$PATH' >> /etc/bash.bashrc

# Install Docker
# See https://docs.docker.com/engine/install/debian/
apt-get remove -y docker docker-engine docker.io containerd runc || true
apt-get install -y ca-certificates curl gnupg lsb-release
mkdir -m 0755 -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
chmod a+r /etc/apt/keyrings/docker.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update
apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
gpasswd -a vagrant docker
