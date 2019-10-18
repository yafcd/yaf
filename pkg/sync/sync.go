package sync

import (
	"os"
	"gopkg.in/src-d/go-git.v4"
	// "gopkg.in/src-d/go-git.v4/storage/memory"
)

type Sync struct {

}

type SyncConfig struct {
	All bool //Sync all the files in the repo
}

//SyncFn syncs
func  SyncFn() {	
	r, err := git.PlainClone("/tmp/tmp-go-siva", false, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva",
		Progress: os.Stdout,
	})

	if err != nil {
		return
	}

	w ,err:= r.Worktree()
	if err != nil {
		return
	}

	w.Checkout(&git.CheckoutOptions{
		Branch: "master",
		// Hash is the hash of the commit to be checked out. If used, HEAD will be
		// in detached mode. If Create is not used, Branch and Hash are mutually
		// exclusive.
		//Hash plumbing.Hash
		// Branch to be checked out, if Branch and Hash are empty is set to `master`.
		
		// Create a new branch named Branch and start it at Hash.
		//Create bool
		// Force, if true when switching branches, proceed even if the index or the
		// working tree differs from HEAD. This is used to throw away local changes
		//Force bool
		// Keep, if true when switching branches, local changes (the index or the
		// working tree changes) will be kept so that they can be committed to the
		// target branch. Force and Keep are mutually exclusive, should not be both
		// set to true.
		//Keep bool
	})

	


}

