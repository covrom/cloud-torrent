package engine

type Config struct {
	AutoStart          bool
	DisableEncryption  bool
	DownloadDirectory  string
	EnableUpload       bool
	EnableSeeding      bool
	IncomingPort       int
	WatchDirs          string
	UploadKbps         int
	DownloadKbps       int
	DeleteAfterMinutes int
}
