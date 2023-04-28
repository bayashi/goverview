package fileinfo

type FileInfo struct {
	Icon         string
	Tag          string
	Descriptions []string
}

type Args struct {
	FilePath string
	ShowAll  bool
	HideTest bool
}
