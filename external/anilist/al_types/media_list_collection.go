package al_types

type ALTitle struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
	Native  string `json:"native"`
}

type ALCoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
	Medium     string `json:"medium"`
	Colour     string `json:"color"`
}

type ALManga struct {
	Id          int          `json:"id"`
	Title       ALTitle      `json:"title"`
	StartDate   ALDate       `json:"startDate"`
	Status      string       `json:"status"`
	Chapters    int          `json:"chapters"`
	Volumes     int          `json:"volumes"`
	CoverImage  ALCoverImage `json:"coverImage"`
	BannerImage string       `json:"bannerImage"`
	Description string       `json:"description"`
	Genres      []string     `json:"genres"`
}

type ALDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type ALProgressEntry struct {
	Progress    int     `json:"progress"`
	CompletedAt ALDate  `json:"completedAt"`
	StartedAt   ALDate  `json:"startedAt"`
	Notes       string  `json:"notes"`
	Score       int     `json:"score"`
	Status      string  `json:"status"`
	Media       ALManga `json:"media"`
}

type ALMediaList struct {
	Name                 string            `json:"name"`
	IsCustomList         bool              `json:"isCustomList"`
	IsSplitCompletedList bool              `json:"isSplitCustomList"`
	Status               string            `json:"status"`
	Entries              []ALProgressEntry `json:"entries"`
}

type ALMediaListCollection struct {
	Lists []ALMediaList `json:"lists"`
}

type ALMediaListCollectionData struct {
	MediaListCollection ALMediaListCollection `json:"MediaListCollection"`
}
