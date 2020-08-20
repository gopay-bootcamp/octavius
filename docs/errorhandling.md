# Error Handling

1. Add `Error` protobuf(message) into your MetadataName protobuf(message) of `pkg/protobuf/process.proto` file.

2. Returning the error in MetadataName
    - import `"octavius/pkg/constant"`
    - return `errorCode`(Integer) according to current binary
       - 0: for No Error
       - 1: for Client
       - 2: for Control Plane
       - 3: for Etcd Database
       - 4: for Executor   
    - return the `errorMessage` using `Constant.<constant-name> `
    - If `<constant-name>` is not in the package `constant`, define your constant with the description in `pkg/constant/constant.go`

3. For mapping error code to it's description
    - import `"octavius/pkg/constant"`
    - Use `Constant.ErrorCode` map for mapping the error code