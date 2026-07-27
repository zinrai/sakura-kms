package main

import (
	"fmt"
	"os"

	kms "github.com/sacloud/kms-api-go"
	v1 "github.com/sacloud/kms-api-go/apis/v1"
	"github.com/sacloud/saclient-go"
)

// NewKMSClient creates a KMS API client for the specified zone.
// Credentials are resolved by saclient-go from environment variables,
// supporting both static API keys and service principals.
func NewKMSClient(zone string) (*v1.Client, error) {
	var sc saclient.Client
	if err := sc.SetEnviron(os.Environ()); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}
	if err := sc.Populate(); err != nil {
		return nil, fmt.Errorf("failed to configure saclient: %w", err)
	}

	apiRootURL := fmt.Sprintf("https://secure.sakura.ad.jp/cloud/zone/%s/api/cloud/1.1", zone)

	client, err := kms.NewClientWithAPIRootURL(&sc, apiRootURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create KMS client: %w", err)
	}

	return client, nil
}
