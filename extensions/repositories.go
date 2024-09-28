package extensions

import (
	"encoding/json"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"tsuki/core"
	"tsuki/helpers"
)

type Repository struct {
	Name           string     `json:"name"`
	ID             string     `json:"id"`
	Logo           string     `json:"logo"`
	URL            string     `json:"url"`
	MangaProviders []Provider `json:"manga_providers"`
	ComicProviders []Provider `json:"comic_providers"`
}

func (r *Repository) Update() error {
	err := InstallRepository(r.URL, true)
	return err
}

func (r *Repository) Uninstall() error {
	err := UninstallRepository(r.ID)
	return err
}

func InstallRepository(jsonUrl string, update bool) error {
	client := http.Client{Timeout: 10 * time.Second}

	request, err := helpers.BuildGetRequest(jsonUrl)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var repository Repository
	err = json.NewDecoder(response.Body).Decode(&repository)
	if err != nil {
		return err
	}
	repository.URL = jsonUrl

	// TODO: Check what happens when someone has a funky id
	if repository.Name == "" || repository.ID == "" {
		return errors.New("Repository did not provide necessary information")
	}

	bytes, err := json.Marshal(repository)
	if err != nil {
		return err
	}
	repositoryJson := string(bytes)

	// Now we want to store the json file in the repositories directory
	repositoryFilename := repository.ID + ".json"
	repositoryLocation := filepath.Join(core.CONFIG.Directories.Repositories, repositoryFilename)

	// Check if the repository already exists
	repositoryExists := helpers.FileExists(repositoryLocation)
	if repositoryExists && !update {
		return errors.New("Repository already exists")
	}

	err = helpers.CreateAndWriteToFile(repositoryLocation, repositoryJson)
	if err != nil {
		return err
	}

	return nil
}

// Returns the location, and an optional error, and mutates given repository.
// If the repository could not be found, it returns an empty string and a nil error.
func GetRepository(repositoryId string, repository *Repository) (string, error) {
	// Build the repository location
	repositoryFile := repositoryId + ".json"
	repositoryLocation := filepath.Join(core.CONFIG.Directories.Repositories, repositoryFile)

	// Check if the file exists
	repositoryExists := helpers.FileExists(repositoryLocation)
	if !repositoryExists {
		return "", nil
	}

	// Read the file
	repositoryJson, err := helpers.ReadFileContents(repositoryLocation)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(repositoryJson), repository)
	if err != nil {
		return "", err
	}

	return repositoryLocation, nil
}

func GetRepositories() ([]Repository, error) {
	var repositories []Repository

	err := filepath.WalkDir(core.CONFIG.Directories.Repositories, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".json" {
			return nil
		}

		repositoryId := strings.SplitN(d.Name(), ".", 2)[0]
		var repository Repository
		_, err = GetRepository(repositoryId, &repository)
		if err != nil {
			return err
		}
		repositories = append(repositories, repository)

		return nil
	})

	if repositories == nil {
		return []Repository{}, nil
	}

	return repositories, err
}

func UninstallRepository(repositoryId string) error {
	repositoryLocation := filepath.Join(core.CONFIG.Directories.Repositories, repositoryId+".json")

	err := os.Remove(repositoryLocation)
	return err
}
