Let's Boot2docker
-----------------

Just to get our docker feet wet. It has no persistance unless your work uploads or saves it's content directly. Once the container is gone; it is gone.

##### TASK
```
boot2docker init
boot2docker start
export DOCKER_HOST=tcp://$(boot2docker ip 2>/dev/null):2375
```

The DOCs suggest that there is supposed to be a popup now. That the popup is a terminal window into the docker session. I'm pretty certain the hint is wrong. So ssh into the instance.

##### TASK
```
boot2docker ssh
```

You can ssh into and back out as often as you want. That's just the lightweight linux instance. It does not actually have any disk so there is no persistance and will not survive a reboot. Do not confuse the boot2docker host OS form the container.

now we are in a linux shell... so run the hello world container.

##### TASK
```
docker run hello-world
```

** one of the things that makes CoreOS nice is that it is immutable where it counts. This theme is extended into docker too. Once you create a container with a Dockerfile you should never change it. If you want to make a change then rebuild the container. As for saving data or persisting information; that is performed using volume mounts points or data-containers.
