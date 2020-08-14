package constant

var ErrorCode  = map[int]string{
	0: "No Error",
	1: "Client",
	2: "Control Plane",
	3: "Etcd Database",
	4: "Executor",
}

const ClientError = "malformed request"
const ServerError = "Something went wrong"

const JobNotFound = "Job not found"
const JobSucceeded = "SUCCEEDED"
const JobFailed = "FAILED"
const JobWaiting = "WAITING"
