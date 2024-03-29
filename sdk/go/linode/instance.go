// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package linode

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Linode Instance resource.  This can be used to create, modify, and delete Linodes.
// For more information, see [Getting Started with Linode](https://linode.com/docs/getting-started/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createLinodeInstance).
// 
// The Linode Guide, [Use Terraform to Provision Linode Environments](https://www.linode.com/docs/applications/configuration-management/how-to-build-your-infrastructure-using-terraform-and-linode/), provides step-by-step guidance and additional examples.
// 
// Linode Instances can also use [provisioners](https://www.terraform.io/docs/provisioners/index.html).
// 
// ## Attributes
// 
// This Linode Instance resource exports the following attributes:
// 
// * `status` - The status of the instance, indicating the current readiness state. (`running`, `offline`, ...)
// 
// * `ip_address` - A string containing the Linode's public IP address.
// 
// * `private_ip_address` - This Linode's Private IPv4 Address, if enabled.  The regional private IP address range is 192.168.128/17 address shared by all Linode Instances in a region.
// 
// * `ipv6` - This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
// 
// * `ipv4` - This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
// 
// * `specs.0.disk` -  The amount of storage space, in GB. this Linode has access to. A typical Linode will divide this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an image through POST /linode/instances.
// 
// * `specs.0.memory` - The amount of RAM, in MB, this Linode has access to. Typically a Linode will choose to boot with all of its available RAM, but this can be configured in a Config profile.
// 
// * `specs.0.vcpus` - The number of vcpus this Linode has access to. Typically a Linode will choose to boot with all of its available vcpus, but this can be configured in a Config Profile.
// 
// * `specs.0.transfer` - The amount of network transfer this Linode is allotted each month.
// 
// * `backups` - Information about this Linode's backups status.
// 
//   * `enabled` - If this Linode has the Backup service enabled.
// 
//   * `schedule`
// 
//     * `day` -  The day of the week that your Linode's weekly Backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as "Scheduling" until the day is automatically selected.
// 
//     * `window` - The window ('W0'-'W22') in which your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur. For example, 'W10' indicates that your backups should be taken between 10:00 and 12:00. If you do not choose a backup window, one will be selected for you automatically.  If not set manually, when backups are initially enabled this may come back as Scheduling until the window is automatically selected.
type Instance struct {
	s *pulumi.ResourceState
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOpt) (*Instance, error) {
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["alerts"] = nil
		inputs["authorizedKeys"] = nil
		inputs["authorizedUsers"] = nil
		inputs["backupId"] = nil
		inputs["backupsEnabled"] = nil
		inputs["bootConfigLabel"] = nil
		inputs["configs"] = nil
		inputs["disks"] = nil
		inputs["group"] = nil
		inputs["image"] = nil
		inputs["label"] = nil
		inputs["privateIp"] = nil
		inputs["region"] = nil
		inputs["rootPass"] = nil
		inputs["stackscriptData"] = nil
		inputs["stackscriptId"] = nil
		inputs["swapSize"] = nil
		inputs["tags"] = nil
		inputs["type"] = nil
		inputs["watchdogEnabled"] = nil
	} else {
		inputs["alerts"] = args.Alerts
		inputs["authorizedKeys"] = args.AuthorizedKeys
		inputs["authorizedUsers"] = args.AuthorizedUsers
		inputs["backupId"] = args.BackupId
		inputs["backupsEnabled"] = args.BackupsEnabled
		inputs["bootConfigLabel"] = args.BootConfigLabel
		inputs["configs"] = args.Configs
		inputs["disks"] = args.Disks
		inputs["group"] = args.Group
		inputs["image"] = args.Image
		inputs["label"] = args.Label
		inputs["privateIp"] = args.PrivateIp
		inputs["region"] = args.Region
		inputs["rootPass"] = args.RootPass
		inputs["stackscriptData"] = args.StackscriptData
		inputs["stackscriptId"] = args.StackscriptId
		inputs["swapSize"] = args.SwapSize
		inputs["tags"] = args.Tags
		inputs["type"] = args.Type
		inputs["watchdogEnabled"] = args.WatchdogEnabled
	}
	inputs["backups"] = nil
	inputs["ipAddress"] = nil
	inputs["ipv4s"] = nil
	inputs["ipv6"] = nil
	inputs["privateIpAddress"] = nil
	inputs["specs"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("linode:index/instance:Instance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceState, opts ...pulumi.ResourceOpt) (*Instance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["alerts"] = state.Alerts
		inputs["authorizedKeys"] = state.AuthorizedKeys
		inputs["authorizedUsers"] = state.AuthorizedUsers
		inputs["backupId"] = state.BackupId
		inputs["backups"] = state.Backups
		inputs["backupsEnabled"] = state.BackupsEnabled
		inputs["bootConfigLabel"] = state.BootConfigLabel
		inputs["configs"] = state.Configs
		inputs["disks"] = state.Disks
		inputs["group"] = state.Group
		inputs["image"] = state.Image
		inputs["ipAddress"] = state.IpAddress
		inputs["ipv4s"] = state.Ipv4s
		inputs["ipv6"] = state.Ipv6
		inputs["label"] = state.Label
		inputs["privateIp"] = state.PrivateIp
		inputs["privateIpAddress"] = state.PrivateIpAddress
		inputs["region"] = state.Region
		inputs["rootPass"] = state.RootPass
		inputs["specs"] = state.Specs
		inputs["stackscriptData"] = state.StackscriptData
		inputs["stackscriptId"] = state.StackscriptId
		inputs["status"] = state.Status
		inputs["swapSize"] = state.SwapSize
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
		inputs["watchdogEnabled"] = state.WatchdogEnabled
	}
	s, err := ctx.ReadResource("linode:index/instance:Instance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Instance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Instance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Instance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Instance) Alerts() *pulumi.Output {
	return r.s.State["alerts"]
}

// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if `image` is provided. *This value can not be imported.* *Changing `authorized_keys` forces the creation of a new Linode Instance.*
func (r *Instance) AuthorizedKeys() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["authorizedKeys"])
}

// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. *This value can not be imported.* *Changing `authorized_users` forces the creation of a new Linode Instance.*
func (r *Instance) AuthorizedUsers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["authorizedUsers"])
}

// A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive. *This value can not be imported.* *Changing `backup_id` forces the creation of a new Linode Instance.*
func (r *Instance) BackupId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["backupId"])
}

// Information about this Linode's backups status.
func (r *Instance) Backups() *pulumi.Output {
	return r.s.State["backups"]
}

// If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
func (r *Instance) BackupsEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["backupsEnabled"])
}

// The Label of the Instance Config that should be used to boot the Linode instance.  If there is only one `config`, the `label` of that `config` will be used as the `boot_config_label`. *This value can not be imported.*
func (r *Instance) BootConfigLabel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bootConfigLabel"])
}

// Configuration profiles define the VM settings and boot behavior of the Linode Instance.
func (r *Instance) Configs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["configs"])
}

func (r *Instance) Disks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["disks"])
}

// The display group of the Linode instance.
func (r *Instance) Group() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["group"])
}

// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use. Examples are `linode/debian9`, `linode/fedora28`, `linode/ubuntu16.04lts`, `linode/arch`, and `private/12345`. *Changing `image` forces the creation of a new Linode Instance.*
func (r *Instance) Image() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["image"])
}

// This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address
// will be used for this field.
func (r *Instance) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single
// private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
func (r *Instance) Ipv4s() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipv4s"])
}

// This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
func (r *Instance) Ipv6() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6"])
}

// The Config's label for display purposes.  Also used by `boot_config_label`.
func (r *Instance) Label() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["label"])
}

// If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region. It can be enabled on an existing Linode but it can't be disabled.
func (r *Instance) PrivateIp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["privateIp"])
}

// This Linode's Private IPv4 Address. The regional private IP address range is 192.168.128/17 address shared by all Linode
// Instances in a region.
func (r *Instance) PrivateIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateIpAddress"])
}

// This is the location where the Linode is deployed. Examples are `"us-east"`, `"us-west"`, `"ap-south"`, etc.  *Changing `region` forces the creation of a new Linode Instance.*.
func (r *Instance) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The initial password for the `root` user account. *This value can not be imported.* *Changing `root_pass` forces the creation of a new Linode Instance.* *If omitted, a random password will be generated but will not be stored in Terraform state.*
func (r *Instance) RootPass() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["rootPass"])
}

func (r *Instance) Specs() *pulumi.Output {
	return r.s.State["specs"]
}

// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.  *This value can not be imported.* *Changing `stackscript_data` forces the creation of a new Linode Instance.*
func (r *Instance) StackscriptData() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["stackscriptData"])
}

// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript. *This value can not be imported.* *Changing `stackscript_id` forces the creation of a new Linode Instance.*
func (r *Instance) StackscriptId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["stackscriptId"])
}

// The status of the instance, indicating the current readiness state.
func (r *Instance) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.
func (r *Instance) SwapSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["swapSize"])
}

// A list of tags applied to this object. Tags are for organizational purposes only.
func (r *Instance) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// The Linode type defines the pricing, CPU, disk, and RAM specs of the instance.  Examples are `"g6-nanode-1"`, `"g6-standard-2"`, `"g6-highmem-16"`, `"g6-dedicated-16"`, etc.
func (r *Instance) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
func (r *Instance) WatchdogEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["watchdogEnabled"])
}

// Input properties used for looking up and filtering Instance resources.
type InstanceState struct {
	Alerts interface{}
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if `image` is provided. *This value can not be imported.* *Changing `authorized_keys` forces the creation of a new Linode Instance.*
	AuthorizedKeys interface{}
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. *This value can not be imported.* *Changing `authorized_users` forces the creation of a new Linode Instance.*
	AuthorizedUsers interface{}
	// A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive. *This value can not be imported.* *Changing `backup_id` forces the creation of a new Linode Instance.*
	BackupId interface{}
	// Information about this Linode's backups status.
	Backups interface{}
	// If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
	BackupsEnabled interface{}
	// The Label of the Instance Config that should be used to boot the Linode instance.  If there is only one `config`, the `label` of that `config` will be used as the `boot_config_label`. *This value can not be imported.*
	BootConfigLabel interface{}
	// Configuration profiles define the VM settings and boot behavior of the Linode Instance.
	Configs interface{}
	Disks interface{}
	// The display group of the Linode instance.
	Group interface{}
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use. Examples are `linode/debian9`, `linode/fedora28`, `linode/ubuntu16.04lts`, `linode/arch`, and `private/12345`. *Changing `image` forces the creation of a new Linode Instance.*
	Image interface{}
	// This Linode's Public IPv4 Address. If there are multiple public IPv4 addresses on this Instance, an arbitrary address
	// will be used for this field.
	IpAddress interface{}
	// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single
	// private IPv4 address if needed. You may need to open a support ticket to get additional IPv4 addresses.
	Ipv4s interface{}
	// This Linode's IPv6 SLAAC addresses. This address is specific to a Linode, and may not be shared.
	Ipv6 interface{}
	// The Config's label for display purposes.  Also used by `boot_config_label`.
	Label interface{}
	// If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region. It can be enabled on an existing Linode but it can't be disabled.
	PrivateIp interface{}
	// This Linode's Private IPv4 Address. The regional private IP address range is 192.168.128/17 address shared by all
	// Linode Instances in a region.
	PrivateIpAddress interface{}
	// This is the location where the Linode is deployed. Examples are `"us-east"`, `"us-west"`, `"ap-south"`, etc.  *Changing `region` forces the creation of a new Linode Instance.*.
	Region interface{}
	// The initial password for the `root` user account. *This value can not be imported.* *Changing `root_pass` forces the creation of a new Linode Instance.* *If omitted, a random password will be generated but will not be stored in Terraform state.*
	RootPass interface{}
	Specs interface{}
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.  *This value can not be imported.* *Changing `stackscript_data` forces the creation of a new Linode Instance.*
	StackscriptData interface{}
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript. *This value can not be imported.* *Changing `stackscript_id` forces the creation of a new Linode Instance.*
	StackscriptId interface{}
	// The status of the instance, indicating the current readiness state.
	Status interface{}
	// When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.
	SwapSize interface{}
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags interface{}
	// The Linode type defines the pricing, CPU, disk, and RAM specs of the instance.  Examples are `"g6-nanode-1"`, `"g6-standard-2"`, `"g6-highmem-16"`, `"g6-dedicated-16"`, etc.
	Type interface{}
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
	WatchdogEnabled interface{}
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	Alerts interface{}
	// A list of SSH public keys to deploy for the root user on the newly created Linode. Only accepted if `image` is provided. *This value can not be imported.* *Changing `authorized_keys` forces the creation of a new Linode Instance.*
	AuthorizedKeys interface{}
	// A list of Linode usernames. If the usernames have associated SSH keys, the keys will be appended to the `root` user's `~/.ssh/authorized_keys` file automatically. *This value can not be imported.* *Changing `authorized_users` forces the creation of a new Linode Instance.*
	AuthorizedUsers interface{}
	// A Backup ID from another Linode's available backups. Your User must have read_write access to that Linode, the Backup must have a status of successful, and the Linode must be deployed to the same region as the Backup. See /linode/instances/{linodeId}/backups for a Linode's available backups. This field and the image field are mutually exclusive. *This value can not be imported.* *Changing `backup_id` forces the creation of a new Linode Instance.*
	BackupId interface{}
	// If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.
	BackupsEnabled interface{}
	// The Label of the Instance Config that should be used to boot the Linode instance.  If there is only one `config`, the `label` of that `config` will be used as the `boot_config_label`. *This value can not be imported.*
	BootConfigLabel interface{}
	// Configuration profiles define the VM settings and boot behavior of the Linode Instance.
	Configs interface{}
	Disks interface{}
	// The display group of the Linode instance.
	Group interface{}
	// An Image ID to deploy the Disk from. Official Linode Images start with linode/, while your Images start with private/. See /images for more information on the Images available for you to use. Examples are `linode/debian9`, `linode/fedora28`, `linode/ubuntu16.04lts`, `linode/arch`, and `private/12345`. *Changing `image` forces the creation of a new Linode Instance.*
	Image interface{}
	// The Config's label for display purposes.  Also used by `boot_config_label`.
	Label interface{}
	// If true, the created Linode will have private networking enabled, allowing use of the 192.168.128.0/17 network within the Linode's region. It can be enabled on an existing Linode but it can't be disabled.
	PrivateIp interface{}
	// This is the location where the Linode is deployed. Examples are `"us-east"`, `"us-west"`, `"ap-south"`, etc.  *Changing `region` forces the creation of a new Linode Instance.*.
	Region interface{}
	// The initial password for the `root` user account. *This value can not be imported.* *Changing `root_pass` forces the creation of a new Linode Instance.* *If omitted, a random password will be generated but will not be stored in Terraform state.*
	RootPass interface{}
	// An object containing responses to any User Defined Fields present in the StackScript being deployed to this Linode. Only accepted if 'stackscript_id' is given. The required values depend on the StackScript being deployed.  *This value can not be imported.* *Changing `stackscript_data` forces the creation of a new Linode Instance.*
	StackscriptData interface{}
	// The StackScript to deploy to the newly created Linode. If provided, 'image' must also be provided, and must be an Image that is compatible with this StackScript. *This value can not be imported.* *Changing `stackscript_id` forces the creation of a new Linode Instance.*
	StackscriptId interface{}
	// When deploying from an Image, this field is optional with a Linode API default of 512mb, otherwise it is ignored. This is used to set the swap disk size for the newly-created Linode.
	SwapSize interface{}
	// A list of tags applied to this object. Tags are for organizational purposes only.
	Tags interface{}
	// The Linode type defines the pricing, CPU, disk, and RAM specs of the instance.  Examples are `"g6-nanode-1"`, `"g6-standard-2"`, `"g6-highmem-16"`, `"g6-dedicated-16"`, etc.
	Type interface{}
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie will give up if there have been more than 5 boot jobs issued within 15 minutes.
	WatchdogEnabled interface{}
}
