package al_types

type ALAvatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type ALViewer struct {
	Name        string   `json:"name"`
	BannerImage string   `json:"bannerImage"`
	Avatar      ALAvatar `json:"avatar"`
}

type ALViewerData struct {
	Viewer ALViewer `json:"Viewer"`
}
