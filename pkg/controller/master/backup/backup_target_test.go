package backup

import (
	"fmt"
	"testing"

	longhornv1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corefake "k8s.io/client-go/kubernetes/fake"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/fake"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/util"
	"github.com/harvester/harvester/pkg/util/fakeclients"
)

func TestTargetHandler_SettingOnChanged(t *testing.T) {
	type input struct {
		key     string
		setting *harvesterv1.Setting
	}
	type output struct {
		backupTarget                           settings.BackupTarget
		s3BackupTargetSecret                   map[string]string
		longhornBackupTargetSettingValue       string
		longhornBackupTargetSecretSettingValue string
		err                                    error
	}
	var testCases = []struct {
		name     string
		given    input
		expected output
	}{
		{
			name: "correct NFS backup target setting, should pass",
			given: input{
				key: settings.BackupTargetSettingName,
				setting: &harvesterv1.Setting{
					ObjectMeta: metav1.ObjectMeta{Name: settings.BackupTargetSettingName},
					Value:      `{"type":"nfs","endpoint":"nfs://longhorn-test-nfs-svc.default:/opt/backupstore","accessKeyId":"","secretAccessKey":"","bucketName":"","bucketRegion":"","cert":"","virtualHostedStyle":false}`,
				},
			},
			expected: output{
				backupTarget: settings.BackupTarget{
					Type:     settings.NFSBackupType,
					Endpoint: "nfs://longhorn-test-nfs-svc.default:/opt/backupstore",
				},
				longhornBackupTargetSettingValue: "nfs://longhorn-test-nfs-svc.default:/opt/backupstore",
			},
		},
		{
			name: "correct S3 backup target setting, should pass",
			given: input{
				key: settings.BackupTargetSettingName,
				setting: &harvesterv1.Setting{
					ObjectMeta: metav1.ObjectMeta{Name: settings.BackupTargetSettingName},
					Value:      `{"type":"s3","endpoint":"","accessKeyId":"FAKE_ACCESS_KEY","secretAccessKey":"FAKE_SECRET_ACCESS_KEY","bucketName":"bucket","bucketRegion":"ap-northeast-1","cert":"","virtualHostedStyle":false}`,
				},
			},
			expected: output{
				backupTarget: settings.BackupTarget{
					Type:         settings.S3BackupType,
					Endpoint:     "s3://bucket@ap-northeast-1/",
					BucketName:   "bucket",
					BucketRegion: "ap-northeast-1",
				},
				s3BackupTargetSecret: map[string]string{
					AWSAccessKey:       "FAKE_ACCESS_KEY",
					AWSSecretKey:       "FAKE_SECRET_ACCESS_KEY",
					AWSCERT:            "",
					VirtualHostedStyle: "false",
				},
				longhornBackupTargetSettingValue:       "s3://bucket@ap-northeast-1/",
				longhornBackupTargetSecretSettingValue: util.BackupTargetSecretName,
			},
		},
		{
			name: "error setting value, should fail",
			given: input{
				key: settings.BackupTargetSettingName,
				setting: &harvesterv1.Setting{
					ObjectMeta: metav1.ObjectMeta{Name: settings.BackupTargetSettingName},
					Value:      "not json",
				},
			},
			expected: output{
				err: fmt.Errorf("decode error"),
			},
		},
	}

	for _, tc := range testCases {
		// track ressources
		var clientset = fake.NewSimpleClientset()
		trackResources := []runtime.Object{&longhornv1.Setting{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: util.LonghornSystemNamespaceName,
				Name:      longhornBackupTargetSettingName,
			},
		}}
		if tc.given.setting != nil {
			trackResources = append(trackResources, tc.given.setting)
		}
		for _, resource := range trackResources {
			var err = clientset.Tracker().Add(resource)
			assert.Nil(t, err, "mock resource should add into fake controller tracker")
		}

		// set default longhorn backup target setting
		var coreclientset = corefake.NewSimpleClientset()

		var handler = &TargetHandler{
			settings:             fakeclients.HarvesterSettingClient(clientset.HarvesterhciV1beta1().Settings),
			longhornSettings:     fakeclients.LonghornSettingClient(clientset.LonghornV1beta1().Settings),
			longhornSettingCache: fakeclients.LonghornSettingCache(clientset.LonghornV1beta1().Settings),
			secrets:              fakeclients.SecretClient(coreclientset.CoreV1().Secrets),
			secretCache:          fakeclients.SecretCache(coreclientset.CoreV1().Secrets),
		}

		var syncActual output
		_, syncActual.err = handler.OnBackupTargetChange(tc.given.key, tc.given.setting)
		if tc.expected.err != nil {
			assert.Error(t, syncActual.err, tc.name)
			continue
		} else {
			assert.NoError(t, syncActual.err, tc.name)
		}

		backupTargetSetting, err := handler.settings.Get(settings.BackupTargetSettingName, metav1.GetOptions{})
		assert.NoError(t, err, tc.name)
		backupTarget, err := settings.DecodeBackupTarget(backupTargetSetting.Value)
		assert.NoError(t, err, tc.name)
		assert.Equal(t, tc.expected.backupTarget, *backupTarget, tc.name)

		longhornBackupTargetSetting, err := handler.longhornSettings.Get(util.LonghornSystemNamespaceName, longhornBackupTargetSettingName, metav1.GetOptions{})
		assert.NoError(t, err, tc.name)
		assert.Equal(t, tc.expected.longhornBackupTargetSettingValue, longhornBackupTargetSetting.Value, tc.name)

		if backupTarget.Type == settings.S3BackupType {
			longhornBackupTargetSecretSetting, err := handler.longhornSettingCache.Get(util.LonghornSystemNamespaceName, longhornBackupTargetSecretSettingName)
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expected.longhornBackupTargetSecretSettingValue, longhornBackupTargetSecretSetting.Value, tc.name)

			secret, err := handler.secretCache.Get(util.LonghornSystemNamespaceName, util.BackupTargetSecretName)
			assert.NoError(t, err, tc.name)
			assert.Equal(t, tc.expected.s3BackupTargetSecret, secret.StringData, tc.name)
		}
	}
}
