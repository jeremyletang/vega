// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package ratelimit

import (
	"encoding/base64"
)

type Key []byte

func (k Key) String() string {
	return base64.StdEncoding.EncodeToString(k)
}

type Rates struct {
	block      int
	requests   int
	perNBlocks int
	entries    map[string][]int
}

func New(requests, perNBlocks int) *Rates {
	return &Rates{
		block:      0,
		requests:   requests,
		perNBlocks: perNBlocks,
		entries:    map[string][]int{},
	}
}

// Count returns the number of requests recorded for a given key
// It returns -1 if the key has been not recorded or evicted.
func (r *Rates) Count(key string) int {
	entry, ok := r.entries[key]
	if !ok {
		return -1
	}

	var count int
	for _, n := range entry {
		count += n
	}
	return count
}

func (r *Rates) NextBlock() {
	// compute the next block index
	r.block = (r.block + 1) % (r.perNBlocks)

	// reset the counters for that particular block index
	for _, c := range r.entries {
		c[r.block] = 0
	}

	// We clean up the entries after finishing the block round
	if r.block != 0 {
		return
	}

	for key := range r.entries {
		if r.Count(key) == 0 {
			delete(r.entries, key)
		}
	}
}

func (r *Rates) Allow(key string) bool {
	entry, ok := r.entries[key]
	if !ok {
		entry = make([]int, r.perNBlocks)
		r.entries[key] = entry
	}
	entry[r.block]++

	return r.Count(key) < r.requests
}
