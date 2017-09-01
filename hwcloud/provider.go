package hwcloud

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for HWCloud.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", nil),
				Description: descriptions["auth_url"],
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_TOKEN", ""),
				Description: descriptions["token"],
			},

			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
					"OS_DOMAIN_ID",
				}, ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_DEFAULT_DOMAIN",
				}, ""),
				Description: descriptions["domain_name"],
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", ""),
				Description: descriptions["insecure"],
			},

			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"swauth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SWAUTH", ""),
				Description: descriptions["swauth"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"hwcloud_images_image_v2":        dataSourceImagesImageV2(),
			"hwcloud_networking_network_v2":  dataSourceNetworkingNetworkV2(),
			"hwcloud_networking_secgroup_v2": dataSourceNetworkingSecGroupV2(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"hwcloud_blockstorage_volume_v1":          resourceBlockStorageVolumeV1(),
			"hwcloud_blockstorage_volume_v2":          resourceBlockStorageVolumeV2(),
			"hwcloud_blockstorage_volume_attach_v2":   resourceBlockStorageVolumeAttachV2(),
			"hwcloud_compute_flavor_v2":               resourceComputeFlavorV2(),
			"hwcloud_compute_instance_v2":             resourceComputeInstanceV2(),
			"hwcloud_compute_keypair_v2":              resourceComputeKeypairV2(),
			"hwcloud_compute_secgroup_v2":             resourceComputeSecGroupV2(),
			"hwcloud_compute_servergroup_v2":          resourceComputeServerGroupV2(),
			"hwcloud_compute_floatingip_v2":           resourceComputeFloatingIPV2(),
			"hwcloud_compute_floatingip_associate_v2": resourceComputeFloatingIPAssociateV2(),
			"hwcloud_compute_volume_attach_v2":        resourceComputeVolumeAttachV2(),
			"hwcloud_dns_recordset_v2":                resourceDNSRecordSetV2(),
			"hwcloud_dns_zone_v2":                     resourceDNSZoneV2(),
			"hwcloud_fw_firewall_v1":                  resourceFWFirewallV1(),
			"hwcloud_fw_policy_v1":                    resourceFWPolicyV1(),
			"hwcloud_fw_rule_v1":                      resourceFWRuleV1(),
			"hwcloud_identity_project_v3":             resourceIdentityProjectV3(),
			"hwcloud_identity_user_v3":                resourceIdentityUserV3(),
			"hwcloud_images_image_v2":                 resourceImagesImageV2(),
			"hwcloud_lb_member_v1":                    resourceLBMemberV1(),
			"hwcloud_lb_monitor_v1":                   resourceLBMonitorV1(),
			"hwcloud_lb_pool_v1":                      resourceLBPoolV1(),
			"hwcloud_lb_vip_v1":                       resourceLBVipV1(),
			"hwcloud_lb_loadbalancer_v2":              resourceLoadBalancerV2(),
			"hwcloud_lb_listener_v2":                  resourceListenerV2(),
			"hwcloud_lb_pool_v2":                      resourcePoolV2(),
			"hwcloud_lb_member_v2":                    resourceMemberV2(),
			"hwcloud_lb_monitor_v2":                   resourceMonitorV2(),
			"hwcloud_networking_network_v2":           resourceNetworkingNetworkV2(),
			"hwcloud_networking_subnet_v2":            resourceNetworkingSubnetV2(),
			"hwcloud_networking_floatingip_v2":        resourceNetworkingFloatingIPV2(),
			"hwcloud_networking_port_v2":              resourceNetworkingPortV2(),
			"hwcloud_networking_router_v2":            resourceNetworkingRouterV2(),
			"hwcloud_networking_router_interface_v2":  resourceNetworkingRouterInterfaceV2(),
			"hwcloud_networking_router_route_v2":      resourceNetworkingRouterRouteV2(),
			"hwcloud_networking_secgroup_v2":          resourceNetworkingSecGroupV2(),
			"hwcloud_networking_secgroup_rule_v2":     resourceNetworkingSecGroupRuleV2(),
			"hwcloud_objectstorage_container_v1":      resourceObjectStorageContainerV1(),
		},

		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The HWCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"endpoint_type": "The catalog endpoint type to use.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"swauth": "Use Swift's authentication system instead of Keystone. Only used for\n" +
			"interaction with Swift.",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		CACertFile:       d.Get("cacert_file").(string),
		ClientCertFile:   d.Get("cert").(string),
		ClientKeyFile:    d.Get("key").(string),
		DomainID:         d.Get("domain_id").(string),
		DomainName:       d.Get("domain_name").(string),
		EndpointType:     d.Get("endpoint_type").(string),
		IdentityEndpoint: d.Get("auth_url").(string),
		Insecure:         d.Get("insecure").(bool),
		Password:         d.Get("password").(string),
		Region:           d.Get("region").(string),
		Swauth:           d.Get("swauth").(bool),
		Token:            d.Get("token").(string),
		TenantID:         d.Get("tenant_id").(string),
		TenantName:       d.Get("tenant_name").(string),
		Username:         d.Get("user_name").(string),
		UserID:           d.Get("user_id").(string),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}