package kubectl

import (
	"fmt"
	"golang.org/x/xerrors"
	"os"
	"os/exec"
	"strings"
)

type Ensure struct {
	// Bin is the path to the kubectl binary used by this resource. Defaults to "kubectl"
	Bin string

	// Version is the version number or the semver version range for the kubectl version to use
	Version string

	// Namespace is the K8s namespace passed to kubectl as the value for the `-n` flag
	Namespace string

	// Resource is the K8s resource kind like `configmap`, `cm`, `pod`, etc.
	Resource string

	// Name is the K8s resource name like metadata.name of a configmap, secret, pod, etc.
	Name string

	// Kubeconfig is the file path to kubeconfig which is set to the KUBECONFIG environment variable on running kubectl
	Kubeconfig string

	Labels      map[string]interface{}
	Annotations map[string]interface{}
}

func NewEnsure(d ResourceRead) (*Ensure, error) {
	f := Ensure{}

	f.Bin = d.Get(KeyBin).(string)

	f.Kubeconfig = d.Get(KeyKubeconfig).(string)

	f.Namespace = d.Get(KeyNamespace).(string)

	f.Resource = d.Get(KeyResource).(string)

	f.Name = d.Get(KeyName).(string)

	f.Version = d.Get(KeyVersion).(string)

	if labels := d.Get(KeyLabels); labels != nil {
		f.Labels = labels.(map[string]interface{})
	}

	if annotations := d.Get(KeyAnnotations); annotations != nil {
		f.Annotations = annotations.(map[string]interface{})
	}

	return &f, nil
}

func NewCommandWithKubeconfigAndNamespace(e *Ensure, args ...string) (*exec.Cmd, error) {
	kubectlBin, err := prepareBinaries(e)
	if err != nil {
		return nil, err
	}

	flags := []string{
		"-n", e.Namespace,
	}

	flags = append(flags, args...)

	logf("Running kubectl %s on %+v", strings.Join(flags, " "), *e)

	cmd := exec.Command(*kubectlBin, flags...)
	cmd.Env = append(cmd.Env, os.Environ()...)

	if e.Kubeconfig != "" {
		cmd.Env = append(cmd.Env, "KUBECONFIG="+e.Kubeconfig)
	} else {
		return nil, fmt.Errorf("[BUG] NewCommandWithKubeconfigAndNamespace must not be called with empty kubeconfig path. args = %s", strings.Join(args, " "))
	}

	logf("[DEBUG] Generated command: args = %s", strings.Join(cmd.Args, " "))

	return cmd, nil
}

func CreateOrUpdate(e *Ensure, d ResourceReadWrite) error {
	logf("[DEBUG] Ensuring labels and annotations...")

	if len(e.Annotations) > 0 {
		args := []string{
			"annotate", "--overwrite", e.Resource, e.Name,
		}

		for k, v := range e.Annotations {
			args = append(args, fmt.Sprintf("%s=%s", k, v))
		}

		cmd, err := NewCommandWithKubeconfigAndNamespace(e, args...)
		if err != nil {
			return err
		}

		state := NewState()
		_, err = runCommand(cmd, state, false)
		if err != nil {
			return xerrors.Errorf("running %s %s: %w", cmd.Path, strings.Join(cmd.Args, " "), err)
		}
	}

	if len(e.Labels) > 0 {
		args := []string{
			"label", "--overwrite", e.Resource, e.Name,
		}

		for k, v := range e.Labels {
			args = append(args, fmt.Sprintf("%s=%s", k, v))
		}

		cmd, err := NewCommandWithKubeconfigAndNamespace(e, args...)
		if err != nil {
			return xerrors.Errorf("running %s %s: %w", cmd.Path, strings.Join(cmd.Args, " "), err)
		}

		state := NewState()
		_, err = runCommand(cmd, state, false)
		if err != nil {
			return fmt.Errorf("running kubectl-label: %w", err)
		}
	}

	return nil

	return nil
}

func Read(e *Ensure, d ResourceReadWrite) error {
	logf("[DEBUG] Reading kubectl ensure resource...")

	if e.Kubeconfig == "" {
		logf("Skipping kubectl-build due to that kubeconfig is empty, which means that this operation has been called on a kubectl resource that depends on in-existent resource")

		return nil
	}

	if _, err := os.Stat(e.Kubeconfig); err != nil {
		return xerrors.Errorf("checking existence of kubeconfig: %w", err)
	}

	return nil
}
