package backend

import (
	"tsuki/backend/extensions"
)

func (a *App) RepositoriesIndex() ([]extensions.Repository, error) {
	repositories, err := extensions.GetRepositories()
	if err != nil {
		return []extensions.Repository{}, err
	}

	return repositories, nil
}

func (a *App) RepositoriesCreate(url string) ([]extensions.Repository, error) {
	var repository extensions.Repository
	if err := extensions.InstallRepository(url, false, &repository); err != nil {
		return []extensions.Repository{}, err
	}

	repositories, err := extensions.GetRepositories()
	if err != nil {
		return []extensions.Repository{}, err
	}

	return repositories, nil
}

func (a *App) RepositoriesUpdate(id string) (extensions.Repository, error) {
	var repository extensions.Repository
	if _, err := extensions.GetRepository(id, &repository); err != nil {
		return extensions.Repository{}, err
	}

	if err := repository.Update(); err != nil {
		return extensions.Repository{}, err
	}

	return repository, nil
}

func (a *App) RepositoriesDestroy(id string) ([]extensions.Repository, error) {
	var repository extensions.Repository
	if _, err := extensions.GetRepository(id, &repository); err != nil {
		return []extensions.Repository{}, err
	}

	if err := repository.Uninstall(); err != nil {
		return []extensions.Repository{}, err
	}

	repositories, err := extensions.GetRepositories()
	if err != nil {
		return []extensions.Repository{}, err
	}

	return repositories, nil
}
