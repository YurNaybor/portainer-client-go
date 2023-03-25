SHELL := /bin/bash
HOST_IP := $(shell ip addr show dev eth0 | grep inet | awk '{print $$2}' | cut -d'/' -f1)

# Start Vagrant environment
.PHONY: vagrant
vagrant: vagrant-up

# Start Vagrant environment
.PHONY: vagrant-up
vagrant-up:
	vagrant up

# Provision Vagrant environment, e.g. (re)run provision scripts
.PHONY: vagrant-provision
vagrant-provision:
	vagrant provision

# Restart Vagrant environment
.PHONY: vagrant-reload
vagrant-reload:
	vagrant reload

# Enter Vagrant environment
.PHONY: vagrant-ssh
vagrant-ssh:
	vagrant ssh

# Remove Vagrant environment
.PHONY: vagrant-destroy
vagrant-destroy:
	vagrant destroy -f

# Setup integration testing environment
.PHONY: integration-test-prepare
integration-test-prepare:
	# Avoid errors and unnecessary redeployments by checking for state first
	if [[ "$(shell docker info --format '{{.Swarm.LocalNodeState}}')" == "inactive" ]]; then \
		docker swarm init; \
	fi
	if ! docker stack ps portainer >/dev/null 2>&1; then \
		docker stack deploy -c test/docker-compose.yml portainer; \
	fi
	until curl -s -o /dev/null -f -k https://localhost:9443; do echo "Waiting for portainer to be reachable"; sleep 10s; done
	@echo
	@echo "Portainer is running now. Please login to https://${HOST_IP}:9443 as admin and create an API token to proceed."
	@echo "To login, the password is '$(shell cat test/secrets/portainer-admin-password.txt)'"

# Remove integration testing environment
.PHONY: integration-test-cleanup
integration-test-cleanup:
	if docker stack ps portainer >/dev/null 2>&1; then \
		docker stack rm portainer; \
		docker volume rm portainer_portainer_data; \
	fi

# Run a dummy app
.PHONY: go-run-main
go-run-main:
	go run cmd/main.go

# Add a dependency
.PHONY: go-get
go-get:
	go get $(MODULE)

# Organize modules
.PHONY: go-mod-tidy
go-get:
	go mod-tidy
