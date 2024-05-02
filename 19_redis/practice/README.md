# Redis Go


Official Docs Tutorial: https://redis.io/docs/latest/develop/connect/clients/go/

### Get client library:

```bash
go get github.com/redis/go-redis/v9
```

### Connect

```go
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        DB:		  0,  // use default DB
    })

```

### Simple operations

```go
ctx := context.Background()

err := client.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}

val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}
fmt.Println("foo", val)
```


### Redis data types 

https://redis.io/docs/latest/develop/data-types/





