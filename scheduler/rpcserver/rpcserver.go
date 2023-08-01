/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpcserver

import (
	"google.golang.org/grpc"

	"github.com/XDTD/Dragonfly2/pkg/rpc/scheduler/server"
	"github.com/XDTD/Dragonfly2/scheduler/config"
	"github.com/XDTD/Dragonfly2/scheduler/networktopology"
	"github.com/XDTD/Dragonfly2/scheduler/resource"
	"github.com/XDTD/Dragonfly2/scheduler/scheduling"
	"github.com/XDTD/Dragonfly2/scheduler/storage"
)

// New returns a new scheduler server from the given options.
func New(
	cfg *config.Config,
	resource resource.Resource,
	scheduling scheduling.Scheduling,
	dynconfig config.DynconfigInterface,
	storage storage.Storage,
	networkTopology networktopology.NetworkTopology,
	opts ...grpc.ServerOption,
) *grpc.Server {
	return server.New(
		newSchedulerServerV1(cfg, resource, scheduling, dynconfig, storage, networkTopology),
		newSchedulerServerV2(cfg, resource, scheduling, dynconfig, storage, networkTopology),
		opts...)
}
