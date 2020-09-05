package constant

// what is the difference between error code and constant error below?

var ErrorCode = map[int]string{
	0: "No Error",
	1: "Client",
	2: "Control Plane",
	3: "Etcd Database",
	4: "Executor",
}

// group the const
const (

	// error messages
	ConfigOctaviusHostMissingError = "Config Error\nMandatory config CP_HOST is missing in Octavius Config file."
	ClientError                    = "malformed request"
	ServerError                    = "Something went wrong"
	NoValueFound                   = "no value found"
	KeyAlreadyPresent              = "key already present"
	EtcdSaveError                  = "error in saving to etcd"
	JobNotFound                    = "Job not found"
	JobSucceeded                   = "SUCCEEDED"
	JobFailed                      = "FAILED"
	JobWaiting                     = "WAITING"
	NullRevision                   = -1

	// LoggerSkipFrameCount is for ...
	LoggerSkipFrameCount = 3

	//etcd namespace prefixes
	MetadataPrefix   = "metadata/"
	PendingJobPrefix = "jobs/pending/"
)
