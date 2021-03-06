![CoreOS](https://raw.githubusercontent.com/rbucker/cododemo/master/coreos-wordmark-vert-color.png)

What is CoreOS?
---------------

CoreOS is a new Linux distribution that has been rearchitected to provide features needed to run modern infrastructure stacks. The strategies and architectures that influence CoreOS allow companies like Google, Facebook and Twitter to run their services at scale with high resilience. -- The CoreOS Site


- Linux distro
- hand assembled
- no package manager, no perl, no python, no ruby
- A/B core borrows from ChromeOS
- immutable
- custom distro channels [releases](http://coreos.com/releases)
- auto update or controlled by enterprise tools (updates are different when installed on CoreOS proper and documented [here](https://github.com/coreos/coreos-vagrant#new-box-versions) on vagrant)

NOTE
----

[Project Atomic](http://www.projectatomic.io/) is from Red Hat and is also awesome. It's based on Fedora with some SELinux sprinkled in. It also has some infrastructure dashboard type stuff and it also borrows from Docker and other projects. [CoreOS vs. Project Atomic: A Review](https://major.io/2014/05/13/coreos-vs-project-atomic-a-review/)

Since Project Atomic is sponsored by Red Hat there is plenty of attention to security. [This](http://sdtimes.com/red-hat-open-source-community-fortifying-docker/) article mentions some of the details.

My opinion is that CoreOS demonstrates strict release discipline.


CoreOS tools
------------

- CoreOS Cluster is defined by 3 or more instances.
- etcd - replicated key/value store using the [raft](http://thesecretlivesofdata.com/raft/) protocol.
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

##### TASK
```
mkdir -p {$HOME}/src/github.com/coreos
cd {$HOME}/src/github.com/coreos
git clone https://github.com/coreos/coreos-vagrant
cd coreos-vagrant
cp config.rb.sample config.rb
cp user-data.sample user-data
```

edit the [config.rb](https://github.com/coreos/coreos-vagrant/blob/master/config.rb.sample) file and make these changes. The lines might need to be uncommented:

##### TASK
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

"CoreOS allows you to declaratively customize various OS-level items, such as network configuration, user accounts, and systemd units."

The first thing you need is key for ETCD so that the cluster can identify itself as part of the cluster. Each core instance of etcd uses this to identify the cluster. (if you ```vagrant destroy``` the cluster you need a ```new``` key)

##### TASK
```
curl https://discovery.etcd.io/new
```

Take the return string and paste it into the user-data file. Notice that the file in the yml format.

##### TASK
```
discovery: https://discovery.etcd.io/<replace this with the from the step above>
```

Start the cluster. Notice that the first instance took a while to create. And the second two very quickly.

##### TASK
```
vagrant up
vagrant status
vagrant ssh-config
```

ssh into an instance.

##### TASK
```
vagrant ssh core-01
vagrant ssh core-02
vagrant ssh core-03
```

Play with etcd
--------------

do this or pick your own combination of tasks

Start here - clean my key, create a key, and verify on core-01

##### TASK
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

##### TASK
```
vagrant ssh core-02
curl -L http://127.0.0.1:4001/v2/machines
curl -L http://127.0.0.1:4001/v2/keys/mykey
curl -L http://127.0.0.1:4001/v2/keys/mykey -XDELETE
```

verify it's been deleted from core-01

##### TASK
```
vagrant ssh core-01
curl -L http://127.0.0.1:4001/v2/keys/mykey
```

verify it's still deleted from core-02

##### TASK
```
vagrant ssh core-02
curl -L http://127.0.0.1:4001/v2/keys/mykey
```

release

##### TASK
```
cat /etc/*release
docker info
```



free space

##### TASK

```
core@core-01 ~ $ df -h
df: '/var/lib/docker/btrfs': Permission denied
Filesystem      Size  Used Avail Use% Mounted on
rootfs           17G  1.3G   15G   9% /
devtmpfs        488M     0  488M   0% /dev
tmpfs           500M     0  500M   0% /dev/shm
tmpfs           500M  348K  499M   1% /run
tmpfs           500M     0  500M   0% /sys/fs/cgroup
/dev/sda9        17G  1.3G   15G   9% /
/dev/sda3      1008M  288M  669M  31% /usr
tmpfs           500M  4.0K  500M   1% /tmp
tmpfs           500M  3.9M  496M   1% /media
/dev/sda6       108M   88K   99M   1% /usr/share/oem
```


[btrfs](https://btrfs.wiki.kernel.org/index.php/Main_Page) is a new copy on write (CoW) filesystem for Linux aimed at implementing advanced features while focusing on fault tolerance, repair and easy administration.

mounted partition

##### TASK
```
core@core-01 ~ $ mount
sysfs on /sys type sysfs (rw,nosuid,nodev,noexec,relatime)
proc on /proc type proc (rw,nosuid,nodev,noexec,relatime)
devtmpfs on /dev type devtmpfs (rw,nosuid,size=499528k,nr_inodes=124882,mode=755)
securityfs on /sys/kernel/security type securityfs (rw,nosuid,nodev,noexec,relatime)
tmpfs on /dev/shm type tmpfs (rw,nosuid,nodev)
devpts on /dev/pts type devpts (rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=000)
tmpfs on /run type tmpfs (rw,nosuid,nodev,mode=755)
tmpfs on /sys/fs/cgroup type tmpfs (ro,nosuid,nodev,noexec,mode=755)
cgroup on /sys/fs/cgroup/systemd type cgroup (rw,nosuid,nodev,noexec,relatime,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd)
pstore on /sys/fs/pstore type pstore (rw,nosuid,nodev,noexec,relatime)
cgroup on /sys/fs/cgroup/cpuset type cgroup (rw,nosuid,nodev,noexec,relatime,cpuset)
cgroup on /sys/fs/cgroup/cpu,cpuacct type cgroup (rw,nosuid,nodev,noexec,relatime,cpu,cpuacct)
cgroup on /sys/fs/cgroup/memory type cgroup (rw,nosuid,nodev,noexec,relatime,memory)
cgroup on /sys/fs/cgroup/devices type cgroup (rw,nosuid,nodev,noexec,relatime,devices)
cgroup on /sys/fs/cgroup/freezer type cgroup (rw,nosuid,nodev,noexec,relatime,freezer)
cgroup on /sys/fs/cgroup/net_cls,net_prio type cgroup (rw,nosuid,nodev,noexec,relatime,net_cls,net_prio)
cgroup on /sys/fs/cgroup/blkio type cgroup (rw,nosuid,nodev,noexec,relatime,blkio)
cgroup on /sys/fs/cgroup/perf_event type cgroup (rw,nosuid,nodev,noexec,relatime,perf_event)
/dev/sda9 on / type btrfs (rw,relatime,space_cache)
/dev/sda3 on /usr type ext4 (ro,relatime)
tmpfs on /tmp type tmpfs (rw)
tmpfs on /media type tmpfs (rw,nosuid,nodev,noexec,relatime)
debugfs on /sys/kernel/debug type debugfs (rw,relatime)
mqueue on /dev/mqueue type mqueue (rw,relatime)
hugetlbfs on /dev/hugepages type hugetlbfs (rw,relatime)
/dev/sda6 on /usr/share/oem type ext4 (rw,nodev,relatime,commit=600,data=ordered)
/dev/sda9 on /var/lib/docker/btrfs type btrfs (rw,relatime,space_cache)
```

Notice ```tmpfs on /media type tmpfs (rw,nosuid,nodev,noexec,relatime)``` is the root folder where the shared volume is stored and that it has the ```noexec``` flag set. This attribute extends into the symlink in the container so that you cannot execute binaries in the volumes. (overrides might be possible)

read-only

##### TASK
```
core@core-01 ~ $ ls -l /etc
. . .
-rw-r--r-- 1 root root    63 Sep  5 03:30 group
-rw------- 1 root root    51 Aug 28 07:55 group-
-rw-r----- 1 root root    51 Sep  5 03:30 gshadow
-rw------- 1 root root    43 Aug 28 07:55 gshadow-
-rw-r--r-- 1 root root     8 Sep  4 21:02 hostname
-rw-r--r-- 1 root root  3580 Aug 28 07:42 idmapd.conf
lrwxrwxrwx 1 root root    31 Aug 28 08:12 inputrc -> ../usr/share/baselayout/inputrc
lrwxrwxrwx 1 root root    12 Aug 28 08:17 issue -> ../run/issue
drwxr-xr-x 1 root root    18 Aug 28 08:16 kernel
-rw-r--r-- 1 root root 17578 Sep  4 21:02 ld.so.cache
lrwxrwxrwx 1 root root    21 Aug 28 08:17 ld.so.conf -> ../usr/lib/ld.so.conf
lrwxrwxrwx 1 root root    26 Aug 28 08:14 limits -> ../usr/share/shadow/limits
-rw-r--r-- 1 root root   109 Aug 28 08:16 locale.conf
lrwxrwxrwx 1 root root    25 Sep  4 21:02 localtime -> ../usr/share/zoneinfo/UTC
lrwxrwxrwx 1 root root    32 Aug 28 08:14 login.access -> ../usr/share/shadow/login.access
lrwxrwxrwx 1 root root    30 Aug 28 08:14 login.defs -> ../usr/share/shadow/login.defs
lrwxrwxrwx 1 root root    31 Aug 28 08:12 lsb-release -> ../usr/share/coreos/lsb-release
drwxr-xr-x 1 root root     0 Sep  4 21:02 lvm
-r--r--r-- 1 root root    33 Sep  4 21:01 machine-id
-rw-r--r-- 1 root root  2689 Aug 28 07:50 mdadm.conf
-rw-r--r-- 1 root root   956 Aug 28 07:50 mke2fs.conf
drwxr-xr-x 1 root root     0 Aug 28 08:16 modules-load.d
lrwxrwxrwx 1 root root    18 Aug 28 08:12 motd -> ../run/coreos/motd
lrwxrwxrwx 1 root root    19 Sep  4 21:02 mtab -> ../proc/self/mounts
drwxr-xr-x 1 root root    22 Aug 28 08:13 mtools

```

How long does it take to boot CoreOS

##### TASK
```
$ dmesg
. . .
[    3.438029] IPv6: ADDRCONF(NETDEV_UP): enp0s8: link is not ready
[    3.440429] e1000: enp0s8 NIC Link is Up 1000 Mbps Full Duplex, Flow Control: RX
[    3.457107] IPv6: ADDRCONF(NETDEV_CHANGE): enp0s8: link becomes ready
[    3.477040] IPv6: ADDRCONF(NETDEV_UP): enp0s3: link is not ready
[    3.479443] e1000: enp0s3 NIC Link is Up 1000 Mbps Full Duplex, Flow Control: RX
[    3.484650] IPv6: ADDRCONF(NETDEV_CHANGE): enp0s3: link becomes ready
```

- Add a 4th node to the local install
- what happens after a reboot of the node?

##### TASK
```
vagrant ssh core-02
sudo reboot  
vagrant ssh core-02
dmesg
fleetctl list-machines
docker ps
```

** it's a trick question

### Cron

Fleet's service files are very much like the systemd service files. I have not found the decoder ring yet, however, this [Man Page](http://www.freedesktop.org/software/systemd/man/systemd.service.html) is useful. You'll have to stitch together the sections.

The CoreOS team calls this "scheduling".

[Link](http://bl.ocks.org/adamalex/b5fb7f6b42caba4c3413) to some units that look like cron configs.

Some oneshot [examples](https://github.com/coreos/fleet/issues/562)


Enterprise tools

- [scale](https://coreos.com/assets/images/media/Fleet-Scheduling.png)
- [managed](https://coreos.com/products/managed-linux/)
- [update](https://coreos.com/products/coreupdate/screenshots/)
- [registry](https://coreos.com/products/enterprise-registry/)

** cool feature ... suspend CoreOS updates while app is upgraded

Interesting use case
--------------------

MemSQL: Running A 107 Node Cluster on CoreOS [Video](https://www.youtube.com/watch?v=uJirOCUg67o)


[return](https://github.com/rbucker/cododemo/blob/master/README.md)
