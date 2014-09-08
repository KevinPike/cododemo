Ooops
-----

```
fleetctl destroy web@{8081..8083}.service
fleetctl start web@{8081..8083}.service
fleetctl list-units
```

now reload the webpage. what happened?

```
fleetctl destroy nginx.service
fleetctl start nginx.service
fleetctl list-units
```

reload the browser. What happened?


What is an Ambassador?
----------------------

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

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
