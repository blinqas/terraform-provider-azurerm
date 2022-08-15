package containerapps

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/tags"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/managedenvironments"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/containerapps/helpers"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type ContainerAppEnvironmentCertificateResource struct{}

type ContainerAppCertificateModel struct {
	Name                 string                 `tfschema:"name"`
	ManagedEnvironmentId string                 `tfschema:"container_app_environment_id"`
	Tags                 map[string]interface{} `tfschema:"tags"`

	// Write only?
	CertificatePassword string `tfschema:"certificate_password"`
	CertificateBlob     string `tfschema:"certificate_blob"`

	// Read Only
	SubjectName    string `tfschema:"subject_name"`
	Issuer         string `tfschema:"issuer"`
	IssueDate      string `tfschema:"issue_date"`
	ExpirationDate string `tfschema:"expiration_date"`
	Thumbprint     string `tfschema:"thumbprint"`
}

var _ sdk.ResourceWithUpdate = ContainerAppEnvironmentCertificateResource{}

func (r ContainerAppEnvironmentCertificateResource) ModelObject() interface{} {
	return &ContainerAppCertificateModel{}
}

func (r ContainerAppEnvironmentCertificateResource) ResourceType() string {
	return "azurerm_container_app_environment_certificate"
}

func (r ContainerAppEnvironmentCertificateResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return certificates.ValidateCertificateID
}

func (r ContainerAppEnvironmentCertificateResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"name": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: helpers.ValidateCertificateName,
			Description:  "The name of the Container Apps Certificate.",
		},

		"container_app_environment_id": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: certificates.ValidateManagedEnvironmentID,
			Description:  "The Container App Managed Environment ID to configure this Certificate on.",
		},

		"certificate_blob": {
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			Description:  "The Certificate Private Key as a base64 encoded PFX or PEM.",
			ValidateFunc: validation.StringIsBase64,
		},

		"certificate_password": {
			Type:        pluginsdk.TypeString,
			Required:    true,
			ForceNew:    true,
			Sensitive:   true,
			Description: "The password for the Certificate.",
		},

		"tags": commonschema.Tags(),
	}
}

func (r ContainerAppEnvironmentCertificateResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"subject_name": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"issuer": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"issue_date": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"expiration_date": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},

		"thumbprint": {
			Type:     pluginsdk.TypeString,
			Computed: true,
		},
	}
}

func (r ContainerAppEnvironmentCertificateResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ContainerApps.CertificatesClient
			environmentsClient := metadata.Client.ContainerApps.ManagedEnvironmentClient

			var cert ContainerAppCertificateModel

			if err := metadata.Decode(&cert); err != nil {
				return err
			}

			envId, err := managedenvironments.ParseManagedEnvironmentID(cert.ManagedEnvironmentId)
			if err != nil {
				return err
			}

			env, err := environmentsClient.Get(ctx, *envId)
			if err != nil || env.Model == nil {
				return fmt.Errorf("reading Managed Environment %s for %s: %+v", envId.EnvironmentName, cert.Name, err)
			}

			id := certificates.NewCertificateID(metadata.Client.Account.SubscriptionId, envId.ResourceGroupName, envId.EnvironmentName, cert.Name)

			model := certificates.Certificate{
				Location: env.Model.Location,
				Name:     utils.String(id.CertificateName),
				Properties: &certificates.CertificateProperties{
					Password: utils.String(cert.CertificatePassword),
					Value:    utils.String(cert.CertificateBlob),
				},
				Tags: tags.Expand(cert.Tags),
			}

			if resp, err := client.CreateOrUpdate(ctx, id, model); err != nil {
				log.Printf("[STEBUG] resp: %+v", resp)
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)

			return nil
		},
	}
}

func (r ContainerAppEnvironmentCertificateResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ContainerApps.CertificatesClient

			id, err := certificates.ParseCertificateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			existing, err := client.Get(ctx, *id)
			if err != nil {
				if response.WasNotFound(existing.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("reading %s: %+v", *id, err)
			}

			var state ContainerAppCertificateModel

			state.Name = id.CertificateName
			state.ManagedEnvironmentId = certificates.NewManagedEnvironmentID(id.SubscriptionId, id.ResourceGroupName, id.EnvironmentName).ID()

			if model := existing.Model; model != nil {
				state.Tags = tags.Flatten(model.Tags)

				// The Certificate Blob and Password are not retrievable in any way, so grab them back from config if we can. Imports will need `ignore_changes`.
				if certBlob, ok := metadata.ResourceData.GetOk("certificate_blob"); ok {
					state.CertificateBlob = certBlob.(string)
				}

				if certPassword, ok := metadata.ResourceData.GetOk("certificate_password"); ok {
					state.CertificatePassword = certPassword.(string)
				}

				if props := model.Properties; props != nil {
					state.Issuer = utils.NormalizeNilableString(props.Issuer)
					state.IssueDate = utils.NormalizeNilableString(props.IssueDate)
					state.ExpirationDate = utils.NormalizeNilableString(props.ExpirationDate)
					state.Thumbprint = utils.NormalizeNilableString(props.Thumbprint)
				}
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ContainerAppEnvironmentCertificateResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ContainerApps.CertificatesClient

			id, err := certificates.ParseCertificateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if resp, err := client.Delete(ctx, *id); err != nil {
				if !response.WasNotFound(resp.HttpResponse) {
					return fmt.Errorf("deleting %s: %+v", *id, err)
				}
			}

			return nil
		},
	}
}

func (r ContainerAppEnvironmentCertificateResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ContainerApps.CertificatesClient

			var cert ContainerAppCertificateModel

			if err := metadata.Decode(&cert); err != nil {
				return err
			}

			id, err := certificates.ParseCertificateID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if metadata.ResourceData.HasChange("tags") {
				patch := certificates.CertificatePatch{
					Tags: tags.Expand(cert.Tags),
				}

				_, err = client.Update(ctx, *id, patch)
				if err != nil {
					return fmt.Errorf("updating tags for %s: %+v", *id, err)
				}
			}

			return nil
		},
	}
}
