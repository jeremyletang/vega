Feature: Verify the order size is correctly cumulated.

  Background:
    Given the insurance pool initial balance for the markets is "0":
    And the execution engine have these markets:
      | name      | quote name | asset | risk model | lamd/long | tau/short              | mu/max move up | r/min move down | sigma | release factor | initial factor | search factor | auction duration | maker fee | infrastructure fee | liquidity fee | p. m. update freq. | p. m. horizons | p. m. probs | p. m. durations | prob. of trading | oracle spec pub. keys | oracle spec property | oracle spec property type | oracle spec binding |
      | ETH/DEC19 | ETH        | ETH   | forward    | 0.001     | 0.00000190128526884174 | 0              | 0.016           | 2.5   | 1.4            | 1.2            | 1.1           | 1                | 0         | 0                  | 0             | 0                  |                |             |                 | 0.1              | 0xDEADBEEF,0xCAFEDOOD | prices.ETH.value     | TYPE_INTEGER              | prices.ETH.value    |
    And oracles broadcast data signed with "0xDEADBEEF":
      | name             | value |
      | prices.ETH.value | 42    |

  Scenario: Order from liquidity provision and from normal order submission are correctly cumulated in order book's total size.

    Given the traders make the following deposits on asset's general account:
      | trader      | asset | amount       |
      | trader1     | ETH   | 10000000     |
      | trader2     | ETH   | 10000000     |
      | trader-lp-1 | ETH   | 100000000000 |
      | trader3     | ETH   | 1000000000   |

    # Trigger an auction to set the mark price
    When traders place the following orders:
      | trader  | market id | side | volume | price    | resulting trades | type       | tif     | reference |
      | trader1 | ETH/DEC19 | buy  | 1      | 11999980 | 0                | TYPE_LIMIT | TIF_GTC | trader1-1 |
      | trader2 | ETH/DEC19 | sell | 1      | 12000020 | 0                | TYPE_LIMIT | TIF_GTC | trader2-1 |
      | trader1 | ETH/DEC19 | buy  | 1      | 12000000 | 0                | TYPE_LIMIT | TIF_GFA | trader1-2 |
      | trader2 | ETH/DEC19 | sell | 1      | 12000000 | 0                | TYPE_LIMIT | TIF_GFA | trader2-2 |
    Then the opening auction period for market "ETH/DEC19" ends
    And the mark price for the market "ETH/DEC19" is "12000000"

    Then debug market data for "ETH/DEC19"


    When the trader submits LP:
      | id  | party       | market id | commitment amount | fee | order side | order reference | order proportion | order offset |reference |
      | lp1 | trader-lp-1 | ETH/DEC19 | 1000000000        | 0.1 | buy        | MID             | 1                | -10          | lp-1-ref |
      | lp1 | trader-lp-1 | ETH/DEC19 | 1000000000        | 0.1 | sell       | MID             | 1                | 10           | lp-1-ref |
    Then I see the LP events:
      | id  | party       | market    | commitment amount | status        |
      | lp1 | trader-lp-1 | ETH/DEC19 | 1000000000        | STATUS_ACTIVE |

    And the trading mode for the market "ETH/DEC19" is "TRADING_MODE_CONTINUOUS"

    And I see the following order events:
      | trader      | market id | side | volume | reference | offset | price    | status        |
      | trader-lp-1 | ETH/DEC19 | buy  | 167    |           | 0      | 11999990 | STATUS_ACTIVE |
      | trader-lp-1 | ETH/DEC19 | sell | 167    |           | 0      | 12000010 | STATUS_ACTIVE |

    When traders place the following orders:
      | trader  | market id | side | volume | price    | resulting trades | type       | tif     | reference |
      | trader3 | ETH/DEC19 | buy  | 167    | 11999990 | 0                | TYPE_LIMIT | TIF_GTC | trader3-1 |
      | trader3 | ETH/DEC19 | sell | 167    | 12000010 | 0                | TYPE_LIMIT | TIF_GTC | trader3-2 |

    Then there's the following volume on the book:
      | market id | side | price    | volume |
      | ETH/DEC19 | buy  | 11999990 | 334    |
      | ETH/DEC19 | sell | 12000010 | 334    |
