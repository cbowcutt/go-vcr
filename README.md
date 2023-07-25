#### Regenerating test fixture stubs

```
protoc -I testing/fixtures --go_out=testing/fixtures --go-grpc_out=testing/fixtures -I$GOOGLE_PROTO_DIR:"testing/fixtures" ./testing/fixtures/test_api.proto
```

```
protoc --go_out=example/time_server/api --go-grpc_out=require_unimplemented_servers=false:example/time_server/api -I$GOOGLE_PROTO_DIR:"example/time_server/api" ./example/time_server/api/*.proto
```