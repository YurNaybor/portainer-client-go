# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "debian/bullseye64"

  config.vm.provision "shell", path: "vagrant-provision.sh"

  config.vm.provider :libvirt do | lv |
    lv.memory = 1024
    lv.cpus = 1
  end
end
