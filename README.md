#### Regenerating test fixture stubs

```
protoc -I testing/fixtures --go_out=testing/fixtures --go-grpc_out=testing/fixtures -I$GOOGLE_PROTO_DIR:"testing/fixtures" ./testing/fixtures/test_api.proto
```