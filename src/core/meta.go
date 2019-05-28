package core

type Images struct {
	Images map[string]*ImageInfo `yaml:"Images"`
}

type ImageInfo struct {
	ImageName      string `yaml:"ImageName"`
	Crontab        string `yaml:"Crontab"`
	RemotePath     string `yaml:"RemotePath"`
	LocalPath      string `yaml:"LocalPath"`
	CostomFileName string `yaml:"CostomFileName"`
}

type Basic struct {
	RemotePath    string
	IndexFileName string
	LocalPath     string
	ImageName     string
}

type Custom struct {
	Basics    *Basic
	ImageName string
}
