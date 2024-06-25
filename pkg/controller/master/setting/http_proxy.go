package setting

import (
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/util"
	"github.com/sirupsen/logrus"
)

const (
	fleetLocalNamespace = "fleet-local"
	localClusterName    = "local"
)

func (h *Handler) syncHTTPProxy(setting *harvesterv1.Setting) error {
	// Add envs to the backup secret used by Longhorn backups
	var httpProxyConfig util.HTTPProxyConfig
	if err := json.Unmarshal([]byte(setting.Value), &httpProxyConfig); err != nil {
		return err
	}
	backupConfig := map[string]string{
		util.HTTPProxyEnv:  httpProxyConfig.HTTPProxy,
		util.HTTPSProxyEnv: httpProxyConfig.HTTPSProxy,
		util.NoProxyEnv:    util.AddBuiltInNoProxy(httpProxyConfig.NoProxy),
	}
	if err := h.updateBackupSecret(backupConfig); err != nil {
		return err
	}
	if err := h.syncHTTPProxySecret(httpProxyConfig); err != nil {
		logrus.WithFields(logrus.Fields{
			"setting": setting.Name,
			"value":   setting.Value,
		}).WithError(err).Error("failed to sync http proxy secret")
		return err
	}

	//redeploy system services. The proxy envs will be injected by the mutation webhook.
	if err := h.redeployDeployment(util.CattleSystemNamespaceName, "rancher"); err != nil {
		return err
	}
	return h.redeployDeployment(h.namespace, "harvester")
}

func (h *Handler) syncHTTPProxySecret(httpProxyConfig util.HTTPProxyConfig) error {
	secret, err := h.secretCache.Get(util.CattleSystemNamespaceName, util.HttpProxySecretName)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}
		_, err = h.secrets.Create(&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: util.CattleSystemNamespaceName,
				Name:      util.HttpProxySecretName,
			},
			Data: map[string][]byte{
				util.HTTPProxyEnv:  []byte(httpProxyConfig.HTTPProxy),
				util.HTTPSProxyEnv: []byte(httpProxyConfig.HTTPSProxy),
				util.NoProxyEnv:    []byte(util.AddBuiltInNoProxy(httpProxyConfig.NoProxy)),
			},
		})
		return err
	}

	secretCopy := secret.DeepCopy()
	secretCopy.Data[util.HTTPProxyEnv] = []byte(httpProxyConfig.HTTPProxy)
	secretCopy.Data[util.HTTPSProxyEnv] = []byte(httpProxyConfig.HTTPSProxy)
	secretCopy.Data[util.NoProxyEnv] = []byte(util.AddBuiltInNoProxy(httpProxyConfig.NoProxy))
	_, err = h.secrets.Update(secretCopy)
	return err
}
