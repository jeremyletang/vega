Feature: Price monitoring test for issue 2668

  Background:
    Given the markets starts on "2020-10-16T00:00:00Z" and expires on "2020-12-31T23:59:59Z"
    And the execution engine have these markets:
      | name      | baseName | quoteName | asset | markprice | risk model | lamd/long | tau/short              | mu/max move up | r/min move down | sigma | release factor | initial factor | search factor | settlementPrice | openAuction | trading mode | makerFee | infrastructureFee | liquidityFee | p. m. update freq. | p. m. horizons | p. m. probs | p. m. durations | Prob of trading | oracleSpecPubKeys     | oracleSpecProperty | oracleSpecPropertyType | oracleSpecBinding |
      | ETH/DEC20 | BTC      | ETH       | ETH   | 5000000   | forward    | 0.000001  | 0.00011407711613050422 | 0              | 0.016           | 0.8   | 1.4            | 1.2            | 1.1           | 42              | 0           | continuous   | 0        | 0                 | 0            | 1                  | 43200          | 0.9999999   | 300             | 0.1             | 0xDEADBEEF,0xCAFEDOOD | prices.ETH.value   | TYPE_INTEGER           | prices.ETH.value  |
    And oracles broadcast data signed with "0xDEADBEEF":
      | name             | value |
      | prices.ETH.value | 42   |
    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

  Scenario: Upper bound breached
    Given the following traders:
      | name    |      amount  |
      | trader1 | 10000000000  |
      | trader2 | 10000000000  |

    Then traders place following orders:
      | trader  | id        | type | volume |    price   | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   5670000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   5670000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "5670000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   4850000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   4850000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "4850000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   6630000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   6630000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   6640000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   6640000  |                0 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min
    Then the time is updated to "2020-10-16T00:05:00Z"

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min01s
    Then the time is updated to "2020-10-16T00:05:01Z"

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

  Scenario: Lower bound breached
    Given the following traders:
      | name    |      amount  |
      | trader1 | 10000000000  |
      | trader2 | 10000000000  |

    Then traders place following orders:
      | trader  | id        | type | volume |    price   | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   5670000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   5670000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "5670000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   4850000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   4850000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "4850000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   6630000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   6630000  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   4840000  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   4840000  |                0 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min
    Then the time is updated to "2020-10-16T00:05:00Z"

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min01s
    Then the time is updated to "2020-10-16T00:05:01Z"

    And the mark price for the market "ETH/DEC20" is "6630000"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Scenario: Upper bound breached (scale prices down by 10000)
    Given the following traders:
      | name    |      amount  |
      | trader1 | 10000000000  |
      | trader2 | 10000000000  |

    Then traders place following orders:
      | trader  | id        | type | volume |    price   | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   567  |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   567  |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "567"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   485      |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   485      |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "485"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   663      |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   663      |                1 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "663"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"

    Then traders place following orders:
      | trader  | id        | type | volume |   price    | resulting trades | type       | tif     |
      | trader1 | ETH/DEC20 | sell |      1 |   664 |                0 | TYPE_LIMIT | TIF_GTC |
      | trader2 | ETH/DEC20 | buy  |      1 |   664 |                0 | TYPE_LIMIT | TIF_FOK |

    And the mark price for the market "ETH/DEC20" is "663"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min
    Then the time is updated to "2020-10-16T00:05:00Z"

    And the mark price for the market "ETH/DEC20" is "663"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_MONITORING_AUCTION"

    # T0 + 5min01s
    Then the time is updated to "2020-10-16T00:05:01Z"

    And the mark price for the market "ETH/DEC20" is "663"

    And the market trading mode for the market "ETH/DEC20" is "TRADING_MODE_CONTINUOUS"
