package client

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	AlertRulesClient *securityinsight.AlertRulesClient
}

func NewClient(o *common.ClientOptions) *Client {
	alertRulesClient := securityinsight.NewAlertRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&alertRulesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		AlertRulesClient: &alertRulesClient,
	}
}
