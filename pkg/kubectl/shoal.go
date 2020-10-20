package kubectl

import (
	"errors"
	"fmt"
	"github.com/mumoshu/shoal"
	"os"
	"path/filepath"
	"sync"
)

var shoalMu sync.Mutex

func prepareBinaries(fs *Ensure) (*string, error) {
	conf := shoal.Config{
		Git: shoal.Git{
			Provider: "go-git",
		},
	}

	rig := "https://github.com/fishworks/fish-food"

	kubectlBin := fs.Bin

	kubectlVersion := fs.Version

	installKubectl := kubectlVersion != ""

	if installKubectl {
		conf.Dependencies = append(conf.Dependencies,
			shoal.Dependency{
				Rig:     rig,
				Food:    "kubectl",
				Version: kubectlVersion,
			},
		)
	}

	shoalMu.Lock()
	defer shoalMu.Unlock()

	s, err := shoal.New()
	if err != nil {
		return nil, err
	}

	if len(conf.Dependencies) > 0 {
		if err := s.Init(); err != nil {
			return nil, fmt.Errorf("initializing shoal: %w", err)
		}

		if err := s.InitGitProvider(conf); err != nil {
			return nil, fmt.Errorf("initializing shoal git provider: %w", err)
		}

		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		// TODO Any better place to do this?
		// This is for letting helm know about the location of helm plugins installed by shoal
		os.Setenv("XDG_DATA_HOME", filepath.Join(wd, ".shoal/Library"))

		if err := s.Sync(conf); err != nil {
			return nil, err
		}
	}

	binPath := s.BinPath()

	if kubectlVersion != "" {
		kubectlBin = filepath.Join(binPath, "kubectl")
	}

	if kubectlBin == "" {
		return nil, errors.New("bug: kubectl_ensure.bin is missing")
	}

	return &kubectlBin, nil
}
