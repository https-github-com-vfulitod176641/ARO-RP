package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"testing"

	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/client-go/config/clientset/versioned/fake"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktesting "k8s.io/client-go/testing"

	"github.com/Azure/ARO-RP/pkg/util/version"
)

func TestUpgradeCluster(t *testing.T) {
	ctx := context.Background()

	newFakecli := func(channel, version string) *fake.Clientset {
		return fake.NewSimpleClientset(&configv1.ClusterVersion{
			ObjectMeta: metav1.ObjectMeta{
				Name: "version",
			},
			Spec: configv1.ClusterVersionSpec{
				Channel: channel,
			},
			Status: configv1.ClusterVersionStatus{
				Desired: configv1.Update{
					Version: version,
				},
			},
		})
	}

	for _, tt := range []struct {
		name        string
		fakecli     *fake.Clientset
		wantUpdated bool
	}{
		{
			name:        "needs update",
			fakecli:     newFakecli("", "0.0.0"),
			wantUpdated: true,
		},
		{
			name:    "right version, no update needed",
			fakecli: newFakecli("", version.OpenShiftVersion),
		},
		{
			name:    "later version, no update needed",
			fakecli: newFakecli("", "99.99.99"),
		},
		{
			name:    "on a channel, no update needed",
			fakecli: newFakecli("my-channel", ""),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var updated bool

			tt.fakecli.PrependReactor("update", "clusterversions", func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
				updated = true
				return false, nil, nil
			})

			i := &Installer{
				log:       logrus.NewEntry(logrus.StandardLogger()),
				configcli: tt.fakecli,
			}

			err := i.upgradeCluster(ctx)
			if err != nil {
				t.Error(err)
			}

			if updated != tt.wantUpdated {
				t.Fatal(updated)
			}

			cv, err := i.configcli.ConfigV1().ClusterVersions().Get("version", metav1.GetOptions{})
			if err != nil {
				t.Error(err)
			}

			if tt.wantUpdated {
				if cv.Spec.DesiredUpdate == nil {
					t.Fatal(cv.Spec.DesiredUpdate)
				}
				if cv.Spec.DesiredUpdate.Version != version.OpenShiftVersion {
					t.Error(cv.Spec.DesiredUpdate.Version)
				}
			}
		})
	}
}
