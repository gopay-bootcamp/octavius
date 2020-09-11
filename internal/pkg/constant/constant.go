package constant

// group the const
const (

	//error source
	Client     = "in client: "
	Controller = "in controller: "
	Etcd       = "in etcd database: "
	Executor   = "in executor: "

	// error messages
	ConfigOctaviusHostMissingError      = "config error\nmandatory config CP_HOST is missing in octavius config file."
	ClientError                         = "malformed request"
	TimeOutError                        = "timeout when waiting job to be available"
	ExecutionKey                        = "octavius"
	OutOfClustor                        = "out-of-cluster"
	ServerError                         = "something went wrong"
	NoValueFound                        = "no value found"
	KeyAlreadyPresent                   = "key already present"
	EtcdSaveError                       = "error in saving to etcd"
	JobNotFound                         = "job not found"
	JobExecutionStatusFetchError        = "job execution status fetch error"
	NoDefinitiveJobExecutionStatusFound = "no definitive job execution status found"
	JobSucceeded                        = "succeeded"
	JobFailed                           = "failed"
	JobWaiting                          = "waiting"
	NullRevision                        = -1

	// LoggerSkipFrameCount is for ...
	LoggerSkipFrameCount = 3

	//etcd namespace prefixes
	MetadataPrefix             = "metadata/"
	JobPendingPrefix           = "jobs/pending/"
	ExecutorRegistrationPrefix = "executor/register/"
	ExecutorStatusPrefix       = "executor/status/"
	ExecutionDataPrefix        = "executor/logs/"
	K8sJobPrefix               = "context/octavius"

	// executor daemon constants
	Received          = "RECEIVED"
	RequirementNotMet = "REQUIREMENT_NOT_MET"
	Created           = "CREATED"
	CreationFailed    = "CREATION_FAILED"
	JobCreationFailed = "JOB_CREATION_FAILED"
	JobReady          = "JOB_READY"
	PodCreationFailed = "POD_CREATION_FAILED"
	PodReady          = "POD_READY"
	PodFailed         = "POD_FAILED"
	FetchPodLogFailed = "FETCH_POD_LOG_FAILED"
	Finished          = "FINISHED"
	IdleState         = "idle"
	RunningState      = "running"
)
