package GoCmd

import (
	"github.com/jfrog/gocmd/executers"
	"github.com/jfrog/jfrog-client-go/artifactory"
)

func RecursivePublish(targetRepo, goModEditMessage string, serviceManager *artifactory.ArtifactoryServicesManager) error {
	return executers.RecursivePublish(targetRepo, goModEditMessage, serviceManager)
}

func RunWithFallbacksAndPublish(goArg, targetRepo string, noRegistry bool, serviceManager *artifactory.ArtifactoryServicesManager) error {
	return executers.RunWithFallbacksAndPublish(goArg, targetRepo, noRegistry, serviceManager)
}

func RunWithFallback(goArg string, url string) error {
	return executers.RunWithFallback(goArg, url)
}