/*
Copyright 2018 The Kubernetes Authors.

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

package options

import (
	stdjson "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	auditv1 "k8s.io/apiserver/pkg/apis/audit/v1"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/server"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	utilfeaturetesting "k8s.io/apiserver/pkg/util/feature/testing"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api/v1"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuditValidOptions(t *testing.T) {
	webhookConfig := makeTmpWebhookConfig(t)
	defer os.Remove(webhookConfig)

	policy := makeTmpPolicy(t)
	defer os.Remove(policy)

	defer utilfeaturetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.DynamicAuditing, true)()

	clientConfig := &restclient.Config{}
	informerFactory := informers.NewSharedInformerFactory(fake.NewSimpleClientset(), 0)
	processInfo := &ProcessInfo{"test", "test"}

	testCases := []struct {
		name     string
		options  func() *AuditOptions
		expected string
	}{{
		name:    "default",
		options: NewAuditOptions,
	}, {
		name: "default log",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			o.PolicyFile = policy
			return o
		},
		expected: "log",
	}, {
		name: "default log no policy",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			return o
		},
		expected: "",
	}, {
		name: "default webhook",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = webhookConfig
			o.PolicyFile = policy
			return o
		},
		expected: "buffered<webhook>",
	}, {
		name: "default webhook no policy",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = webhookConfig
			return o
		},
		expected: "",
	}, {
		name: "default union",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			o.WebhookOptions.ConfigFile = webhookConfig
			o.PolicyFile = policy
			return o
		},
		expected: "union[log,buffered<webhook>]",
	}, {
		name: "custom",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.BatchOptions.Mode = ModeBatch
			o.LogOptions.Path = "/audit"
			o.WebhookOptions.BatchOptions.Mode = ModeBlocking
			o.WebhookOptions.ConfigFile = webhookConfig
			o.PolicyFile = policy
			return o
		},
		expected: "union[buffered<log>,webhook]",
	}, {
		name: "default webhook with truncating",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = webhookConfig
			o.WebhookOptions.TruncateOptions.Enabled = true
			o.PolicyFile = policy
			return o
		},
		expected: "truncate<buffered<webhook>>",
	}, {
		name: "dynamic",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			return o
		},
		expected: "dynamic[]",
	}, {
		name: "dynamic with truncating",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			o.WebhookOptions.TruncateOptions.Enabled = true
			return o
		},
		expected: "truncate<dynamic[]>",
	}, {
		name: "dynamic with log",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			o.LogOptions.Path = "/audit"
			o.PolicyFile = policy
			return o
		},
		expected: "union[enforced<log>,dynamic[]]",
	}, {
		name: "dynamic with truncating and webhook",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			o.WebhookOptions.TruncateOptions.Enabled = true
			o.WebhookOptions.ConfigFile = webhookConfig
			o.PolicyFile = policy
			return o
		},
		expected: "truncate<union[enforced<buffered<webhook>>,dynamic[]]>",
	}, {
		name: "dynamic with truncating and webhook and log",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			o.WebhookOptions.TruncateOptions.Enabled = true
			o.WebhookOptions.ConfigFile = webhookConfig
			o.PolicyFile = policy
			o.LogOptions.Path = "/audit"
			return o
		},
		expected: "union[enforced<log>,truncate<union[enforced<buffered<webhook>>,dynamic[]]>]",
	},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options := tc.options()
			require.NotNil(t, options)

			// Verify flags don't change defaults.
			fs := pflag.NewFlagSet("Test", pflag.PanicOnError)
			options.AddFlags(fs)
			require.NoError(t, fs.Parse(nil))
			assert.Equal(t, tc.options(), options, "Flag defaults should match default options.")

			assert.Empty(t, options.Validate(), "Options should be valid.")
			config := &server.Config{}
			require.NoError(t, options.ApplyTo(config, clientConfig, informerFactory, processInfo, nil))
			if tc.expected == "" {
				assert.Nil(t, config.AuditBackend)
			} else {
				assert.Equal(t, tc.expected, fmt.Sprintf("%s", config.AuditBackend))
			}
		})
	}
}

func TestAuditInvalidOptions(t *testing.T) {
	testCases := []struct {
		name    string
		options func() *AuditOptions
	}{{
		name: "invalid log format",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			o.LogOptions.Format = "foo"
			return o
		},
	}, {
		name: "invalid log mode",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			o.LogOptions.BatchOptions.Mode = "foo"
			return o
		},
	}, {
		name: "invalid log buffer size",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.LogOptions.Path = "/audit"
			o.LogOptions.BatchOptions.Mode = "batch"
			o.LogOptions.BatchOptions.BatchConfig.BufferSize = -3
			return o
		},
	}, {
		name: "invalid webhook mode",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = "/audit"
			o.WebhookOptions.BatchOptions.Mode = "foo"
			return o
		},
	}, {
		name: "invalid webhook buffer throttle qps",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = "/audit"
			o.WebhookOptions.BatchOptions.Mode = "batch"
			o.WebhookOptions.BatchOptions.BatchConfig.ThrottleQPS = -1
			return o
		},
	}, {
		name: "invalid webhook truncate max event size",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = "/audit"
			o.WebhookOptions.TruncateOptions.Enabled = true
			o.WebhookOptions.TruncateOptions.TruncateConfig.MaxEventSize = -1
			return o
		},
	}, {
		name: "invalid webhook truncate max batch size",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.WebhookOptions.ConfigFile = "/audit"
			o.WebhookOptions.TruncateOptions.Enabled = true
			o.WebhookOptions.TruncateOptions.TruncateConfig.MaxEventSize = 2
			o.WebhookOptions.TruncateOptions.TruncateConfig.MaxBatchSize = 1
			return o
		},
	}, {
		name: "invalid dynamic flag group",
		options: func() *AuditOptions {
			o := NewAuditOptions()
			o.DynamicOptions.Enabled = true
			return o
		},
	},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			options := tc.options()
			require.NotNil(t, options)
			assert.NotEmpty(t, options.Validate(), "Options should be invalid.")
		})
	}
}

func makeTmpWebhookConfig(t *testing.T) string {
	config := v1.Config{
		Clusters: []v1.NamedCluster{
			{Cluster: v1.Cluster{Server: "localhost", InsecureSkipTLSVerify: true}},
		},
	}
	f, err := ioutil.TempFile("", "k8s_audit_webhook_test_")
	require.NoError(t, err, "creating temp file")
	require.NoError(t, stdjson.NewEncoder(f).Encode(config), "writing webhook kubeconfig")
	require.NoError(t, f.Close())
	return f.Name()
}

func makeTmpPolicy(t *testing.T) string {
	pol := auditv1.Policy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "audit.k8s.io/v1",
		},
		Rules: []auditv1.PolicyRule{
			{
				Level: auditv1.LevelRequestResponse,
			},
		},
	}
	f, err := ioutil.TempFile("", "k8s_audit_policy_test_")
	require.NoError(t, err, "creating temp file")
	require.NoError(t, stdjson.NewEncoder(f).Encode(pol), "writing policy file")
	require.NoError(t, f.Close())
	return f.Name()
}
