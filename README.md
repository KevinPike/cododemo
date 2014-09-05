CoreOS / Docker - demo (cododemo)
======================

This repo is meant for a small CoreOS + Docker workshop that I'm presenting by collecting bits from all over. I will provide references when possible. I appreciate those who came first.

Requirements
------------

[Boot2docker](http://boot2docker.io/)

[VirtualBox](https://www.virtualbox.org/)

[Vagrant](https://www.vagrantup.com/)

[CoreOS](https://coreos.com/)

NOTE
----

[Project Atomic](http://www.projectatomic.io/) is from Red Hat and is also awesome. It's based on Fedora with some SELinux sprinkled in. It also has some infrastructure dashboard type stuff and it also borrows from Docker and other projects.


Install
-------

[Boot2docker for OSX Installer Downloads](https://github.com/boot2docker/osx-installer/releases)

[VirtualBox Downloads](https://www.virtualbox.org/wiki/Downloads) **I'm not certain you need to install virtualbox. It might have been installed when boot2docker was installed.

[Vagrant](https://www.vagrantup.com/downloads)

[CoreOS Install with Vagrant](https://coreos.com/docs/running-coreos/platforms/vagrant/)

Let's Boot2docker
-----------------

Just to get our docker feet wet.

```
boot2docker init
boot2docker start
export DOCKER_HOST=tcp://$(boot2docker ip 2>/dev/null):2375
```

The DOCs suggest that there is supposed to be a popup now. That the popup is a terminal window into the docker session. I'm pretty certain the hint is wrong. So ssh into the instance.

```
boot2docker ssh
```

You can ssh into and back out as often as you want. That's just the lightweight linux instance. It does not actually have any disk so there is no persistance and will not survive a reboot. Do not confuse the boot2docker host OS form the container.

now we are in a linux shell... so run the hello world container.

```
docker run hello-world
```

** one of the things that makes CoreOS nice is that it is immutable where it counts. This theme is extended into docker too. Once you create a container with a Dockerfile you should never change it. If you want to make a change then rebuild the container. As for saving data or persisting information; that is performed using volume mounts points or data-containers.

Deploy a 3 CoreOS cluster
-------------------------

Thw CoreOS/Vagrant installer includes sample files that need to properly configures.

```
cd ${HOME}/src/github.com/coreos/coreos-vagrant
cp config.rb.sample config.rb
cp user-data.sample user-data
```

edit the config.rb file and make these changes. The lines might need to be uncommented:

```
# Size of the CoreOS cluster created by Vagrant
$num_instances=3

# Official CoreOS channel from which updates should be downloaded
$update_channel='alpha'

# Enable port forwarding of Docker TCP socket
# Set to the TCP port you want exposed on the *host* machine, default is 2375
# If 2375 is used, Vagrant will auto-increment (e.g. in the case of $num_instances > 1)
# You can then use the docker tool locally by setting the following env var:
#   export DOCKER_HOST='tcp://127.0.0.1:2375'
$expose_docker_tcp=2375

# Setting for VirtualBox VMs
$vb_gui = false
$vb_memory = 1024
$vb_cpus = 1

```

The user-data is a little tricky. The file is modelled a after the cloud-config file.

The first thing you need is key for ETCD so that the cluster can identify itself as part of the cluster. Each core instance of etcd uses this to identify the cluster. (if you ```vagrant destroy``` the cluster you need a ```new``` key)

```
curl https://discovery.etcd.io/new
```

Take the return string and paste it into the user-data file. Notice that the file in the yml format.

```
discovery: https://discovery.etcd.io/<replace this with the from the step above>
```

Start the cluster. Notice that the first instance took a while to create. And the second two very quickly.

```
vagrant up
vagrant status
```

ssh into an instance.

```
vagrant ssh core-01
vagrant ssh core-02
vagrant ssh core-03
```

What is a Dockerfile
--------------------

tbd

What is a registry / private registry
-------------------------------------

tbd

** find a way to block writes to the public registry so that we are not leaking intellectual property


Docker Commands
---------------

build
run
commit

Docker Cleanup
--------------

containers (running or exit)
images
flatten

** watch out for persistent or data-containers.

Where is it?
------------

** Docker artifacts are stored here

** watch your disk usage

Docker Limits
-------------

there was once a 42 image limit for a container

devbox
------

tbd

Hello World webserver
---------------------

tbd

HA Hello World
--------------

[docs](https://coreos.com/docs/launching-containers/launching/launching-containers-fleet/)

Shipyard
--------

tbd

Redis Client and Server on the same instance
--------------------------------------------

tbd

Redis Client and Server on different instances
----------------------------------------------

tbd

Redis Server Failure and restart
--------------------------------

tbd

** docker auto restart can be dangerous, especially when combined with CoreOS.

Redis Server Failure and restart (ambassador)
---------------------------------------------

tbd

Redis Server Failure and fleet
------------------------------

tbd

Drone
-----

tbd


References
----------

So far everything I have written comes from one of the 4 sources I've already identified as requirements.

[Boot2docker](http://boot2docker.io/)

[VirtualBox](https://www.virtualbox.org/)

[Vagrant](https://www.vagrantup.com/)

[CoreOS](https://coreos.com/)

License
-------
No license is offered.
