package version

// Version used to indicates where we are.
type Version struct {
	APIVersion string `json:"api_version"`
	BuildDate  string `json:"build_date"`
	GitCommit  string `json:"git_commit"`
}

// APIVersion indicates which API version of the binary is running.
var APIVersion string

// GitCommit indicates which git commit of the binary is running.
var GitCommit string

// BuildDate indicates the build time of the binary is running.
var BuildDate string
