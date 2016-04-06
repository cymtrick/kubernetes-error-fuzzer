/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type set struct {
	*flags.DatacenterFlag

	types.Permission

	role string
}

func init() {
	cli.Register("permissions.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.Principal, "principal", "", "User or group for which the permission is defined")
	f.BoolVar(&cmd.Group, "group", false, "True, if principal refers to a group name; false, for a user name")
	f.BoolVar(&cmd.Propagate, "propagate", true, "Whether or not this permission propagates down the hierarchy to sub-entities")
	f.StringVar(&cmd.role, "role", "Admin", "Permission role name")
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *set) Usage() string {
	return "[PATH]..."
}

func (cmd *set) Description() string {
	return `Set the permissions managed entities.
Example:
govc permissions.set -principal root -role Admin
`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	refs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m := object.NewAuthorizationManager(c)
	rl, err := m.RoleList(ctx)
	if err != nil {
		return err
	}

	role := rl.ByName(cmd.role)
	if role == nil {
		return fmt.Errorf("role '%s' not found", cmd.role)
	}
	cmd.Permission.RoleId = role.RoleId

	perms := []types.Permission{cmd.Permission}

	for _, ref := range refs {
		err = m.SetEntityPermissions(ctx, ref, perms)
		if err != nil {
			return err
		}
	}

	return nil
}
