/*
Copyright 2019 The Tekton Authors.

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

/*
Poll Kubernetes resources

After creating Kubernetes resources or making changes to them, you will need to
wait for the system to realize those changes. You can use polling methods to
check the resources reach the desired state.

The `WaitFor*` functions use the Kubernetes
[`wait` package](https://godoc.org/k8s.io/apimachinery/pkg/util/wait). For
polling they use
[`PollImmediate`](https://godoc.org/k8s.io/apimachinery/pkg/util/wait#PollImmediate)
with a [`ConditionFunc`](https://godoc.org/k8s.io/apimachinery/pkg/util/wait#ConditionFunc)
callback function, which returns a `bool` to indicate if the polling should stop
and an `error` to indicate if there was an error.
*/

package test

import (
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
)

const (
	interval = 1 * time.Second
	timeout  = 30 * time.Second
)

// Await polls the k8s rest client for the existence/non-existence of the resource specified by the requestUri.
// Errors other than metav1.StatusReasonNotFound errors are returned.
func Await(kubeClient rest.Interface, requestUri string, exists bool) error {
	return wait.PollImmediate(interval, timeout, func() (bool, error) {
		err := kubeClient.Get().
			RequestURI(requestUri).
			SetHeader("Content-Type", "application/json").
			Do().
			Error()
		if err != nil {
			if errors.IsNotFound(err) {
				return !exists, nil
			} else {
				return false, err
			}
		}
		return true, nil
	})
}
