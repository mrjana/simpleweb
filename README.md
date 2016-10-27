# simpleweb

Build binary.

```
go build --ldflags '-linkmode external -extldflags "-static"'
```

Build image.

```
docker build .
docker tag xxx simpleweb
```

Run container.

```
docker run -d -p 5000:5000 simpleweb simpleweb

dchen@vm2:~/docker-1.13$ docker ps
CONTAINER ID        IMAGE                       COMMAND                  CREATED             STATUS                       PORTS                    NAMES
1cddbb33ec6d        simpleweb	                 "simpleweb"              18 minutes ago      Up 18 minutes (healthy)      0.0.0.0:5000->5000/tcp   silly_mahavira
```
