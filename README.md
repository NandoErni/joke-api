# joke-api

A REST API for Swiss Jokes

Run the server:

```
go run main.go
```

Test it [here](http://api.faultieroase.cloud:8080/jokes/)

Build the docker image

```
docker build -t joke-api .
```

Upload to dockerhub

```
docker tag joke-api nandoerni/joke-api:latest
docker push nandoerni/joke-api:latest
```
