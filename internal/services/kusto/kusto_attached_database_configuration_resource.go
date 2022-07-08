package kusto

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2021-08-27/attacheddatabaseconfigurations"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/kusto/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/kusto/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceKustoAttachedDatabaseConfiguration() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceKustoAttachedDatabaseConfigurationCreateUpdate,
		Read:   resourceKustoAttachedDatabaseConfigurationRead,
		Update: resourceKustoAttachedDatabaseConfigurationCreateUpdate,
		Delete: resourceKustoAttachedDatabaseConfigurationDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.AttachedDatabaseConfigurationID(id)
			return err
		}),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(60 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(60 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.DataConnectionName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"cluster_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.ClusterName,
			},

			"database_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.Any(validate.DatabaseName, validation.StringInSlice([]string{"*"}, false)),
			},

			"cluster_resource_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"attached_database_names": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"default_principal_modification_kind": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Default:  attacheddatabaseconfigurations.DefaultPrincipalsModificationKindNone,
				ValidateFunc: validation.StringInSlice([]string{
					string(attacheddatabaseconfigurations.DefaultPrincipalsModificationKindNone),
					string(attacheddatabaseconfigurations.DefaultPrincipalsModificationKindReplace),
					string(attacheddatabaseconfigurations.DefaultPrincipalsModificationKindUnion),
				}, false),
			},

			"sharing": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"external_tables_to_exclude": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"external_tables_to_include": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"materialized_views_to_exclude": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"materialized_views_to_include": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"tables_to_exclude": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"tables_to_include": {
							Type:     pluginsdk.TypeSet,
							Optional: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceKustoAttachedDatabaseConfigurationCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.AttachedDatabaseConfigurationsClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := attacheddatabaseconfigurations.NewAttachedDatabaseConfigurationID(subscriptionId, d.Get("resource_group_name").(string), d.Get("cluster_name").(string), d.Get("name").(string))
	if d.IsNewResource() {
		resp, err := client.Get(ctx, id)
		if err != nil {
			if !response.WasNotFound(resp.HttpResponse) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
		}

		if !response.WasNotFound(resp.HttpResponse) {
			return tf.ImportAsExistsError("azurerm_kusto_attached_database_configuration", id.ID())
		}
	}

	configurationRequest := attacheddatabaseconfigurations.AttachedDatabaseConfiguration{
		Location:   utils.String(location.Normalize(d.Get("location").(string))),
		Properties: expandKustoAttachedDatabaseConfigurationProperties(d),
	}

	if err := client.CreateOrUpdateThenPoll(ctx, id, configurationRequest); err != nil {
		return fmt.Errorf("creating/updating %s: %+v", id, err)
	}

	d.SetId(id.ID())
	return resourceKustoAttachedDatabaseConfigurationRead(d, meta)
}

func resourceKustoAttachedDatabaseConfigurationRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.AttachedDatabaseConfigurationsClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := attacheddatabaseconfigurations.ParseAttachedDatabaseConfigurationID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving %s: %+v", id, err)
	}

	d.Set("name", id.AttachedDatabaseConfigurationName)
	d.Set("resource_group_name", id.ResourceGroupName)
	d.Set("cluster_name", id.ClusterName)

	if model := resp.Model; model != nil {

		d.Set("location", location.NormalizeNilable(model.Location))

		if props := model.Properties; props != nil {
			d.Set("cluster_resource_id", props.ClusterResourceId)
			d.Set("database_name", props.DatabaseName)
			d.Set("default_principal_modification_kind", props.DefaultPrincipalsModificationKind)
			d.Set("attached_database_names", props.AttachedDatabaseNames)
			d.Set("sharing", flattenAttachedDatabaseConfigurationTableLevelSharingProperties(props.TableLevelSharingProperties))
		}
	}

	return nil
}

func resourceKustoAttachedDatabaseConfigurationDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Kusto.AttachedDatabaseConfigurationsClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := attacheddatabaseconfigurations.ParseAttachedDatabaseConfigurationID(d.Id())
	if err != nil {
		return err
	}

	if err = client.DeleteThenPoll(ctx, *id); err != nil {
		return fmt.Errorf("deleting %s: %+v", id, err)
	}

	return nil
}

func expandKustoAttachedDatabaseConfigurationProperties(d *pluginsdk.ResourceData) *attacheddatabaseconfigurations.AttachedDatabaseConfigurationProperties {
	AttachedDatabaseConfigurationProperties := &attacheddatabaseconfigurations.AttachedDatabaseConfigurationProperties{}

	if clusterResourceID, ok := d.GetOk("cluster_resource_id"); ok {
		AttachedDatabaseConfigurationProperties.ClusterResourceId = clusterResourceID.(string)
	}

	if databaseName, ok := d.GetOk("database_name"); ok {
		AttachedDatabaseConfigurationProperties.DatabaseName = databaseName.(string)
	}

	if defaultPrincipalModificationKind, ok := d.GetOk("default_principal_modification_kind"); ok {
		AttachedDatabaseConfigurationProperties.DefaultPrincipalsModificationKind = attacheddatabaseconfigurations.DefaultPrincipalsModificationKind(defaultPrincipalModificationKind.(string))
	}

	AttachedDatabaseConfigurationProperties.TableLevelSharingProperties = expandAttachedDatabaseConfigurationTableLevelSharingProperties(d.Get("sharing").([]interface{}))

	return AttachedDatabaseConfigurationProperties
}

func expandAttachedDatabaseConfigurationTableLevelSharingProperties(input []interface{}) *attacheddatabaseconfigurations.TableLevelSharingProperties {
	if len(input) == 0 {
		return nil
	}
	v := input[0].(map[string]interface{})
	return &attacheddatabaseconfigurations.TableLevelSharingProperties{
		TablesToInclude:            utils.ExpandStringSlice(v["tables_to_include"].(*pluginsdk.Set).List()),
		TablesToExclude:            utils.ExpandStringSlice(v["tables_to_exclude"].(*pluginsdk.Set).List()),
		ExternalTablesToInclude:    utils.ExpandStringSlice(v["external_tables_to_include"].(*pluginsdk.Set).List()),
		ExternalTablesToExclude:    utils.ExpandStringSlice(v["external_tables_to_exclude"].(*pluginsdk.Set).List()),
		MaterializedViewsToInclude: utils.ExpandStringSlice(v["materialized_views_to_include"].(*pluginsdk.Set).List()),
		MaterializedViewsToExclude: utils.ExpandStringSlice(v["materialized_views_to_exclude"].(*pluginsdk.Set).List()),
	}
}

func flattenAttachedDatabaseConfigurationTableLevelSharingProperties(input *attacheddatabaseconfigurations.TableLevelSharingProperties) []interface{} {
	if input == nil {
		return make([]interface{}, 0)
	}

	return []interface{}{
		map[string]interface{}{
			"external_tables_to_exclude":    utils.FlattenStringSlice(input.ExternalTablesToExclude),
			"external_tables_to_include":    utils.FlattenStringSlice(input.ExternalTablesToInclude),
			"materialized_views_to_exclude": utils.FlattenStringSlice(input.MaterializedViewsToExclude),
			"materialized_views_to_include": utils.FlattenStringSlice(input.MaterializedViewsToInclude),
			"tables_to_exclude":             utils.FlattenStringSlice(input.TablesToExclude),
			"tables_to_include":             utils.FlattenStringSlice(input.TablesToInclude),
		},
	}
}
