package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceNodeCreate),
		Read:   wrapCrudOperation(resourceNodeRead),
		Update: wrapCrudOperation(resourceNodeUpdate),
		Delete: wrapCrudOperation(resourceNodeDelete),
		Schema: map[string]*schema.Schema{
			"relay": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Relay represents a StrongDM CLI installation running in relay mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Relay. Generated if not provided on create.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
						},
						"token": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gateway": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway represents a StrongDM CLI installation running in gateway mode.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique human-readable name of the Gateway. Generated if not provided on create.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
						},
						"listen_address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The public hostname/port tuple at which the gateway will be accessible to clients.",
						},
						"bind_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hostname/port tuple which the gateway daemon will bind to.\n If not provided on create, set to \"0.0.0.0:<listen_address_port>\".",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return new == ""
							},
						},
						"token": {
							Type:     schema.TypeString,
							Computed: true,
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
func nodeFromResourceData(d *schema.ResourceData) sdm.Node {
	if list := d.Get("relay").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Relay{}
		}
		return &sdm.Relay{
			ID:   d.Id(),
			Name: stringFromMap(raw, "name"),
		}
	}
	if list := d.Get("gateway").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Gateway{}
		}
		return &sdm.Gateway{
			ID:            d.Id(),
			Name:          stringFromMap(raw, "name"),
			ListenAddress: stringFromMap(raw, "listen_address"),
			BindAddress:   stringFromMap(raw, "bind_address"),
		}
	}
	return nil
}

func resourceNodeCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Nodes().Create(ctx, nodeFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Node %s: %w", "", err)
	}
	d.SetId(resp.Node.GetID())
	switch v := resp.Node.(type) {
	case *sdm.Relay:
		d.Set("relay", []map[string]interface{}{
			{
				"name":  v.Name,
				"token": resp.Token,
			},
		})
	case *sdm.Gateway:
		d.Set("gateway", []map[string]interface{}{
			{
				"name":           v.Name,
				"listen_address": v.ListenAddress,
				"bind_address":   v.BindAddress,
				"token":          resp.Token,
			},
		})
	}
	return nil
}

func resourceNodeRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Nodes().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Node %s: %w", d.Id(), err)
	}
	switch v := resp.Node.(type) {
	case *sdm.Relay:
		d.Set("relay", []map[string]interface{}{
			{
				"name":  v.Name,
				"token": d.Get("relay.0.token"),
			},
		})
	case *sdm.Gateway:
		d.Set("gateway", []map[string]interface{}{
			{
				"name":           v.Name,
				"listen_address": v.ListenAddress,
				"bind_address":   v.BindAddress,
				"token":          d.Get("gateway.0.token"),
			},
		})
	}
	return nil
}
func resourceNodeUpdate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Nodes().Update(ctx, nodeFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Node %s: %w", d.Id(), err)
	}
	d.SetId(resp.Node.GetID())
	return resourceNodeRead(d, cc)
}
func resourceNodeDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.Nodes().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
