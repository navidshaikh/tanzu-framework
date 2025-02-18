// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package catalog implements catalog management functions
package catalog

import (
	cliapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/cli/v1alpha1"
)

// Catalog is the interface that maintains an index of the installed plugins as well as the active plugins.
type Catalog interface {
	// Upsert inserts/updates the given plugin.
	Upsert(plugin cliapi.PluginDescriptor)

	// Get looks up the descriptor of a plugin given its name.
	Get(pluginName string) (cliapi.PluginDescriptor, bool)

	// List returns the list of active plugins.
	// Active plugin means the plugin that are available to the user
	// based on the current logged-in server.
	List() []cliapi.PluginDescriptor

	// Delete deletes the given plugin from the catalog, but it does not delete the installation.
	Delete(plugin string)
}
