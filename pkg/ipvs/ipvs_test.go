// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipvs

import (
	"net"
	"testing"
)

var want = []string{
	`apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  name: kube-lvscare
  namespace: kube-system
spec:
  containers:
  - args:
    - care
    - --vs
    - 10.107.2.1:6443
    - --health-path
    - /healthz
    - --health-schem
    - https
    - --rs
    - 172.16.228.157:6443
    - --rs
    - 172.16.228.158:6443
    - --rs
    - 172.16.228.159:6443
    command:
    - /usr/bin/lvscare
    image: fdfadf
    imagePullPolicy: IfNotPresent
    name: kube-lvscare
    resources: {}
    securityContext:
      privileged: true
    volumeMounts:
    - mountPath: /lib/modules
      name: lib-modules
      readOnly: true
  hostNetwork: true
  volumes:
  - hostPath:
      path: /lib/modules
      type: ""
    name: lib-modules
status: {}
`,
	`apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  name: kube-lvscare
  namespace: kube-system
spec:
  containers:
  - args:
    - care
    - --vs
    - 10.107.2.1:6443
    - --health-path
    - /healthz
    - --health-schem
    - https
    - --rs
    - 172.16.228.157:6443
    command:
    - /usr/bin/lvscare
    image: fdfadf
    imagePullPolicy: IfNotPresent
    name: kube-lvscare
    resources: {}
    securityContext:
      privileged: true
    volumeMounts:
    - mountPath: /lib/modules
      name: lib-modules
      readOnly: true
  hostNetwork: true
  volumes:
  - hostPath:
      path: /lib/modules
      type: ""
    name: lib-modules
status: {}
`,
}

func TestLvsStaticPodYaml(t *testing.T) {
	type args struct {
		vip     net.IP
		masters []net.IP
		image   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				vip: net.ParseIP("10.107.2.1"),
				masters: []net.IP{
					net.ParseIP("172.16.228.157"),
					net.ParseIP("172.16.228.158"),
					net.ParseIP("172.16.228.159")},
				image: "fdfadf",
			},
			want: want[0],
		},
		{
			args: args{
				vip:     net.ParseIP("10.107.2.1"),
				masters: []net.IP{net.ParseIP("172.16.228.157")},
				image:   "fdfadf",
			},
			want: want[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := LvsStaticPodYaml(tt.args.vip, tt.args.masters, tt.args.image); got != tt.want {
				t.Errorf("LvsStaticPodYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}
