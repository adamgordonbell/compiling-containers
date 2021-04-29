## code
find the code
`git remote --verbose`

## build simple c program
` clang main.cc -o main `

` ./main`

## build simple docker file

`docker build .-t agbell/demo1`

`docker run -it agbell/demo1 /bin/sh`

`ls README.md`
`cat build.txt`

## build with buildkit

build same but with buildkit directly
 * change the echo script

```javascript
brew install buildkit
```
```
docker run --rm --privileged -d \
    --name buildkit moby/buildkit
docker ps
export BUILDKIT_HOST=docker-container://buildkit
```

```javascript
buildctl build \
    --frontend=dockerfile.v0 \
    --local context=. \
    --local dockerfile=. \
    --output type=image,name=docker.io/agbell/demo2,push=true
```
```javascript
docker run -it --pull always agbell/demo2 /bin/sh
cat build.txt
```

## Programmatically build images

 * Frontend -> LLB -> Buildkit -> LLB -> Backend
 * Why GO : https://pkg.go.dev/github.com/moby/buildkit/client/llb
 * writellb.go explain


`go run ./writellb/writellb.go`

``` 
   go run ./writellb/writellb.go | \
   buildctl debug dump-llb | \
   jq .
```

```
  go run ./writellb/writellb.go | \
    buildctl build \
      --local context=. \
      --output type=image,name=docker.io/agbell/demo3,push=true
```

```javascript
docker run -it --pull always agbell/demo3 /bin/sh
cat build.txt
```
This is cool because you can use all abstraction and control from and libraries that a programming language offers you to build docker files. You can raise the abstraction level.

## Build own frontend

* a true frontend is not just code, but `Text -> Tokenize -> Parse -> IR `
* you can do this using #syntax={mydockerfile}
* needs to be somewhere docker can pull it from
* needs to listen on GRPC for the text input
* needs to generate LLB

 I did this: 
 ICK file - trivial programming langauge
Let's build it:
```
docker build . -f ickfile.Dockerfile -t agbell/demo4
```
```
docker push agbell/demo4
```

This is my ick file
```
docker build . -f ./ickfile/Ickfile -t demo5
```

```
docker run -it demo4 /bin/sh
```
- This code is all available online if you want to use it. 
- These images, even made with ICK files will work with any container runtime
- look at command.go
