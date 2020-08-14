# Error Handling

1. Add `Error` protobuf into your response protobuf of `process.proto` file.
2. For mapping error code to it's description
    - import `"octavius/pkg/constant"`
    - Use `Constant.ErrorCode` map for mapping the error code