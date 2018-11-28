/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package bootstrap

import (
	"context"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/controller/common"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/pkg/errors"
)

type NetworkBootstrapper struct {
	certificate         core.Certificate
	sessionManager      *SessionManager
	bootstrapper        *Bootstrapper
	authController      *AuthorizationController
	challengeController *ChallengeResponseController
}

func (nb *NetworkBootstrapper) Bootstrap(ctx context.Context) error {
	if OriginIsDiscovery(nb.certificate) {
		return nb.bootstrapDiscovery(ctx)
	}
	return nb.bootstrapJoiner(ctx)
}

func (nb *NetworkBootstrapper) Start() {
	nb.bootstrapper.Start()
}

type DiscoveryNode struct {
	Host *host.Host
	Node core.BootstrapNode
}

func (nb *NetworkBootstrapper) bootstrapJoiner(ctx context.Context) error {
	discoveryNode, err := nb.bootstrapper.Bootstrap(ctx)
	if err != nil {
		return errors.Wrap(err, "Error bootstrapping to discovery node")
	}
	sessionID, err := nb.authController.Authorize(ctx, discoveryNode, nb.certificate)
	if err != nil {
		return errors.Wrap(err, "Error authorizing on discovery node")
	}
	data, err := nb.challengeController.Execute(ctx, discoveryNode, sessionID)
	if err != nil {
		return errors.Wrap(err, "Error executing double challenge response")
	}
	// TODO: use data
	print(data.AssignShortID)
	return nb.authController.Register(ctx, discoveryNode, sessionID)
}

func (nb *NetworkBootstrapper) bootstrapDiscovery(ctx context.Context) error {
	return nb.bootstrapper.BootstrapDiscovery(ctx)
}

func NewNetworkBootstrapper(options *common.Options, cert core.Certificate, transport network.InternalTransport) *NetworkBootstrapper {
	nb := &NetworkBootstrapper{}
	nb.certificate = cert
	nb.sessionManager = NewSessionManager()
	nb.bootstrapper = NewBootstrapper(options, cert, transport)
	nb.authController = NewAuthorizationController(options, transport, nb.sessionManager)
	// nb.challengeController = NewChallengeController()
	return nb
}
