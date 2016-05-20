# TODO

* log to file/container in data volume
* set up nginx to serve console output
* create endpoint to create pipeline
* build react frontend
* pull/clone from git repo
* check out socket.io / node.js for streaming console output
* check out / migrate to goji

# FIXME

* [ ] pull image via go docker client if not available (no such image error)
* [x] containers shouldn't have access to other containers output, i.e. mount only relevant output dir

# Data

    /input
      + <git_url>
    /output
      + <build_id>
    /logs
      + <container_id>

* Output should be kept for 24h?
* Logs should be kept forever

# Docker

## Digitalocean as docker host

    export DIGITALOCEAN_ACCESS_TOKEN=...
    dm create -d digitalocean --digitalocean-region lon1 --digitalocean-size 4gb --digitalocean-image docker droplet
    dm create -d digitalocean --digitalocean-region lon1 --digitalocean-size 1gb --digitalocean-image docker droplet

## Mounting data containers

    # Create data containers
    docker create -v /output --name output tianon/true
    docker create -v /input --name input tianon/true
    docker create -v /input -v /output -v /logs --name data tianon/true

    # Spin up container
    docker run -ti --rm --volumes-from=input:ro --volumes-from=output ubuntu:15.10 /bin/bash

Note: In order to map the data volumes to a different directory inside the
container, see [this Stackoveflow question][1] and [this post][2].

## Actual directory of mounted volume on docker host

Example: Find directory for `/input` in `data` container:

    docker inspect -f '{{ range .Mounts }}{{ if eq .Destination "/input" }}{{ .Source }}{{ end }}{{ end }}' data

## Mounting host directories

    # Create host directories
    mkdir -p /tmp/input/x /tmp/output/x

    # Spin up container
    docker run -ti --rm -v /tmp/input/x:/input:ro -v /tmp/output/x:/output ubuntu:15.10 /bin/bash

Note: the host directories must be created on the docker host, not on the host
where you run the `docker-machine` command from.

## Overlaying

This could be useful to separate input and output directories.
For example: input -> git repo (read-only), output -> build output.

    mount -t overlayfs none -o lowerdir=/input,upperdir=/output /work
    cd /work

Note: - You must run the container in `--privileged` mode.
Note: - Using the `-w, --workdir` conveniently creates a work directory.

Examples:

    # with data containers
    docker run --privileged -ti --rm --volumes-from=input:ro --volumes-from=output -w /work ubuntu:15.10 /bin/bash

    # with host directories
    docker run --privileged -ti --rm -v /tmp/input/x:/input:ro -v /tmp/output/x:/output -w /work ubuntu:15.10 /bin/bash

## Compiling some Java

    docker run --privileged --rm -v /tmp/input/x:/input:ro -v /tmp/output/x:/output -w /work java:8 /bin/bash -c \
      "mount -t overlayfs none -o lowerdir=/input,upperdir=/output /work && cd /work && javac -verbose Hello.java"

# ElasticSearch

## Searching for all messages

`curl "http://$(dm ip droplet):9200/_search?pretty"`

## Searching for all messages from a given container id

`curl "http://$(dm ip droplet):9200/_search?q=docker.cid:000f0c1f0f1b&pretty"`

# Redis

## CLI

    docker run -it --link contdinfra_redis_1:redis --rm redis sh -c 'exec redis-cli -h "$REDIS_PORT_6379_TCP_ADDR" -p "$REDIS_PORT_6379_TCP_PORT"'


<!-- Links -->
[1]: http://stackoverflow.com/questions/23137544/how-to-map-volume-paths-using-dockers-volumes-from
[2]: https://martinvanbeurden.nl/blog/parsing-docker-1-8-volume-info/
