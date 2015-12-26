package venue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestDownloader struct {
	GetResp []byte
}

func (t TestDownloader) GetJSON(path string) ([]byte, error) {
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
		v := Venue{
			Downloader: downloader,
			Name:       "TESTVENUE",
		}
		Convey("IsUp is True", func() {
			up, err := v.IsUP()
			if err != nil {
				t.Errorf("IsUP returned an error, %e", err)
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
		v := Venue{
			Downloader: downloader,
			Name:       "TESTVENUE",
		}

		Convey("It extracts the bids", func() {
			bids, err := v.GetBids("stock")
			if err != nil {
				t.Errorf("GetBids failed with error, %e", err)
			}
			if len(bids) != 1 {
				t.Error("Didn't get the right number of bids")
			}
		})

	})
}
