package setting

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	corev1 "k8s.io/api/core/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/settings"
)

func (h *Handler) syncGeneralJobImage(setting *harvesterv1.Setting) error {
	if setting.Name != "tmp-general-job" {
		return nil
	}

	getValues := action.NewGetValues(&h.helmConfiguration)
	getValues.AllValues = true
	helmValues, err := getValues.Run("harvester")
	if err != nil {
		logrus.WithError(err).Error("failed to get harvester values")
		return err
	}

	if helmValues["generalJob"] == nil {
		logrus.Debug("general job is not set")
		return nil
	}

	generalJobValues := helmValues["generalJob"].(map[string]interface{})
	if generalJobValues["image"] == nil {
		logrus.Debug("general job image is not set")
		return nil
	}

	imageValues := generalJobValues["image"].(map[string]interface{})
	repository := imageValues["repository"].(string)
	tag := imageValues["tag"].(string)
	imagePullPolicy := imageValues["imagePullPolicy"].(string)
	if repository == "" || tag == "" || imagePullPolicy == "" {
		logrus.Debug("some fields in general job image is not set")
	}

	image := &settings.Image{
		Repository:      repository,
		Tag:             tag,
		ImagePullPolicy: corev1.PullPolicy(imagePullPolicy),
	}
	imageStr, err := json.Marshal(image)
	if err != nil {
		logrus.WithField("image", image).WithError(err).Error("failed to marshal general job image")
		return err
	}

	if setting.Default == string(imageStr) {
		return nil
	}

	settingCopy := setting.DeepCopy()
	settingCopy.Default = string(imageStr)
	_, err = h.settings.Update(settingCopy)
	return err
}
