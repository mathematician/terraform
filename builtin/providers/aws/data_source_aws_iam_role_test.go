package aws

import (
    //    "fmt"
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
    //    "github.com/hashicorp/terraform/terraform"
)

func TestAccAWSDataSourceIAMRole_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAccAwsIAMRoleConfig,
                Check: resource.ComposeTestCheckFunc(
                    //                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "arn", "arn:aws:iam::120082580449:role/testpath/TestRole"),
                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "assume_role_policy_document", "%7B%22Version%22%3A%222012-10-17%22%2C%22Statement%22%3A%5B%7B%22Sid%22%3A%22%22%2C%22Effect%22%3A%22Allow%22%2C%22Principal%22%3A%7B%22AWS%22%3A%22arn%3Aaws%3Aiam%3A%3A120082580449%3Atestpath%2FTestRole%22%7D%2C%22Action%22%3A%22sts%3AAssumeRole%22%7D%5D%7D"),
                    //                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "create_date", "2017-02-10 09:22 CDT"),
                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "path", "/testpath/"),
                    //                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "role_id", "AROAIDJMYO7TVEMT2BX32"),
                    resource.TestCheckResourceAttr("data.aws_iam_role.test", "role_name", "TestRole"),
                ),
            },
        },
    })
}

const testAccAwsIAMRoleConfig = `
resource "aws_iam_role" "test_role" {
  name = "TestRole"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

path = "/testpath/"
}

data "aws_iam_role" "test" {
    role_name = "TestRole"
}
`
