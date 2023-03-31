package volume

import (
	"fmt"

	lhv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	ctllonghornv1 "github.com/harvester/harvester/pkg/generated/controllers/longhorn.io/v1beta1"
)

type AttachController struct {
	volumes ctllonghornv1.VolumeClient
}

// Attach volumes if workload is pending.
func (c *AttachController) AttachVolumesOnChange(_ string, volume *lhv1beta1.Volume) (*lhv1beta1.Volume, error) {
	if volume == nil || volume.DeletionTimestamp != nil {
		return volume, nil
	}

	if isVolumeAttached(volume) {
		return volume, nil
	}

	for _, workload := range volume.Status.KubernetesStatus.WorkloadsStatus {
		if workload.PodStatus != string(corev1.PodPending) {
			continue
		}

		logrus.Infof("Attach volume %s to node %s", volume.Name, volume.Status.OwnerID)
		volCpy := volume.DeepCopy()
		volCpy.Spec.NodeID = volCpy.Status.OwnerID
		if _, err := c.volumes.Update(volCpy); err != nil {
			return volume, fmt.Errorf("can't attach volume %s, err: %w", volume.Name, err)
		}
		break
	}

	return volume, nil
}
