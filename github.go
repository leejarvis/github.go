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

type GistFile struct {
	Type     string
	Filename string
	Content  string
	RawURL   string `json:"raw_url"`
	Language string
	Size     int
}

func (gh *GistFile) String() string {
	return fmt.Sprintf("%s (%s)", gh.Filename, gh.Language)
}

type Gist struct {
	ID           string
	URL          string `json:"html_url"`
	Public       bool
	CommentCount int `json:"comments"`
	Description  string
	Files        map[string]*GistFile
	Content      string
}

func (g *Gist) String() string {
	return fmt.Sprintf("%s @ %s", g.ID, g.URL)
}

type Issue struct {
	ID           int
	Title        string
	CommentCount int    `json:"comments"`
	HtmlURL      string `json:"html_url"`
	State        string
	Number       int
	Body         string
}

func (i *Issue) String() string {
	return fmt.Sprintf("[%d] %s (%s)", i.ID, i.Title, i.State)
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

func GetGist(id string) (*Gist, error) {
	gist := new(Gist)
	path := fmt.Sprintf("gists/%s", id)
	err := get(path, gist)

	if err != nil {
		return nil, err
	}

	if len(gist.Files) == 1 {
		for _, gf := range gist.Files {
			gist.Content = gf.Content
		}
	}

	return gist, nil
}

func GetIssues(reponame string) ([]*Issue, error) {
	var issues []*Issue
	path := fmt.Sprintf("repos/%s/issues", reponame)
	err := get(path, &issues)

	if err != nil {
		return nil, err
	}

	return issues, nil
}
