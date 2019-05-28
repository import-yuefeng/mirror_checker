package core

type Images struct {
	Images map[string]*ImageInfo `yaml:"Images"`
}

type ImageInfo struct {
	ImageName      string `yaml:"ImageName"`
	Crontab        string `yaml:"Crontab"`
	CostomPath     string `yaml:"CostomPath"`
	CostomFileName string `yaml:"CostmonFileName"`
}

type Basic struct {
	Path     string
	FileName string
}
