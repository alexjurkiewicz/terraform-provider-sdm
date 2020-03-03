package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceResource() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceResourceCreate),
		Read:   wrapCrudOperation(resourceResourceRead),
		Update: wrapCrudOperation(resourceResourceUpdate),
		Delete: wrapCrudOperation(resourceResourceDelete),
		Schema: map[string]*schema.Schema{
			"athena": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"output": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"big_query": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"private_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"project": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cassandra": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"druid": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"dynamo_db": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"amazon_es": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"elastic": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"http_basic_auth": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"http_no_auth": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"http_auth": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"healthcheck_path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_header": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"headers_blacklist": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"kubernetes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_certificate_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_key_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"kubernetes_basic_auth": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"kubernetes_service_account": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"token": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"amazon_eks": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"secret_access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"cluster_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"google_gke": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"endpoint": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_account_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"service_account_key_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"aks": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"certificate_authority": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"certificate_authority_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_certificate_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"client_key_filename": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"aks_basic_auth": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"aks_service_account": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"token": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"memcached": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mongo_legacy_host": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mongo_legacy_replicaset": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"connect_to_replica": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mongo_host": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mongo_replica_set": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"auth_database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"replica_set": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"connect_to_replica": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mysql": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"aurora_mysql": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"clustrix": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"maria": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"memsql": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"oracle": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"postgres": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"aurora_postgres": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"greenplum": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cockroach": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"redshift": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"presto": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"rdp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"redis": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"elasticache_redis": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tls_required": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"snowflake": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"schema": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
					},
				},
			},
			"sql_server": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"database": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"schema": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"override_database": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"public_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"sybase": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"sybase_iq": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"teradata": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Resource.",
						},
						"hostname": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "",
						},
						"port_override": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}
func resourceFromResourceData(d *schema.ResourceData) sdm.Resource {
	if list := d.Get("athena").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Athena{}
		}
		return &sdm.Athena{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			AccessKey:       stringFromMap(raw, "access_key"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Output:          stringFromMap(raw, "output"),
			PortOverride:    int32FromMap(raw, "port_override"),
			Region:          stringFromMap(raw, "region"),
		}
	}
	if list := d.Get("big_query").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.BigQuery{}
		}
		return &sdm.BigQuery{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			PrivateKey:   stringFromMap(raw, "private_key"),
			Project:      stringFromMap(raw, "project"),
			PortOverride: int32FromMap(raw, "port_override"),
			Endpoint:     stringFromMap(raw, "endpoint"),
			Username:     stringFromMap(raw, "username"),
		}
	}
	if list := d.Get("cassandra").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Cassandra{}
		}
		return &sdm.Cassandra{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("druid").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Druid{}
		}
		return &sdm.Druid{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("dynamo_db").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.DynamoDB{}
		}
		return &sdm.DynamoDB{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			AccessKey:       stringFromMap(raw, "access_key"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Region:          stringFromMap(raw, "region"),
			Endpoint:        stringFromMap(raw, "endpoint"),
			PortOverride:    int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("amazon_es").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AmazonES{}
		}
		return &sdm.AmazonES{
			ID:              d.Id(),
			Name:            stringFromMap(raw, "name"),
			Region:          stringFromMap(raw, "region"),
			SecretAccessKey: stringFromMap(raw, "secret_access_key"),
			Endpoint:        stringFromMap(raw, "endpoint"),
			AccessKey:       stringFromMap(raw, "access_key"),
			PortOverride:    int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("elastic").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Elastic{}
		}
		return &sdm.Elastic{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("http_basic_auth").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.HTTPBasicAuth{}
		}
		return &sdm.HTTPBasicAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("http_no_auth").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.HTTPNoAuth{}
		}
		return &sdm.HTTPNoAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("http_auth").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.HTTPAuth{}
		}
		return &sdm.HTTPAuth{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Url:              stringFromMap(raw, "url"),
			HealthcheckPath:  stringFromMap(raw, "healthcheck_path"),
			AuthHeader:       stringFromMap(raw, "auth_header"),
			HeadersBlacklist: stringFromMap(raw, "headers_blacklist"),
			DefaultPath:      stringFromMap(raw, "default_path"),
		}
	}
	if list := d.Get("kubernetes").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Kubernetes{}
		}
		return &sdm.Kubernetes{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Hostname:                     stringFromMap(raw, "hostname"),
			Port:                         int32FromMap(raw, "port"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			ClientCertificate:            stringFromMap(raw, "client_certificate"),
			ClientCertificateFilename:    stringFromMap(raw, "client_certificate_filename"),
			ClientKey:                    stringFromMap(raw, "client_key"),
			ClientKeyFilename:            stringFromMap(raw, "client_key_filename"),
		}
	}
	if list := d.Get("kubernetes_basic_auth").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.KubernetesBasicAuth{}
		}
		return &sdm.KubernetesBasicAuth{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Port:     int32FromMap(raw, "port"),
			Username: stringFromMap(raw, "username"),
			Password: stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("kubernetes_service_account").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.KubernetesServiceAccount{}
		}
		return &sdm.KubernetesServiceAccount{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Port:     int32FromMap(raw, "port"),
			Token:    stringFromMap(raw, "token"),
		}
	}
	if list := d.Get("amazon_eks").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AmazonEKS{}
		}
		return &sdm.AmazonEKS{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Endpoint:                     stringFromMap(raw, "endpoint"),
			AccessKey:                    stringFromMap(raw, "access_key"),
			SecretAccessKey:              stringFromMap(raw, "secret_access_key"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			Region:                       stringFromMap(raw, "region"),
			ClusterName:                  stringFromMap(raw, "cluster_name"),
		}
	}
	if list := d.Get("google_gke").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.GoogleGKE{}
		}
		return &sdm.GoogleGKE{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Endpoint:                     stringFromMap(raw, "endpoint"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			ServiceAccountKey:            stringFromMap(raw, "service_account_key"),
			ServiceAccountKeyFilename:    stringFromMap(raw, "service_account_key_filename"),
		}
	}
	if list := d.Get("aks").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AKS{}
		}
		return &sdm.AKS{
			ID:                           d.Id(),
			Name:                         stringFromMap(raw, "name"),
			Hostname:                     stringFromMap(raw, "hostname"),
			Port:                         int32FromMap(raw, "port"),
			CertificateAuthority:         stringFromMap(raw, "certificate_authority"),
			CertificateAuthorityFilename: stringFromMap(raw, "certificate_authority_filename"),
			ClientCertificate:            stringFromMap(raw, "client_certificate"),
			ClientCertificateFilename:    stringFromMap(raw, "client_certificate_filename"),
			ClientKey:                    stringFromMap(raw, "client_key"),
			ClientKeyFilename:            stringFromMap(raw, "client_key_filename"),
		}
	}
	if list := d.Get("aks_basic_auth").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AKSBasicAuth{}
		}
		return &sdm.AKSBasicAuth{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Port:     int32FromMap(raw, "port"),
			Username: stringFromMap(raw, "username"),
			Password: stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("aks_service_account").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AKSServiceAccount{}
		}
		return &sdm.AKSServiceAccount{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Port:     int32FromMap(raw, "port"),
			Token:    stringFromMap(raw, "token"),
		}
	}
	if list := d.Get("memcached").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Memcached{}
		}
		return &sdm.Memcached{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("mongo_legacy_host").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.MongoLegacyHost{}
		}
		return &sdm.MongoLegacyHost{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			AuthDatabase: stringFromMap(raw, "auth_database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
			ReplicaSet:   stringFromMap(raw, "replica_set"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_legacy_replicaset").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.MongoLegacyReplicaset{}
		}
		return &sdm.MongoLegacyReplicaset{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			AuthDatabase:     stringFromMap(raw, "auth_database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Port:             int32FromMap(raw, "port"),
			ReplicaSet:       stringFromMap(raw, "replica_set"),
			ConnectToReplica: boolFromMap(raw, "connect_to_replica"),
			TlsRequired:      boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_host").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.MongoHost{}
		}
		return &sdm.MongoHost{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			AuthDatabase: stringFromMap(raw, "auth_database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mongo_replica_set").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.MongoReplicaSet{}
		}
		return &sdm.MongoReplicaSet{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			AuthDatabase:     stringFromMap(raw, "auth_database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Port:             int32FromMap(raw, "port"),
			ReplicaSet:       stringFromMap(raw, "replica_set"),
			ConnectToReplica: boolFromMap(raw, "connect_to_replica"),
			TlsRequired:      boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("mysql").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Mysql{}
		}
		return &sdm.Mysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("aurora_mysql").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AuroraMysql{}
		}
		return &sdm.AuroraMysql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("clustrix").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Clustrix{}
		}
		return &sdm.Clustrix{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("maria").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Maria{}
		}
		return &sdm.Maria{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("memsql").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Memsql{}
		}
		return &sdm.Memsql{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("oracle").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Oracle{}
		}
		return &sdm.Oracle{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Port:         int32FromMap(raw, "port"),
			PortOverride: int32FromMap(raw, "port_override"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("postgres").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Postgres{}
		}
		return &sdm.Postgres{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("aurora_postgres").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.AuroraPostgres{}
		}
		return &sdm.AuroraPostgres{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("greenplum").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Greenplum{}
		}
		return &sdm.Greenplum{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("cockroach").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Cockroach{}
		}
		return &sdm.Cockroach{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("redshift").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Redshift{}
		}
		return &sdm.Redshift{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("presto").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Presto{}
		}
		return &sdm.Presto{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			Username:     stringFromMap(raw, "username"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("rdp").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.RDP{}
		}
		return &sdm.RDP{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("redis").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Redis{}
		}
		return &sdm.Redis{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("elasticache_redis").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.ElasticacheRedis{}
		}
		return &sdm.ElasticacheRedis{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			PortOverride: int32FromMap(raw, "port_override"),
			Password:     stringFromMap(raw, "password"),
			Port:         int32FromMap(raw, "port"),
			TlsRequired:  boolFromMap(raw, "tls_required"),
		}
	}
	if list := d.Get("snowflake").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Snowflake{}
		}
		return &sdm.Snowflake{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			Database:     stringFromMap(raw, "database"),
			Schema:       stringFromMap(raw, "schema"),
			PortOverride: int32FromMap(raw, "port_override"),
		}
	}
	if list := d.Get("sql_server").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.SQLServer{}
		}
		return &sdm.SQLServer{
			ID:               d.Id(),
			Name:             stringFromMap(raw, "name"),
			Hostname:         stringFromMap(raw, "hostname"),
			Username:         stringFromMap(raw, "username"),
			Password:         stringFromMap(raw, "password"),
			Database:         stringFromMap(raw, "database"),
			PortOverride:     int32FromMap(raw, "port_override"),
			Schema:           stringFromMap(raw, "schema"),
			Port:             int32FromMap(raw, "port"),
			OverrideDatabase: boolFromMap(raw, "override_database"),
		}
	}
	if list := d.Get("ssh").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.SSH{}
		}
		return &sdm.SSH{
			ID:       d.Id(),
			Name:     stringFromMap(raw, "name"),
			Hostname: stringFromMap(raw, "hostname"),
			Username: stringFromMap(raw, "username"),
			Port:     int32FromMap(raw, "port"),
		}
	}
	if list := d.Get("sybase").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Sybase{}
		}
		return &sdm.Sybase{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			Password:     stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("sybase_iq").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.SybaseIQ{}
		}
		return &sdm.SybaseIQ{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
			Password:     stringFromMap(raw, "password"),
		}
	}
	if list := d.Get("teradata").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Teradata{}
		}
		return &sdm.Teradata{
			ID:           d.Id(),
			Name:         stringFromMap(raw, "name"),
			Hostname:     stringFromMap(raw, "hostname"),
			Username:     stringFromMap(raw, "username"),
			Password:     stringFromMap(raw, "password"),
			PortOverride: int32FromMap(raw, "port_override"),
			Port:         int32FromMap(raw, "port"),
		}
	}
	return nil
}

func resourceResourceCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Resources().Create(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Resource %s: %w", "", err)
	}
	d.SetId(resp.Resource.GetID())
	switch v := resp.Resource.(type) {
	case *sdm.Athena:
		d.Set("athena", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"output":            v.Output,
				"port_override":     v.PortOverride,
				"region":            v.Region,
			},
		})
	case *sdm.BigQuery:
		d.Set("big_query", []map[string]interface{}{
			{
				"name":          v.Name,
				"private_key":   v.PrivateKey,
				"project":       v.Project,
				"port_override": v.PortOverride,
				"endpoint":      v.Endpoint,
				"username":      v.Username,
			},
		})
	case *sdm.Cassandra:
		d.Set("cassandra", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Druid:
		d.Set("druid", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *sdm.DynamoDB:
		d.Set("dynamo_db", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"region":            v.Region,
				"endpoint":          v.Endpoint,
				"port_override":     v.PortOverride,
			},
		})
	case *sdm.AmazonES:
		d.Set("amazon_es", []map[string]interface{}{
			{
				"name":              v.Name,
				"region":            v.Region,
				"secret_access_key": v.SecretAccessKey,
				"endpoint":          v.Endpoint,
				"access_key":        v.AccessKey,
				"port_override":     v.PortOverride,
			},
		})
	case *sdm.Elastic:
		d.Set("elastic", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.HTTPBasicAuth:
		d.Set("http_basic_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"username":          v.Username,
				"password":          v.Password,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.HTTPNoAuth:
		d.Set("http_no_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.HTTPAuth:
		d.Set("http_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"auth_header":       v.AuthHeader,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.Kubernetes:
		d.Set("kubernetes", []map[string]interface{}{
			{
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			},
		})
	case *sdm.KubernetesBasicAuth:
		d.Set("kubernetes_basic_auth", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			},
		})
	case *sdm.KubernetesServiceAccount:
		d.Set("kubernetes_service_account", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"token":    v.Token,
			},
		})
	case *sdm.AmazonEKS:
		d.Set("amazon_eks", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"access_key":                     v.AccessKey,
				"secret_access_key":              v.SecretAccessKey,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"region":                         v.Region,
				"cluster_name":                   v.ClusterName,
			},
		})
	case *sdm.GoogleGKE:
		d.Set("google_gke", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"service_account_key":            v.ServiceAccountKey,
				"service_account_key_filename":   v.ServiceAccountKeyFilename,
			},
		})
	case *sdm.AKS:
		d.Set("aks", []map[string]interface{}{
			{
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			},
		})
	case *sdm.AKSBasicAuth:
		d.Set("aks_basic_auth", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			},
		})
	case *sdm.AKSServiceAccount:
		d.Set("aks_service_account", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"token":    v.Token,
			},
		})
	case *sdm.Memcached:
		d.Set("memcached", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.MongoLegacyHost:
		d.Set("mongo_legacy_host", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"replica_set":   v.ReplicaSet,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.MongoLegacyReplicaset:
		d.Set("mongo_legacy_replicaset", []map[string]interface{}{
			{
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			},
		})
	case *sdm.MongoHost:
		d.Set("mongo_host", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.MongoReplicaSet:
		d.Set("mongo_replica_set", []map[string]interface{}{
			{
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			},
		})
	case *sdm.Mysql:
		d.Set("mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.AuroraMysql:
		d.Set("aurora_mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Clustrix:
		d.Set("clustrix", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Maria:
		d.Set("maria", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Memsql:
		d.Set("memsql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Oracle:
		d.Set("oracle", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
				"port_override": v.PortOverride,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Postgres:
		d.Set("postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.AuroraPostgres:
		d.Set("aurora_postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Greenplum:
		d.Set("greenplum", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Cockroach:
		d.Set("cockroach", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Redshift:
		d.Set("redshift", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Presto:
		d.Set("presto", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"username":      v.Username,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.RDP:
		d.Set("rdp", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Redis:
		d.Set("redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *sdm.ElasticacheRedis:
		d.Set("elasticache_redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Snowflake:
		d.Set("snowflake", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"schema":        v.Schema,
				"port_override": v.PortOverride,
			},
		})
	case *sdm.SQLServer:
		d.Set("sql_server", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"schema":            v.Schema,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.SSH:
		d.Set("ssh", []map[string]interface{}{
			{
				"name":       v.Name,
				"hostname":   v.Hostname,
				"username":   v.Username,
				"port":       v.Port,
				"public_key": v.PublicKey,
			},
		})
	case *sdm.Sybase:
		d.Set("sybase", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			},
		})
	case *sdm.SybaseIQ:
		d.Set("sybase_iq", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			},
		})
	case *sdm.Teradata:
		d.Set("teradata", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	}
	return nil
}

func resourceResourceRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Resources().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Resource %s: %w", d.Id(), err)
	}
	switch v := resp.Resource.(type) {
	case *sdm.Athena:
		d.Set("athena", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"output":            v.Output,
				"port_override":     v.PortOverride,
				"region":            v.Region,
			},
		})
	case *sdm.BigQuery:
		d.Set("big_query", []map[string]interface{}{
			{
				"name":          v.Name,
				"private_key":   v.PrivateKey,
				"project":       v.Project,
				"port_override": v.PortOverride,
				"endpoint":      v.Endpoint,
				"username":      v.Username,
			},
		})
	case *sdm.Cassandra:
		d.Set("cassandra", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Druid:
		d.Set("druid", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *sdm.DynamoDB:
		d.Set("dynamo_db", []map[string]interface{}{
			{
				"name":              v.Name,
				"access_key":        v.AccessKey,
				"secret_access_key": v.SecretAccessKey,
				"region":            v.Region,
				"endpoint":          v.Endpoint,
				"port_override":     v.PortOverride,
			},
		})
	case *sdm.AmazonES:
		d.Set("amazon_es", []map[string]interface{}{
			{
				"name":              v.Name,
				"region":            v.Region,
				"secret_access_key": v.SecretAccessKey,
				"endpoint":          v.Endpoint,
				"access_key":        v.AccessKey,
				"port_override":     v.PortOverride,
			},
		})
	case *sdm.Elastic:
		d.Set("elastic", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.HTTPBasicAuth:
		d.Set("http_basic_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"username":          v.Username,
				"password":          v.Password,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.HTTPNoAuth:
		d.Set("http_no_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.HTTPAuth:
		d.Set("http_auth", []map[string]interface{}{
			{
				"name":              v.Name,
				"url":               v.Url,
				"healthcheck_path":  v.HealthcheckPath,
				"auth_header":       v.AuthHeader,
				"headers_blacklist": v.HeadersBlacklist,
				"default_path":      v.DefaultPath,
			},
		})
	case *sdm.Kubernetes:
		d.Set("kubernetes", []map[string]interface{}{
			{
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			},
		})
	case *sdm.KubernetesBasicAuth:
		d.Set("kubernetes_basic_auth", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			},
		})
	case *sdm.KubernetesServiceAccount:
		d.Set("kubernetes_service_account", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"token":    v.Token,
			},
		})
	case *sdm.AmazonEKS:
		d.Set("amazon_eks", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"access_key":                     v.AccessKey,
				"secret_access_key":              v.SecretAccessKey,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"region":                         v.Region,
				"cluster_name":                   v.ClusterName,
			},
		})
	case *sdm.GoogleGKE:
		d.Set("google_gke", []map[string]interface{}{
			{
				"name":                           v.Name,
				"endpoint":                       v.Endpoint,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"service_account_key":            v.ServiceAccountKey,
				"service_account_key_filename":   v.ServiceAccountKeyFilename,
			},
		})
	case *sdm.AKS:
		d.Set("aks", []map[string]interface{}{
			{
				"name":                           v.Name,
				"hostname":                       v.Hostname,
				"port":                           v.Port,
				"certificate_authority":          v.CertificateAuthority,
				"certificate_authority_filename": v.CertificateAuthorityFilename,
				"client_certificate":             v.ClientCertificate,
				"client_certificate_filename":    v.ClientCertificateFilename,
				"client_key":                     v.ClientKey,
				"client_key_filename":            v.ClientKeyFilename,
			},
		})
	case *sdm.AKSBasicAuth:
		d.Set("aks_basic_auth", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"username": v.Username,
				"password": v.Password,
			},
		})
	case *sdm.AKSServiceAccount:
		d.Set("aks_service_account", []map[string]interface{}{
			{
				"name":     v.Name,
				"hostname": v.Hostname,
				"port":     v.Port,
				"token":    v.Token,
			},
		})
	case *sdm.Memcached:
		d.Set("memcached", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.MongoLegacyHost:
		d.Set("mongo_legacy_host", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"replica_set":   v.ReplicaSet,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.MongoLegacyReplicaset:
		d.Set("mongo_legacy_replicaset", []map[string]interface{}{
			{
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			},
		})
	case *sdm.MongoHost:
		d.Set("mongo_host", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"auth_database": v.AuthDatabase,
				"port_override": v.PortOverride,
				"username":      v.Username,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.MongoReplicaSet:
		d.Set("mongo_replica_set", []map[string]interface{}{
			{
				"name":               v.Name,
				"hostname":           v.Hostname,
				"auth_database":      v.AuthDatabase,
				"port_override":      v.PortOverride,
				"username":           v.Username,
				"password":           v.Password,
				"port":               v.Port,
				"replica_set":        v.ReplicaSet,
				"connect_to_replica": v.ConnectToReplica,
				"tls_required":       v.TlsRequired,
			},
		})
	case *sdm.Mysql:
		d.Set("mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.AuroraMysql:
		d.Set("aurora_mysql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Clustrix:
		d.Set("clustrix", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Maria:
		d.Set("maria", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Memsql:
		d.Set("memsql", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Oracle:
		d.Set("oracle", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"port":          v.Port,
				"port_override": v.PortOverride,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Postgres:
		d.Set("postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.AuroraPostgres:
		d.Set("aurora_postgres", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Greenplum:
		d.Set("greenplum", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Cockroach:
		d.Set("cockroach", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Redshift:
		d.Set("redshift", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.Presto:
		d.Set("presto", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"password":      v.Password,
				"database":      v.Database,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"username":      v.Username,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.RDP:
		d.Set("rdp", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	case *sdm.Redis:
		d.Set("redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
			},
		})
	case *sdm.ElasticacheRedis:
		d.Set("elasticache_redis", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"port_override": v.PortOverride,
				"password":      v.Password,
				"port":          v.Port,
				"tls_required":  v.TlsRequired,
			},
		})
	case *sdm.Snowflake:
		d.Set("snowflake", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"database":      v.Database,
				"schema":        v.Schema,
				"port_override": v.PortOverride,
			},
		})
	case *sdm.SQLServer:
		d.Set("sql_server", []map[string]interface{}{
			{
				"name":              v.Name,
				"hostname":          v.Hostname,
				"username":          v.Username,
				"password":          v.Password,
				"database":          v.Database,
				"port_override":     v.PortOverride,
				"schema":            v.Schema,
				"port":              v.Port,
				"override_database": v.OverrideDatabase,
			},
		})
	case *sdm.SSH:
		d.Set("ssh", []map[string]interface{}{
			{
				"name":       v.Name,
				"hostname":   v.Hostname,
				"username":   v.Username,
				"port":       v.Port,
				"public_key": v.PublicKey,
			},
		})
	case *sdm.Sybase:
		d.Set("sybase", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			},
		})
	case *sdm.SybaseIQ:
		d.Set("sybase_iq", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"port_override": v.PortOverride,
				"port":          v.Port,
				"password":      v.Password,
			},
		})
	case *sdm.Teradata:
		d.Set("teradata", []map[string]interface{}{
			{
				"name":          v.Name,
				"hostname":      v.Hostname,
				"username":      v.Username,
				"password":      v.Password,
				"port_override": v.PortOverride,
				"port":          v.Port,
			},
		})
	}
	return nil
}
func resourceResourceUpdate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Resources().Update(ctx, resourceFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Resource %s: %w", d.Id(), err)
	}
	d.SetId(resp.Resource.GetID())
	return resourceResourceRead(d, cc)
}
func resourceResourceDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.Resources().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
