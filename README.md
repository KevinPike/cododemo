CoreOS / Docker - demo (cododemo)
======================

This repo is meant for a small CoreOS + Docker workshop that I'm presenting by collecting bits from all over. I will provide references when possible. I appreciate those who came first.

Warning
-----------

There is a lot of vaporware and junkware out there that is representing itself as orchestration or composition tools for CoreOS and/or Docker, however, while there is some effort involved in getting an environment into production it's almost better to execute and understand in the underlying scaffolding than to immediately defer to some opinionated framework.

Opinionated Preface
-------------------

When considering both CoreOS and Docker; they are viewed as codependent opinionated environments. Here are some nuggest:
- CoreOS should be installed on bare metal
- CoreOS should be the only OS on the machine if a hypervisor is present for ease of mgmt
- CoreOS is mostly immutable; there are places to store user apps but there is no package manager ...
- CoreOS wants user apps to run on docker and user their systemd, fleetd, etcd ecosystem (more tools coming)
- CoreOS does not ship with perl, python or ruby. THANK GOODNESS!
- Docker containers are meant to look like a standalone machine but use the hosts OS kernel
- Docker containers are immutable
- Docker container persistance is accomplished by using volumes or data-links
- Docker wants one process per container
- Docker containers require links in order to communicate between peers
- Docker links have behaviors (see link ambassadors)
- Docker wants you to build, test and deploy the same container instance
- each Docker instance refers to itself as localhost and 127.0.0.1 (new)
- Docker wants you to run your apps in the container foreground

** neither CoreOS nor Docker are going to create a transparent environment or experience. Getting to that level would potentially compromise security or create operational issues that the environments intend to prevent (FUD)

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


Fun fact about CoreOS
---------------------

Looking at the [release](https://coreos.com/releases/) page at CoreOS you'll that the latest alpha version is 423.0.0.  When the alpha version is promoted to beta or stable it is that exact image that is promoted. There is no additional build that takes place.

Fun fact about Docker
---------------------

Docker wants you to do the same thing.  Build the container, test the container, move the container to the next stage in the pipeline until it get's to production.


Table of Contents
-----------------

- [CoreOS](https://github.com/rbucker/cododemo/blob/master/CoreOS.md)
- [Docker](https://github.com/rbucker/cododemo/blob/master/Docker.md)
- [Cluster](https://github.com/rbucker/cododemo/blob/master/Cluster.md)




References
----------

So far everything I have written comes from one of the 4 sources I've already identified as requirements.

[Boot2docker](http://boot2docker.io/)

[VirtualBox](https://www.virtualbox.org/)

[Vagrant](https://www.vagrantup.com/)

[CoreOS](https://coreos.com/)

Stuff
-----

##### CoreOS
- What is it?
- Where does it fit?
- Compared to ChromeOS?
- How is it managed?
- How/where is it deployed?
- Windows, Mac
- VMware, VirtualBox
- channels (alpha, beta, stable)
- releases
- upgrades, rollback

##### CoreOS Enterprise
- Managed Linux
- Managed Cluster
- Core Update
- Enterprise Registry

##### CoreOS Idioms
- TBD

##### Docker
- What is it?
- Where does it fit?
- Compared to VM?
- How is it managed?
- How is it deployed?
- Where is it deployed?
- CoreOS, ProjectAtomic, Other Linux
- persistence
- linked data
- linked network
- ambassadors
- composition vs orchestration

##### Docker Idioms
- TBD

##### Build
- travis-ci
- drone.io
- circleci

##### Frameworks
- kubernetes
- mesos, marathon
- deis
- flynn
- fig
- boot2docker
- panamax
- VMware docker
- OpenStack Docker

#### Tools
- datadog
- chef, puppet, ansible, saltstack
- elastic beanstalk
- dokku
- shipyard

##### Benefits
- same container through pipeline
- microservices
- TCO

##### Risks
- microservices
- J2EE
- visualization
- lxc vs libcontainer
- CoreOS vs ESXi

Thanks
------

Special thanks to the CoreOS team Alex, Alex, and Brian (from CoreOS) who peeked over my shoulder while I wrote this. :)

License
-------
No license is offered.
