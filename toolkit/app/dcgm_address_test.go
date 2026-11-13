// SPDX-FileCopyrightText: Copyright (c) 2026, NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package app

import "testing"

func TestDCGMEndpoint(t *testing.T) {
	for _, test := range []struct {
		env, address string
		socket       bool
	}{
		{env: "", address: ""},
		{env: "   ", address: ""},
		{env: "10.0.3.7:5555", address: "10.0.3.7:5555"},
		{env: " nvidia-dcgm.gpu-operator:5555 ", address: "nvidia-dcgm.gpu-operator:5555"},
		{env: "unix:/run/nvidia-dcgm/dcgm.sock", address: "/run/nvidia-dcgm/dcgm.sock", socket: true},
	} {
		address, socket := dcgmEndpoint(func(name string) string {
			if name != DCGMAddressEnv {
				t.Fatalf("read %s, want %s", name, DCGMAddressEnv)
			}
			return test.env
		})
		if address != test.address || socket != test.socket {
			t.Errorf("%q: address %q socket %v, want %q %v", test.env, address, socket, test.address, test.socket)
		}
	}
}
