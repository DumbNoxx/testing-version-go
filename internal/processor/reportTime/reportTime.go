package reporttime

import (
	"time"

	"github.com/DumbNoxx/testing-version-go/internal/options"
	"github.com/DumbNoxx/testing-version-go/internal/processor"
	"github.com/DumbNoxx/testing-version-go/internal/utils"
)

func GetReportFileTime() {
	processor.TickerReportFile.Stop()
	processor.TickerReportFile.Reset(utils.TimeReportFile)
}

func GetReportPartialTime() {
	processor.TimeReport = time.Duration(options.Config.ReportInterval * float64(time.Minute))
	processor.Ticker.Stop()
	processor.Ticker.Reset(processor.TimeReport)
}
