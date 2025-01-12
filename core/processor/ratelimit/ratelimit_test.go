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

package ratelimit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code.vegaprotocol.io/vega/core/processor/ratelimit"
)

// runN executes the given `fn` func, `n` times.
func runN(n int, fn func()) {
	for {
		if n == 0 {
			return
		}
		n--
		fn()
	}
}

func TestRateLimits(t *testing.T) {
	t.Run("Single Block", func(t *testing.T) {
		r := ratelimit.New(10, 10) // 10 requests in the last 10 blocks

		// make 9 requests, all should be allowed
		runN(9, func() {
			ok := r.Allow("test")
			assert.True(t, ok)
		})

		// request number 10 should not be allowed
		ok := r.Allow("test")
		assert.False(t, ok)
	})

	t.Run("Even Blocks", func(t *testing.T) {
		r := ratelimit.New(10, 10) // 10 requests in the last 10 blocks

		// this will make 1 request and move to the next block.
		runN(9, func() {
			ok := r.Allow("test")
			assert.True(t, ok)
			r.NextBlock()
		})

		ok := r.Allow("test")
		assert.False(t, ok)
	})

	t.Run("Uneven Blocks", func(t *testing.T) {
		r := ratelimit.New(10, 3) // 10 request in the last 3 blocks

		// let's fill the rate limiter
		runN(100, func() {
			_ = r.Allow("test")
		})
		require.False(t, r.Allow("test"))

		r.NextBlock()
		assert.False(t, r.Allow("test"))

		r.NextBlock()
		assert.False(t, r.Allow("test"))

		r.NextBlock()
		assert.True(t, r.Allow("test"))
	})

	t.Run("Clean up", func(t *testing.T) {
		r := ratelimit.New(10, 10)
		runN(10, func() {
			r.Allow("test")
		})
		require.Equal(t, 10, r.Count("test"))

		runN(10, func() {
			r.NextBlock()
		})
		require.Equal(t, -1, r.Count("test"))
	})
}
