CoreOS / Docker - demo (cododemo)
======================

This repo is meant for a small CoreOS + Docker workshop that I'm presenting by collecting bits from all over. I will provide references when possible. I appreciate those who came first.

Preface
-------

[Link](http://www.ctlab.org/documents/How%20Complex%20Systems%20Fail.pdf) to an article from the University of Chicago entitled "How Complex Systems Fail (Being a Short Treatise on the Nature of Failure; How Failure is Evaluated; How Failure is Attributed to Proximate Cause; and the Resulting New Understanding of Patient Safety)". 

The author does not say anything about monolithic applications vs microservices and so it's a topic for the reader as to which system is more complex.

Warning
-----------

There is a lot of vaporware and junkware out there that is representing itself as orchestration or composition tools for CoreOS and/or Docker, however, while there is some effort involved in getting an environment into production it's almost better to execute and understand in the underlying scaffolding than to immediately defer to some opinionated framework.

Additionally, event though Docker is 1.X it is still being developed. Plenty of new features are being added regularly.

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
- Docker container [persistance](https://raw.githubusercontent.com/rbucker/cododemo/master/coreos%20docker%20volumes.jpg) is accomplished by using volumes or data-links
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

Install
-------

[Boot2docker for OSX Installer Downloads](https://github.com/boot2docker/osx-installer/releases)
   ** run the hello-world command

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
- preface (this document)
  - warning
  - opinionated
  - install install install
  - fun facts
- [CoreOS](https://github.com/rbucker/cododemo/blob/master/CoreOS.md)
  - major components discuss etcd(raft), fleetd, systemd
  - vagrant start CoreOS
  - play with etcd
  - tasks (add 4th node, reboot)
- [Docker](https://github.com/rbucker/cododemo/blob/master/Docker.md)
  - compare vm to container; what is a container
  - Dockerfile
  - registry
  - Docker commands
  - run hello-world (compare to boot2docker)
  - devbox
  - hello-world part 2
  - hello-world webserver
  - IP address
  - monitoring
  - backup
- [Cluster](https://github.com/rbucker/cododemo/blob/master/Cluster.md)
  - hello-world cluster
  - fleet
  - service/unit files
  - HA hello-world webserver
  - logging
  - sidekick
  - clustering tools like kubernetes sometimes called scheduling, composition or orchestration
- [Nats](https://github.com/rbucker/cododemo/blob/master/nats.md)
  - deploy the server
  - deploy 2 sub
  - deploy 1 pub
  - send a message
- [Ambassador](https://github.com/rbucker/cododemo/blob/master/Ambassador.md)
 

References
----------

too many to summarize here... links are provided in the text.

Tools
-----

[Vessel](http://awvessel.github.io/)

Kubernetes - [CoreOS example](https://coreos.com/blog/running-kubernetes-example-on-CoreOS-part-2/)

Panamax

cAdvisor

Shipyard + shipyard-agent


Contributors
------------

TBD

Thanks
------

Special thanks to the CoreOS team Alex, Alex, and Brian (from CoreOS) who peeked over my shoulder while I wrote this. :)

License
-------
No license is offered.
