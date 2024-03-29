// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a Linode Domain resource.  This can be used to create, modify, and delete Linode Domains through Linode's managed DNS service.
 * For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomain).
 * 
 * The Linode Guide, [Deploy a WordPress Site Using Terraform and Linode StackScripts](https://www.linode.com/docs/applications/configuration-management/deploy-a-wordpress-site-using-terraform-and-linode-stackscripts/), demonstrates the management of Linode Domain resources in the context of Linode Instance running WordPress.
 * 
 * ## Example Usage
 * 
 * The following example shows how one might use this resource to configure a Domain Record attached to a Linode Domain.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as linode from "@pulumi/linode";
 * 
 * const foobarDomain = new linode.Domain("foobar", {
 *     domain: "foobar.example",
 *     soaEmail: "example@foobar.example",
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 * });
 * const foobarDomainRecord = new linode.DomainRecord("foobar", {
 *     domainId: foobarDomain.id,
 *     name: "www",
 *     recordType: "CNAME",
 *     target: "foobar.example",
 * });
 * ```
 * 
 * ## Attributes
 * 
 * This resource exports no additional attributes, however `status` may reflect degraded states.
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /**
     * The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
     */
    public readonly axfrIps: pulumi.Output<string[] | undefined>;
    /**
     * A description for this Domain. This is for display purposes only.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
     */
    public readonly domain: pulumi.Output<string>;
    /**
     * The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    public readonly expireSec: pulumi.Output<number | undefined>;
    /**
     * The group this Domain belongs to. This is for display purposes only.
     */
    public readonly group: pulumi.Output<string | undefined>;
    /**
     * The IP addresses representing the master DNS for this Domain.
     */
    public readonly masterIps: pulumi.Output<string[] | undefined>;
    /**
     * The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    public readonly refreshSec: pulumi.Output<number | undefined>;
    /**
     * The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    public readonly retrySec: pulumi.Output<number | undefined>;
    /**
     * Start of Authority email address. This is required for master Domains.
     */
    public readonly soaEmail: pulumi.Output<string | undefined>;
    /**
     * Used to control whether this Domain is currently being rendered (defaults to "active").
     */
    public readonly status: pulumi.Output<string>;
    /**
     * A list of tags applied to this object. Tags are for organizational purposes only.
     */
    public readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    public readonly ttlSec: pulumi.Output<number | undefined>;
    /**
     * If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
     */
    public readonly type: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainState = argsOrState as DomainState | undefined;
            inputs["axfrIps"] = state ? state.axfrIps : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["expireSec"] = state ? state.expireSec : undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["masterIps"] = state ? state.masterIps : undefined;
            inputs["refreshSec"] = state ? state.refreshSec : undefined;
            inputs["retrySec"] = state ? state.retrySec : undefined;
            inputs["soaEmail"] = state ? state.soaEmail : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["ttlSec"] = state ? state.ttlSec : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if (!args || args.domain === undefined) {
                throw new Error("Missing required property 'domain'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["axfrIps"] = args ? args.axfrIps : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["expireSec"] = args ? args.expireSec : undefined;
            inputs["group"] = args ? args.group : undefined;
            inputs["masterIps"] = args ? args.masterIps : undefined;
            inputs["refreshSec"] = args ? args.refreshSec : undefined;
            inputs["retrySec"] = args ? args.retrySec : undefined;
            inputs["soaEmail"] = args ? args.soaEmail : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ttlSec"] = args ? args.ttlSec : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        super("linode:index/domain:Domain", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
     */
    readonly axfrIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for this Domain. This is for display purposes only.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly expireSec?: pulumi.Input<number>;
    /**
     * The group this Domain belongs to. This is for display purposes only.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The IP addresses representing the master DNS for this Domain.
     */
    readonly masterIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly refreshSec?: pulumi.Input<number>;
    /**
     * The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly retrySec?: pulumi.Input<number>;
    /**
     * Start of Authority email address. This is required for master Domains.
     */
    readonly soaEmail?: pulumi.Input<string>;
    /**
     * Used to control whether this Domain is currently being rendered (defaults to "active").
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A list of tags applied to this object. Tags are for organizational purposes only.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly ttlSec?: pulumi.Input<number>;
    /**
     * If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
     */
    readonly axfrIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for this Domain. This is for display purposes only.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly expireSec?: pulumi.Input<number>;
    /**
     * The group this Domain belongs to. This is for display purposes only.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The IP addresses representing the master DNS for this Domain.
     */
    readonly masterIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly refreshSec?: pulumi.Input<number>;
    /**
     * The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly retrySec?: pulumi.Input<number>;
    /**
     * Start of Authority email address. This is required for master Domains.
     */
    readonly soaEmail?: pulumi.Input<string>;
    /**
     * Used to control whether this Domain is currently being rendered (defaults to "active").
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A list of tags applied to this object. Tags are for organizational purposes only.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
     */
    readonly ttlSec?: pulumi.Input<number>;
    /**
     * If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
     */
    readonly type: pulumi.Input<string>;
}
