package repofiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	repo := NewRepo("charlieegan3", "dotfiles", "efba099dae2acc7daac7d0638ba1be39a45914ab")
	expected := Repo{user: "charlieegan3", repo: "dotfiles", revision: "efba099dae2acc7daac7d0638ba1be39a45914ab"}
	assert.Equal(t, expected, repo, "should be equal")
}

func TestList(t *testing.T) {
	repo := NewRepo("charlieegan3", "dotfiles", "efba099dae2acc7daac7d0638ba1be39a45914ab")
	list := repo.List()
	var files []string
	for _, v := range list.Tree {
		files = append(files, v.Path)
	}
	expectedFiles :=
		[]string{".bashrc", ".gitconfig", ".gitignore", "com.googlecode.iterm2.plist"}
	assert.Equal(t, expectedFiles, files, "should be equal")
}

func TestFiles(t *testing.T) {
	repo := NewRepo("charlieegan3", "dotfiles", "efba099dae2acc7daac7d0638ba1be39a45914ab")
	repo.List()
	repo.Files()

	assert.Equal(t, repo.files[0].name, ".bashrc", "should be equal")
}
