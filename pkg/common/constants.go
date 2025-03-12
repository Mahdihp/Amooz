package common

//type Collection string

const (
	Exchnage_Bybit   string = "bybit"
	Exchnage_Chainup string = "chainup"

	Collection_CryptoCurrency     string = "CryptoCurrency"
	Collection_FearGreed          string = "FearGreed"
	Collection_CoinMarketCapMap   string = "CoinMarketCapMap"
	Collection_ListingsLatest_CMC string = "ListingsLatestCMC"
	Collection_BullishBearish     string = "BullishBearish"
	Collection_PairSpot           string = "PairSpot"
	Collection_TrendingLatest     string = "TrendingLatest"
	Collection_LocalCurrency      string = "LocalCurrency"
	Collection_CryptoInfo         string = "CryptoInfo"
)

const (
	BaseRoute string = "/api"
	VerRoute  string = BaseRoute + "/v1"

	//OrderBook string = VerRoute + "/orderbook"
	//Order     string = VerRoute + "/order"
	//History   string = VerRoute + "/history"

	Market_All            string = "all"
	Market_RiskLimit      string = "risklimit"
	Market_InstrumentInfo string = "instrumentinfo"
	Market_Ticker         string = "ticker"

	Market_Spot    string = "spot"
	Market_Linear  string = "linear"
	Market_Inverse string = "inverse"
	Market_Option  string = "option"
)

const (
	And_Opt                 string = "$and" // and
	Or_Opt                  string = "$or"  // or
	Equal_Opt               string = "$eq"  // ==
	GreaterThan_Opt         string = "$gt"  // >
	GreaterThanEqual_Opt    string = "$gte" // >=
	LessThan_Opt            string = "$lt"  // <
	LessThanEqual_Opt       string = "$lte" // <=
	LessThanEqual_Not_Equal string = "$ne"  // <=

	Field_Search_Id           string = "id"
	Field_Search_ID           string = "ID"
	Field_Search_Symbol       string = "symbol"
	Field_Search_ApiKey       string = "apiKey"
	Field_Search_SecretKey    string = "apiSecret"
	Field_Search_UserId       string = "userId"
	Field_Search_ExecTime     string = "execTime"
	Field_Search_CreatedTime  string = "createdTime"
	Field_Search_LastUpdated  string = "lastUpdated"
	Field_Search_Category     string = "category"
	Field_Search_OrderId      string = "orderID"
	Field_Search_ExecId       string = "execID"
	Field_Search_QuoteTxID    string = "quoteTxID"
	Field_Search_BaseCoin     string = "baseCoin"
	Field_Search_Exchange     string = "exchange"
	Field_Search_UpdateTime   string = "updateTime"
	Field_Search_Bullish      string = "bullish"
	Field_Search_Bearish      string = "bearish"
	Field_Search_LastPrice    string = "lastPrice"
	Field_Search_Price24HPcnt string = "price24HPcnt"
	Field_Search_Rank         string = "rank"
)

const (
	KEY_DEVELOPMENT      string = "development"
	KEY_CODE_ZERO        int32  = 0
	KEY_CODE_NEGATIVE    int32  = -1
	KEY_NOT_FOUND        string = "not found"
	KEY_SYMBOL_NOT_FOUND string = "symbol not found"
	KEY_OPT_NOT_SUCCESS  string = "operation not success"
	KEY_OK               string = "ok"
	KEY_SUCCESS          string = "success"
	KEY_FAIL             string = "fail"
	KEY_TIMEOUT          string = "timeout"
	ZIPKIN_SPAN_ERROR    string = "error"
	ZIPKIN_SPAN_INFO     string = "info"
)

const (
	TOPIC_ORDER_SPOT_HISTORY    string = "topic-order-spot-history"
	TOPIC_ORDER_LINEAR_HISTORY  string = "topic-order-linear-history"
	TOPIC_TRADE_SPOT_HISTORY    string = "topic-trade-spot-history"
	TOPIC_TRADE_LINEAR_HISTORY  string = "topic-trade-linear-history"
	TOPIC_PNL_HISTORY           string = "topic-pnl-history"
	TOPIC_REQUEST_QUOTE_HISTORY string = "topic-request_quote"
)
