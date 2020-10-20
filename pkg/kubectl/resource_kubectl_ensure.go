package kubectl

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

const KeyAnnotations = "annotations"
const KeyLabels = "labels"
const KeyBin = "bin"
const KeyNamespace = "namespace"
const KeyName = "name"
const KeyVersion = "version"
const KeyResource = "resource"
const KeyKubeconfig = "kubeconfig"

var EnsureSchema = map[string]*schema.Schema{
	KeyAnnotations: {
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     schema.TypeString,
	},
	KeyLabels: {
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     schema.TypeString,
	},
	KeyKubeconfig: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
	},
	KeyNamespace: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
	},
	KeyResource: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
	},
	KeyName: {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
	},
	KeyBin: {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: false,
		Default:  "kubectl",
	},
	KeyVersion: {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: false,
		Default:  "",
	},
}

func resourceKubectlEnsure() *schema.Resource {
	return &schema.Resource{
		Create: resourceReleaseSetCreate,
		Read:   resourceReleaseSetRead,
		Update: resourceReleaseSetUpdate,
		Delete: func(_ *schema.ResourceData, _ interface{}) error {
			return nil
		},
		Schema: EnsureSchema,
	}
}

//helpers to unwravel the recursive bits by adding a base condition
func resourceReleaseSetCreate(d *schema.ResourceData, _ interface{}) error {
	fs, err := NewEnsure(d)
	if err != nil {
		return err
	}

	if err := CreateOrUpdate(fs, d); err != nil {
		return fmt.Errorf("creating release set: %w", err)
	}

	d.MarkNewResource()

	d.SetId(newId())

	return nil
}

func newId() string {
	//create random uuid for the id
	id := xid.New().String()

	return id
}

func resourceReleaseSetRead(d *schema.ResourceData, _ interface{}) error {
	fs, err := NewEnsure(d)
	if err != nil {
		return err
	}

	if err := Read(fs, d); err != nil {
		return fmt.Errorf("reading release set: %w", err)
	}

	return nil
}

func resourceReleaseSetUpdate(d *schema.ResourceData, meta interface{}) error {
	fs, err := NewEnsure(d)
	if err != nil {
		return err
	}

	if err := CreateOrUpdate(fs, d); err != nil {
		return xerrors.Errorf("updating resource: %w", err)
	}

	return nil
}
