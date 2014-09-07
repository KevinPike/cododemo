



Hello World webserver Cluster
--------------

Repeat the steps above on core-02 and core-03 and just for completeness run and verify the webserver from the previous step.

** another fun fact. I'm writing this workshop on my home network which is experiencing stability problems. During the build on core-02 the network crashed and the build stopped. I simply restarted the build and it resumed from where it had left off. (a) no fuss or muss (b) no worries if it was complete or not.

```kill``` the container currently running on core-01.

```
mkdir -p ~/src/github.com/rbucker
cd ~/src/github.com/rbucker
git clone https://github.com/rbucker/cododemo
cd cododemo
```


Get a list of the machines in the cluster
```
fleetctl list-machines
```

Load the service file into fleet
```
fleetctl list-unit-files
fleetctl submit web@.service
fleetctl list-unit-files
fleetctl list-units
```

Start the units (it's going to take a while)
```
fleetctl start web@{8081..8083}.service
```
And then you can test the status
```
fleetctl list-unit-files
fleetctl list-units
```

Get the port number
```
docker ps
```

then in a browser
```
http://<ipaddr>/bar
```



HA Hello World webserver
--------------

```
fleetctl submit nginx.service
fleetctl start nginx.service
```

Get the IP for the instance
```
grep COREOS_PUBLIC_IPV4 /etc/environment | awk 'BEGIN{FS = "="} {print $2}'
```

then in a browser
```
http://<ipaddr>/bar
```

** reload a few times and watch the ID

[docs](https://coreos.com/docs/launching-containers/launching/launching-containers-fleet/)


Logging
-------

It would be great to execut ALL of these commands from the same or any system. Try each one. What do you suppose is the issue?

```
fleetctl journal -f web@8081.service
fleetctl journal -f web@8082.service
fleetctl journal -f web@8083.service
fleetctl journal -f nginx.service
```

How would you correct this problem?




Hello World cloud-config
------------------------

tbd

What is a sidekick app?
-----------------------

In this example the nginx instance (1) watches etcd, (2) reconstructs the nginx config with the changes, (3) restarts nginx. Much of this is triggered by fleet and the nginx container. In the case of a sidekick, there would be a special container that is linked to the main app in separate fleet unit files. Then when the sidekick determined that an efent was occuring it would notify some other service (including etcd) that a change had ocurred; and thus the change would be reflected upstream.

Just as in this example. "sidekick" is a different, yet idiomatic, way to perform the same function.

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
