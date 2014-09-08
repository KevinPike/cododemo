What is CoreOS?
---------------

CoreOS is a new Linux distribution that has been rearchitected to provide features needed to run modern infrastructure stacks. The strategies and architectures that influence CoreOS allow companies like Google, Facebook and Twitter to run their services at scale with high resilience. -- The CoreOS Site

exactly:
- Linux distro
- hand assembled
- no package manager, no perl, no python, no ruby
- A/B core
- immutable
- custom distro channels
- auto update or controlled by enterprise tools

NOTE
----

[Project Atomic](http://www.projectatomic.io/) is from Red Hat and is also awesome. It's based on Fedora with some SELinux sprinkled in. It also has some infrastructure dashboard type stuff and it also borrows from Docker and other projects. [CoreOS vs. Project Atomic: A Review](https://major.io/2014/05/13/coreos-vs-project-atomic-a-review/)


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

edit the [config.rb](https://github.com/coreos/coreos-vagrant/blob/master/config.rb.sample) file and make these changes. The lines might need to be uncommented:

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

The [user-data](https://github.com/coreos/coreos-vagrant/blob/master/user-data.sample) is a little tricky. The file is modelled a after the cloud-config file.

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

Play with etcd
--------------

do this or pick your own combination of tasks

Start here - clean my key, create a key, and verify on core-01
```
vagrant ssh core-01
curl -L http://127.0.0.1:4001/version
curl -L http://127.0.0.1:4001/v2/machines
curl -L http://127.0.0.1:4001/v2/leader
curl -L http://127.0.0.1:4001/v2/keys/mykey -XDELETE
curl -L http://127.0.0.1:4001/v2/keys/mykey -XPUT -d value="this is awesome"
curl -L http://127.0.0.1:4001/v2/keys/mykey
```

Check that the key was replicated to core-02
```
vagrant ssh core-02
curl -L http://127.0.0.1:4001/v2/machines
curl -L http://127.0.0.1:4001/v2/keys/mykey
curl -L http://127.0.0.1:4001/v2/keys/mykey -XDELETE
```

verify it's been deleted from core-01
```
vagrant ssh core-01
curl -L http://127.0.0.1:4001/v2/keys/mykey
```

verify it's still deleted from core-02
```
vagrant ssh core-02
curl -L http://127.0.0.1:4001/v2/keys/mykey
```

release
```
cat /etc/*release
```

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
