package al_types

type ALMediaList struct {
	Progress    int     `json:"progress"`
	CompletedAt ALDate  `json:"completedAt"`
	StartedAt   ALDate  `json:"startedAt"`
	Notes       string  `json:"notes"`
	Score       int     `json:"score"`
	Status      string  `json:"status"`
	Media       ALManga `json:"media"`
}

type ALMediaListData struct {
	MediaList ALMediaList `json:"MediaList"`
}
