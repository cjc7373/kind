/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliep.
See the License for the specific language governing permissions and
limitations under the License.
*/

package docker

// clusterLabelKey is applied to each "node" docker container for identification
const clusterLabelKey = "io.x-k8s.kind.cluster"

// nodeRoleLabelKey is applied to each "node" docker container for categorization
// of nodes by role
const nodeRoleLabelKey = "io.x-k8s.kind.role"

// Docker supports the following restart modes:
// - no
// - on-failure[:max-retries]
// - unless-stopped
// - always
// https://docs.docker.com/engine/reference/commandline/run/#restart-policies---restart
//
// What we desire is:
// - restart on host / dockerd reboot
// - don't restart for any other reason
//
// This means:
// - no is out of the question ... it never restarts
// - always is a poor choice, we'll keep trying to restart nodes that were
// never going to work
// - unless-stopped will also retry failures indefinitely, similar to always
// except that it won't restart when the container is `docker stop`ed
// - on-failure is not great, we're only interested in restarting on
// reboots, not failures. *however* we can limit the number of retries
// *and* it forgets all state on dockerd restart and retries anyhow.
// - on-failure:0 is what we want .. restart on failures, except max
// retries is 0, so only restart on reboots.
// however this _actually_ means the same thing as always
// so the closest thing is on-failure:1, which will retry *once*
const dockerRestartPolicyStartArg = "--restart=on-failure:1"
const dockerRestartPolicyStopArg = "--restart=no"
