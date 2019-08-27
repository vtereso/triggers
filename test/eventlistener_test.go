// +build e2e

/*
Copyright 2019 The Tekton Authors

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

package test

import (
	"fmt"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	reconciler "github.com/tektoncd/triggers/pkg/reconciler/v1alpha1/eventlistener"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	knativetest "knative.dev/pkg/test"
	"testing"
)

func createServiceAccount(t *testing.T, namespace string) *corev1.ServiceAccount {
	t.Helper()
	sa, err := c.KubeClient.CoreV1().ServiceAccounts(namespace).Create(
		&corev1.ServiceAccount{
			ObjectMeta: metav1.ObjectMeta{
				Name: "sa",
			}
		}
	)
	if err != nil {
		t.Fatalf("Failed to create ServiceAccount: %s", err)
	}
	t.Logf("Created ServiceAccount %s in namespace %s", sa.Name, sa.Namespace)
	return sa
	}
}

func TestEventListener(t *testing.T) {
	c, namespace := setup(t)
	t.Parallel()

	defer tearDown(t, c, namespace)
	knativetest.CleanupOnInterrupt(func() { tearDown(t, c, namespace) }, t.Logf)
	t.Log("Start EventListener e2e test")
	sa := createServiceAccount(t, namespace)

	// Create EventListener
	el, err := c.TriggersClient.TektonV1alpha1().EventListeners(namespace).Create(
		&v1alpha1.EventListener{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-eventlistener",
			},
			Spec: v1alpha1.EventListenerSpec{
				ServiceAccountName: sa.Name,
				Triggers: []v1alpha1.Trigger{
					v1alpha1.Trigger{
						TriggerBinding: v1alpha1.TriggerBindingRef{
							Name: "some-trigger-binding",
						},
						TriggerTemplate: v1alpha1.TriggerTemplateRef{
							Name: "some-trigger-template",
						},
					},
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("Failed to create EventListener: %s", err)
	}
	t.Logf("Created EventListener %s in namespace %s", el.Name, el.Namespace)

	verifyMap := map[string]string{
		"Role": fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/namespaces/%s/roles/%s", namespace, fmt.Sprintf("%s%s", el.Name, reconciler.RolePostfix)),
		"RoleBinding": fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/namespaces/%s/rolebindings/%s", namespace, fmt.Sprintf("%s%s", el.Name, reconciler.RoleBindingPostfix)),
		"Deployment": fmt.Sprintf("/apis/apps/v1/namespaces/%s/deployments/%s", namespace, el.Name),
		"Service": fmt.Sprintf("/api/v1/namespaces/%s/services/%s", namespace, el.Name),
	}

	// Verify creation
	for kind, checker := range verifyMap {
		t.Logf("Awaiting %s creation", kind)
		if err = Await(checker.client, checker.uri, true); err != nil {
			t.Fatalf("Failed to create EventListener's %s: %s", kind, err)
		}
		t.Logf("Found EventListener's %s", kind)
	}

	// Delete EventListener
	err = c.TriggersClient.TektonV1alpha1().EventListeners(namespace).Delete(el.Name, &metav1.DeleteOptions{})
	if err != nil {
		t.Fatalf("Failed to delete EventListener: %s", err)
	}
	t.Log("Deleted EventListener")

	// Verify deletion
	for kind, checker := range verifyMap {
		t.Logf("Awaiting %s deletion", kind)
		if err = Await(checker.client, checker.uri, false, nil); err != nil {
			t.Fatalf("Failed to delete EventListener's %s: %s", kind, err)
		}
		t.Logf("Deleted EventListener's %s", kind)
	}
}
