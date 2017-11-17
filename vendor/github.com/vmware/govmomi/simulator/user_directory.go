/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"strings"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var DefaultUserGroup = []*types.UserSearchResult{
	{FullName: "root", Group: true, Principal: "root"},
	{FullName: "root", Group: false, Principal: "root"},
	{FullName: "administrator", Group: false, Principal: "admin"},
}

type UserDirectory struct {
	mo.UserDirectory

	userGroup []*types.UserSearchResult
}

func NewUserDirectory(ref types.ManagedObjectReference) object.Reference {
	u := &UserDirectory{}

	u.Self = ref
	u.userGroup = DefaultUserGroup

	return u
}

func (u *UserDirectory) RetrieveUserGroups(req *types.RetrieveUserGroups) soap.HasFault {
	compare := compareFunc(req.SearchStr, req.ExactMatch)

	var res []types.BaseUserSearchResult
	for _, ug := range u.userGroup {
		if req.FindUsers && !ug.Group || req.FindGroups && ug.Group {
			if compare(ug.Principal) {
				res = append(res, ug)
			}
		}
	}

	body := &methods.RetrieveUserGroupsBody{
		Res: &types.RetrieveUserGroupsResponse{
			Returnval: res,
		},
	}

	return body
}

func compareFunc(compared string, exactly bool) func(string) bool {
	return func(s string) bool {
		if exactly {
			return s == compared
		}
		return strings.Contains(strings.ToLower(s), strings.ToLower(compared))
	}
}
