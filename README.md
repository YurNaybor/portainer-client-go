# Go Client for Portainer API

A small wrapper around existing Portainer API / structures to provide a Golang
API for Portainer with focus on Docker (Swarm) management.


## Background

Portainer's API sadly does not not provide structures for stuff like Secrets or Networks, which are platform specific (e.g. such requests are passed to the platform API, like Docker). To have a generic API, I wanted to provide those structures.

After all, this module is driven by the implementation of the Terraform Portainer Provider.


## Current state

The module is in a very early state. Also, my Golang skills are basic, to be frank. For now, we've got the development environment and some basic read operations, to get data from the Portainer API.


## TODO

- Stacks from Git with webhook, environment variables pointing to secrets
  -> it appears secrets are to be managed via Docker api passtrough -> endpoints docker secrets create, same for networks
- Secrets with generated names
- Networks (macvlan, overlay)
- Implement automated tests (generate api token automatically..)


## Development

The development environment is based on Vagrant (with libvirt provider) and uses Make for some automation. Within Vagrant, Docker is installed and a Docker Swarm initialized. Then, Portainer is deployed as a Swarm Stack. This is the (integration) test environment. An API token for Portainer must be created manually, for now.


```bash
# Setup the development environment
$ make vagrant
$ vagrant ssh
~ cd /vagrant
# Now do your development

```
