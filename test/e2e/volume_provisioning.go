/*
Copyright 2016 The Kubernetes Authors.

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

package e2e

import (
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/kubernetes/pkg/api/v1"
	rbacv1beta1 "k8s.io/kubernetes/pkg/apis/rbac/v1beta1"
	storage "k8s.io/kubernetes/pkg/apis/storage/v1"
	storageutil "k8s.io/kubernetes/pkg/apis/storage/v1/util"
	storagebeta "k8s.io/kubernetes/pkg/apis/storage/v1beta1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	"k8s.io/kubernetes/test/e2e/framework"
)

type storageClassTest struct {
	name           string
	cloudProviders []string
	provisioner    string
	parameters     map[string]string
	claimSize      string
	expectedSize   string
	pvCheck        func(volume *v1.PersistentVolume) error
}

const (
	// Plugin name of the external provisioner
	externalPluginName = "example.com/nfs"
)

func testDynamicProvisioning(t storageClassTest, client clientset.Interface, claim *v1.PersistentVolumeClaim, class *storage.StorageClass) {
	if class != nil {
		By("creating a StorageClass " + class.Name)
		class, err := client.Storage().StorageClasses().Create(class)
		defer func() {
			framework.Logf("deleting storage class %s", class.Name)
			client.Storage().StorageClasses().Delete(class.Name, nil)
		}()
		Expect(err).NotTo(HaveOccurred())
	}

	By("creating a claim")
	claim, err := client.Core().PersistentVolumeClaims(claim.Namespace).Create(claim)
	defer func() {
		framework.Logf("deleting claim %s/%s", claim.Namespace, claim.Name)
		client.Core().PersistentVolumeClaims(claim.Namespace).Delete(claim.Name, nil)
	}()
	Expect(err).NotTo(HaveOccurred())
	err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, client, claim.Namespace, claim.Name, framework.Poll, framework.ClaimProvisionTimeout)
	Expect(err).NotTo(HaveOccurred())

	By("checking the claim")
	// Get new copy of the claim
	claim, err = client.Core().PersistentVolumeClaims(claim.Namespace).Get(claim.Name, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred())

	// Get the bound PV
	pv, err := client.Core().PersistentVolumes().Get(claim.Spec.VolumeName, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred())

	// Check sizes
	expectedCapacity := resource.MustParse(t.expectedSize)
	pvCapacity := pv.Spec.Capacity[v1.ResourceName(v1.ResourceStorage)]
	Expect(pvCapacity.Value()).To(Equal(expectedCapacity.Value()), "pvCapacity is not equal to expectedCapacity")

	requestedCapacity := resource.MustParse(t.claimSize)
	claimCapacity := claim.Spec.Resources.Requests[v1.ResourceName(v1.ResourceStorage)]
	Expect(claimCapacity.Value()).To(Equal(requestedCapacity.Value()), "claimCapacity is not equal to requestedCapacity")

	// Check PV properties
	By("checking the PV")
	Expect(pv.Spec.PersistentVolumeReclaimPolicy).To(Equal(v1.PersistentVolumeReclaimDelete))
	expectedAccessModes := []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}
	Expect(pv.Spec.AccessModes).To(Equal(expectedAccessModes))
	Expect(pv.Spec.ClaimRef.Name).To(Equal(claim.ObjectMeta.Name))
	Expect(pv.Spec.ClaimRef.Namespace).To(Equal(claim.ObjectMeta.Namespace))

	// Run the checker
	if t.pvCheck != nil {
		err = t.pvCheck(pv)
		Expect(err).NotTo(HaveOccurred())
	}

	// We start two pods:
	// - The first writes 'hello word' to the /mnt/test (= the volume).
	// - The second one runs grep 'hello world' on /mnt/test.
	// If both succeed, Kubernetes actually allocated something that is
	// persistent across pods.
	By("checking the created volume is writable")
	runInPodWithVolume(client, claim.Namespace, claim.Name, "echo 'hello world' > /mnt/test/data")

	By("checking the created volume is readable and retains data")
	runInPodWithVolume(client, claim.Namespace, claim.Name, "grep 'hello world' /mnt/test/data")

	By("deleting the claim")
	framework.ExpectNoError(client.Core().PersistentVolumeClaims(claim.Namespace).Delete(claim.Name, nil))

	// Wait for the PV to get deleted. Technically, the first few delete
	// attempts may fail, as the volume is still attached to a node because
	// kubelet is slowly cleaning up the previous pod, however it should succeed
	// in a couple of minutes. Wait 20 minutes to recover from random cloud
	// hiccups.
	framework.ExpectNoError(framework.WaitForPersistentVolumeDeleted(client, pv.Name, 5*time.Second, 20*time.Minute))
}

// checkAWSEBS checks properties of an AWS EBS. Test framework does not
// instantiate full AWS provider, therefore we need use ec2 API directly.
func checkAWSEBS(volume *v1.PersistentVolume, volumeType string, encrypted bool) error {
	diskName := volume.Spec.AWSElasticBlockStore.VolumeID

	var client *ec2.EC2

	tokens := strings.Split(diskName, "/")
	volumeID := tokens[len(tokens)-1]

	zone := framework.TestContext.CloudConfig.Zone
	if len(zone) > 0 {
		region := zone[:len(zone)-1]
		cfg := aws.Config{Region: &region}
		framework.Logf("using region %s", region)
		client = ec2.New(session.New(), &cfg)
	} else {
		framework.Logf("no region configured")
		client = ec2.New(session.New())
	}

	request := &ec2.DescribeVolumesInput{
		VolumeIds: []*string{&volumeID},
	}
	info, err := client.DescribeVolumes(request)
	if err != nil {
		return fmt.Errorf("error querying ec2 for volume %q: %v", volumeID, err)
	}
	if len(info.Volumes) == 0 {
		return fmt.Errorf("no volumes found for volume %q", volumeID)
	}
	if len(info.Volumes) > 1 {
		return fmt.Errorf("multiple volumes found for volume %q", volumeID)
	}

	awsVolume := info.Volumes[0]
	if awsVolume.VolumeType == nil {
		return fmt.Errorf("expected volume type %q, got nil", volumeType)
	}
	if *awsVolume.VolumeType != volumeType {
		return fmt.Errorf("expected volume type %q, got %q", volumeType, *awsVolume.VolumeType)
	}
	if encrypted && awsVolume.Encrypted == nil {
		return fmt.Errorf("expected encrypted volume, got no encryption")
	}
	if encrypted && !*awsVolume.Encrypted {
		return fmt.Errorf("expected encrypted volume, got %v", *awsVolume.Encrypted)
	}
	return nil
}

func checkGCEPD(volume *v1.PersistentVolume, volumeType string) error {
	cloud, err := framework.GetGCECloud()
	if err != nil {
		return err
	}
	diskName := volume.Spec.GCEPersistentDisk.PDName
	disk, err := cloud.GetDiskByNameUnknownZone(diskName)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(disk.Type, volumeType) {
		return fmt.Errorf("unexpected disk type %q, expected suffix %q", disk.Type, volumeType)
	}
	return nil
}

var _ = framework.KubeDescribe("Dynamic Provisioning", func() {
	f := framework.NewDefaultFramework("volume-provisioning")

	// filled in BeforeEach
	var c clientset.Interface
	var ns string

	BeforeEach(func() {
		c = f.ClientSet
		ns = f.Namespace.Name
	})

	framework.KubeDescribe("DynamicProvisioner", func() {
		// This test checks that dynamic provisioning can provision a volume
		// that can be used to persist data among pods.

		tests := []storageClassTest{
			{
				"should provision SSD PD on GCE/GKE",
				[]string{"gce", "gke"},
				"kubernetes.io/gce-pd",
				map[string]string{
					"type": "pd-ssd",
					// Check that GCE can parse "zone" parameter, however
					// we can't create PDs in different than default zone
					// as we don't know if we're running with Multizone=true
					"zone": framework.TestContext.CloudConfig.Zone,
				},
				"1.5Gi",
				"2Gi",
				func(volume *v1.PersistentVolume) error {
					return checkGCEPD(volume, "pd-ssd")
				},
			},
			{
				"should provision HDD PD on GCE/GKE",
				[]string{"gce", "gke"},
				"kubernetes.io/gce-pd",
				map[string]string{
					"type": "pd-standard",
				},
				"1.5Gi",
				"2Gi",
				func(volume *v1.PersistentVolume) error {
					return checkGCEPD(volume, "pd-standard")
				},
			},
			// AWS
			{
				"should provision gp2 EBS on AWS",
				[]string{"aws"},
				"kubernetes.io/aws-ebs",
				map[string]string{
					"type": "gp2",
					// Check that AWS can parse "zone" parameter, however
					// we can't create PDs in different than default zone
					// as we don't know zone names
					"zone": framework.TestContext.CloudConfig.Zone,
				},
				"1.5Gi",
				"2Gi",
				func(volume *v1.PersistentVolume) error {
					return checkAWSEBS(volume, "gp2", false)
				},
			},
			{
				"should provision io1 EBS on AWS",
				[]string{"aws"},
				"kubernetes.io/aws-ebs",
				map[string]string{
					"type":      "io1",
					"iopsPerGB": "50",
				},
				"3.5Gi",
				"4Gi", // 4 GiB is minimum for io1
				func(volume *v1.PersistentVolume) error {
					return checkAWSEBS(volume, "io1", false)
				},
			},
			{
				"should provision sc1 EBS on AWS",
				[]string{"aws"},
				"kubernetes.io/aws-ebs",
				map[string]string{
					"type": "sc1",
				},
				"500Gi", // minimum for sc1
				"500Gi",
				func(volume *v1.PersistentVolume) error {
					return checkAWSEBS(volume, "sc1", false)
				},
			},
			{
				"should provision st1 EBS on AWS",
				[]string{"aws"},
				"kubernetes.io/aws-ebs",
				map[string]string{
					"type": "st1",
				},
				"500Gi", // minimum for st1
				"500Gi",
				func(volume *v1.PersistentVolume) error {
					return checkAWSEBS(volume, "st1", false)
				},
			},
			{
				"should provision encrypted EBS on AWS",
				[]string{"aws"},
				"kubernetes.io/aws-ebs",
				map[string]string{
					"encrypted": "true",
				},
				"1Gi",
				"1Gi",
				func(volume *v1.PersistentVolume) error {
					return checkAWSEBS(volume, "gp2", true)
				},
			},
			// OpenStack generic tests (works on all OpenStack deployments)
			{
				"should provision generic Cinder volume on OpenStack",
				[]string{"openstack"},
				"kubernetes.io/cinder",
				map[string]string{},
				"1.5Gi",
				"2Gi",
				nil, // there is currently nothing to check on OpenStack
			},
			{
				"should provision Cinder volume with empty volume type and zone on OpenStack",
				[]string{"openstack"},
				"kubernetes.io/cinder",
				map[string]string{
					"type":         "",
					"availability": "",
				},
				"1.5Gi",
				"2Gi",
				nil, // there is currently nothing to check on OpenStack
			},
			// vSphere generic test
			{
				"should provision generic vSphere",
				[]string{"vsphere"},
				"kubernetes.io/vsphere-volume",
				map[string]string{},
				"1.5Gi",
				"1.5Gi",
				nil,
			},
		}

		var betaTest *storageClassTest
		for i, t := range tests {
			// Beware of clojure, use local variables instead of those from
			// outer scope
			test := t
			suffix := fmt.Sprintf("%d", i)
			It(test.name+" [Slow] [Volume]", func() {
				if len(t.cloudProviders) > 0 {
					framework.SkipUnlessProviderIs(test.cloudProviders...)
				}

				// Remember the last supported test for subsequent test of beta API
				betaTest = &test

				class := newStorageClass(test, ns, suffix)
				claim := newClaim(test, ns, suffix)
				claim.Spec.StorageClassName = &class.Name
				testDynamicProvisioning(test, c, claim, class)
			})
		}

		// Run the last test with storage.k8s.io/v1beta1 and beta annotation on pvc
		if betaTest != nil {
			It("Beta "+betaTest.name+" [Slow] [Volume]", func() {
				class := newBetaStorageClass(*betaTest, "beta")
				// we need to create the class manually, testDynamicProvisioning does not accept beta class
				class, err := c.StorageV1beta1().StorageClasses().Create(class)
				defer deleteStorageClass(c, class.Name)
				Expect(err).NotTo(HaveOccurred())

				claim := newClaim(*betaTest, ns, "beta")
				claim.Annotations = map[string]string{
					v1.BetaStorageClassAnnotation: class.Name,
				}
				testDynamicProvisioning(*betaTest, c, claim, nil)
			})
		}

		// NOTE: Slow!  The test will wait up to 5 minutes (framework.ClaimProvisionTimeout) when there is
		// no regression.
		It("should not provision a volume in an unmanaged GCE zone. [Slow] [Volume]", func() {
			framework.SkipUnlessProviderIs("gce", "gke")
			var suffix string = "unmananged"

			By("Discovering an unmanaged zone")
			allZones := sets.NewString()     // all zones in the project
			managedZones := sets.NewString() // subset of allZones

			gceCloud, err := framework.GetGCECloud()
			Expect(err).NotTo(HaveOccurred())

			// Get all k8s managed zones
			managedZones, err = gceCloud.GetAllZones()
			Expect(err).NotTo(HaveOccurred())

			// Get a list of all zones in the project
			zones, err := gceCloud.GetComputeService().Zones.List(framework.TestContext.CloudConfig.ProjectID).Do()
			for _, z := range zones.Items {
				allZones.Insert(z.Name)
			}

			// Get the subset of zones not managed by k8s
			var unmanagedZone string
			var popped bool
			unmanagedZones := allZones.Difference(managedZones)
			// And select one of them at random.
			if unmanagedZone, popped = unmanagedZones.PopAny(); !popped {
				framework.Skipf("No unmanaged zones found.")
			}

			By("Creating a StorageClass for the unmanaged zone")
			test := storageClassTest{
				name:        "unmanaged_zone",
				provisioner: "kubernetes.io/gce-pd",
				parameters:  map[string]string{"zone": unmanagedZone},
				claimSize:   "1Gi",
			}
			sc := newStorageClass(test, ns, suffix)
			sc, err = c.StorageV1().StorageClasses().Create(sc)
			Expect(err).NotTo(HaveOccurred())
			defer deleteStorageClass(c, sc.Name)

			By("Creating a claim and expecting it to timeout")
			pvc := newClaim(test, ns, suffix)
			pvc.Spec.StorageClassName = &sc.Name
			pvc, err = c.Core().PersistentVolumeClaims(ns).Create(pvc)
			Expect(err).NotTo(HaveOccurred())
			defer framework.DeletePersistentVolumeClaim(c, pvc.Name, ns)

			// The claim should timeout phase:Pending
			err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, c, ns, pvc.Name, 2*time.Second, framework.ClaimProvisionTimeout)
			Expect(err).To(HaveOccurred())
			framework.Logf(err.Error())
		})

		It("should test that deleting a claim before the volume is provisioned deletes the volume. [Volume]", func() {
			// This case tests for the regressions of a bug fixed by PR #21268
			// REGRESSION: Deleting the PVC before the PV is provisioned can result in the PV
			// not being deleted.
			// NOTE:  Polls until no PVs are detected, times out at 5 minutes.

			framework.SkipUnlessProviderIs("openstack", "gce", "aws", "gke", "vsphere")

			const raceAttempts int = 100
			var residualPVs []*v1.PersistentVolume
			By(fmt.Sprintf("Creating and deleting PersistentVolumeClaims %d times", raceAttempts))
			test := storageClassTest{
				name:        "deletion race",
				provisioner: "", // Use a native one based on current cloud provider
				claimSize:   "1Gi",
			}

			class := newStorageClass(test, ns, "race")
			class, err := c.Storage().StorageClasses().Create(class)
			Expect(err).NotTo(HaveOccurred())
			defer deleteStorageClass(c, class.Name)

			// To increase chance of detection, attempt multiple iterations
			for i := 0; i < raceAttempts; i++ {
				suffix := fmt.Sprintf("race-%d", i)
				claim := newClaim(test, ns, suffix)
				claim.Spec.StorageClassName = &class.Name
				tmpClaim := framework.CreatePVC(c, ns, claim)
				framework.DeletePersistentVolumeClaim(c, tmpClaim.Name, ns)
			}

			By(fmt.Sprintf("Checking for residual PersistentVolumes associated with StorageClass %s", class.Name))
			residualPVs, err = waitForProvisionedVolumesDeleted(c, class.Name)
			if err != nil {
				// Cleanup the test resources before breaking
				deleteProvisionedVolumesAndDisks(c, residualPVs)
				Expect(err).NotTo(HaveOccurred())
			}

			// Report indicators of regression
			if len(residualPVs) > 0 {
				framework.Logf("Remaining PersistentVolumes:")
				for i, pv := range residualPVs {
					framework.Logf("\t%d) %s", i+1, pv.Name)
				}
				framework.Failf("Expected 0 PersistentVolumes remaining. Found %d", len(residualPVs))
			}
			framework.Logf("0 PersistentVolumes remain.")
		})
	})

	framework.KubeDescribe("DynamicProvisioner Alpha", func() {
		It("should create and delete alpha persistent volumes [Slow] [Volume]", func() {
			framework.SkipUnlessProviderIs("openstack", "gce", "aws", "gke", "vsphere")

			By("creating a claim with an alpha dynamic provisioning annotation")
			test := storageClassTest{
				name:         "alpha test",
				claimSize:    "2Gi",
				expectedSize: "2Gi",
			}

			claim := newClaim(test, ns, "alpha")
			claim.Annotations = map[string]string{
				v1.AlphaStorageClassAnnotation: "true",
			}
			testDynamicProvisioning(test, c, claim, nil)
		})
	})

	framework.KubeDescribe("DynamicProvisioner External", func() {
		It("should let an external dynamic provisioner create and delete persistent volumes [Slow] [Volume]", func() {
			// external dynamic provisioner pods need additional permissions provided by the
			// persistent-volume-provisioner role
			framework.BindClusterRole(c.Rbac(), "system:persistent-volume-provisioner", ns,
				rbacv1beta1.Subject{Kind: rbacv1beta1.ServiceAccountKind, Namespace: ns, Name: "default"})

			err := framework.WaitForAuthorizationUpdate(c.AuthorizationV1beta1(),
				serviceaccount.MakeUsername(ns, "default"),
				"", "get", schema.GroupResource{Group: "storage.k8s.io", Resource: "storageclasses"}, true)
			framework.ExpectNoError(err, "Failed to update authorization: %v", err)

			By("creating an external dynamic provisioner pod")
			pod := startExternalProvisioner(c, ns)
			defer framework.DeletePodOrFail(c, ns, pod.Name)

			By("creating a StorageClass")
			test := storageClassTest{
				name:         "external provisioner test",
				provisioner:  externalPluginName,
				claimSize:    "1500Mi",
				expectedSize: "1500Mi",
			}
			class := newStorageClass(test, ns, "external")
			className := class.Name
			claim := newClaim(test, ns, "external")
			// the external provisioner understands Beta only right now, see
			// https://github.com/kubernetes-incubator/external-storage/issues/37
			// claim.Spec.StorageClassName = &className
			claim.Annotations = map[string]string{
				v1.BetaStorageClassAnnotation: className,
			}

			By("creating a claim with a external provisioning annotation")
			testDynamicProvisioning(test, c, claim, class)
		})
	})

	framework.KubeDescribe("DynamicProvisioner Default", func() {
		It("should create and delete default persistent volumes [Slow] [Volume]", func() {
			framework.SkipUnlessProviderIs("openstack", "gce", "aws", "gke", "vsphere", "azure")

			By("creating a claim with no annotation")
			test := storageClassTest{
				name:         "default",
				claimSize:    "2Gi",
				expectedSize: "2Gi",
			}
			claim := newClaim(test, ns, "default")
			defer func() {
				framework.DeletePersistentVolumeClaim(c, claim.Name, ns)
			}()
			testDynamicProvisioning(test, c, claim, nil)
		})

		// Modifying the default storage class can be disruptive to other tests that depend on it
		It("should be disabled by changing the default annotation[Slow] [Serial] [Disruptive] [Volume]", func() {
			framework.SkipUnlessProviderIs("openstack", "gce", "aws", "gke", "vsphere")
			scName := getDefaultStorageClassName(c)
			test := storageClassTest{
				name:      "default",
				claimSize: "2Gi",
			}

			By("setting the is-default StorageClass annotation to false")
			verifyDefaultStorageClass(c, scName, true)
			defer updateDefaultStorageClass(c, scName, "true")
			updateDefaultStorageClass(c, scName, "false")

			By("creating a claim with default storageclass and expecting it to timeout")
			claim := newClaim(test, ns, "default")
			defer func() {
				framework.DeletePersistentVolumeClaim(c, claim.Name, ns)
			}()
			claim, err := c.Core().PersistentVolumeClaims(ns).Create(claim)
			Expect(err).NotTo(HaveOccurred())

			// The claim should timeout phase:Pending
			err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, c, ns, claim.Name, 2*time.Second, framework.ClaimProvisionTimeout)
			Expect(err).To(HaveOccurred())
			framework.Logf(err.Error())
			claim, err = c.Core().PersistentVolumeClaims(ns).Get(claim.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(claim.Status.Phase).To(Equal(v1.ClaimPending))
		})

		// Modifying the default storage class can be disruptive to other tests that depend on it
		It("should be disabled by removing the default annotation[Slow] [Serial] [Disruptive] [Volume]", func() {
			framework.SkipUnlessProviderIs("openstack", "gce", "aws", "gke", "vsphere")
			scName := getDefaultStorageClassName(c)
			test := storageClassTest{
				name:      "default",
				claimSize: "2Gi",
			}

			By("removing the is-default StorageClass annotation")
			verifyDefaultStorageClass(c, scName, true)
			defer updateDefaultStorageClass(c, scName, "true")
			updateDefaultStorageClass(c, scName, "")

			By("creating a claim with default storageclass and expecting it to timeout")
			claim := newClaim(test, ns, "default")
			defer func() {
				framework.DeletePersistentVolumeClaim(c, claim.Name, ns)
			}()
			claim, err := c.Core().PersistentVolumeClaims(ns).Create(claim)
			Expect(err).NotTo(HaveOccurred())

			// The claim should timeout phase:Pending
			err = framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, c, ns, claim.Name, 2*time.Second, framework.ClaimProvisionTimeout)
			Expect(err).To(HaveOccurred())
			framework.Logf(err.Error())
			claim, err = c.Core().PersistentVolumeClaims(ns).Get(claim.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(claim.Status.Phase).To(Equal(v1.ClaimPending))
		})
	})
})

func getDefaultStorageClassName(c clientset.Interface) string {
	list, err := c.StorageV1().StorageClasses().List(metav1.ListOptions{})
	if err != nil {
		framework.Failf("Error listing storage classes: %v", err)
	}
	var scName string
	for _, sc := range list.Items {
		if storageutil.IsDefaultAnnotation(sc.ObjectMeta) {
			if len(scName) != 0 {
				framework.Failf("Multiple default storage classes found: %q and %q", scName, sc.Name)
			}
			scName = sc.Name
		}
	}
	if len(scName) == 0 {
		framework.Failf("No default storage class found")
	}
	framework.Logf("Default storage class: %q", scName)
	return scName
}

func verifyDefaultStorageClass(c clientset.Interface, scName string, expectedDefault bool) {
	sc, err := c.StorageV1().StorageClasses().Get(scName, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred())
	Expect(storageutil.IsDefaultAnnotation(sc.ObjectMeta)).To(Equal(expectedDefault))
}

func updateDefaultStorageClass(c clientset.Interface, scName string, defaultStr string) {
	sc, err := c.StorageV1().StorageClasses().Get(scName, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred())

	if defaultStr == "" {
		delete(sc.Annotations, storageutil.BetaIsDefaultStorageClassAnnotation)
		delete(sc.Annotations, storageutil.IsDefaultStorageClassAnnotation)
	} else {
		if sc.Annotations == nil {
			sc.Annotations = make(map[string]string)
		}
		sc.Annotations[storageutil.BetaIsDefaultStorageClassAnnotation] = defaultStr
		sc.Annotations[storageutil.IsDefaultStorageClassAnnotation] = defaultStr
	}

	sc, err = c.StorageV1().StorageClasses().Update(sc)
	Expect(err).NotTo(HaveOccurred())

	expectedDefault := false
	if defaultStr == "true" {
		expectedDefault = true
	}
	verifyDefaultStorageClass(c, scName, expectedDefault)
}

func newClaim(t storageClassTest, ns, suffix string) *v1.PersistentVolumeClaim {
	claim := v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "pvc-",
			Namespace:    ns,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.ReadWriteOnce,
			},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceName(v1.ResourceStorage): resource.MustParse(t.claimSize),
				},
			},
		},
	}

	return &claim
}

// runInPodWithVolume runs a command in a pod with given claim mounted to /mnt directory.
func runInPodWithVolume(c clientset.Interface, ns, claimName, command string) {
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "pvc-volume-tester-",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "volume-tester",
					Image:   "gcr.io/google_containers/busybox:1.24",
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "my-volume",
							MountPath: "/mnt/test",
						},
					},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "my-volume",
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: claimName,
							ReadOnly:  false,
						},
					},
				},
			},
		},
	}
	pod, err := c.Core().Pods(ns).Create(pod)
	defer func() {
		framework.DeletePodOrFail(c, ns, pod.Name)
	}()
	framework.ExpectNoError(err, "Failed to create pod: %v", err)
	framework.ExpectNoError(framework.WaitForPodSuccessInNamespaceSlow(c, pod.Name, pod.Namespace))
}

func getDefaultPluginName() string {
	switch {
	case framework.ProviderIs("gke"), framework.ProviderIs("gce"):
		return "kubernetes.io/gce-pd"
	case framework.ProviderIs("aws"):
		return "kubernetes.io/aws-ebs"
	case framework.ProviderIs("openstack"):
		return "kubernetes.io/cinder"
	case framework.ProviderIs("vsphere"):
		return "kubernetes.io/vsphere-volume"
	}
	return ""
}

func newStorageClass(t storageClassTest, ns string, suffix string) *storage.StorageClass {
	pluginName := t.provisioner
	if pluginName == "" {
		pluginName = getDefaultPluginName()
	}
	if suffix == "" {
		suffix = "sc"
	}
	return &storage.StorageClass{
		TypeMeta: metav1.TypeMeta{
			Kind: "StorageClass",
		},
		ObjectMeta: metav1.ObjectMeta{
			// Name must be unique, so let's base it on namespace name
			Name: ns + "-" + suffix,
		},
		Provisioner: pluginName,
		Parameters:  t.parameters,
	}
}

// TODO: remove when storage.k8s.io/v1beta1 and beta storage class annotations
// are removed.
func newBetaStorageClass(t storageClassTest, suffix string) *storagebeta.StorageClass {
	pluginName := t.provisioner

	if pluginName == "" {
		pluginName = getDefaultPluginName()
	}
	if suffix == "" {
		suffix = "default"
	}

	return &storagebeta.StorageClass{
		TypeMeta: metav1.TypeMeta{
			Kind: "StorageClass",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: suffix + "-",
		},
		Provisioner: pluginName,
		Parameters:  t.parameters,
	}
}

func startExternalProvisioner(c clientset.Interface, ns string) *v1.Pod {
	podClient := c.Core().Pods(ns)

	provisionerPod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "external-provisioner-",
		},

		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nfs-provisioner",
					Image: "quay.io/kubernetes_incubator/nfs-provisioner:v1.0.3",
					SecurityContext: &v1.SecurityContext{
						Capabilities: &v1.Capabilities{
							Add: []v1.Capability{"DAC_READ_SEARCH"},
						},
					},
					Args: []string{
						"-provisioner=" + externalPluginName,
						"-grace-period=0",
					},
					Ports: []v1.ContainerPort{
						{Name: "nfs", ContainerPort: 2049},
						{Name: "mountd", ContainerPort: 20048},
						{Name: "rpcbind", ContainerPort: 111},
						{Name: "rpcbind-udp", ContainerPort: 111, Protocol: v1.ProtocolUDP},
					},
					Env: []v1.EnvVar{
						{
							Name: "POD_IP",
							ValueFrom: &v1.EnvVarSource{
								FieldRef: &v1.ObjectFieldSelector{
									FieldPath: "status.podIP",
								},
							},
						},
					},
					ImagePullPolicy: v1.PullIfNotPresent,
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "export-volume",
							MountPath: "/export",
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: "export-volume",
					VolumeSource: v1.VolumeSource{
						EmptyDir: &v1.EmptyDirVolumeSource{},
					},
				},
			},
		},
	}
	provisionerPod, err := podClient.Create(provisionerPod)
	framework.ExpectNoError(err, "Failed to create %s pod: %v", provisionerPod.Name, err)

	framework.ExpectNoError(framework.WaitForPodRunningInNamespace(c, provisionerPod))

	By("locating the provisioner pod")
	pod, err := podClient.Get(provisionerPod.Name, metav1.GetOptions{})
	framework.ExpectNoError(err, "Cannot locate the provisioner pod %v: %v", provisionerPod.Name, err)

	return pod
}

// waitForProvisionedVolumesDelete is a polling wrapper to scan all PersistentVolumes for any associated to the test's
// StorageClass.  Returns either an error and nil values or the remaining PVs and their count.
func waitForProvisionedVolumesDeleted(c clientset.Interface, scName string) ([]*v1.PersistentVolume, error) {
	var remainingPVs []*v1.PersistentVolume

	err := wait.Poll(10*time.Second, 300*time.Second, func() (bool, error) {
		remainingPVs = []*v1.PersistentVolume{}

		allPVs, err := c.Core().PersistentVolumes().List(metav1.ListOptions{})
		if err != nil {
			return true, err
		}
		for _, pv := range allPVs.Items {
			if v1.GetPersistentVolumeClass(&pv) == scName {
				remainingPVs = append(remainingPVs, &pv)
			}
		}
		if len(remainingPVs) > 0 {
			return false, nil // Poll until no PVs remain
		} else {
			return true, nil // No PVs remain
		}
	})
	return remainingPVs, err
}

// deleteStorageClass deletes the passed in StorageClass and catches errors other than "Not Found"
func deleteStorageClass(c clientset.Interface, className string) {
	err := c.Storage().StorageClasses().Delete(className, nil)
	if err != nil && !apierrs.IsNotFound(err) {
		Expect(err).NotTo(HaveOccurred())
	}
}

// deleteProvisionedVolumes [gce||gke only]  iteratively deletes persistent volumes and attached GCE PDs.
func deleteProvisionedVolumesAndDisks(c clientset.Interface, pvs []*v1.PersistentVolume) {
	for _, pv := range pvs {
		framework.DeletePDWithRetry(pv.Spec.PersistentVolumeSource.GCEPersistentDisk.PDName)
		framework.DeletePersistentVolume(c, pv.Name)
	}
}
