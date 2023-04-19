// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
		},
		{
			Factory:  DataSourceRepository,
			TypeName: "aws_codecommit_repository",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
		},
		{
			Factory:  ResourceApprovalRuleTemplateAssociation,
			TypeName: "aws_codecommit_approval_rule_template_association",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_codecommit_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTrigger,
			TypeName: "aws_codecommit_trigger",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeCommit
}

var ServicePackage = &servicePackage{}
