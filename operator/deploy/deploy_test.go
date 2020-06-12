package deploy

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"os"
	"testing"

	utillog "github.com/Azure/ARO-RP/pkg/util/log"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

func TestAroOperatorImage(t *testing.T) {
	testlog := utillog.GetLogger()
	tests := []struct {
		name      string
		isDev     bool
		image     string
		imageTag  string
		gitCommit string
		want      string
		wantErr   bool
	}{
		{
			name:  "(dev) ARO_IMAGE set",
			image: "quay.io/freddyZ/arrow:latest",
			isDev: true,
			want:  "quay.io/freddyZ/arrow:latest",
		},
		{
			name:     "(dev) ARO_IMAGE_TAG set",
			imageTag: "9f7f4282",
			isDev:    true,
			want:     "arointsvc.azurecr.io/aro:9f7f4282",
		},
		{
			name:      "use commit",
			gitCommit: "be925d16",
			want:      "arointsvc.azurecr.io/aro:be925d16",
		},
		{
			name:      "non-dev no commit",
			gitCommit: version.GitCommitUnknown,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.gitCommit != "" {
				version.GitCommit = tt.gitCommit
			}
			if tt.image == "" {
				os.Unsetenv("ARO_IMAGE")
			} else {
				os.Setenv("ARO_IMAGE", tt.image)
			}
			if tt.imageTag == "" {
				os.Unsetenv("ARO_IMAGE_TAG")
			} else {
				os.Setenv("ARO_IMAGE_TAG", tt.imageTag)
			}
			got, err := aroOperatorImage(testlog, tt.isDev, "arointsvc.azurecr.io")
			if (err != nil) != tt.wantErr {
				t.Errorf("aroOperatorImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("aroOperatorImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
