// Code generated by goctl. DO NOT EDIT.
package types

type Empty struct {
}

type UpsertServiceConfigReq struct {
	Key        string `json:"key"`         //交易对
	ConfigData string `json:"config_data"` //配置
}

type GetServiceConfigReq struct {
	Key string `json:"key"`
}

type GetServiceConfigResp struct {
	ConfigData string `json:"config_data"`
}

type UploadTemplateFileReq struct {
	ServiceName string `form:"service_name"` //服务名
	Symbol      string `form:"symbol"`       //交易对
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Username string `json:"username"`
	UserId   uint32 `json:"user_id"`
}

type LoginResp struct {
	Token    string    `json:"token"`
	Expire   int64     `json:"expire"`
	UserInfo *UserInfo `json:"user_info"`
}

type RegisterReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"comfirm_password"`
}

type GetSymbolListReq struct {
	PageNo   int64 `form:"page_no"`
	PageSize int64 `form:"page_size"`
}

type SymbolInfo struct {
	ID            uint32 `json:"id"`
	SymbolName    string `json:"symbol_name"`
	SymbolId      int32  `json:"symbol_id"`
	BaseCoinID    uint32 `json:"base_coin_id"`
	BaseCoinName  string `json:"base_coin_name"`
	BaseCoinPrec  int32  `json:"base_coin_prec"`
	QuoteCoinID   int32  `json:"quote_coin_id"`
	QuoteCoinName string `json:"quote_coin_name"`
	QuotePrec     int32  `json:"quote_prec"`
}

type GetSymbolListResp struct {
	List  []*SymbolInfo `json:"list"`
	Total int64         `json:"total"`
}

type AddSymbolReq struct {
	BaseCoinID  uint32 `json:"base_coin_id"`
	QuoteCoinID int32  `json:"quote_coin_id"`
	SymbolId    int32  `json:"symbol_id"`
}

type AddSymbolResp struct {
}

type AddCoinReq struct {
	CoinName string `json:"coin_name"`
	Prec     int32  `json:"prec"`
	CoinId   int32  `json:"coin_id"`
}

type AddCoinResp struct {
}

type GetCoinListReq struct {
	PageNo   int64 `form:"page_no"`
	PageSize int64 `form:"page_size"`
}

type CoinInfo struct {
	ID       uint32 `json:"id"`
	CoinName string `json:"coin_name"`
	CoinId   int32  `json:"coin_id"`
	Prec     int32  `json:"prec"`
}

type GetCoinListResp struct {
	List  []*CoinInfo `json:"list"`
	Total int64       `json:"total"`
}

type UpdateSymbolReq struct {
	SymbolName    string `json:"symbol_name"`
	BaseCoinID    uint32 `json:"base_coin_id"`
	BaseCoinName  string `json:"base_coin_name"`
	BaseCoinPrec  int32  `json:"base_coin_prec"`
	QuoteCoinID   int32  `json:"quote_coin_id"`
	QuoteCoinName string `json:"quote_coin_name"`
	QuotePrec     int32  `json:"quote_prec"`
}

type UpdateSymbolResp struct {
}

type AddErrorCodeReq struct {
	ErrorCodeName string `json:"error_code_name"`
	ErrorCodeId   int32  `json:"error_code_id"`
	Language      string `json:"language"`
}

type GetErrorCodeListReq struct {
	Language string `form:"language"`
}

type ErrorCode struct {
	Id            uint32 `json:"id"`
	ErrorCodeName string `json:"error_code_name"`
	ErrorCodeId   int32  `json:"error_code_id"`
	Language      string `json:"language"`
}

type GetErrorCodeListResp struct {
	List  []*ErrorCode `json:"list"`
	Total int64        `json:"total"`
}

type UpdateErrorCodeReq struct {
	Id            uint32 `json:"id"`
	ErrorCodeName string `json:"error_code_name"`
	ErrorCodeId   int32  `json:"error_code_id"`
	Language      string `json:"language"`
}

type UpdateCoinReq struct {
	ID       uint32 `json:"id"`
	CoinName string `json:"coin_name"`
	Prec     int32  `json:"prec"`
	CoinId   int32  `json:"coin_id"`
}
