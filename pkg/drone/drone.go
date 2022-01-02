package drone

import (
	"errors"
	"os"
	"strings"

	"github.com/yaskinny/discord-notif/pkg/models"
)

var DroneEnvs = map[string]string{
	// https://docs.drone.io/pipeline/environment/reference/
	"DroneBranch":             os.Getenv("DRONE_BRANCH"),
	"DroneBuildAction":        os.Getenv("DRONE_BUILD_ACTION"),
	"DroneBuildCreated":       os.Getenv("DRONE_BUILD_CREATED"),
	"DroneBuildEvent":         os.Getenv("DRONE_BUILD_EVENT"),
	"DroneBuildFinished":      os.Getenv("DRONE_BUILD_FINISHED"),
	"DroneBuildNumber":        os.Getenv("DRONE_BUILD_NUMBER"),
	"DroneBuildParent":        os.Getenv("DRONE_BUILD_PARENT"),
	"DroneBuildStarted":       os.Getenv("DRONE_BUILD_STARTED"),
	"DroneBuildStatus":        os.Getenv("DRONE_BUILD_STATUS"),
	"DroneCalver":             os.Getenv("DRONE_CALVER"),
	"DroneCommit":             os.Getenv("DRONE_COMMIT"),
	"DroneCommitAfter":        os.Getenv("DRONE_COMMIT_AFTER"),
	"DroneCommitAuthor":       os.Getenv("DRONE_COMMIT_AUTHOR"),
	"DroneCommitAuthorAvatar": os.Getenv("DRONE_COMMIT_AUTHOR_AVATAR"),
	"DroneCommitAuthorEmail":  os.Getenv("DRONE_COMMIT_AUTHOR_EMAIL"),
	"DroneCommitAuthorName":   os.Getenv("DRONE_COMMIT_AUTHOR_NAME"),
	"DroneCommitBefore":       os.Getenv("DRONE_COMMIT_BEFORE"),
	"DroneCommitBranch":       os.Getenv("DRONE_COMMIT_BRANCH"),
	"DroneCommitLink":         os.Getenv("DRONE_COMMIT_LINK"),
	"DroneCommitMessage":      os.Getenv("DRONE_COMMIT_MESSAGE"),
	"DroneCommitRef":          os.Getenv("DRONE_COMMIT_REF"),
	"DroneCommitSha":          os.Getenv("DRONE_COMMIT_SHA"),
	"DroneDeployTo":           os.Getenv("DRONE_DEPLOY_TO"),
	"DroneFailedStages":       os.Getenv("DRONE_FAILED_STAGES"),
	"DroneFailedSteps":        os.Getenv("DRONE_FAILED_STEPS"),
	"DroneGitHttpUrl":         os.Getenv("DRONE_GIT_HTTP_URL"),
	"DroneGitSshUrl":          os.Getenv("DRONE_GIT_SSH_URL"),
	"DronePullRequest":        os.Getenv("DRONE_PULL_REQUEST"),
	"DroneRemoteUrl":          os.Getenv("DRONE_REMOTE_URL"),
	"DroneRepo":               os.Getenv("DRONE_REPO"),
	"DroneRepoBranch":         os.Getenv("DRONE_REPO_BRANCH"),
	"DroneRepoLink":           os.Getenv("DRONE_REPO_LINK"),
	"DroneRepoName":           os.Getenv("DRONE_REPO_NAME"),
	"DroneRepoNamespace":      os.Getenv("DRONE_REPO_NAMESPACE"),
	"DroneRepoOwner":          os.Getenv("DRONE_REPO_OWNER"),
	"DroneRepoPrivate":        os.Getenv("DRONE_REPO_PRIVATE"),
	"DroneRepoScm":            os.Getenv("DRONE_REPO_SCM"),
	"DroneRepoVisibility":     os.Getenv("DRONE_REPO_VISIBILITY"),
	"DroneSemver":             os.Getenv("DRONE_SEMVER"),
	"DroneSemverBuild":        os.Getenv("DRONE_SEMVER_BUILD"),
	"DroneSemverError":        os.Getenv("DRONE_SEMVER_ERROR"),
	"DroneSemverMajor":        os.Getenv("DRONE_SEMVER_MAJOR"),
	"DroneSemverMinor":        os.Getenv("DRONE_SEMVER_MINOR"),
	"DroneSemverPatch":        os.Getenv("DRONE_SEMVER_PATCH"),
	"DroneSemverPrerelease":   os.Getenv("DRONE_SEMVER_PRERELEASE"),
	"DroneSemverShort":        os.Getenv("DRONE_SEMVER_SHORT"),
	"DroneSourceBranch":       os.Getenv("DRONE_SOURCE_BRANCH"),
	"DroneStageArch":          os.Getenv("DRONE_STAGE_ARCH"),
	"DroneStageDependsOn":     os.Getenv("DRONE_STAGE_DEPENDS_ON"),
	"DroneStageFinished":      os.Getenv("DRONE_STAGE_FINISHED"),
	"DroneStageKind":          os.Getenv("DRONE_STAGE_KIND"),
	"DroneStageMachine":       os.Getenv("DRONE_STAGE_MACHINE"),
	"DroneStageName":          os.Getenv("DRONE_STAGE_NAME"),
	"DroneStageNumber":        os.Getenv("DRONE_STAGE_NUMBER"),
	"DroneStageOs":            os.Getenv("DRONE_STAGE_OS"),
	"DroneStageStarted":       os.Getenv("DRONE_STAGE_STARTED"),
	"DroneStageStatus":        os.Getenv("DRONE_STAGE_STATUS"),
	"DroneStageType":          os.Getenv("DRONE_STAGE_TYPE"),
	"DroneStageVariant":       os.Getenv("DRONE_STAGE_VARIANT"),
	"DroneStepName":           os.Getenv("DRONE_STEP_NAME"),
	"DroneStepNumber":         os.Getenv("DRONE_STEP_NUMBER"),
	"DroneSystemHost":         os.Getenv("DRONE_SYSTEM_HOST"),
	"DroneSystemHostname":     os.Getenv("DRONE_SYSTEM_HOSTNAME"),
	"DroneSystemProto":        os.Getenv("DRONE_SYSTEM_PROTO"),
	"DroneSystemVersion":      os.Getenv("DRONE_SYSTEM_VERSION"),
	"DroneTag":                os.Getenv("DRONE_TAG"),
	"DroneTargetBranch":       os.Getenv("DRONE_TARGET_BRANCH"),
}

type Drone struct {
	//
}

// goes through NOTIF_FIELDS environment variable and set the ones user set for template
func (d Drone) Setter() ([]models.Field, error) {
	var fields []models.Field
	droneEnvs := os.Getenv(`NOTIF_FIELDS`)
	if droneEnvs == "" {
		// maybe set a default base template instead of returing errors?
		return nil, errors.New("There is not tag set to create template, please set `NOTIF_FIELDS`")
	}
	droneTags := strings.Split(droneEnvs, ",")
	for _, v := range droneTags {
		for key, value := range DroneEnvs {
			if v == key {
				if DroneEnvs[key] != "" {
					fields = append(fields, models.Field{
						Name:   key,
						Value:  value,
						Inline: "false",
					})
				}
			}
		}
	}
	return fields, nil
}
