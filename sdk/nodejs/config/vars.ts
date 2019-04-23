// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("linode");

/**
 * The token that allows you access to your Linode account
 */
export let token: string = utilities.requireWithDefault(() => __config.require("token"), utilities.getEnv("LINODE_TOKEN", "LINODE_API_TOKEN"));
