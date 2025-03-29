package backend

import (
	"errors"
	"tsuki/backend/anilist"
	"tsuki/backend/anilist/types"
)

type MangaShowResponse struct {
	MediaList *types.ALMediaList `json:"media_list"`
}

func (a *App) MangaIndex() ([]types.ALMediaListGroup, error) {
	hasAnilistToken := a.verifyAnilistToken()
	if !hasAnilistToken {
		return []types.ALMediaListGroup{}, errors.New("Anilist token is missing")
	}

	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: a.CurrentAccount.AnilistName,
	}
	listCollection, err := anilist.BuildAndSendRequest[types.ALMediaListCollectionData](
		"media_list_collection",
		a.CurrentAccount.AnilistToken,
		nil,
		varUserName,
	)
	if err != nil {
		return []types.ALMediaListGroup{}, errors.New("There was an issue when trying to retrieve the list collection.")
	}

	// for _, groupList := range listCollection.MediaListCollection.Lists {
	// 	for _, list := range groupList.Entries {
	// 		list.SetMangaMapping(*account)
	// 	}
	// }

	return listCollection.MediaListCollection.Lists, nil
}

func (a *App) MangaShow(id int) (*MangaShowResponse, error) {
	hasAnilistToken := a.verifyAnilistToken()
	if !hasAnilistToken {
		return nil, errors.New("Anilist token is missing")
	}

	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: a.CurrentAccount.AnilistName,
	}
	varMediaId := anilist.GraphQLVariable{
		Key:   "mediaId",
		Value: id,
	}
	mediaList, err := anilist.BuildAndSendRequest[types.ALMediaListData](
		"media_list",
		a.CurrentAccount.AnilistToken,
		nil,
		varUserName,
		varMediaId,
	)
	if err != nil {
		return nil, errors.New("There was an issue when trying to retrieve the list.")
	}

	// mediaList.MediaList.SetMangaMapping(*account)
	//
	// Get the chapter list and send it in the response
	// chapterList := []models.Chapter{}
	// if mediaList.MediaList.Mapping != nil {
	// 	// TODO: This should not be doing this. We should instead hit the database directly.
	// 	chapterList, _ = jobs.RetrieveChaptersForMapping(*mediaList.MediaList.Mapping)
	// }

	data := MangaShowResponse{
		MediaList: &mediaList.MediaList,
	}

	return &data, nil
}

func (a *App) verifyAnilistToken() bool {
	if a.CurrentAccount == nil || a.CurrentAccount.AnilistToken == "" {
		return false
	}
	return true
}
