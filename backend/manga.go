package backend

import (
	"errors"
	"tsuki/backend/anilist"
	"tsuki/backend/anilist/types"
)

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

func (a *App) verifyAnilistToken() bool {
	if a.CurrentAccount == nil || a.CurrentAccount.AnilistToken == "" {
		return false
	}
	return true
}
