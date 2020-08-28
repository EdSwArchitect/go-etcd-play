package main

// TheWriters type
type TheWriters interface{}

// CommonLog json
type CommonLog struct {
	DisableLog4J  bool       `json:"disableLog4J"`
	MaxHistorical int        `json:"maxHistorical"`
	Writers       TheWriters `json:"writers"`
}

// FileType the file type stuff
type FileType struct {
	Delay          int    `json:"delay"`
	Enabled        bool   `json:"enabled"`
	FilenamePrefix string `json:"filenamePrefix"`
	FilenameSuffix string `json:"filenameSuffix"`
	FolderPath     string `json:"folderPath"`
	Levels         string `json:"levels"`
	RollOverDays   int    `json:"rollOverDays"`
	TheType        string `json:"type"`
}
