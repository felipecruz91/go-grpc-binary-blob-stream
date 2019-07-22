# grpc-binary-blob-stream

Exercise from the blog post [Chunking large messages with gRPC](https://jbrandhorst.com/post/grpc-binary-blob-stream/) by [Johan Brandhorst
](https://twitter.com/JohanBrandhorst)

Generate proto file:

Inside `..\src\protos` run the following statement:

    λ protoc --go_out=plugins=grpc:. chunker\chunker.proto