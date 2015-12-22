package repofiles

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Repo struct {
	user     string
	repo     string
	revision string
	list     List
	files    []File
}

type List struct {
	Sha  string
	Url  string
	Tree []struct {
		Path   string
		Mode   string
		Format string
		Sha    string
		Size   int
		Url    string
	}
	Trunctated bool
}

type File struct {
	name     string
	contents string
}

func (r *Repo) List() List {
	resp, _ := http.Get(r.listEndpoint())
	respBody, _ := ioutil.ReadAll(resp.Body)

	list := List{}
	json.Unmarshal([]byte(respBody), &list)

	r.list = list

	return list
}

func (r *Repo) Files() []File {
	type jsonFile struct {
		Sha      string
		Size     int
		Url      string
		Content  string
		Encoding string
	}
	var parsedFile jsonFile
	for _, v := range r.list.Tree {
		resp, _ := http.Get(v.Url)
		respBody, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal([]byte(respBody), &parsedFile)
		contents, _ := base64.StdEncoding.DecodeString(parsedFile.Content)
		r.files = append(r.files, File{name: v.Path, contents: string(contents)})
	}
	return r.files
}

func (r *Repo) listEndpoint() string {
	segments := []string{
		"https://api.github.com/repos", r.user, r.repo,
		"git/trees/" + r.revision + "?recursive=1",
	}
	return strings.Join(segments, "/")
}

func NewRepo(user string, repo string, revision string) Repo {
	return Repo{user: user, repo: repo, revision: revision}
}
