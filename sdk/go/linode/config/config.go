// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/pulumi/pulumi/sdk/go/pulumi/config"
)

// The token that allows you access to your Linode account
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "linode:token")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "LINODE_TOKEN", "LINODE_API_TOKEN").(string); ok {
		return dv
	}
	panic(err.Error())
	return v
}
