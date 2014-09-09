Nats
====

NATS is an open-source, lightweight cloud messaging system.

NATS was created by Derek Collison, Founder/CEO of Apcera who has spent 20+ years designing, building, and using publish-subscribe messaging systems. Unlike traditional enterprise messaging systems, NATS has an always-on dial tone that does whatever it takes to remain available. This forms a great base for building modern, reliable, and scalable cloud and distributed systems. - [Full site](http://nats.io)


#### Install and run gnatsd server on core-01

```
docker run --name my_gnatsd -d apcera/gnatsd
```

#### create a publisher example app on Core-01

NATS Client source is [here](https://github.com/apcera/nats)


#### create a publisher example app on Core-02

#### create a subscriber example app on Core-03



[return](https://github.com/rbucker/cododemo/blob/master/README.md)
