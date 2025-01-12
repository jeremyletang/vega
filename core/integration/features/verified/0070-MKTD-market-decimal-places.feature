Feature: Allow markets to be specified with a smaller number of decimal places than the underlying settlement asset

    Background:
        Given the following network parameters are set:
            | name                                          | value |
            | market.stake.target.timeWindow                | 24h   |
            | market.stake.target.scalingFactor             | 1     |
            | market.liquidity.bondPenaltyParameter         | 0.2   |
            | market.liquidity.targetstake.triggering.ratio | 0.1   |
        And the following assets are registered:
            | id  | decimal places |
            | ETH | 5              |
            | USD | 2              |
        And the average block duration is "1"
        And the log normal risk model named "log-normal-risk-model-1":
            | risk aversion | tau | mu | r | sigma |
            | 0.000001      | 0.1 | 0  | 0 | 1.0   |
        And the fees configuration named "fees-config-1":
            | maker fee | infrastructure fee |
            | 0.004     | 0.001              |
        And the price monitoring named "price-monitoring-1":
            | horizon | probability | auction extension |
            | 1       | 0.99        | 300               |
        And the markets:
            | id        | quote name | asset | risk model              | margin calculator         | auction duration | fees          | price monitoring   | oracle config          | decimal places | position decimal places |
            | ETH/MAR22 | ETH        | USD   | log-normal-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future | 0              | 0                       |
            | USD/DEC19 | USD        | ETH   | log-normal-risk-model-1 | default-margin-calculator | 1                | default-none  | price-monitoring-1 | default-usd-for-future | 3              | 3                       |
            | USD/DEC20 | USD        | ETH   | log-normal-risk-model-1 | default-margin-calculator | 1                | default-none  | price-monitoring-1 | default-usd-for-future | 5              | 5                       |
            | USD/DEC21 | USD        | ETH   | log-normal-risk-model-1 | default-margin-calculator | 1                | default-none  | price-monitoring-1 | default-usd-for-future | 5              | 3                       |
        And the parties deposit on asset's general account the following amount:
            | party  | asset | amount    |
            | party0 | USD   | 5000000   |
            | party0 | ETH   | 5000000   |
            | party1 | USD   | 100000000 |
            | party1 | ETH   | 100000000 |
            | party2 | USD   | 100000000 |
            | party2 | ETH   | 100000000 |
            | party3 | USD   | 100000000 |
            | lpprov | ETH   | 100000000 |
            | lpprov | USD   | 100000000 |

    Scenario: Markets with different precisions trade at the same price

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp0 | party0 | USD/DEC20 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp0 | party0 | USD/DEC20 | 1000              | 0.001 | buy  | BID              | 100        | -20    | submission |
            | lp1 | party0 | USD/DEC21 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC21 | 1000              | 0.001 | buy  | BID              | 100        | -20    | submission |
            | lp2 | party0 | USD/DEC19 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp2 | party0 | USD/DEC19 | 1000              | 0.001 | buy  | BID              | 100        | -20    | submission |
            | lp3 | lpprov | USD/DEC20 | 4000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp3 | lpprov | USD/DEC20 | 4000              | 0.001 | buy  | BID              | 100        | -20    | submission |
            | lp4 | lpprov | USD/DEC21 | 4000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp4 | lpprov | USD/DEC21 | 4000              | 0.001 | buy  | BID              | 100        | -20    | submission |
            | lp5 | lpprov | USD/DEC19 | 4000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp5 | lpprov | USD/DEC19 | 4000              | 0.001 | buy  | BID              | 100        | -20    | submission |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC21 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC21 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party1 | USD/DEC20 | buy  | 1000   | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC20 | sell | 1000   | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party1 | USD/DEC19 | buy  | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC19 | sell | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party0 | USD/DEC21 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party0 | USD/DEC21 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party0 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party0 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party0 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party0 | USD/DEC19 | sell | 1      | 1100   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2a |

    Then the market data for the market "USD/DEC19" should be:
      | target stake | supplied stake |
      | 3556         | 5000           |
    Then the market data for the market "USD/DEC20" should be:
      | target stake | supplied stake |
      | 3556         | 5000           |
    Then the market data for the market "USD/DEC21" should be:
      | target stake | supplied stake |
      | 3556         | 5000           |

        When the opening auction period ends for market "USD/DEC21"
        And the opening auction period ends for market "USD/DEC20"
        And the opening auction period ends for market "USD/DEC19"

        Then the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "USD/DEC21"
        And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "USD/DEC20"
        And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "USD/DEC19"

        Then the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond |
            | party0 | ETH   | USD/DEC21 | 8536   | 4972160  | 1000 |
            | party1 | ETH   | USD/DEC21 | 1081   | 99996757 | 0    |
            | party2 | ETH   | USD/DEC21 | 4388   | 99986836 | 0    |
            | party0 | ETH   | USD/DEC20 | 7768   | 4972160  | 1000 |
            | party1 | ETH   | USD/DEC20 | 1081   | 99996757 | 0    |
            | party2 | ETH   | USD/DEC20 | 4388   | 99986836 | 0    |
            | party0 | ETH   | USD/DEC19 | 8536   | 4972160  | 1000 |
            | party1 | ETH   | USD/DEC19 | 1081   | 99996757 | 0    |
            | party2 | ETH   | USD/DEC19 | 4388   | 99986836 | 0    |

    Scenario: Users engage in a USD market auction, (0070-MKTD-003, 0070-MKTD-008)
        Given the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | ETH/MAR22 | 35569             | 0.001 | sell | ASK              | 500        | 20     | submission |
            | lp1 | party0 | ETH/MAR22 | 35569             | 0.001 | buy  | BID              | 500        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference  |
            | party1 | ETH/MAR22 | buy  | 1      | 9     | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | ETH/MAR22 | buy  | 1      | 9     | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | ETH/MAR22 | buy  | 10     | 10    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
            | party2 | ETH/MAR22 | sell | 10     | 10    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3 |
            | party2 | ETH/MAR22 | sell | 1      | 10    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
            | party2 | ETH/MAR22 | sell | 1      | 11    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

        When the opening auction period ends for market "ETH/MAR22"
        Then the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR22"
        And the auction ends with a traded volume of "10" at a price of "10"
        And the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond |
            | party0 | USD   | ETH/MAR22 | 307317 | 4657114  | 35569  |
            | party1 | USD   | ETH/MAR22 | 12730  | 99987270 | 0    |
            | party2 | USD   | ETH/MAR22 | 51819  | 99948181 | 0    |
        And the following trades should be executed:
            | buyer  | price | size | seller |
            | party1 | 10    | 10   | party2 |

    Scenario: Users engage in an ETH market auction, (0070-MKTD-003, 0070-MKTD-008)
        Given the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC19 | 50000             | 0.001 | sell | ASK              | 500        | 20     | submission |
            | lp1 | party0 | USD/DEC19 | 50000             | 0.001 | buy  | BID              | 500        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference  |
            | party1 | USD/DEC19 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC19 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC19 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
            | party2 | USD/DEC19 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3 |
            | party2 | USD/DEC19 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
            | party2 | USD/DEC19 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

        When the opening auction period ends for market "USD/DEC19"
        Then the auction ends with a traded volume of "10" at a price of "1000"
        And the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond  |
            | party0 | ETH   | USD/DEC19 | 422988 | 4527012  | 50000 |
            | party1 | ETH   | USD/DEC19 | 1273   | 99998727 | 0     |
            | party2 | ETH   | USD/DEC19 | 5188   | 99994812 | 0     |
        And the following trades should be executed:
            | buyer  | price | size | seller |
            | party1 | 1000  | 10   | party2 |

    Scenario: Users engage in an ETH market auction with full decimal places, (0070-MKTD-003, 0070-MKTD-008)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC20 | 500               | 0.001 | sell | ASK              | 500        | 20     | submission |
            | lp1 | party0 | USD/DEC20 | 500               | 0.001 | buy  | BID              | 500        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference  |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC20 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
            | party2 | USD/DEC20 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3 |
            | party2 | USD/DEC20 | sell | 1      | 101000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
            | party2 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

        When the opening auction period ends for market "USD/DEC20"
        Then the auction ends with a traded volume of "10" at a price of "100000"
        And the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond |
            | party0 | ETH   | USD/DEC20 | 4230   | 4995270  | 500  |
            | party1 | ETH   | USD/DEC20 | 13     | 99999987 | 0    |
            | party2 | ETH   | USD/DEC20 | 52     | 99999948 | 0    |
        And the following trades should be executed:
            | buyer  | price  | size | seller |
            | party1 | 100000 | 10   | party2 |

    Scenario: User tops up markets with differing precisions with the same asset + amount, should result in identical margin changes, (0070-MKTD-004)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC20 | 100000            | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC20 | 100000            | 0.001 | buy  | BID              | 100        | -20    | amendment  |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC20 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party2 | USD/DEC20 | sell | 1      | 101000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1  |
            | party2 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC19 | sell | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party2 | USD/DEC19 | sell | 1      | 1010   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1a |
            | party2 | USD/DEC19 | sell | 1      | 1100   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2a |

    Then the market data for the market "USD/DEC19" should be:
      | target stake | supplied stake |
      | 3556         | 5000           |

        When the opening auction period ends for market "USD/DEC20"
        And the opening auction period ends for market "USD/DEC19"

        # party1 has position +10 and mark price 100000

        Then the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond   |
            | party0 | ETH   | USD/DEC20 | 845206 | 4007110  | 100000 |
            | party1 | ETH   | USD/DEC20 | 13     | 99998714 | 0      |
            | party2 | ETH   | USD/DEC20 | 52     | 99994760 | 0      |
            | party0 | ETH   | USD/DEC19 | 42684  | 4007110  | 5000   |
            | party1 | ETH   | USD/DEC19 | 1273   | 99998714 | 0      |
            | party2 | ETH   | USD/DEC19 | 5188   | 99994760 | 0      |

        When the parties deposit on asset's general account the following amount:
            | party  | asset | amount |
            | party0 | ETH   | 1000   |
            | party1 | ETH   | 1000   |
            | party2 | ETH   | 1000   |
        Then the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond   |
            | party0 | ETH   | USD/DEC20 | 845206 | 4008110  | 100000 |
            | party1 | ETH   | USD/DEC20 | 13     | 99999714 | 0      |
            | party2 | ETH   | USD/DEC20 | 52     | 99995760 | 0      |
            | party0 | ETH   | USD/DEC19 | 42684  | 4008110  | 5000   |
            | party1 | ETH   | USD/DEC19 | 1273   | 99999714 | 0      |
            | party2 | ETH   | USD/DEC19 | 5188   | 99995760 | 0      |

    Scenario: User checks prices after opening auction, (0070-MKTD-005)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC20 | 100000            | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC20 | 100000            | 0.001 | buy  | BID              | 100        | -20    | amendment  |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC20 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party2 | USD/DEC20 | sell | 1      | 101000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1  |
            | party2 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC19 | sell | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party2 | USD/DEC19 | sell | 1      | 1010   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1a |
            | party2 | USD/DEC19 | sell | 1      | 1100   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2a |

        When the opening auction period ends for market "USD/DEC20"
        And the opening auction period ends for market "USD/DEC19"
        Then the mark price should be "100000" for the market "USD/DEC20"
        And the mark price should be "1000" for the market "USD/DEC19"

    Scenario: Offsets are calculated in market units, (0070-MKTD-007)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC20 | 5000              | 0.001 | sell | MID              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC20 | 5000              | 0.001 | buy  | MID              | 100        | -20    | amendment  |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | sell | MID              | 100        | 20     | submission |
            | lp2 | party0 | USD/DEC19 | 5000              | 0.001 | buy  | MID              | 100        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party1 | USD/DEC20 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC20 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party2 | USD/DEC20 | sell | 1      | 101000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1  |
            | party2 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party1 | USD/DEC19 | buy  | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC19 | sell | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party2 | USD/DEC19 | sell | 1      | 1010   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1a |
            | party2 | USD/DEC19 | sell | 1      | 1100   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2a |

        When the opening auction period ends for market "USD/DEC20"
        And the opening auction period ends for market "USD/DEC19"
        And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "USD/DEC19"
        When the parties place the following pegged orders:
            | party  | market id | side | volume | pegged reference | offset |
            | party1 | USD/DEC19 | sell | 5      | ASK              | 5      |
            | party1 | USD/DEC20 | sell | 5      | ASK              | 5      |
        Then the pegged orders should have the following states:
            | party  | market id | side | volume | reference | offset | price  | status        |
            | party1 | USD/DEC20 | sell | 5      | ASK       | 5      | 101005 | STATUS_ACTIVE |
            | party1 | USD/DEC19 | sell | 5      | ASK       | 5      | 1015   | STATUS_ACTIVE |

    Scenario: Price monitoring bounds are calculated at asset precision but displayed rounded, (0070-MKTD-006)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp1 | party0 | USD/DEC20 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC20 | 1000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |
            | lp1 | party0 | USD/DEC21 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp1 | party0 | USD/DEC21 | 1000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |
            | lp2 | party0 | USD/DEC19 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp2 | party0 | USD/DEC19 | 1000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price  | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC21 | buy  | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC21 | sell | 10     | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party1 | USD/DEC20 | buy  | 1000   | 100000 | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2   |
            | party2 | USD/DEC20 | sell | 1000   | 100000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3  |
            | party1 | USD/DEC19 | buy  | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC19 | sell | 10     | 1000   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party0 | USD/DEC21 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party0 | USD/DEC21 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party0 | USD/DEC20 | buy  | 1      | 90000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1   |
            | party0 | USD/DEC20 | sell | 1      | 110000 | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2  |
            | party0 | USD/DEC19 | buy  | 1      | 900    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party0 | USD/DEC19 | sell | 1      | 1100   | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2a |

        When the opening auction period ends for market "USD/DEC20"
        And the opening auction period ends for market "USD/DEC19"

        Then the price monitoring bounds for the market "USD/DEC19" should be:
            | min bound | max bound |
            | 1000      | 1000      |
        And the price monitoring bounds for the market "USD/DEC20" should be:
            | min bound | max bound |
            | 99955     | 100045    |
        And the price monitoring bounds for the market "USD/DEC21" should be:
            | min bound | max bound |
            | 99955     | 100045    |
