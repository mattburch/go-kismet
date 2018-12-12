/*Package kismet parses Kismet XML data into a similary formed struct.*/

package kismet

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
)

type KismetRun struct {
	Version    string     `xml:"kismet-version,attr"`
	StartTime  string     `xml:"start-time,attr"`
	CardSource CardSource `xml:"card-source"`
	Networks   []Network  `xml:"wireless-network"`
}

type CardSource struct {
	UUID     string `xml:"uuid,attr"`
	Source   string `xml:"card-source"`
	Name     string `xml:"card-name"`
	Type     string `xml:"card-type"`
	Packets  int    `xml:"card-packets"`
	Hop      bool   `xml:"card-hop"`
	Channels string `xml:"card-channels"`
}

type Network struct {
	Number       string    `xml:"number,attr"`
	Type         string    `xml:"type,attr"`
	FirstTime    string    `xml:"first-time,attr"`
	LastTime     string    `xml:"last-time,attr"`
	SSID         SSID      `xml:"SSID"`
	BSSID        string    `xml:"BSSID"`
	Manufacturer string    `xml:"manuf"`
	Channel      int       `xml:"channel"`
	MHZ          string    `xml:"freqmhz"`
	MaxSeen      int       `xml:"maxseenrate"`
	Packets      Packets   `xml:"packets"`
	DataSize     int       `xml:"datasize"`
	SNRInfo      SNRInfo   `xml:"snr-info"`
	BSSTimeStamp string    `xml:"bsstimestamp"`
	CDPDevice    string    `xml:"cdp-device"`
	CDPPortID    string    `xml:"cdp-portid"`
	SeenCard     SeenCard  `xml:"seen-card"`
	Clients      []Clients `xml:"wireless-clients"`
}

type SSID struct {
	FirstTime  string   `xml:"first-time,attr"`
	LastTime   string   `xml:"last-time,attr"`
	Type       string   `xml:"type"`
	MaxRate    string   `xml:"max-rate"`
	Packets    int      `xml:"packets"`
	Encryption []string `xml:"encryption"`
	ESSID      ESSID    `xml:"essid"`
}

type ESSID struct {
	Cloaked bool   `xml:"cloaked,attr"`
	ESSID   string `xml:"essid"`
}

type Packets struct {
	LLC       int `xml:"LLC"`
	Data      int `xml:"data"`
	Crypt     int `xml:"crypt"`
	Total     int `xml:"total"`
	Fragments int `xml:"fragments"`
	Retries   int `xml:"retries"`
}

type SNRInfo struct {
	LastSigDBM    int `xml:"last_signal_dbm"`
	LastNoiseDBM  int `xml:"last_noise_dbm"`
	LastSigRSSI   int `xml:"last_signal_rssi"`
	LastNoiseRSSI int `xml:"last_noise_rssi"`
	MinSigDBM     int `xml:"min_signal_dbm"`
	MinNoiseDBM   int `xml:"min_noise_dbm"`
	MinSignalRSSI int `xml:"min_signal_rssi"`
	MinNoiseRSSI  int `xml:"min_noise_rssi"`
	MaxSigDBM     int `xml:"max_signal_dbm"`
	MaxNoiseDBM   int `xml:"max_noise_dbm"`
	MaxSigRSSI    int `xml:"max_signal_rssi"`
	MaxNoiseRSSI  int `xml:"max_noise_rssi"`
}

type SeenCard struct {
	SeenUUID    string `xml:"seen-uuid"`
	SeenTime    string `xml:"seen-time"`
	SeenPackets int    `xml:"seen-packets"`
}

type Clients struct {
	ClientMAC   string   `xml:"client-mac"`
	ClientManuf string   `xml:"client-manuf"`
	Channel     int      `xml:"channel"`
	FreqMHZ     string   `xml:"freqmhz"`
	MaxSeenRate int      `xml:"maxseenrate"`
	Packets     Packets  `xml:"packets"`
	DataSize    int      `xml:"datasize"`
	SNRInfo     SNRInfo  `xml:"snr-info"`
	SeenCard    SeenCard `xml:"snr-info"`
}

// Parse takes a byte array of sslscan xml data and unmarshals it into an
// KismetRun struct. All elements are returned as strings, it is up to the caller
// to check and cast them to the proper type.
func Parse(content io.Reader) (*KismetRun, error) {
	r := &KismetRun{}
	// Convert Charset to UTF-8
	decoder := xml.NewDecoder(content)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
