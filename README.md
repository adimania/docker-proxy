docker-proxy
============

A simple proxy for docker client to bind to and run commands. May be I'll add ACLs someday.

Why?
----
* Docker, by default, listens on the Unix socket file. I don't want to run it on tcp socket forever but I don't like to restart docker daemon whenever I want to access docker over network.
* At some point of time, I would like the proxy to have features like MySQL grant system. This proxy will enable me to intercept the request and allow or deny on the basis of predefined ACLs.

How do I use it?
----------------
This is a transparent proxy. So one can just run it on the docker server as a member of docker group or as root (former is preferred) and the use docker to fire commands normally with an additional -H flah.

Example:
```
$ nohup go run docker-proxy.go --port 4321 &
$ docker -H tcp://docker_host:4321 ps
```
What is wrong?
--------------
* If the reply by docker daemon is long, then the response is chunked and right now handling of that is really screwed up.
* I am a golang novice. The code works but might be unoptimized. 
