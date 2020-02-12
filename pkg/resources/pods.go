//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	corev1 "k8s.io/api/core/v1"
)

var podSecurity = &corev1.PodSecurityContext{
	RunAsNonRoot: &runAsNonRoot,
	RunAsUser:    &runAsUser,
	FSGroup:      &fsgroup,
}

var certManagerControllerPod = corev1.PodSpec{
	ImagePullSecrets: []corev1.LocalObjectReference{{
		Name: ImagePullSecret,
	}},
	ServiceAccountName: ServiceAccount,
	SecurityContext:    podSecurity,
	Containers: []corev1.Container{
		controllerContainer,
	},
}

var certManagerWebhookPod = corev1.PodSpec{
	ImagePullSecrets: []corev1.LocalObjectReference{{
		Name: ImagePullSecret,
	}},
	HostNetwork:        TrueVar,
	ServiceAccountName: ServiceAccount,
	SecurityContext:    podSecurity,
	Containers: []corev1.Container{
		webhookContainer,
	},
	Volumes: []corev1.Volume{
		{
			Name: "certs",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: "cert-manager-webhook-tls",
				},
			},
		},
	},
}

var certManagerCainjectorPod = corev1.PodSpec{
	ImagePullSecrets: []corev1.LocalObjectReference{{
		Name: ImagePullSecret,
	}},
	ServiceAccountName: ServiceAccount,
	SecurityContext:    podSecurity,
	Containers: []corev1.Container{
		cainjectorContainer,
	},
}

var configmapWatcherPod = corev1.PodSpec{
	ImagePullSecrets: []corev1.LocalObjectReference{{
		Name: ImagePullSecret,
	}},
	ServiceAccountName: ServiceAccount,
	SecurityContext:    podSecurity,
	Containers: []corev1.Container{
		configmapWatcherContainer,
	},
}