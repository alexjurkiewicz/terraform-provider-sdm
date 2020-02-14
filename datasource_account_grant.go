package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	apiv1 "github.com/strongdm/strongdm-sdk-go"
)

func dataSourceAccountGrant() *schema.Resource {
	return &schema.Resource{
		Read: wrapCrudOperation(dataSourceAccountGrantList),
		Schema: map[string]*schema.Schema{

			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the AccountGrant.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the composite role of this AccountGrant.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the attached role of this AccountGrant.",
			},
			"account_grants": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unique identifier of the AccountGrant.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The id of the composite role of this AccountGrant.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The id of the attached role of this AccountGrant.",
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

func accountGrantFilterFromResourceData(d *schema.ResourceData) (string, []interface{}) {
	filter := ""
	args := []interface{}{}
	if v, ok := d.GetOk("id"); ok {
		filter += "id:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("resource_id"); ok {
		filter += "resourceid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("account_id"); ok {
		filter += "accountid:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("start_from"); ok {
		filter += "startfrom:? "
		args = append(args, v)
	}
	if v, ok := d.GetOk("valid_until"); ok {
		filter += "validuntil:? "
		args = append(args, v)
	}
	return filter, args
}

func dataSourceAccountGrantList(d *schema.ResourceData, cc *apiv1.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout(schema.TimeoutRead))
	defer cancel()
	filter, args := accountGrantFilterFromResourceData(d)
	resp, err := cc.AccountGrants().List(ctx, filter, args...)
	if err != nil {
		return fmt.Errorf("cannot list AccountGrants %s: %w", d.Id(), err)
	}
	type entity = map[string]interface{}
	output := make([]entity, 0)
	for resp.Next() {
		v := resp.Value()
		output = append(output,
			entity{
				"id":          v.ID,
				"resource_id": v.ResourceID,
				"account_id":  v.AccountID,
			})
	}
	if resp.Err() != nil {
		return fmt.Errorf("failure during list: %w", resp.Err())
	}

	err = d.Set("account_grants", output)
	if err != nil {
		return fmt.Errorf("cannot set vList: %w", err)
	}
	d.SetId("AccountGrant" + filter + fmt.Sprint(args...))
	return nil
}
