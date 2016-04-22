package stockfighter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	sf "github.com/wwkeyboard/stockfighter"
)

type TestDownloader struct {
	GetResp []byte
}

func (t TestDownloader) GetJSON(path string) ([]byte, error) {
	return t.GetResp, nil
}

func (t TestDownloader) PostJSON(path string, payload []byte) ([]byte, error) {
	return t.GetResp, nil
}

func TestHeartbeatUnpack(t *testing.T) {
	Convey("Given the heartbeat JSON", t, func() {
		json := `{
  "ok": true,
  "error": ""
}`
		downloader := TestDownloader{
			GetResp: []byte(json),
		}
		v := sf.Venue{
			Downloader: downloader,
			Name:       "TESTVENUE",
		}
		Convey("IsUp is True", func() {
			up, err := v.IsUP()
			if err != nil {
				t.Errorf("IsUP returned an error, %v", err)
			}
			if up != true {
				t.Error("IsUP didn't return true")
			}
		})

	})
}

func TestOrderBookUnpack(t *testing.T) {
	Convey("Given the orderbook JSON", t, func() {
		json := `{
    "ok": true,
    "venue": "TESTEX",
    "symbol": "FOOBAR",
    "ts": "2015-12-21T19:39:42.613795468Z",
    "bids": [
        {
            "price": 30640,
            "qty": 109,
            "isBuy": true
        }
    ],
    "asks": [
        {
            "price": 10000,
            "qty": 638,
            "isBuy": true
        }
    ]
}`

		downloader := TestDownloader{
			GetResp: []byte(json),
		}
		v := sf.Venue{
			Downloader: downloader,
			Name:       "TESTVENUE",
		}

		Convey("It extracts the bids", func() {
			bids, err := v.GetBids("stock")
			if err != nil {
				t.Errorf("GetBids failed with error, %v", err)
			}
			if len(bids) != 1 {
				t.Error("Didn't get the right number of bids")
			}
		})

	})
}

func TestQuoteUnpack(t *testing.T) {
	Convey("Given the Quote JSON", t, func() {
		json := `{
    "ok": true,
    "symbol": "FAC",
    "venue": "OGEX",
    "bid": 5100,
    "ask": 5125,
    "bidSize": 392,
    "askSize": 711,
    "bidDepth": 2748,
    "askDepth": 2237,
    "last": 5125,
    "lastSize": 52,
    "lastTrade": "2015-07-13T05:38:17.33640392Z",
    "quoteTime": "2015-07-13T05:38:17.33640392Z"
}`
		downloader := TestDownloader{
			GetResp: []byte(json),
		}
		v := sf.Venue{
			Downloader: downloader,
			Name:       "TESTVENUE",
		}
		Convey("ok is true", func() {
			up, err := v.GetQuote()
			if err != nil {
				t.Errorf("IsUP returned an error, %v", err)
			}
			if up != true {
				t.Error("IsUP didn't return true")
			}
		})

	})
}
