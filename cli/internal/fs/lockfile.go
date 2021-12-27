package fs

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ReadLockfile will read `yarn.lock` into memory (either from the cache or fresh)
func ReadLockfile(cacheDir string) (*YarnLockfile, error) {
	var prettyLockFile = YarnLockfile{}

	hash, err := HashFile("yarn.lock")
	if err != nil {
		return &YarnLockfile{}, fmt.Errorf("failed to hash lockfile: %w", err)
	}

	contentsOfLock, err := ioutil.ReadFile(filepath.Join(cacheDir, fmt.Sprintf("%v-turbo-lock.yaml", hash)))
	if err != nil {

		prettyLockFile, err := parseLockfile("yarn.lock")
		if err != nil {
			return &YarnLockfile{}, fmt.Errorf("failed to read lockfile: %w", err)
		}

		better, err := yaml.Marshal(&prettyLockFile)
		if err != nil {
			fmt.Println(err.Error())
			return &YarnLockfile{}, err
		}
		if err = EnsureDir(cacheDir); err != nil {
			fmt.Println(err.Error())
			return &YarnLockfile{}, err
		}
		if err = EnsureDir(filepath.Join(cacheDir, fmt.Sprintf("%v-turbo-lock.yaml", hash))); err != nil {
			fmt.Println(err.Error())
			return &YarnLockfile{}, err
		}
		if err = ioutil.WriteFile(filepath.Join(cacheDir, fmt.Sprintf("%v-turbo-lock.yaml", hash)), []byte(better), 0644); err != nil {
			fmt.Println(err.Error())
			return &YarnLockfile{}, err
		}
	} else {
		if contentsOfLock != nil {
			err = yaml.Unmarshal(contentsOfLock, &prettyLockFile)
			if err != nil {
				return &YarnLockfile{}, fmt.Errorf("could not unmarshal yaml: %w", err)
			}
		}
	}

	return &prettyLockFile, nil
}
