# ICKFILE : An INTERCAL frontend for Docker Build

This is a syntax frontend for BuildKit that allows you to write Dockerfiles using a syntax based on [INTERCAL](https://en.wikipedia.org/wiki/INTERCAL)

## Example Ickfile
```
#syntax=agbell/ick
COME_FROM alpine

PLEASE echo "Finally a great syntax for creating docker images"
DO ["/bin/sh"]
ARE_YOU_OK CMD ls 
```

## Building
```
➜  docker build . -f ./Ickfile -t ick

[+] Building Ickfile                            
 => [internal] load build definition from Ickfile         0.0s
 => => transferring dockerfile: 258B                      0.0s
 => [internal] load .dockerignore                         0.0s
 => => transferring context: 2B                           0.0s
 => resolve image config for docker.io/agbell/ick:latest  0.0s
 => docker-image://docker.io/agbell/ick:latest            0.0s
 => => resolve docker.io/agbell/ick:latest                0.0s
 => [internal] load metadata for docker.io/library/alpin  0.0s
 => [1/3] COME_FROM docker.io/library/alpine              0.0s
 => => resolve docker.io/library/alpine:latest            0.0s
 => [2/3] <- /src                                         0.0s
 => [3/3] PLEASE echo "Finally a great syntax for creati  0.3s
 => exporting to image                                    0.0s
 => => exporting layers                                   0.0s
 => => writing image sha256:cbb3269cdb870eb2bf06fb3650e0  0.0s
```
## Running
Run the same as any other container image:
```
➜  docker run -it ick
/src # echo "It Works!
It Works!
```

## Instructions:

### COME FROM
This is one of INTERCAL's [most innovative features](https://en.wikipedia.org/wiki/COMEFROM) and is now available for `docker build`. In an Ickfile, it causes the container image to inherit from the listed base image.  The base image need not be created with an Ickfile, although that is preferred.

```
COME_FROM alpine
```

### PLEASE
Execute something
```
PLEASE ["/bin/bash", "-c", "echo hello"]
```

### PLEASE DO
Configure the command to execute when the container starts. It is similar but confusingly different from `DO`
```
PLEASE DO ["/bin/sh"]
```

### DO
The default command to execute when the container starts. 
```
DO ["/bin/sh"]
```
### MYSTERY
This option is deliberately undocumented.

### ARE YOU OK

This command tells Docker how to test that the container is working.

```
ARE_YOU_OK --interval=5s --timeout=3s --retries=3 CMD ls --quiet
```

### STASH
Stash files in the container image.
```
STASH hom* /mydir/
```

# Usage
No installation is necessary. Just add this line to your Dockerfile:
```
#syntax=agbell/ick
```
And enable BuildKit and build:
```
DOCKER_BUILDKIT=1 docker build .
``` 