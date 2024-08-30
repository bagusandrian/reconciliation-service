package config

import (
	"fmt"
	"os"
	"path/filepath"

	cons "github.com/bagusandrian/reconciliation-service/internals/constant"
	"github.com/bagusandrian/reconciliation-service/internals/model"
	yaml "gopkg.in/yaml.v2"
)

func New(repoName string) (*model.Config, error) {
	filename := getConfigFile(repoName)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg model.Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// getConfigFile get  config file name
// - files/etc/repo_name/repo_name.development.yaml in dev
// - otherwise /etc/repo_name/repo_name.{TKPENV}.yaml
func getConfigFile(repoName string) string {
	var (
		SysEnv   = getEnv()
		filename = fmt.Sprintf("%s.%s.yaml", repoName, SysEnv)
	)

	// for non dev env, use config in /etc
	if SysEnv == cons.TestingEnv {
		return filepath.Join("/etc", repoName, fmt.Sprintf("%s.%s.yaml", repoName, cons.DevelopmentEnv))
	}
	if SysEnv == cons.DevelopmentEnv || SysEnv == "" {
		// use local files in dev
		repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/bagusandrian", repoName)
		return filepath.Join(repoPath, "files/etc", repoName, filename)
	}
	return filepath.Join("/etc", repoName, string(SysEnv), filename)
}

func getEnv() string {
	env := os.Getenv("SysEnv")
	if env == "" {
		return "development"
	}
	return env
}
