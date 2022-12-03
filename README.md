# simplekv

A simple distributed key-value store with strong consistency and redundant copies of data on
multiple nodes

# Components of the system

## Routing Server
The routing server is what accepts HTTP requests from users and determines where to store and
retrieve keys in the storage servers. Code located in `pkg/routing/`

## Storage Server
Simple IO server that accepts requests from the routing server and stores/retrieves values.
Code located in `pkg/storage/`

The core of the Storage Server is the StorageEngine interface which tells the storage server how to
store and retreive the value for a key. See `pkg/storage/storageEngine.go`


# API
`PUT /:key`   Saves the value in the request body to a key  
`GET /:key`   Gets the value of a key  


# Test locally with Docker containers
Run `docker-compose up` to spin up 4 Docker containers
- 1 Routing server running on `http://localhost:8080`
- 3 Storage nodes

And each key will be saved to 2 storage nodes (numReplicas = 2)

#### Send local requests to the routing server:

- Write:
```
curl -i -X PUT -H 'Content-type: text/plain' \
  --data-binary 'hello world' \
  http://localhost:8080/exampleKey
```

- Get:
```
curl http://localhost:8080/exampleKey
```

# CLI entrypoint
TODO: publish a single binary to start either type of service

### Start routingServer:
`go run cmd/routingServer/main.go`

Options:  
- `--replicas <num nodes to save each key to>`    
- `--storage <storage node URL>` (include multiple times to include more than 1 storage server)

### Start storageServer:
`go run cmd/storageServer/main.go`


# ROADMAP:
- [ ] API to delete keys
- [ ] API to list all keys
- [ ] Add/Remove storage nodes while running
- [ ] Rebalance keys when nodes are added or removed
