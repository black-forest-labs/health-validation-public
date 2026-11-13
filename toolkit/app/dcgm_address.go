// SPDX-FileCopyrightText: Copyright (c) 2026, NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package app

import "strings"

// DCGMAddressEnv names the environment variable that selects the DCGM
// hostengine the CLI reads from: "host:port" for a TCP endpoint, or
// "unix:<path>" for a Unix domain socket. Unset or empty keeps the default,
// the hostengine on localhost.
//
// Set it when the hostengine does not listen on the node's localhost, for
// example when DCGM runs in a pod (the NVIDIA GPU Operator's nvidia-dcgm) and
// the CLI reaches it at the pod's address, or when the node runs more than
// one hostengine and the CLI must use a given one.
const DCGMAddressEnv = "HV_DCGM_ADDRESS"

// dcgmEndpoint reads DCGMAddressEnv: the hostengine address, and whether it
// is a Unix socket path. An empty address keeps the DCGM source's default.
func dcgmEndpoint(getenv func(string) string) (address string, isSocket bool) {
	v := strings.TrimSpace(getenv(DCGMAddressEnv))
	if path, ok := strings.CutPrefix(v, "unix:"); ok {
		return path, true
	}
	return v, false
}
