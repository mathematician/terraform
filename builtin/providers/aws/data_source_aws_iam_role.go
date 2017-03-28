pckage aws

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsIAMRole() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsIAMRoleRead,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,	
			},
			"assume_role_policy_document": {
				Type: schema.TypeString,
				Computed: true,
			},
			"create_date": {
				Type: schema.TypeString,
				Computed: true,
			},
			"path": {
				Type: schema.TypeString,
				Computed: true,
			},
			"role_id": {
				Type: schema.TypeString,
				Computed: true,	
			},
			"role_name": {
				Type: schema.TypeString,
				Computed: true,	
			},
			"tags": dataSourceTagsSchema(),
		},
	}
}

func dataSourceAwsIAMRoleRead(d *schema.ResourceData, meta interface{}) error {
	iamconn := meta.(*AWSClient).iamconn

	req := &iam.ListRoleInput{}
	resp, err := conn.ListRole(req)
    if err != nil {
        return err
    }
		if resp == nil || len(resp.Arn) == 0 {
        return fmt.Errorf("no IAM role found")
    }

	role := resp.Roles[0]
	d.setId(role.RoleId)
	d.set("arn", role.Arn)
	d.set("assume_role_policy_document", role.AssumeRolePolicyDocument)
	d.set("create_date", role.CreateDate)
	d.set("role_id", role.RoleId)
	d.set("role_name", role.RoleName)
	d.set("tags", tagsToMap(role.Tags))

	return nil
}
