# Graphql Go Sample

## Start Go

```
docker-compose up -d
go run main.go
```

## Test

```
wrk -c 100 -d 1s -s post.lua http://localhost:4000/graphql/
```
