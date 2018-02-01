package models

import (
	//	"errors"
	"github.com/astaxie/beego/logs"
	//	"github.com/kubernetes/kompose/pkg/app"
	"github.com/zhanglianx111/kompose/pkg/kobject"
	//	"strconv"
	//	"time"
)

type File struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

func Convert(f File) string {
	logs.Debug(f)
	return f.Name
}

func InitConvertOpt(file []string) kobject.ConvertOptions {
	kobj := kobject.ConvertOptions{}

	kobj.ToStdout = true
	kobj.CreateD = false
	kobj.CreateRC = false
	kobj.CreateDS = false
	kobj.CreateDeploymentConfig = true

	kobj.BuildRepo = ""
	kobj.BuildBranch = ""
	kobj.Build = "none"

	kobj.CreateChart = false
	kobj.GenerateYaml = false
	kobj.GenerateJSON = true
	kobj.EmptyVols = false
	kobj.Volumes = "persistentVolumeClaim"

	kobj.InsecureRepository = false
	kobj.Replicas = 1
	kobj.InputFiles = file
	kobj.OutFile = ""
	kobj.Provider = "openshift"
	kobj.Namespace = ""
	kobj.Controller = ""
	kobj.IsDeploymentFlag = false
	kobj.IsDaemonSetFlag = false
	kobj.IsReplicationControllerFlag = false
	kobj.IsReplicaSetFlag = false
	kobj.IsDeploymentConfigFlag = false
	kobj.IsNamespaceFlag = false

	return kobj
}
