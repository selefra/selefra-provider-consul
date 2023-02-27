package provider

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/selefra_terraform_schema"
	"sort"
	"strings"
)

// terraform resource: consul_prepared_query. S
func GetResource_consul_prepared_query() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_prepared_query",
		TerraformResourceName: "consul_prepared_query",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			prepareQuerys, _, err := client.ConsulClient.PreparedQuery().List(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, query := range prepareQuerys {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: query.ID,
					ArgumentMap: map[string]any{
						"name":    query.Name,
						"service": query.Service.Service,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_config_entry. S
func GetResource_consul_config_entry() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_config_entry",
		TerraformResourceName: "consul_config_entry",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			configEntries, _, err := client.ConsulClient.ConfigEntries().List(client.conf.ConfigEntryKind, &consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}
			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, entry := range configEntries {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: fmt.Sprintf("%s-%s", entry.GetKind(), entry.GetName()),
					ArgumentMap: map[string]any{
						"kind": entry.GetKind(),
						"name": entry.GetName(),
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_agent_service. S
func GetResource_consul_agent_service() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_agent_service",
		TerraformResourceName: "consul_agent_service",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			agentServices, err := client.ConsulClient.Agent().Services()
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for name, service := range agentServices {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: service.ID,
					ArgumentMap: map[string]any{
						"name": name,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_catalog_entry. S
func GetResource_consul_catalog_entry() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_catalog_entry",
		TerraformResourceName: "consul_catalog_entry",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)

			nodes, _, err := client.ConsulClient.Catalog().Nodes(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}
			for _, node := range nodes {
				services, _, err := client.ConsulClient.Catalog().Services(&consulapi.QueryOptions{})
				if err != nil {
					return nil, schema.NewDiagnostics().AddError(err)
				}
				serviceIDs := make([]string, 0)
				for name, _ := range services {
					serviceIDs = append(serviceIDs, name)
				}
				sort.Strings(serviceIDs)
				serviceIDsJoined := strings.Join(serviceIDs, ",")
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: fmt.Sprintf("%s-%s-[%s]", node.Node, node.Address, serviceIDsJoined),
					ArgumentMap: map[string]any{
						"node":    node.Node,
						"address": node.Address,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_admin_partition (enterprise only)
func GetResource_consul_admin_partition() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_admin_partition",
		TerraformResourceName: "consul_admin_partition",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_keys
func GetResource_consul_keys() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_keys",
		TerraformResourceName: "consul_keys",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			//client := taskClient.(*Client)
			//
			//kvs, _, err := client.ConsulClient.KV().List("", &consulapi.QueryOptions{})
			//if err != nil {
			//	return nil, schema.NewDiagnostics().AddError(err)
			//}
			//
			//resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			//for _, kv := range kvs {
			//	resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
			//		ID: "consul",
			//		ArgumentMap: map[string]any{
			//			//"key": []string{fmt.Sprintf(`[path: %s, value: %s]`, kv.Key, string(kv.Value))},
			//			//"key": fmt.Sprintf("[path: %s, value: %s]", kv.Key, string(kv.Value)),
			//			"key": map[string]any{
			//				"path":  kv.Key,
			//				"value": string(kv.Value),
			//			},
			//		},
			//	})
			//}
			//
			//return resourceRequestParamSlice, nil

			// TODO

			return nil, nil
		},
	}
}

// terraform resource: consul_acl_token. S
func GetResource_consul_acl_token() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_token",
		TerraformResourceName: "consul_acl_token",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			tokens, _, err := client.ConsulClient.ACL().TokenList(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, token := range tokens {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: token.AccessorID,
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_license  (enterprise only)
func GetResource_consul_license() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_license",
		TerraformResourceName: "consul_license",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			license, err := client.ConsulClient.Operator().LicenseGet(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}
			resourceRequestParamSlice := []*selefra_terraform_schema.ResourceRequestParam{
				&selefra_terraform_schema.ResourceRequestParam{
					ID: license.License.LicenseID,
					//ArgumentMap: map[string]any{
					//	"license": license.License.Product
					//},
				},
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_acl_policy. S
func GetResource_consul_acl_policy() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_policy",
		TerraformResourceName: "consul_acl_policy",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			entries, _, err := client.ConsulClient.ACL().PolicyList(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, entry := range entries {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: entry.ID,
					ArgumentMap: map[string]any{
						"name": entry.Name,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_namespace_role_attachment (enterprise only)
func GetResource_consul_namespace_role_attachment() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_namespace_role_attachment",
		TerraformResourceName: "consul_namespace_role_attachment",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_peering
func GetResource_consul_peering() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_peering",
		TerraformResourceName: "consul_peering",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {

			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_certificate_authority. S
func GetResource_consul_certificate_authority() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_certificate_authority",
		TerraformResourceName: "consul_certificate_authority",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			config, _, err := client.ConsulClient.Connect().CAGetConfig(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: "consul-ca",
				ArgumentMap: map[string]any{
					"connect_provider": config.Provider,
					"config":           config.Config,
				},
			})

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_acl_role. S
func GetResource_consul_acl_role() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_role",
		TerraformResourceName: "consul_acl_role",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			roles, _, err := client.ConsulClient.ACL().RoleList(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, role := range roles {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: role.ID,
					ArgumentMap: map[string]any{
						"name": role.Name,
					},
				})
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_acl_token_role_attachment. S
func GetResource_consul_acl_token_role_attachment() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_token_role_attachment",
		TerraformResourceName: "consul_acl_token_role_attachment",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			tokens, _, err := client.ConsulClient.ACL().TokenList(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, token := range tokens {
				for _, role := range token.Roles {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: fmt.Sprintf("%s:%s", token.AccessorID, role.Name),
						ArgumentMap: map[string]any{
							"token_id": token.AccessorID,
							"role_id":  role.ID,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_key_prefix
func GetResource_consul_key_prefix() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_key_prefix",
		TerraformResourceName: "consul_key_prefix",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO

			return nil, nil
		},
	}
}

// terraform resource: consul_node. S
func GetResource_consul_node() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_node",
		TerraformResourceName: "consul_node",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			nodes, _, err := client.ConsulClient.Catalog().Nodes(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, node := range nodes {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: fmt.Sprintf("%s-%s", node.Node, node.Address),
					ArgumentMap: map[string]any{
						"address": node.Address,
						"name":    node.Node,
					},
				})

			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_intention. S
func GetResource_consul_intention() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_intention",
		TerraformResourceName: "consul_intention",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			intentions, _, err := client.ConsulClient.Connect().Intentions(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, t := range intentions {
				resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
					ID: t.ID,
					ArgumentMap: map[string]any{
						"source_name":      t.SourceName,
						"destination_name": t.DestinationName,
						"action":           t.Action,
					},
				})

			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_service. S
func GetResource_consul_service() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_service",
		TerraformResourceName: "consul_service",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			serviceNames, _, err := client.ConsulClient.Catalog().Services(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for name, _ := range serviceNames {
				services, _, err := client.ConsulClient.Catalog().Service(name, "", &consulapi.QueryOptions{})
				if err != nil {
					return nil, schema.NewDiagnostics().AddError(err)
				}

				for _, s := range services {

					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: s.ServiceID,
						ArgumentMap: map[string]any{
							"name": s.ServiceName,
							"node": s.Node,
						},
					})

				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_acl_binding_rule
func GetResource_consul_acl_binding_rule() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_binding_rule",
		TerraformResourceName: "consul_acl_binding_rule",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_namespace (enterprise only)
func GetResource_consul_namespace() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_namespace",
		TerraformResourceName: "consul_namespace",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_acl_auth_method
func GetResource_consul_acl_auth_method() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_auth_method",
		TerraformResourceName: "consul_acl_auth_method",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_network_area  (enterprise only)
func GetResource_consul_network_area() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_network_area",
		TerraformResourceName: "consul_network_area",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_namespace_policy_attachment (enterprise only)
func GetResource_consul_namespace_policy_attachment() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_namespace_policy_attachment",
		TerraformResourceName: "consul_namespace_policy_attachment",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_peering_token
func GetResource_consul_peering_token() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_peering_token",
		TerraformResourceName: "consul_peering_token",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			// TODO
			return nil, nil
		},
	}
}

// terraform resource: consul_acl_token_policy_attachment. S
func GetResource_consul_acl_token_policy_attachment() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_acl_token_policy_attachment",
		TerraformResourceName: "consul_acl_token_policy_attachment",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			tokens, _, err := client.ConsulClient.ACL().TokenList(&consulapi.QueryOptions{})
			if err != nil {
				return nil, schema.NewDiagnostics().AddError(err)
			}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			for _, token := range tokens {
				for _, policy := range token.Policies {
					resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
						ID: fmt.Sprintf("%s:%s", token.AccessorID, policy.Name),
						ArgumentMap: map[string]any{
							"token_id": token.AccessorID,
							"policy":   policy.Name,
						},
					})
				}
			}

			return resourceRequestParamSlice, nil
		},
	}
}

// terraform resource: consul_autopilot_config. S
func GetResource_consul_autopilot_config() *selefra_terraform_schema.SelefraTerraformResource {
	return &selefra_terraform_schema.SelefraTerraformResource{
		SelefraTableName:      "consul_autopilot_config",
		TerraformResourceName: "consul_autopilot_config",
		Description:           "",
		SubTables:             nil,
		ListResourceParamsFunc: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) ([]*selefra_terraform_schema.ResourceRequestParam, *schema.Diagnostics) {
			client := taskClient.(*Client)

			//conf, err := client.ConsulClient.Operator().AutopilotGetConfiguration(&consulapi.QueryOptions{})
			//if err != nil {
			//	return nil, schema.NewDiagnostics().AddError(err)
			//}

			resourceRequestParamSlice := make([]*selefra_terraform_schema.ResourceRequestParam, 0)
			resourceRequestParamSlice = append(resourceRequestParamSlice, &selefra_terraform_schema.ResourceRequestParam{
				ID: fmt.Sprintf("consul-autopilot-%s", client.conf.Datacenter),
			})

			return resourceRequestParamSlice, nil
		},
	}
}
