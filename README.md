# compiling-containers

# front end
build my front end:

```
earthly -i --push +build  
```

```
docker build -f ./dockerfile/cmd/dockerfile-frontend/Dockerfile . -t agbell/frontend 
```

```
go build ./dockerfile/cmd/dockerfile-frontend/main.go
```


## make LLB

Output:
```
 go run ./writellb/writellb.go | buildctl debug dump-llb | jq .
```
Start buildkit:
```
docker run --rm --privileged -d --name buildkit moby/buildkit
export BUILDKIT_HOST=docker-container://buildkit
```



