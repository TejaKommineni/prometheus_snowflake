// Copyright 2022 The Prometheus Authors
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

// This file is generated by "make plugins".

package plugins

import (
	// Register aws plugin.
	_ "github.com/prometheus/prometheus/discovery/aws"

	// Register azure plugin.
	_ "github.com/prometheus/prometheus/discovery/azure"

	// Register consul plugin.
	_ "github.com/prometheus/prometheus/discovery/consul"

	// Register digitalocean plugin.
	_ "github.com/prometheus/prometheus/discovery/digitalocean"

	// Register dns plugin.
	_ "github.com/prometheus/prometheus/discovery/dns"

	// Register eureka plugin.
	_ "github.com/prometheus/prometheus/discovery/eureka"

	// Register gce plugin.
	_ "github.com/prometheus/prometheus/discovery/gce"

	// Register hetzner plugin.
	_ "github.com/prometheus/prometheus/discovery/hetzner"

	// Register ionos plugin.
	_ "github.com/prometheus/prometheus/discovery/ionos"

	// Register kubernetes plugin.
	_ "github.com/prometheus/prometheus/discovery/kubernetes"

	// Register linode plugin.
	_ "github.com/prometheus/prometheus/discovery/linode"

	// Register marathon plugin.
	_ "github.com/prometheus/prometheus/discovery/marathon"

	// Register moby plugin.
	_ "github.com/prometheus/prometheus/discovery/moby"

	// Register nomad plugin.
	_ "github.com/prometheus/prometheus/discovery/nomad"

	// Register openstack plugin.
	_ "github.com/prometheus/prometheus/discovery/openstack"

	// Register ovhcloud plugin.
	_ "github.com/prometheus/prometheus/discovery/ovhcloud"

	// Register puppetdb plugin.
	_ "github.com/prometheus/prometheus/discovery/puppetdb"

	// Register scaleway plugin.
	_ "github.com/prometheus/prometheus/discovery/scaleway"

	// Register snowflake plugin.
	_ "github.com/prometheus/prometheus/discovery/snowflake"

	// Register triton plugin.
	_ "github.com/prometheus/prometheus/discovery/triton"

	// Register uyuni plugin.
	_ "github.com/prometheus/prometheus/discovery/uyuni"

	// Register vultr plugin.
	_ "github.com/prometheus/prometheus/discovery/vultr"

	// Register xds plugin.
	_ "github.com/prometheus/prometheus/discovery/xds"

	// Register zookeeper plugin.
	_ "github.com/prometheus/prometheus/discovery/zookeeper"
)
