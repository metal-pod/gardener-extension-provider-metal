// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package worker

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/metal-pod/gardener-extension-provider-metal/pkg/metal"

	"github.com/gardener/gardener/pkg/utils/chart"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

var (
	mcmChart = &chart.Chart{
		Name:   metal.MachineControllerManagerName,
		Path:   filepath.Join(metal.InternalChartsPath, metal.MachineControllerManagerName, "seed"),
		Images: []string{metal.MachineControllerManagerImageName},
		Objects: []*chart.Object{
			{Type: &appsv1.Deployment{}, Name: metal.MachineControllerManagerName},
			{Type: &corev1.Service{}, Name: metal.MachineControllerManagerName},
			{Type: &corev1.ServiceAccount{}, Name: metal.MachineControllerManagerName},
			{Type: &corev1.Secret{}, Name: metal.MachineControllerManagerName},
		},
	}

	mcmShootChart = &chart.Chart{
		Name: metal.MachineControllerManagerName,
		Path: filepath.Join(metal.InternalChartsPath, metal.MachineControllerManagerName, "shoot"),
		Objects: []*chart.Object{
			{Type: &rbacv1.ClusterRole{}, Name: fmt.Sprintf("extensions.gardener.cloud:%s:%s", metal.Name, metal.MachineControllerManagerName)},
			{Type: &rbacv1.ClusterRoleBinding{}, Name: fmt.Sprintf("extensions.gardener.cloud:%s:%s", metal.Name, metal.MachineControllerManagerName)},
		},
	}
)

func (w *workerDelegate) GetMachineControllerManagerChartValues(ctx context.Context) (map[string]interface{}, error) {
	namespace := &corev1.Namespace{}
	if err := w.client.Get(ctx, kutil.Key(w.worker.Namespace), namespace); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"providerName": metal.Name,
		"namespace": map[string]interface{}{
			"uid": namespace.UID,
		},
	}, nil
}

func (w *workerDelegate) GetMachineControllerManagerShootChartValues(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"providerName": metal.Name,
	}, nil
}
