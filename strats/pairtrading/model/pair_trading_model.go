package model

import (
	"github.com/HaoxuanXu/TradingBot/strats/pairtrading/updater"
	"github.com/HaoxuanXu/TradingBot/tools/readwrite"
	"github.com/HaoxuanXu/TradingBot/tools/repeater"
)

type PairTradingModel struct {
	ExpensiveStockSymbol                              string
	CheapStockSymbol                                  string
	EntryNetValue                                     float64
	ExitNetValue                                      float64
	LoserNums                                         int
	MinProfitThreshold                                float64
	PriceRatioThreshold                               float64
	CheapStockEntryVolume                             float64
	ExpensiveStockEntryVolume                         float64
	ExpensiveStockFilledQuantity                      float64
	CheapStockFilledQuantity                          float64
	ExpensiveStockFilledPrice                         float64
	CheapStockFilledPrice                             float64
	ExpensiveStockShortQuotePrice                     float64
	ExpensiveStockLongQuotePrice                      float64
	CheapStockShortQuotePrice                         float64
	CheapStockLongQuotePrice                          float64
	IsShortExpensiveStockLongCheapStock               bool
	IsLongExpensiveStockShortCheapStock               bool
	ShortExpensiveStockLongCheapStockPriceRatio       float64
	LongExpensiveStockShortCheapStockPriceRatio       float64
	ShortExpensiveStockLongCheapStockPreviousRatio    float64
	LongExpensiveStockShortCheapStockPreviousRatio    float64
	ShortExpensiveStockLongCheapStockRepeatNumber     int
	LongExpensiveStockShortCheapStockRepeatNumber     int
	ShortExpensiveStockLongCheapStockPriceRatioRecord []float64
	LongExpensiveStockShortCheapStockPriceRatioRecord []float64
	RepeatArray                                       []int
	RepeatNumThreshold                                int
	DefaultRepeatArrayLength                          int
	DefaultPriceRatioArrayLength                      int
}

func (model *PairTradingModel) getStockSymbols(assetType string) (string, string) {
	if assetType == "gold" {
		return "GLD", "IAU"
	}
	return "", ""
}

func GetModel(assetType, shortLongPath, longShortPath, repeatNumPath string) *PairTradingModel {
	dataModel := &PairTradingModel{}
	dataModel.initialize(assetType, shortLongPath, longShortPath, repeatNumPath)
	return dataModel
}

func (model *PairTradingModel) initialize(assetType, shortLongPath, longShortPath, repeatNumPath string) {
	model.ExpensiveStockSymbol, model.CheapStockSymbol = model.getStockSymbols(assetType)
	model.ShortExpensiveStockLongCheapStockPriceRatioRecord = readwrite.ReadRecordFloat(shortLongPath)
	model.LongExpensiveStockShortCheapStockPriceRatioRecord = readwrite.ReadRecordFloat(longShortPath)
	model.RepeatArray = readwrite.ReadRecordInt(repeatNumPath)
	model.ShortExpensiveStockLongCheapStockRepeatNumber = 0
	model.LongExpensiveStockShortCheapStockRepeatNumber = 0
	model.LongExpensiveStockShortCheapStockPriceRatio = 0.0
	model.ShortExpensiveStockLongCheapStockPriceRatio = 0.0
	model.LongExpensiveStockShortCheapStockPreviousRatio = 0.0
	model.ShortExpensiveStockLongCheapStockPreviousRatio = 0.0
	model.IsLongExpensiveStockShortCheapStock = false
	model.IsShortExpensiveStockLongCheapStock = false
	model.CheapStockLongQuotePrice = 0.0
	model.CheapStockShortQuotePrice = 0.0
	model.ExpensiveStockLongQuotePrice = 0.0
	model.ExpensiveStockShortQuotePrice = 0.0
	model.CheapStockFilledPrice = 0.0
	model.ExpensiveStockFilledPrice = 0.0
	model.CheapStockFilledQuantity = 0.0
	model.ExpensiveStockFilledQuantity = 0.0
	model.ExpensiveStockEntryVolume = 0.0
	model.CheapStockEntryVolume = 0.0
	model.PriceRatioThreshold = updater.UpdatePriceRatioThreshold(
		model.LongExpensiveStockShortCheapStockPriceRatioRecord,
		model.ShortExpensiveStockLongCheapStockPriceRatioRecord,
	)
	model.RepeatNumThreshold = repeater.CalculateOptimalRepeatNum(model.RepeatArray)
	model.DefaultRepeatArrayLength = len(model.RepeatArray)
	model.DefaultPriceRatioArrayLength = len(model.ShortExpensiveStockLongCheapStockPriceRatioRecord)
	model.EntryNetValue = 0.0
	model.ExitNetValue = 0.0
	model.LoserNums = 0
}
