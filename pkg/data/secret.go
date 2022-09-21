package data

import (
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/harvester/harvester/pkg/config"
	"github.com/harvester/harvester/pkg/util"
)

func createSecrets(mgmt *config.Management) error {
	secrets := mgmt.CoreFactory.Core().V1().Secret()

	// Initializing the secret for Plan cattle-system/sync-additional-ca and cattle-system/sync-rke2-registries,
	// so plans don't fail to mount secrets to jobs.
	defaultSecrets := []corev1.Secret{
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: util.CattleSystemNamespaceName,
				Name:      util.AdditionalCASecretName,
			},
			Data: map[string][]byte{
				util.AdditionalCAFileName: []byte(""),
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: util.CattleSystemNamespaceName,
				Name:      util.ContainerdRegistrySecretName,
			},
			Data: map[string][]byte{
				util.ContainerdRegistryFileName: []byte(""),
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: util.CattleSystemNamespaceName,
				Name:      util.ContainerdRegistryUpdateScriptSecretName,
			},
			StringData: map[string]string{
				util.ContainerdRegistryUpdateScriptName: `#!/bin/sh
				RESTART=true
				if [ -f "/etc/rancher/rke2/registries.yaml" ]; then
					if [ cmp -s "/run/system-upgrade/secrets/registries/registries.yaml" "/etc/rancher/rke2/registries.yaml" ]; then
					  RESTART=false
					  echo "/etc/rancher/rke2/registries.yaml is not changed"
					fi
				else
					echo "File /etc/rancher/rke2/registries.yaml doesn't exist"
					if [ ! -s "/run/system-upgrade/secrets/registries/registries.yaml" ]; then
					  RESTART=false
					  echo "New registries.yaml is empty"
					fi
				fi
				if [ "$RESTART" == true ]; then
					echo "Update /etc/rancher/rke2 ..."
					cp /run/system-upgrade/secrets/registries/registries.yaml /etc/rancher/rke2
					echo "Restart RKE2 ..."
					kill $(pgrep rke2)
					echo "Done"
				else
					echo "Registry content doesn't change"
				fi`,
			},
		},
	}
	for _, defaultSecret := range defaultSecrets {
		if _, err := secrets.Create(&defaultSecret); err != nil && !apierrors.IsAlreadyExists(err) {
			return errors.Wrapf(err, "Failed to create secret %s/%s", defaultSecret.Namespace, defaultSecret.Name)
		}
	}

	return nil
}
