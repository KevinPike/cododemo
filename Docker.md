![Docker](https://raw.githubusercontent.com/rbucker/cododemo/master/docker.png)

What is Docker?
--------------

- docker is the wrapper for the toolchain that orchestrates everything between the docker container(s) and the host OS.
- docker containers/images are described as specialized tarfiles that are immutable and represent the userspace needed for the app that runs in the container
- the container should only run one process per container instance
- communication between containers is done with ```links```
- since containers are immutable persistence is implemented through ```volumes``` or ```links``` to storage-containers
- storage-containers that are not persisted via a ```volume``` will be lost when the reference count is zero
- there are 3rd party tools for backing up containers and moving them around the cluster; and then there is the registry
 
Preface
-------

[Link](http://blog.xen.org/index.php/2014/09/08/xen-docker-made-for-each-other/?utm_source=rss&utm_medium=rss&utm_campaign=xen-docker-made-for-each-other&utm_source=twitterfeed&utm_medium=twitter) to an article that compares VM (hypervisors) to containers.

A [press release](http://cto.vmware.com/vmware-docker-better-together/) from VMware talks aboult how VMware is going to bolt Docker onto their ecosystem.

This [post](https://wiki.openstack.org/wiki/Docker) from OpenStack presents a similar and more complete understanding of their implementation. One of the advantages of a Docker container is that their is a density of containers in a single host. When combined with etcd, fleetd and other tools those containers can talk to local and remote containers as the infrastructure is located in the host. The post, linked here, suggests that their is a 1:1 between the hypervisor and the container. There are many missing elements as you have or will see in this example. The exact details of the Docker Virt Driver are TBD.

The original Docker implementation depended on LXC containers and AUFS. Since then the Docker team has been developing their own container and has also experimenting with different filesystems.

The Docker Stack
----------------

![image stack](https://docs.docker.com/terms/images/docker-filesystems-multilayer.png)

** Introduction to [layers](https://docs.docker.com/terms/layer/#ufs-def)

Docker tools
------------

- registry, private registry (docker hub)
- gitreceive - webhooks for CI build of containers
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

#### Docker commands

```pull``` - pull an image from the repository. It will remain dormant until ```run``` or ```start```.

```commit``` - push a local image to the remote registry. The <name> is critical and can effect where it is private or public.

```rmi``` - remove an image

```images``` - list all the local images


Docker Commands
---------------

```build``` - Build a new image from the source code (Dockerfile). Each task in the Dockerfile creates a separate image file. Using the ```-rm``` flag deletes the intermediate images saving space but subsequent builds will take longer.

```run``` - Run a command in a new container. If you want the command to run in the background then you need to set the interactive flag. There is a new (1.2.0) ```restart``` flag that tells docker to restart the environment. When deploying a cluster then use use with care.

```commit``` - Create a new image from a container's changes (save it in the repo)

```stop``` - Stop a running container by sending SIGTERM and then SIGKILL after a grace period. (assuming that you initiated a command by calling the ```run``` command with the interactive flag.

```start``` - Restart a ```stop```ped container. You cannot start a command that has exited. (docker provides a restart flag) Auto-restart must be considered carefully when working with CoreOS.

Docker Cleanup
--------------

I ran my hello world command multiple times. Something like. Since the command was just printing out hello and some other text before exiting; container is in a terminal state. Not running. Running the ```docker ps``` command will return an empty list. However, after running the command ```docker run hello-world``` multiple times the docker container did leave some breadcrumbs.

##### TASK
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

##### TASK
```
vagrant ssh core-01
docker run hello-world
```

notice that the output from this hello-world is the exact same as the boot2docker version. That's because they are the same container image constructed (built) from the same Dockerfile. [Here](https://registry.hub.docker.com/u/library/hello-world/) is the registry where the hello-world image lives.

devbox
------

This dockerfile is a "standardized" way to deploy a dev environment. The team should decide on what tools everyone needs so that they are constructing from the same toolchain.

[shykes/devbox](https://registry.hub.docker.com/u/shykes/devbox/) - is a container the allows the user to create a proper development environment. (build the container, run the container with a command, execute your shell commands etc...) He has a [link](https://github.com/shykes/devbox) to the github source.

rbucker/devbox(not registered) - I have created a Dockerfile with a little more tooling and some documentation on bitbucket [here](https://bitbucket.org/rbucker/devbox/src). The different commands and dependencies are included.


create a host user with the same uid and gid as the container

boot2docker (broken because docker already uses the uid that I was intending to use in my devbox)
```
adduser .....

```

Setup Our CoreOS devbox
----------------------------

The original container had some dependencies that I decided to work with instead of against. For example the default user is "dev"... in order to share persistence between the container and the host they need the same uid/gid.

##### perform the following tasks on each node

CoreOS

##### TASK
```
sudo groupadd -g 1000 dev
sudo useradd -d /home/dev -g 1000 -m -s /bin/bash -u 1000 dev
```

create a storage volume

##### TASK
```
sudo mkdir -p /media/state/shared/{bin,db,src,.ssh}
sudo touch /media/state/shared/.bash_history
sudo touch /media/state/shared/.maintainercfg
sudo chown -R dev:dev /media/state/shared/.*
sudo chown -R dev:dev /media/state/shared/
```

The dev user on the host OS is not currently configured with the capability to do anything priviledged except maybe modify the files or folders assigned to is.

##### TASK
```
mkdir -p ${HOME}/src/bitbucket.org/rbucker
cd ${HOME}/src/bitbucket.org/rbucker
git clone https://bitbucket.org/rbucker/devbox.git
cd devbox
docker build --rm -t=rbucker/devbox .
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
```

** this last command can be considered an "interactive shell" (with the ```-it``` flags.

** fun fact: when the ```build``` is in progress the docker folks would prefer that you select a modern and active distro as the base and therefore you should not have to execute the ```apt-get update``` etc... this simply creates an unnecessary set of delta changes consuming disk and performance.

#### Play with etcd again

##### TASK
```
# advanced - connecting to the host from inside the container (somewhat unreliable)
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
export DOCKERHOST=`netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}'`
curl -L http://${DOCKERHOST}:4001/v2/machines
curl -L http://${DOCKERHOST}:4001/v2/keys/mykey -XPUT -d value="this is awesome"
curl -L http://${DOCKERHOST}:4001/v2/keys/mykey
```

Hello World Part 2
------------------

If you are not already running in a container then launch a container

##### TASK
```
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
```

##### TASK
```
mkdir -p ${HOME}/src/github.com/rbucker
cd ${HOME}/src/github.com/rbucker
git clone https://github.com/rbucker/cododemo
cd cododemo
```

run the hello.go program through the go compiler/runner

##### TASK
```
go run hello.go 
```

##### TASK
- exit the container
- run the container
- go back to the hello source ```cd ${HOME}/src/github.com/rbucker/cododemo```
- run hello again
- build hello instead of run ```go build hello.go```
- get a long list from this folder and notice the flags on the executable ```hello```
- run hello from the executable ```./hello```
- what did you get?
- why?
- copy the executable to your home and execute ```cp ./hello ~/. && ~/hello```
- what did you get?
- why?

Hello World webserver
---------------------

start the container up
```
docker run -it -v /media/state/shared/:/var/shared/ -p 8080:8080 rbucker/devbox /bin/bash
cd ${HOME}/src/github.com/rbucker/cododemo
go run web.go
```

in the run command the ```-p 8080:8080``` param tells docker to redirect the public port 8080 to the private port 8080 of this container. If the application uses multiple ports then you can include multiple ```-p``` params.

Get the IP for the instance
```
grep COREOS_PUBLIC_IPV4 /etc/environment | awk 'BEGIN{FS = "="} {print $2}'
```

then in a browser
```
http://<ipaddr>:8080/bar
```

IP Addresses
------------

- the host has a public IP address (see /etc/environment)
- the host has a localhost too
- the container has an IP address (see /etc/environment, ifconfig)
- the container has a localhost too

##### QUESTION
What happened here?

Alternately (the ```-d``` flag indicates that this is detached)
```
docker run -d -v /media/state/shared/:/var/shared/ -p 8080:8080 rbucker/devbox /bin/sh -c "cd ~/src/github.com/rbucker/cododemo && go run web.go"
```

Now you can get the docker process stack
```
docker ps
```

and you can stop the container if you want.

cAdvisor Monitoring (single node)
-------------------

```
sudo docker run \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=cadvisor \
  google/cadvisor:latest
```

##### Backup a Docker Container

- https://github.com/discordianfish/docker-backup
- https://docs.docker.com/userguide/dockervolumes/
 
Tools
-----

##### Weave

[Project](https://github.com/zettio/weave/) to implement a virtual network to across hosts.

This is my overarching concern: [link](http://www.zett.io/)

##### Rudder

From the CoreOS project [this](https://github.com/coreos/rudder) "Rudder is an overlay network that gives a subnet to each machine for use with Kubernetes."
 
[return](https://github.com/rbucker/cododemo/blob/master/README.md)
