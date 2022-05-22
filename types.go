package flink

import "time"

type State string

const (
	RUNNING   State = "RUNNING"
	SUSPENDED State = "SUSPENDED"
	CANCELLED State = "CANCELLED"
)

type Status string

const (
	StatusTRANSITIONING Status = "TRANSITIONING"
	StatusCANCELLED     Status = "CANCELLED"
	StatusRUNNING       Status = "RUNNING"
)

type Deployment struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Annotations struct {
			ComDataartisansAppmanagerControllerDeploymentSpecVersion        string `json:"com.dataartisans.appmanager.controller.deployment.spec.version"`
			ComDataartisansAppmanagerControllerDeploymentTransitioning      string `json:"com.dataartisans.appmanager.controller.deployment.transitioning"`
			ComDataartisansAppmanagerControllerDeploymentTransitioningSince string `json:"com.dataartisans.appmanager.controller.deployment.transitioning.since"`
			Comment                                                         string `json:"comment"`
			Creator                                                         string `json:"creator"`
			CreatorName                                                     string `json:"creatorName"`
			DeploymentHasDeployed                                           string `json:"deployment.has.deployed"`
			DeploymentReferencedDeploymentDraftHistoryId                    string `json:"deployment.referenced.deployment.draft.history.id"`
			DeploymentReferencedDeploymentDraftHistoryVersion               string `json:"deployment.referenced.deployment.draft.history.version"`
			DeploymentReferencedDeploymentDraftId                           string `json:"deployment.referenced.deployment.draft.id"`
			DeploymentSourceType                                            string `json:"deployment.source-type"`
			Modifier                                                        string `json:"modifier"`
			ModifierName                                                    string `json:"modifierName"`
			OneTimeRetry                                                    string `json:"one-time.retry"`
			Taker                                                           string `json:"taker"`
		} `json:"annotations"`
		CreatedAt time.Time `json:"createdAt"`
		Id        string    `json:"id"`
		Labels    struct {
		} `json:"labels"`
		ModifiedAt      time.Time `json:"modifiedAt"`
		Name            string    `json:"name"`
		Namespace       string    `json:"namespace"`
		ResourceVersion int       `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		DeploymentTargetId           string `json:"deploymentTargetId"`
		MaxJobCreationAttempts       int    `json:"maxJobCreationAttempts"`
		MaxSavepointCreationAttempts int    `json:"maxSavepointCreationAttempts"`
		RestoreStrategy              struct {
			Kind string `json:"kind"`
		} `json:"restoreStrategy"`
		State    State `json:"state"`
		Template struct {
			Metadata struct {
				Annotations struct {
					FlinkQueryableStateEnabled string `json:"flink.queryable-state.enabled"`
				} `json:"annotations"`
			} `json:"metadata"`
			Spec struct {
				Artifact struct {
					FlinkImageRegistry   string `json:"flinkImageRegistry"`
					FlinkImageRepository string `json:"flinkImageRepository"`
					FlinkImageTag        string `json:"flinkImageTag"`
					FlinkVersion         string `json:"flinkVersion"`
					ImageUserDefined     bool   `json:"imageUserDefined"`
					JarUri               string `json:"jarUri"`
					Kind                 string `json:"kind"`
					VersionName          string `json:"versionName"`
				} `json:"artifact"`
				BatchMode          bool `json:"batchMode"`
				FlinkConfiguration struct {
					ExecutionCheckpointingInterval string `json:"execution.checkpointing.interval"`
					ExecutionCheckpointingMinPause string `json:"execution.checkpointing.min-pause"`
				} `json:"flinkConfiguration"`
				Logging struct {
					Log4JLoggers struct {
						Field1 string `json:""`
					} `json:"log4jLoggers"`
					LogReservePolicy struct {
						ExpirationDays int  `json:"expirationDays"`
						OpenHistory    bool `json:"openHistory"`
					} `json:"logReservePolicy"`
					LoggingProfile string `json:"loggingProfile"`
				} `json:"logging"`
				Parallelism int `json:"parallelism"`
				Resources   struct {
					Jobmanager struct {
						Cpu    int    `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"jobmanager"`
					Taskmanager struct {
						Cpu    int    `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"taskmanager"`
				} `json:"resources"`
			} `json:"spec"`
		} `json:"template"`
		UpgradeStrategy struct {
			Kind string `json:"kind"`
		} `json:"upgradeStrategy"`
	} `json:"spec"`
	Status struct {
		EndTime   int    `json:"endTime"`
		StartTime int64  `json:"startTime"`
		State     Status `json:"state"`
	} `json:"status"`
}
