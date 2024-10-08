# finance-type-site

## tradestie.com

exposes a public api endpoint `https://tradestie.com/api/v1/apps/reddit` that returns the top 50 stocks being discussed on r/WallstreetBets.

example response:

```json
[
    {
        "no_of_comments": 361,
        "sentiment": "Bullish",
        "sentiment_score": 0.038,
        "ticker": "NVDA"
    },
    {
        "no_of_comments": 50,
        "sentiment": "Bullish",
        "sentiment_score": 0.185,
        "ticker": "AI"
    },
    {
        "no_of_comments": 31,
        "sentiment": "Bullish",
        "sentiment_score": 0.007,
        "ticker": "SMCI"
    },
    {
        "no_of_comments": 22,
        "sentiment": "Bullish",
        "sentiment_score": 0.036,
        "ticker": "AAPL"
    },
    {
        "no_of_comments": 16,
        "sentiment": "Bullish",
        "sentiment_score": 0.095,
        "ticker": "INTC"
    },
    {
        "no_of_comments": 14,
        "sentiment": "Bearish",
        "sentiment_score": -0.046,
        "ticker": "DG"
    },
    {
        "no_of_comments": 10,
        "sentiment": "Bearish",
        "sentiment_score": -0.102,
        "ticker": "QQQ"
    },
    {
        "no_of_comments": 8,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "OXY"
    },
    {
        "no_of_comments": 7,
        "sentiment": "Bearish",
        "sentiment_score": -0.098,
        "ticker": "AMD"
    },
    {
        "no_of_comments": 4,
        "sentiment": "Bearish",
        "sentiment_score": -0.243,
        "ticker": "TSLA"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bullish",
        "sentiment_score": 0.251,
        "ticker": "TQQQ"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bullish",
        "sentiment_score": 0.116,
        "ticker": "BBY"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bullish",
        "sentiment_score": 0.14,
        "ticker": "EOD"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bearish",
        "sentiment_score": -0.099,
        "ticker": "CVNA"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "VS"
    },
    {
        "no_of_comments": 3,
        "sentiment": "Bullish",
        "sentiment_score": 0.173,
        "ticker": "TLT"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bearish",
        "sentiment_score": -0.279,
        "ticker": "ANF"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.06,
        "ticker": "AFRM"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.433,
        "ticker": "PLTR"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.357,
        "ticker": "MSFT"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "CORE"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.811,
        "ticker": "LULU"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.167,
        "ticker": "GOOD"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.236,
        "ticker": "CHWY"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.038,
        "ticker": "RTX"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.229,
        "ticker": "CC"
    },
    {
        "no_of_comments": 2,
        "sentiment": "Bullish",
        "sentiment_score": 0.734,
        "ticker": "MC"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "SQQQ"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.34,
        "ticker": "UPS"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "AAP"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "REAL"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.552,
        "ticker": "RIVN"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.629,
        "ticker": "SMH"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.691,
        "ticker": "BY"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.945,
        "ticker": "ONE"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bullish",
        "sentiment_score": 0.527,
        "ticker": "GE"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "LEAP"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bullish",
        "sentiment_score": 0.3,
        "ticker": "TSM"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bullish",
        "sentiment_score": 0.224,
        "ticker": "HEAR"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.459,
        "ticker": "UK"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "COST"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.282,
        "ticker": "TGT"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "TA"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": -0.417,
        "ticker": "DARE"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bullish",
        "sentiment_score": 0.891,
        "ticker": "ROIC"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "ON"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "MDB"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "OKTA"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "EYE"
    },
    {
        "no_of_comments": 1,
        "sentiment": "Bearish",
        "sentiment_score": 0.0,
        "ticker": "CME"
    }
]
```

tradestie provides [this dashboard](https://tradestie.com/apps/reddit/wallstreetbets/) displaying the data and utilizing other endpoints to display samples of the comments being analyzed.

- [reliability information](https://www.freepublicapis.com/reddit-stocks)
- [api documentation](https://tradestie.com/apps/reddit/api/)