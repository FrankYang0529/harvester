package volume

import (
	"context"

	"github.com/harvester/harvester/pkg/config"
)

const (
	volumeControllerDetachVolume = "detach-volume-controller"
	volumeControllerAttachVolume = "attach-volume-controller"
)

func Register(ctx context.Context, management *config.Management, options config.Options) error {
	var (
		podClient     = management.CoreFactory.Core().V1().Pod()
		podCache      = podClient.Cache()
		pvcCache      = management.CoreFactory.Core().V1().PersistentVolumeClaim().Cache()
		volumeClient  = management.LonghornFactory.Longhorn().V1beta1().Volume()
		volumeCache   = volumeClient.Cache()
		snapshotCache = management.SnapshotFactory.Snapshot().V1beta1().VolumeSnapshot().Cache()
	)

	// registers controllers
	var detachCtrl = &DetachController{
		podCache:         podCache,
		podController:    podClient,
		pvcCache:         pvcCache,
		volumes:          volumeClient,
		volumeController: volumeClient,
		volumeCache:      volumeCache,
		snapshotCache:    snapshotCache,
	}
	var attachCtrl = &AttachController{
		volumes: volumeClient,
	}
	volumeClient.OnChange(ctx, volumeControllerDetachVolume, detachCtrl.DetachVolumesOnChange)
	volumeClient.OnChange(ctx, volumeControllerAttachVolume, attachCtrl.AttachVolumesOnChange)

	return nil
}
