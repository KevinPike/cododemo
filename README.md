CoreOS / Docker - demo (cododemo)
======================

This repo is meant for a small CoreOS + Docker workshop that I'm presenting by collecting bits from all over. I will provide references when possible. I appreciate those who came first.

Warning
-----------

There is a lot of vaporware and junkware out there that is representing itself as orchestration or composition tools for CoreOS and/or Docker, however, while there is some effort involved in getting an environment into production it's almost better to execute and understand in the underlying scaffolding than to immediately defer to some opinionated framework.

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

Just to get our docker feet wet. It has no persistance unless your work uploads or saves it's content directly. Once the container is gone; it is gone.

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

Fun fact about CoreOS
---------------------

Looking at the [release](https://coreos.com/releases/) page at CoreOS you'll that the latest alpha version is 423.0.0.  When the alpha version is promoted to beta or stable it is that exact image that is promoted. There is no additional build that takes place.

Fun fact about Docker
---------------------

Docker wants you to do the same thing.  Build the container, test the container, move the container to the next stage in the pipeline until it get's to production.

CoreOS tools
------------

- CoreOS Cluster is defined by 3 or more instances.
- etcd - replicated key/value store using the raft protocol.
- fleetd - cluster manager
- journald - aggregated logging
- systemd - startup
  - cron-like scheduling
  - cloud-config
- other logging
- locksmith - upgrades
- rudder - dynamic networks

Deploy a 3 CoreOS cluster
-------------------------

The CoreOS/Vagrant installer includes sample files that need to properly configures.

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

What is Docker?
--------------

- docker is the wrapper for the toolchain that orchestrates everything between the docker container(s) and the host OS.
- docker containers/images are described as specialized tarfiles that are immutable and represent the userspace needed for the app that runs in the container
- the container should only run one process per container instance
- communication between containers is done with ```links```
- since containers are immutable persistence is implemented through ```volumes``` or ```links``` to storage-containers
- storage-containers that are not persisted via a ```volume``` will be lost when the reference count is zero
- there are 3rd party tools for backing up containers and moving them around the cluster; and then there is the registry
- 

Docker tools
------------

- registry, private registry (docker hub)
- gitreceive
- libcontainer - replacement for the lxc container
- libswarm - orchestration APIs
- libchan - golang RPC replacement for remote channel calls
- fig
- gordon


What is a Dockerfile
--------------------

A ```Dockerfile``` is a docker container configuration file that can be considered similar to a ```Makefile```. The Dockerfile is used by the docker CLI tool in order to construct the container instance called an ```image```. The format of the Dockerfile is defined by its' DSL. As each step is performed an image is captured on the host OS and a signature assigned. (a random name is also assigned; and when the final step is completed the user defined name is assigned to the last image as an alias). All of the intermediate images remain unless the ```rm``` flag is applied to the ```build``` command.

** each image created by a Dockerfile is called a layer. Historically there isa limit to 42 layers in a single Dockerfile. The good news is that are ways to combines tasks. (think about ```apt-get``` multiple packages at once)

The docker images are stored here: ```/var/lib/docker/graph/<id>/layer```

There is plenty of discussion suggesting that Dockerfiles devalue chef, puppet, ansible, saltstack.

[docs](https://docs.docker.com/reference/builder/)

What is a registry / private registry
-------------------------------------

The Dockerfile is held locally but can be stored in a public or private registry. There are also public and private service providers who will autobuild your container from your docker file so that it's always ready.

(see the docker ```commit``` command)

** find a way to block writes to the public registry so that we are not leaking intellectual property

The registry.hub.docker.com is full of projects. There are many more community contributions than there are curated. ```Stackbrew``` is the username assigned to the docker team so their images can be trusted (don't take my word for it). There are some other users that are considered "trusted" but I'm not certain about the certification process. I prefer private repositories but that takes planning and storage.


Docker Commands
---------------

```build``` - Build a new image from the source code (Dockerfile). Each task in the Dockerfile creates a separate image file. Using the ```-rm``` flag deletes the intermediate images saving space but subsequent builds will take longer.

```run``` - Run a command in a new container. If you want the command to run in the background then you need to set the interactive flag.

```commit``` - Create a new image from a container's changes (save it in the repo)

```stop``` - Stop a running container by sending SIGTERM and then SIGKILL after a grace period. (assuming that you initiated a command by calling the ```run``` command with the interactive flag.

```start``` - Restart a ```stop```ped container. You cannot start a command that has exited. (docker provides a restart flag) Auto-restart must be considered carefully when working with CoreOS.

Docker Cleanup
--------------

I ran my hello world command multiple times. Something like. Since the command was just printing out hello and some other text before exiting; container is in a terminal state. Not running. Running the ```docker ps``` command will return an empty list. However, after running the command ```docker run hello-world``` multiple times the docker container did leave some breadcrumbs.

```
$ docker ps -a
CONTAINER ID        IMAGE                COMMAND             CREATED             STATUS                         PORTS               NAMES
20f0c6dd1dc9        hello-world:latest   "/hello"            15 seconds ago      Exited (0) 14 seconds ago                          berserk_wright      
c0a8f556973d        hello-world:latest   "/hello"            21 minutes ago      Exited (0) 21 minutes ago                          angry_babbage       
6df23e68e893        hello-world:latest   "/hello"            About an hour ago   Exited (0) About an hour ago                       naughty_bohr        
2306f2880ff4        hello-world:latest   "/hello"            About an hour ago   Exited (0) About an hour ago                       boring_nobel        
5086487ecf20        hello-world:latest   "/hello"            2 hours ago         Exited (0) 2 hours ago                             naughty_poincare    
4e73bb80cc8d        hello-world:latest   "/hello"            2 hours ago         Exited (0) 2 hours ago                             jolly_bell    
```

Remove all stopped containers: ```docker rm $(docker ps -a -q)```

Removed all untagged images: ```docker rmi $(docker images | grep "^<none>" | awk "{print $3}")```

** watch out for persistent or data-containers. As I've already said if the reference count reaches zero the data will be lost.

Where is it?
------------

Docker artifacts are stored here ```/var/lib/docker```

watch your disk usage... it's easy to fill up the host drive.

Docker Limits
-------------

there was once a 42 image limit for a container. At some point you have to flatten the layers.

Hello World
-----------

```
vagrant ssh core-01
docker run hello-world
```

notice that the output from this hello-world is the exact same as the boot2docker version. That's because they are the same container image constructed (built) from the same Dockerfile. [Here](https://registry.hub.docker.com/u/library/hello-world/) is the registry where the hello-world image lives.

devbox
------

[shykes/devbox](https://registry.hub.docker.com/u/shykes/devbox/) - is a container the allows the user to create a proper development environment. (build the container, run the container with a command, execute your shell commands etc...) He has a [link](https://github.com/shykes/devbox) to the github source.

rbucker/devbox - I have created a Dockerfile with a little more tooling and some documentation on bitbucket [here](https://bitbucket.org/rbucker/devbox). The different commands and dependencies are included.


create a host user with the same uid and gid as the container

boot2docker (broken because docker already uses the uid that I was intending to use in my devbox)
```
adduser 

```

CoreOS
```
sudo groupadd -g 1000 dev
sudo useradd -d /home/dev -g 1000 -m -s /bin/bash -u 1000 dev
sudo su - dev
```

create a storage volume

```
sudo mkdir -p /media/state/shared/{bin,db,src,.ssh}
sudo touch /media/state/shared/.bash_history
sudo touch /media/state/shared/.maintainercfg
sudo chown -R dev:dev /media/state/shared/.*
sudo chown -R dev:dev /media/state/shared/
```

The dev user on the host OS is not currently configured with the capability to do anything priviledged except maybe modify the files or folders assigned to is.

```
mkdir -p ${HOME}/src/bitbucket.org/rbucker
cd ${HOME}/src/bitbucket.org/rbucker
git clone https://bitbucket.org/rbucker/devbox.git
cd devbox
docker build --rm -t=rbucker/devbox .
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
```


Hello World webserver
---------------------

tbd

HA Hello World
--------------

[docs](https://coreos.com/docs/launching-containers/launching/launching-containers-fleet/)

Hello World cloud-config
------------------------

tbd

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

Thanks
------

Special thanks to the CoreOS team Alex, Alex, and Brian (from CoreOS) who peeked over my shoulder while I wrote this. :)

License
-------
No license is offered.
