/*
Copyright 2014 The Kubernetes Authors.

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

package cmd

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kubernetes/pkg/kubectl"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

var patchTypes = map[string]types.PatchType{"json": types.JSONPatchType, "merge": types.MergePatchType, "strategic": types.StrategicMergePatchType}

// PatchOptions is the start of the data required to perform the operation.  As new fields are added, add them here instead of
// referencing the cmd.Flags()
type PatchOptions struct {
	resource.FilenameOptions
	RecordFlags *genericclioptions.RecordFlags

	Local  bool
	DryRun bool

	Recorder genericclioptions.Recorder

	OutputFormat string
}

var (
	patchLong = templates.LongDesc(i18n.T(`
		Update field(s) of a resource using strategic merge patch, a JSON merge patch, or a JSON patch.

		JSON and YAML formats are accepted.

		Please refer to the models in https://htmlpreview.github.io/?https://github.com/kubernetes/kubernetes/blob/HEAD/docs/api-reference/v1/definitions.html to find if a field is mutable.`))

	patchExample = templates.Examples(i18n.T(`
		# Partially update a node using a strategic merge patch. Specify the patch as JSON.
		kubectl patch node k8s-node-1 -p '{"spec":{"unschedulable":true}}'

		# Partially update a node using a strategic merge patch. Specify the patch as YAML.
		kubectl patch node k8s-node-1 -p $'spec:\n unschedulable: true'

		# Partially update a node identified by the type and name specified in "node.json" using strategic merge patch.
		kubectl patch -f node.json -p '{"spec":{"unschedulable":true}}'

		# Update a container's image; spec.containers[*].name is required because it's a merge key.
		kubectl patch pod valid-pod -p '{"spec":{"containers":[{"name":"kubernetes-serve-hostname","image":"new image"}]}}'

		# Update a container's image using a json patch with positional arrays.
		kubectl patch pod valid-pod --type='json' -p='[{"op": "replace", "path": "/spec/containers/0/image", "value":"new image"}]'`))
)

func NewPatchOptions() *PatchOptions {
	return &PatchOptions{
		RecordFlags: genericclioptions.NewRecordFlags(),

		Recorder: genericclioptions.NoopRecorder{},
	}
}

func NewCmdPatch(f cmdutil.Factory, out io.Writer) *cobra.Command {
	o := NewPatchOptions()
	validArgs := cmdutil.ValidArgList(f)

	cmd := &cobra.Command{
		Use: "patch (-f FILENAME | TYPE NAME) -p PATCH",
		DisableFlagsInUseLine: true,
		Short:   i18n.T("Update field(s) of a resource using strategic merge patch"),
		Long:    patchLong,
		Example: patchExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(f, cmd))
			cmdutil.CheckErr(o.RunPatch(f, out, cmd, args))
		},
		ValidArgs:  validArgs,
		ArgAliases: kubectl.ResourceAliases(validArgs),
	}

	o.RecordFlags.AddFlags(cmd)

	cmd.Flags().StringP("patch", "p", "", "The patch to be applied to the resource JSON file.")
	cmd.MarkFlagRequired("patch")
	cmd.Flags().String("type", "strategic", fmt.Sprintf("The type of patch being provided; one of %v", sets.StringKeySet(patchTypes).List()))
	cmdutil.AddPrinterFlags(cmd)
	cmdutil.AddDryRunFlag(cmd)

	usage := "identifying the resource to update"
	cmdutil.AddFilenameOptionFlags(cmd, &o.FilenameOptions, usage)

	cmd.Flags().BoolVar(&o.Local, "local", o.Local, "If true, patch will operate on the content of the file, not the server-side resource.")

	return cmd
}

func (o *PatchOptions) Complete(f cmdutil.Factory, cmd *cobra.Command) error {
	var err error

	o.RecordFlags.Complete(f.Command(cmd, false))
	o.Recorder, err = o.RecordFlags.ToRecorder()
	if err != nil {
		return err
	}

	o.OutputFormat = cmdutil.GetFlagString(cmd, "output")
	o.DryRun = cmdutil.GetFlagBool(cmd, "dry-run")

	return err
}

func (o *PatchOptions) RunPatch(f cmdutil.Factory, out io.Writer, cmd *cobra.Command, args []string) error {
	switch {
	case o.Local && len(args) != 0:
		return fmt.Errorf("cannot specify --local and server resources")
	}

	cmdNamespace, enforceNamespace, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	patchType := types.StrategicMergePatchType
	patchTypeString := strings.ToLower(cmdutil.GetFlagString(cmd, "type"))
	if len(patchTypeString) != 0 {
		ok := false
		patchType, ok = patchTypes[patchTypeString]
		if !ok {
			return cmdutil.UsageErrorf(cmd, "--type must be one of %v, not %q",
				sets.StringKeySet(patchTypes).List(), patchTypeString)
		}
	}

	patch := cmdutil.GetFlagString(cmd, "patch")
	if len(patch) == 0 {
		return cmdutil.UsageErrorf(cmd, "Must specify -p to patch")
	}
	patchBytes, err := yaml.ToJSON([]byte(patch))
	if err != nil {
		return fmt.Errorf("unable to parse %q: %v", patch, err)
	}

	r := f.NewBuilder().
		Unstructured().
		ContinueOnError().
		NamespaceParam(cmdNamespace).DefaultNamespace().
		FilenameParam(enforceNamespace, &o.FilenameOptions).
		ResourceTypeOrNameArgs(false, args...).
		Flatten().
		Do()
	err = r.Err()
	if err != nil {
		return err
	}

	count := 0
	err = r.Visit(func(info *resource.Info, err error) error {
		if err != nil {
			return err
		}
		name, namespace := info.Name, info.Namespace
		mapping := info.ResourceMapping()
		client, err := f.UnstructuredClientForMapping(mapping)
		if err != nil {
			return err
		}

		if !o.Local && !o.DryRun {
			helper := resource.NewHelper(client, mapping)
			patchedObj, err := helper.Patch(namespace, name, patchType, patchBytes)
			if err != nil {
				return err
			}

			didPatch := !reflect.DeepEqual(info.Object, patchedObj)

			// if the recorder makes a change, compute and create another patch
			if mergePatch, err := o.Recorder.MakeRecordMergePatch(patchedObj); err != nil {
				glog.V(4).Infof("error recording current command: %v", err)
			} else if len(mergePatch) > 0 {
				if recordedObj, err := helper.Patch(info.Namespace, info.Name, types.MergePatchType, mergePatch); err != nil {
					glog.V(4).Infof("error recording reason: %v", err)
				} else {
					patchedObj = recordedObj
				}
			}
			count++

			// After computing whether we changed data, refresh the resource info with the resulting object
			if err := info.Refresh(patchedObj, true); err != nil {
				return err
			}

			if len(o.OutputFormat) > 0 && o.OutputFormat != "name" {
				return cmdutil.PrintObject(cmd, info.Object, out)
			}
			cmdutil.PrintSuccess(o.OutputFormat == "name", out, info.Object, false, patchOperation(didPatch))

			// if object was not successfully patched, exit with error code 1
			if !didPatch {
				return cmdutil.ErrExit
			}

			return nil
		}

		count++

		originalObjJS, err := runtime.Encode(unstructured.UnstructuredJSONScheme, info.Object)
		if err != nil {
			return err
		}

		originalPatchedObjJS, err := getPatchedJSON(patchType, originalObjJS, patchBytes, mapping.GroupVersionKind, scheme.Scheme)
		if err != nil {
			return err
		}

		targetObj, err := runtime.Decode(unstructured.UnstructuredJSONScheme, originalPatchedObjJS)
		if err != nil {
			return err
		}

		didPatch := !reflect.DeepEqual(info.Object, targetObj)

		// TODO: if we ever want to go generic, this allows a clean -o yaml without trying to print columns or anything
		// rawExtension := &runtime.Unknown{
		//	Raw: originalPatchedObjJS,
		// }
		if didPatch {
			if err := info.Refresh(targetObj, true); err != nil {
				return err
			}
		}

		if len(o.OutputFormat) > 0 && o.OutputFormat != "name" {
			return cmdutil.PrintObject(cmd, info.Object, out)
		}

		cmdutil.PrintSuccess(o.OutputFormat == "name", out, info.Object, o.DryRun, patchOperation(didPatch))
		return nil
	})
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no objects passed to patch")
	}
	return nil
}

func getPatchedJSON(patchType types.PatchType, originalJS, patchJS []byte, gvk schema.GroupVersionKind, creater runtime.ObjectCreater) ([]byte, error) {
	switch patchType {
	case types.JSONPatchType:
		patchObj, err := jsonpatch.DecodePatch(patchJS)
		if err != nil {
			return nil, err
		}
		return patchObj.Apply(originalJS)

	case types.MergePatchType:
		return jsonpatch.MergePatch(originalJS, patchJS)

	case types.StrategicMergePatchType:
		// get a typed object for this GVK if we need to apply a strategic merge patch
		obj, err := creater.New(gvk)
		if err != nil {
			return nil, fmt.Errorf("cannot apply strategic merge patch for %s locally, try --type merge", gvk.String())
		}
		return strategicpatch.StrategicMergePatch(originalJS, patchJS, obj)

	default:
		// only here as a safety net - go-restful filters content-type
		return nil, fmt.Errorf("unknown Content-Type header for patch: %v", patchType)
	}
}

func patchOperation(didPatch bool) string {
	if didPatch {
		return "patched"
	}
	return "not patched"
}
