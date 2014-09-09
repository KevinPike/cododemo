Nats
====

NATS is an open-source, lightweight cloud messaging system.

NATS was created by Derek Collison, Founder/CEO of Apcera who has spent 20+ years designing, building, and using publish-subscribe messaging systems. Unlike traditional enterprise messaging systems, NATS has an always-on dial tone that does whatever it takes to remain available. This forms a great base for building modern, reliable, and scalable cloud and distributed systems. - [Full site](http://nats.io)

NATS Client source is [here](https://github.com/apcera/nats)

NATS [Community](http://nats.io/community/)

#### Install and run gnatsd server on core-01

```
docker run --name my_gnatsd -p 4222:4222 -p 8333:8333 -d apcera/gnatsd
```

#### create a publisher example app on Core-01


```
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
mkdir -p ${HOME}/src/github.com/apcera
cd ${HOME}/src/github.com/apcera
git clone https://github.com/apcera/nats.git
export GOPATH=${HOME}
cd nats/examples
go build nats-sub.go
cp nats-sub ${HOME}/.
${HOME}/nats-sub -s nats://172.17.8.101:4222 "ntest"
```


#### create a publisher example app on Core-02

```
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
mkdir -p ${HOME}/src/github.com/apcera
cd ${HOME}/src/github.com/apcera
git clone https://github.com/apcera/nats.git
export GOPATH=${HOME}
cd nats/examples
go build nats-sub.go
cp nats-sub ${HOME}/.
${HOME}/nats-sub -s nats://172.17.8.101:4222 "ntest"
```

#### create a subscriber example app on Core-03

```
docker run -it -v /media/state/shared/:/var/shared/ rbucker/devbox /bin/bash
mkdir -p ${HOME}/src/github.com/apcera
cd ${HOME}/src/github.com/apcera
git clone https://github.com/apcera/nats.git
export GOPATH=${HOME}
cd nats/examples
go build nats-pub.go
cp nats-pub ${HOME}/.
${HOME}/nats-pub -s nats://172.17.8.101:4222 "ntest" "test message"
```

Now that you have created two subscribers and sent one message via the one publisher... review the other two consoles and notice that they received the incoming test message.

[return](https://github.com/rbucker/cododemo/blob/master/README.md)
