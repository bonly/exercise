package main

import (
"github.com/libgit2/git2go"
)

func main(){
	TestCloneWithCallback();
}

func TestCloneWithCallback() {
	testPayload := 0

	repo := createTestRepo(t)
	defer cleanupTestRepo(t, repo)

	seedTestRepo(t, repo)

	path, err := ioutil.TempDir("", "git2go")
	checkFatal(t, err)

	opts := CloneOptions{
		Bare: true,
		RemoteCreateCallback: func(r *Repository, name, url string) (*Remote, ErrorCode) {
			testPayload += 1

			remote, err := r.Remotes.Create(REMOTENAME, url)
			if err != nil {
				return nil, ErrGeneric
			}

			return remote, ErrOk
		},
	}

	repo2, err := Clone(repo.Path(), path, &opts)
	defer cleanupTestRepo(t, repo2)

	checkFatal(t, err)

	if testPayload != 1 {
		fmt.Println("Payload's value has not been changed")
	}

	remote, err := repo2.Remotes.Lookup(REMOTENAME)
	if err != nil || remote == nil {
		fmt.Println("Remote was not created properly")
	}
}
