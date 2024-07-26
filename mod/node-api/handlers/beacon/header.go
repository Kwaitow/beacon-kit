// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package beacon

import (
	"strconv"

	beacontypes "github.com/berachain/beacon-kit/mod/node-api/handlers/beacon/types"
	"github.com/berachain/beacon-kit/mod/node-api/handlers/utils"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
)

func (h *Handler[
	BeaconBlockHeaderT, ContextT, _, _,
]) GetBlockHeaders(c ContextT) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.GetBlockHeadersRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	slot, err := strconv.ParseUint(req.Slot, 10, 64)
	if err != nil {
		return nil, err
	}
	header, err := h.backend.BlockHeaderAtSlot(slot)
	if err != nil {
		return nil, err
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data: &beacontypes.BlockHeaderResponse[BeaconBlockHeaderT]{
			Root:      header.GetBodyRoot(),
			Canonical: true,
			Header: &beacontypes.BlockHeader[BeaconBlockHeaderT]{
				Message:   header,
				Signature: bytes.B48{}, // TODO: implement
			},
		},
	}, nil
}

func (h *Handler[
	BeaconBlockHeaderT, ContextT, _, _,
]) GetBlockHeaderByID(c ContextT) (any, error) {
	req, err := utils.BindAndValidate[beacontypes.GetBlockHeaderRequest](
		c, h.Logger(),
	)
	if err != nil {
		return nil, err
	}
	slot, err := utils.SlotFromBlockID(req.BlockID, h.backend)
	if err != nil {
		return nil, err
	}
	header, err := h.backend.BlockHeaderAtSlot(slot)
	if err != nil {
		return nil, err
	}
	return beacontypes.ValidatorResponse{
		ExecutionOptimistic: false, // stubbed
		Finalized:           false, // stubbed
		Data: &beacontypes.BlockHeaderResponse[BeaconBlockHeaderT]{
			Root:      header.GetBodyRoot(),
			Canonical: true,
			Header: &beacontypes.BlockHeader[BeaconBlockHeaderT]{
				Message:   header,
				Signature: bytes.B48{}, // TODO: implement
			},
		},
	}, nil
}