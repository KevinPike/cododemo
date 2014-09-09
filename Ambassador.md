What is an Ambassador?
----------------------

The ambassador pattern is like the CoreOS-Sidekick-confd example except that the sidekick is implemented as a stub app and fleet dependencies. They both share etcd in order to replicate confiuration info. 

The mission for the ambassador pattern is to (a) support multiple network topologies (b) survive network partitioning (c) and some HA strategies.

First Example
-------------

Cross-Host linking using Ambassador Containers - [Full Article](https://docs.docker.com/articles/ambassador_pattern_linking/)

The ambassador is implemented by redirecting pipes [socat](http://linux.die.net/man/1/socat) and a little sed.

### Exercise

#### on the server
```
docker run -d -name redis crosbymichael/redis
docker run -t -i --link redis:redis --name redis_ambassador -p 6379:6379 svendowideit/ambassador
```

#### on the client
```
docker run -d --name redis_ambassador --expose 6379 -e REDIS_PORT_6379_TCP=tcp://172.17.8.104:6379 svendowideit/ambassador
docker run -i -t --rm --link redis_ambassador:redis relateiq/redis-cli
```

#### Question
are there any criticisms of this ambassador?

####Task:
- How would you improve on this example?
- what happens when the server crashes?


Second Example
--------------

"The ambassador pattern is a novel way to deploy sets of containers that are configured at runtime via the Docker Links feature."  - [Full Article](https://coreos.com/blog/docker-dynamic-ambassador-powered-by-etcd/)

### Exercise

#### on the server
#### on the client

#### Question
are there any criticisms of this ambassador?

####Task:
- How would you improve on this example?
- what happens when the server crashes?


[return](https://github.com/rbucker/cododemo/blob/master/README.md)
