/*
SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors
SPDX-License-Identifier: Apache-2.0
*/

package target

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/gardener/gardenctl-v2/internal/util"
	"github.com/gardener/gardenctl-v2/pkg/cmd/base"
)

// NewCmdKubeconfig returns a new version command.
func NewCmdKubeconfig(f util.Factory, o *KubeconfigOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubeconfig",
		Short: "Show the commands to configure kubectl to use the kubeconfig of the current target",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Validate(); err != nil {
				return err
			}

			return runKubeconfigCommand(f, o)
		},
	}

	o.AddFlags(cmd.Flags())

	return cmd
}

func runKubeconfigCommand(f util.Factory, opt *KubeconfigOptions) error {
	m, err := f.Manager()
	if err != nil {
		return err
	}

	currentTarget, err := m.CurrentTarget()
	if err != nil {
		return fmt.Errorf("failed to get current target: %v", err)
	}

	ctx := f.Context()

	kubeconfig, err := m.Kubeconfig(ctx, currentTarget)
	if err != nil {
		return fmt.Errorf("failed to get kubeconfig for current target: %w", err)
	}

	tmpFile, err := ioutil.TempFile("", "kubeconfig*")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(tmpFile.Name(), kubeconfig, 0600); err != nil {
		return fmt.Errorf("failed to write kubeconfig file to %s: %w", tmpFile.Name(), err)
	}

	_, err = fmt.Fprintf(opt.IOStreams.Out, "export KUBECONFIG=%s", tmpFile.Name())

	return err
}

// KubeconfigOptions is a struct to support version command
type KubeconfigOptions struct {
	base.Options
}

// NewKubeconfigOptions returns initialized KubeconfigOptions
func NewKubeconfigOptions(ioStreams util.IOStreams) *KubeconfigOptions {
	return &KubeconfigOptions{
		Options: base.Options{
			IOStreams: ioStreams,
		},
	}
}
