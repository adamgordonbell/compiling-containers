# compiling-containers

# front end
build my front end:
```
docker build -f ./dockerfile/cmd/dockerfile-frontend/Dockerfile . -t agbell/frontend 
```

```
go build ./dockerfile/cmd/dockerfile-frontend/main.go
```


## make LLB
