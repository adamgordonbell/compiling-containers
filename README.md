# compiling-containers

# front end
build my front end:

```
earthly -i --push +ick-frontend  
```

```
docker build -f ./dockerfile/cmd/dockerfile-frontend/Dockerfile . -t agbell/frontend 
```

```
go build ./dockerfile/cmd/dockerfile-frontend/main.go
```


## make LLB
Start buildkit:
```
docker run --rm --privileged -d --name buildkit moby/buildkit
export BUILDKIT_HOST=docker-container://buildkit
```

Output LLB:
```
 go run ./writellb/writellb.go | buildctl debug dump-llb | jq .
```

Build the Image:
```
  go run ./writellb/writellb.go | buildctl build --local context=. --output type=image,name=docker.io/agbell/test,push=true
```

Run the image
```

docker run -it agbell/test:latest /bin/sh
```

```
➜  compiling-containers git:(main) ✗ go run ./writellb/writellb.go | buildctl build --no-cache --local context=. --output type=image,name=docker.io/agbell/test,push=true
[+] Building 1.9s (7/7) FINISHED                                                                                                                                                                                   
 => local://context                                                                                                                                                                                           0.1s
 => => transferring context: 48.26kB                                                                                                                                                                          0.1s
 => docker-image://docker.io/library/alpine:latest                                                                                                                                                            0.0s
 => => resolve docker.io/library/alpine:latest                                                                                                                                                                0.6s
 => [auth] library/alpine:pull token for registry-1.docker.io                                                                                                                                                 0.0s
 => copy /README.md /README.md                                                                                                                                                                                0.0s
 => /bin/sh -c echo "programmatically built" > /built.txt                                                                                                                                                     0.1s
 => exporting to image                                                                                                                                                                                        1.7s
 => => exporting layers                                                                                                                                                                                       0.1s
 => => exporting manifest sha256:7b613d1abdd023824a709c4c97d59087dbdd6368dbccb6b22ed0305d68f083da                                                                                                             0.0s
 => => exporting config sha256:ca6a7e01c880886fd55f0e0e7056d63731c164a85e82e0c4833cc84d07b716da                                                                                                               0.0s
 => => pushing layers                                                                                                                                                                                         1.4s
 => => pushing manifest for docker.io/agbell/test:latest                                                                                                                                                      0.3s
 => [auth] agbell/test:pull,push token for registry-1.docker.io                                                                                                                                               0.0s
➜  compiling-containers git:(main) ✗ docker run -it --pull always agbell/test:latest /bin/sh                                                                             
latest: Pulling from agbell/test
ba3557a56b15: Already exists 
e631696b715d: Already exists 
9868e2d03655: Pull complete 
Digest: sha256:7b613d1abdd023824a709c4c97d59087dbdd6368dbccb6b22ed0305d68f083da
Status: Downloaded newer image for agbell/test:latest
/ # cat built.txt 
programmatically built
/ # ls README.md
README.md
```
