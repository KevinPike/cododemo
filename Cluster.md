



Hello World webserver Cluster
--------------

Repeat the steps above on core-02 and core-03 and just for completeness run and verify the webserver from the previous step.

** another fun fact. I'm writing this workshop on my home network which is experiencing stability problems. During the build on core-02 the network crashed and the build stopped. I simply restarted the build and it resumed from where it had left off. (a) no fuss or muss (b) no worries if it was complete or not.

```kill``` the container currently running on core-01.

Get a list of the machines in the cluster
```
fleetctl list-machines
```

```
fleetctl list-unit-files
```

```
fleetctl list-units
```



HA Hello World webserver
--------------

[docs](https://coreos.com/docs/launching-containers/launching/launching-containers-fleet/)

Hello World cloud-config
------------------------

tbd


Get the IP for the instance
```
grep COREOS_PUBLIC_IPV4 /etc/environment | awk 'BEGIN{FS = "="} {print $2}'
```

then in a browser
```
http://<ipaddr>:8080/bar
```

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
