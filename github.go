package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	API_HOST = "https://api.github.com/%s"
)

type User struct {
	ID            int
	Login         string
	Email         string
	Name          string
	RepoCount     int `json:"public_repos"`
	FollowerCount int `json:"followers"`
}

func (u *User) String() string {
	return fmt.Sprintf("[%d] %s", u.ID, u.Login)
}

type Repo struct {
	ID            int
	Name          string
	Homepage      string
	CloneURL      string
	GitURL        string
	HtmlURL       string `json:"html_url"`
	WatchersCount int    `json:"watchers_count"`
	Description   string
	Language      string
}

func (r *Repo) String() string {
	return fmt.Sprintf("[%d] %s @ %s", r.ID, r.Name, r.HtmlURL)
}

func get(path string, resource interface{}) error {
	url := fmt.Sprintf(API_HOST, path)
	res, err := http.Get(url)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(resource)

	if err != nil {
		return err
	}

	return nil
}

func GetUser(username string) (*User, error) {
	user := new(User)
	path := fmt.Sprintf("users/%s", username)
	err := get(path, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetRepo(reponame string) (*Repo, error) {
	repo := new(Repo)
	path := fmt.Sprintf("repos/%s", reponame)
	err := get(path, repo)

	if err != nil {
		return nil, err
	}

	return repo, nil
}
