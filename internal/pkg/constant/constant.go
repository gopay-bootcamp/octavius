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
	ConfigOctaviusHostMissingError = "config error: mandatory config CP_HOST is missing in octavius config file."
	ClientError                    = "malformed request"
	ServerError                    = "something went wrong"
	NoValueFound                   = "no value found"
	KeyAlreadyPresent              = "key already present"
	EtcdSaveError                  = "error in saving to etcd"
	JobNotFound                    = "job not found"
	JobSucceeded                   = "succeeded"
	JobFailed                      = "failed"
	JobWaiting                     = "waiting"

	// LoggerSkipFrameCount is for ...
	LoggerSkipFrameCount = 3
)
