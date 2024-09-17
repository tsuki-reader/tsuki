package anilist

type Avatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type Viewer struct {
	Name        string `json:"name"`
	BannerImage string `json:"bannerImage"`
	Avatar      Avatar `json:"avatar"`
}

type ViewerData struct {
	Viewer Viewer `json:"Viewer"`
}
