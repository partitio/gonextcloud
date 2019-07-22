package gonextcloud

import (
	"os"
	"path/filepath"
	"sort"

	"gitlab.bertha.cloud/adphi/gowebdav"
)

type webDav struct {
	*gowebdav.Client
}

func newWebDav(url string, user string, password string) *webDav {
	wb := gowebdav.NewClient(url, user, password)
	return &webDav{Client: wb}
}

// Implementation adapted from filepath.Walk

// Walk walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root. All errors that arise visiting files
// and directories are filtered by walkFn. The files are walked in lexical
// order, which makes the output deterministic but means that for very
// large directories Walk can be inefficient.
// Walk does not follow symbolic links.
func (wd *webDav) Walk(root string, walkFn filepath.WalkFunc) error {
	info, err := wd.Stat(root)
	if err != nil {
		err = walkFn(root, nil, err)
	} else {
		err = wd.walk(root, info, walkFn)
	}
	if err == filepath.SkipDir {
		return nil
	}
	return err
}

// walk recursively descends path, calling walkFn.
func (wd *webDav) walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	if !info.IsDir() {
		return walkFn(path, info, nil)
	}
	fis, err := wd.readDir(path)
	err1 := walkFn(path, info, err)
	// If err != nil, walk can't walk into this directory.
	// err1 != nil means walkFn want walk to skip this directory or stop walking.
	// Therefore, if one of err and err1 isn't nil, walk will return.
	if err != nil || err1 != nil {
		// The caller's behavior is controlled by the return value, which is decided
		// by walkFn. walkFn may ignore err and return nil.
		// If walkFn returns SkipDir, it will be handled by the caller.
		// So walk should return whatever walkFn returns.
		return err1
	}

	for _, fi := range fis {
		filename := filepath.Join(path, fi.Name())
		err = wd.walk(filename, fi, walkFn)
		if err != nil {
			if !fi.IsDir() || err != filepath.SkipDir {
				return err
			}
		}
	}
	return nil
}

// readDir reads the directory and returns
// a sorted list of directory entries.
func (wd *webDav) readDir(dirname string) ([]os.FileInfo, error) {
	fs, err := wd.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	sort.Slice(fs, func(i, j int) bool {
		return sort.StringsAreSorted([]string{fs[i].Name(), fs[j].Name()})
	})
	return fs, nil
}
