package gitlab

import (
	"fmt"
	"os"
	"strings"

	"github.com/yaskinny/discord-notif/pkg/models"
)

var GitlabEnvs = map[string]string{
	"CiApiV4Url":                              os.Getenv(`CI_API_V4_URL`),
	"CiBuildsDir":                             os.Getenv(`CI_BUILDS_DIR`),
	"CiCommitAuthor":                          os.Getenv(`CI_COMMIT_AUTHOR`),
	"CiCommitBeforeSha":                       os.Getenv(`CI_COMMIT_BEFORE_SHA`),
	"CiCommitBranch":                          os.Getenv(`CI_COMMIT_BRANCH`),
	"CiCommitDescription":                     os.Getenv(`CI_COMMIT_DESCRIPTION`),
	"CiCommitMessage":                         os.Getenv(`CI_COMMIT_MESSAGE`),
	"CiCommitRefProtected":                    os.Getenv(`CI_COMMIT_REF_PROTECTED`),
	"CiCommitRefSlug":                         os.Getenv(`CI_COMMIT_REF_SLUG`),
	"CiCommitSha":                             os.Getenv(`CI_COMMIT_SHA`),
	"CiCommitShortSha":                        os.Getenv(`CI_COMMIT_SHORT_SHA`),
	"CiCommitTag":                             os.Getenv(`CI_COMMIT_TAG`),
	"CiCommitTimestamp":                       os.Getenv(`CI_COMMIT_TIMESTAMP`),
	"CiCommitTitle":                           os.Getenv(`CI_COMMIT_TITLE`),
	"CiConcurrentId":                          os.Getenv(`CI_CONCURRENT_ID`),
	"CiConcurrentProjectId":                   os.Getenv(`CI_CONCURRENT_PROJECT_ID`),
	"CiConfigPath":                            os.Getenv(`CI_CONFIG_PATH`),
	"CiDebugTrace":                            os.Getenv(`CI_DEBUG_TRACE`),
	"CiDefaultBranch":                         os.Getenv(`CI_DEFAULT_BRANCH`),
	"CiDependencyProxyGroupImagePrefix":       os.Getenv(`CI_DEPENDENCY_PROXY_GROUP_IMAGE_PREFIX`),
	"CiDependencyProxyDirectGroupImagePrefix": os.Getenv(`CI_DEPENDENCY_PROXY_DIRECT_GROUP_IMAGE_PREFIX`),
	"CiDependencyProxyPassword":               os.Getenv(`CI_DEPENDENCY_PROXY_PASSWORD`),
	"CiDependencyProxyServer":                 os.Getenv(`CI_DEPENDENCY_PROXY_SERVER`),
	"CiDependencyProxyUser":                   os.Getenv(`CI_DEPENDENCY_PROXY_USER`),
	"CiDeployFreeze":                          os.Getenv(`CI_DEPLOY_FREEZE`),
	"CiDeployPassword":                        os.Getenv(`CI_DEPLOY_PASSWORD`),
	"CiDeployUser":                            os.Getenv(`CI_DEPLOY_USER`),
	"CiDisposableEnvironment":                 os.Getenv(`CI_DISPOSABLE_ENVIRONMENT`),
	"CiEnvironmentName":                       os.Getenv(`CI_ENVIRONMENT_NAME`),
	"CiEnvironmentSlug":                       os.Getenv(`CI_ENVIRONMENT_SLUG`),
	"CiEnvironmentUrl":                        os.Getenv(`CI_ENVIRONMENT_URL`),
	"CiEnvironmentAction":                     os.Getenv(`CI_ENVIRONMENT_ACTION`),
	"CiEnvironmentTier":                       os.Getenv(`CI_ENVIRONMENT_TIER`),
	"CiHasOpenRequirements":                   os.Getenv(`CI_HAS_OPEN_REQUIREMENTS`),
	"CiJobId":                                 os.Getenv(`CI_JOB_ID`),
	"CiJobImage":                              os.Getenv(`CI_JOB_IMAGE`),
	"CiJobJwtV1":                              os.Getenv(`CI_JOB_JWT_V1`),
	"CiJobJwt":                                os.Getenv(`CI_JOB_JWT`),
	"CiJobJwtV2":                              os.Getenv(`CI_JOB_JWT_V2`),
	"CiJobManual":                             os.Getenv(`CI_JOB_MANUAL`),
	"CiJobName":                               os.Getenv(`CI_JOB_NAME`),
	"CiJobStage":                              os.Getenv(`CI_JOB_STAGE`),
	"CiJobStatus":                             os.Getenv(`CI_JOB_STATUS`),
	"CiJobToken":                              os.Getenv(`CI_JOB_TOKEN`),
	"CiJobUrl":                                os.Getenv(`CI_JOB_URL`),
	"CiJobStartedAt":                          os.Getenv(`CI_JOB_STARTED_AT`),
	"CiKubernetesActive":                      os.Getenv(`CI_KUBERNETES_ACTIVE`),
	"CiNodeIndex":                             os.Getenv(`CI_NODE_INDEX`),
	"CiNodeTotal":                             os.Getenv(`CI_NODE_TOTAL`),
	"CiOpenMergeRequests":                     os.Getenv(`CI_OPEN_MERGE_REQUESTS`),
	"CiPagesUrl":                              os.Getenv(`CI_PAGES_URL`),
	"CiPagesDomain":                           os.Getenv(`CI_PAGES_DOMAIN`),
	"CiPipelineId":                            os.Getenv(`CI_PIPELINE_ID`),
	"CiPipelineIid":                           os.Getenv(`CI_PIPELINE_IID`),
	"CiPipelineSource":                        os.Getenv(`CI_PIPELINE_SOURCE`),
	"CiPipelineTriggered":                     os.Getenv(`CI_PIPELINE_TRIGGERED`),
	"CiPipelineUrl":                           os.Getenv(`CI_PIPELINE_URL`),
	"CiPipelineCreatedAt":                     os.Getenv(`CI_PIPELINE_CREATED_AT`),
	"CiProjectConfigPath":                     os.Getenv(`CI_PROJECT_CONFIG_PATH`),
	"CiProjectDir":                            os.Getenv(`CI_PROJECT_DIR`),
	"CiProjectId":                             os.Getenv(`CI_PROJECT_ID`),
	"CiProjectName":                           os.Getenv(`CI_PROJECT_NAME`),
	"CiProjectPathSlug":                       os.Getenv(`CI_PROJECT_PATH_SLUG`),
	"CiProjectPath":                           os.Getenv(`CI_PROJECT_PATH`),
	"CiProjectRepositoryLanguages":            os.Getenv(`CI_PROJECT_REPOSITORY_LANGUAGES`),
	"CiProjectRootNamespace":                  os.Getenv(`CI_PROJECT_ROOT_NAMESPACE`),
	"CiProjectNamespace":                      os.Getenv(`CI_PROJECT_NAMESPACE`),
	"CiProjectTitle":                          os.Getenv(`CI_PROJECT_TITLE`),
	"CiProjectUrl":                            os.Getenv(`CI_PROJECT_URL`),
	"CiProjectVisibility":                     os.Getenv(`CI_PROJECT_VISIBILITY`),
	"CiProjectClassificationLabel":            os.Getenv(`CI_PROJECT_CLASSIFICATION_LABEL`),
	"CiRegistryImage":                         os.Getenv(`CI_REGISTRY_IMAGE`),
	"CiRegistryPassword":                      os.Getenv(`CI_REGISTRY_PASSWORD`),
	"CiRegistryUser":                          os.Getenv(`CI_REGISTRY_USER`),
	"CiRegistry":                              os.Getenv(`CI_REGISTRY`),
	"CiRepositoryUrl":                         os.Getenv(`CI_REPOSITORY_URL`),
	"CiRunnerDescription":                     os.Getenv(`CI_RUNNER_DESCRIPTION`),
	"CiRunnerExecutableArch":                  os.Getenv(`CI_RUNNER_EXECUTABLE_ARCH`),
	"CiRunnerId":                              os.Getenv(`CI_RUNNER_ID`),
	"CiRunnerRevision":                        os.Getenv(`CI_RUNNER_REVISION`),
	"CiRunnerShortToken":                      os.Getenv(`CI_RUNNER_SHORT_TOKEN`),
	"CiRunnerTags":                            os.Getenv(`CI_RUNNER_TAGS`),
	"CiRunnerVersion":                         os.Getenv(`CI_RUNNER_VERSION`),
	"CiServerHost":                            os.Getenv(`CI_SERVER_HOST`),
	"CiServerName":                            os.Getenv(`CI_SERVER_NAME`),
	"CiServerPort":                            os.Getenv(`CI_SERVER_PORT`),
	"CiServerProtocol":                        os.Getenv(`CI_SERVER_PROTOCOL`),
	"CiServerRevision":                        os.Getenv(`CI_SERVER_REVISION`),
	"CiServerUrl":                             os.Getenv(`CI_SERVER_URL`),
	"CiServerVersionMajor":                    os.Getenv(`CI_SERVER_VERSION_MAJOR`),
	"CiServerVersionMinor":                    os.Getenv(`CI_SERVER_VERSION_MINOR`),
	"CiServerVersionPatch":                    os.Getenv(`CI_SERVER_VERSION_PATCH`),
	"CiServerVersion":                         os.Getenv(`CI_SERVER_VERSION`),
	"CiServer":                                os.Getenv(`CI_SERVER`),
	"CiSharedEnvironment":                     os.Getenv(`CI_SHARED_ENVIRONMENT`),
	"CiMergeRequestApproved":                  os.Getenv(`CI_MERGE_REQUEST_APPROVED`),
	"CiMergeRequestAssignees":                 os.Getenv(`CI_MERGE_REQUEST_ASSIGNEES`),
	"CiMergeRequestId":                        os.Getenv(`CI_MERGE_REQUEST_ID`),
	"CiMergeRequestIid":                       os.Getenv(`CI_MERGE_REQUEST_IID`),
	"CiMergeRequestLabels":                    os.Getenv(`CI_MERGE_REQUEST_LABELS`),
	"CiMergeRequestMilestone":                 os.Getenv(`CI_MERGE_REQUEST_MILESTONE`),
	"CiMergeRequestProjectId":                 os.Getenv(`CI_MERGE_REQUEST_PROJECT_ID`),
	"CiMergeRequestProjectPath":               os.Getenv(`CI_MERGE_REQUEST_PROJECT_PATH`),
	"CiMergeRequestProjectUrl":                os.Getenv(`CI_MERGE_REQUEST_PROJECT_URL`),
	"CiMergeRequestRefPath":                   os.Getenv(`CI_MERGE_REQUEST_REF_PATH`),
	"CiMergeRequestSourceBranchName":          os.Getenv(`CI_MERGE_REQUEST_SOURCE_BRANCH_NAME`),
	"CiMergeRequestSourceBranchSha":           os.Getenv(`CI_MERGE_REQUEST_SOURCE_BRANCH_SHA`),
	"CiMergeRequestSourceProjectId":           os.Getenv(`CI_MERGE_REQUEST_SOURCE_PROJECT_ID`),
	"CiMergeRequestSourceProjectPath":         os.Getenv(`CI_MERGE_REQUEST_SOURCE_PROJECT_PATH`),
	"CiMergeRequestSourceProjectUrl":          os.Getenv(`CI_MERGE_REQUEST_SOURCE_PROJECT_URL`),
	"CiMergeRequestTargetBranchName":          os.Getenv(`CI_MERGE_REQUEST_TARGET_BRANCH_NAME`),
	"CiMergeRequestTargetBranchSha":           os.Getenv(`CI_MERGE_REQUEST_TARGET_BRANCH_SHA`),
	"CiMergeRequestTitle":                     os.Getenv(`CI_MERGE_REQUEST_TITLE`),
	"CiMergeRequestEventType":                 os.Getenv(`CI_MERGE_REQUEST_EVENT_TYPE`),
	"CiMergeRequestDiffId":                    os.Getenv(`CI_MERGE_REQUEST_DIFF_ID`),
	"CiMergeRequestDiffBaseSha":               os.Getenv(`CI_MERGE_REQUEST_DIFF_BASE_SHA`),
	"CiExternalPullRequestIid":                os.Getenv(`CI_EXTERNAL_PULL_REQUEST_IID`),
	"CiExternalPullRequestSourceRepository":   os.Getenv(`CI_EXTERNAL_PULL_REQUEST_SOURCE_REPOSITORY`),
	"CiExternalPullRequestTargetRepository":   os.Getenv(`CI_EXTERNAL_PULL_REQUEST_TARGET_REPOSITORY`),
	"CiExternalPullRequestSourceBranchName":   os.Getenv(`CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_NAME`),
	"CiExternalPullRequestSourceBranchSha":    os.Getenv(`CI_EXTERNAL_PULL_REQUEST_SOURCE_BRANCH_SHA`),
	"CiExternalPullRequestTargetBranchName":   os.Getenv(`CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_NAME`),
	"CiExternalPullRequestTargetBranchSha":    os.Getenv(`CI_EXTERNAL_PULL_REQUEST_TARGET_BRANCH_SHA`),
}

type Gitlab struct {
	//
}

// goes through NOTIF_FIELDS environment variable and set the ones user set for template
func (g Gitlab) Setter() ([]models.Field, error) {
	var fields []models.Field
	gitlabEnvs := os.Getenv(`NOTIF_FIELDS`)
	if gitlabEnvs == "" {
		fmt.Fprintf(os.Stderr, "`NOTF_FIELDS` is empty, default template used.\n")
		fields = []models.Field{
			{
				Name:   "Branch:",
				Value:  GitlabEnvs["CiCommitBranch"],
				Inline: "false",
			},
			{
				Name:   "Commit Hash:",
				Value:  GitlabEnvs["CiCommitSha"],
				Inline: "false",
			},
			{
				Name:   "Author:",
				Value:  GitlabEnvs["CiCommitAuthor"],
				Inline: "false",
			},
			{
				Name:   "Commit Message:",
				Value:  GitlabEnvs["CiCommitMessage"],
				Inline: "false",
			},
			{
				Name:   "Repo URL:",
				Value:  GitlabEnvs["CiServerUrl"],
				Inline: "false",
			},
		}
		return fields, nil
	}
	gitlabTags := strings.Split(gitlabEnvs, ",")
	for _, v := range gitlabTags {
		for key, value := range GitlabEnvs {
			if v == key {
				if GitlabEnvs[key] != "" {
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
