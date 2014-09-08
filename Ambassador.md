What is an Ambassador?
----------------------

"The ambassador pattern is a novel way to deploy sets of containers that are configured at runtime via the Docker Links feature."  - [Full Article](https://coreos.com/blog/docker-dynamic-ambassador-powered-by-etcd/)

![Ambassador](https://docs.docker.com/articles/ambassador_pattern_linking/)

The ambassador pattern is like the CoreOS-Sidekick-confd example except that the sidekick is implemented as a stub app and fleet dependencies. They both share etcd in order to replicate confiuration info.


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

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
