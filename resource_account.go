package sdm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	sdm "github.com/strongdm/strongdm-sdk-go"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: wrapCrudOperation(resourceAccountCreate),
		Read:   wrapCrudOperation(resourceAccountRead),
		Update: wrapCrudOperation(resourceAccountUpdate),
		Delete: wrapCrudOperation(resourceAccountDelete),
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A User can connect to resources they are granted directly, or granted\n via roles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The User's email address. Must be unique.",
						},
						"first_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The User's first name.",
						},
						"last_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The User's last name.",
						},
						"suspended": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The User's suspended state.",
						},
					},
				},
			},
			"service": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Service is a service account that can connect to resources they are granted\n directly, or granted via roles. Services are typically automated jobs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique human-readable name of the Service.",
						},
						"suspended": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Service's suspended state.",
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
func accountFromResourceData(d *schema.ResourceData) sdm.Account {
	if list := d.Get("user").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.User{}
		}
		return &sdm.User{
			ID:        d.Id(),
			Email:     stringFromMap(raw, "email"),
			FirstName: stringFromMap(raw, "first_name"),
			LastName:  stringFromMap(raw, "last_name"),
			Suspended: boolFromMap(raw, "suspended"),
		}
	}
	if list := d.Get("service").([]interface{}); len(list) > 0 {
		raw, ok := list[0].(map[string]interface{})
		if !ok {
			return &sdm.Service{}
		}
		return &sdm.Service{
			ID:        d.Id(),
			Name:      stringFromMap(raw, "name"),
			Suspended: boolFromMap(raw, "suspended"),
		}
	}
	return nil
}

func resourceAccountCreate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutCreate))
	defer cancel()
	resp, err := cc.Accounts().Create(ctx, accountFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot create Account %s: %w", "", err)
	}
	d.SetId(resp.Account.GetID())
	switch v := resp.Account.(type) {
	case *sdm.User:
		d.Set("user", []map[string]interface{}{
			{
				"email":      v.Email,
				"first_name": v.FirstName,
				"last_name":  v.LastName,
				"suspended":  v.Suspended,
			},
		})
	case *sdm.Service:
		d.Set("service", []map[string]interface{}{
			{
				"name":      v.Name,
				"suspended": v.Suspended,
				"token":     resp.Token,
			},
		})
	}
	return nil
}

func resourceAccountRead(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	resp, err := cc.Accounts().Get(ctx, d.Id())
	var errNotFound *sdm.NotFoundError
	if err != nil && errors.As(err, &errNotFound) {
		d.SetId("")
		return nil
	} else if err != nil {
		return fmt.Errorf("cannot read Account %s: %w", d.Id(), err)
	}
	switch v := resp.Account.(type) {
	case *sdm.User:
		d.Set("user", []map[string]interface{}{
			{
				"email":      v.Email,
				"first_name": v.FirstName,
				"last_name":  v.LastName,
				"suspended":  v.Suspended,
			},
		})
	case *sdm.Service:
		d.Set("service", []map[string]interface{}{
			{
				"name":      v.Name,
				"suspended": v.Suspended,
				"token":     d.Get("service.0.token"),
			},
		})
	}
	return nil
}
func resourceAccountUpdate(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutUpdate))
	defer cancel()
	resp, err := cc.Accounts().Update(ctx, accountFromResourceData(d))
	if err != nil {
		return fmt.Errorf("cannot update Account %s: %w", d.Id(), err)
	}
	d.SetId(resp.Account.GetID())
	return resourceAccountRead(d, cc)
}
func resourceAccountDelete(d *schema.ResourceData, cc *sdm.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutDelete))
	defer cancel()
	var errNotFound *sdm.NotFoundError
	_, err := cc.Accounts().Delete(ctx, d.Id())
	if err != nil && errors.As(err, &errNotFound) {
		return nil
	}
	return err
}
