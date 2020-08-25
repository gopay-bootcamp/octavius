# Error Handling

1. Create custom error message using package octaviusErrors while returning the error
    - `errMsg := Errors.New(<error-code>, <error>)`
    - for <error-code> refer the below map
        - 0: for No Error
        - 1: for Client
        - 2: for Control Plane
        - 3: for Etcd Database
        - 4: for Executor  		 

2. Print or Log error message using Error() or directly print `errMsg`

- NOTE:- Use this method only while using any third party functions.