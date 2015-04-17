/*Package kismet parses Kismet XML data into a similary formed struct.*/

package kismet

import (
	"encoding/xml"
)

type KismetRun struct {
	Version    string     `json:"kismet-version,attr"`
	StartTime  string     `json:"start-time,attr"`
	CardSource CardSource `json:"card-source"`
	Networks   []Network  `json:"wireless-network"`
}

type CardSource struct {
	UUID     string `json:"uuid,attr"`
	Source   string `json:"card-source"`
	Name     string `json:"card-name"`
	Type     string `json:"card-type"`
	Packets  int    `json:"card-packets"`
	Hop      bool   `json:"card-hop"`
	Channels string `json:"card-channels"`
}

type Network struct {
	Number       string    `json:"number,attr"`
	Type         string    `json:"type,attr"`
	FirstTime    string    `json:"first-time,attr"`
	LastTime     string    `json:"last-time,attr"`
	SSID         SSID      `json:"SSID"`
	BSSID        string    `json:"BSSID"`
	Manufacturer string    `json:"manuf"`
	Channel      int       `json:"channel"`
	MHZ          string    `json:"freqmhz"`
	MaxSeen      int       `json:"maxseenrate"`
	Packets      Packets   `json:"packets"`
	DataSize     int       `json:"datasize"`
	SNRInfo      SNRInfo   `json:"snr-info"`
	BSSTimeStamp string    `json:"bsstimestamp"`
	CDPDevice    string    `json:"cdp-device"`
	CDPPortID    string    `json:"cdp-portid"`
	SeenCard     SeenCard  `json:"seen-card"`
	Clients      []Clients `json:"wireless-clients"`
}

type SSID struct {
	FirstTime  string   `json:"first-time,attr"`
	LastTime   string   `json:"last-time,attr"`
	Type       string   `json:"type"`
	MaxRate    string   `json:"max-rate"`
	Packets    int      `json:"packets"`
	Encryption []string `json:"encryption"`
	ESSID      ESSID    `json:"essid"`
}

type ESSID struct {
	Cloaked bool   `json:"cloaked,attr"`
	ESSID   string `json:"essid"`
}

type Packets struct {
	LLC       int `json:"LLC"`
	Data      int `json:"data"`
	Crypt     int `json:"crypt"`
	Total     int `json:"total"`
	Fragments int `json:"fragments"`
	Retries   int `json:"retries"`
}

type SNRInfo struct {
	LastSigDBM    uint `json:"last_signal_dbm"`
	LastNoiseDBM  uint `json:"last_noise_dbm"`
	LastSigRSSI   uint `json:"last_signal_rssi"`
	LastNoiseRSSI uint `json:"last_noise_rssi"`
	MinSigDBM     uint `json:"min_signal_dbm"`
	MinNoiseDBM   uint `json:"min_noise_dbm"`
	MinSignalRSSI uint `json:"min_signal_rssi"`
	MinNoiseRSSI  uint `json:"min_noise_rssi"`
	MaxSigDBM     uint `json:"max_signal_dbm"`
	MaxNoiseDBM   uint `json:"max_noise_dbm"`
	MaxSigRSSI    uint `json:"max_signal_rssi"`
	MaxNoiseRSSI  uint `json:"max_noise_rssi"`
}

type SeenCard struct {
	SeenUUID    string `json:"seen-uuid"`
	SeenTime    string `json:"seen-time"`
	SeenPackets int    `json:"seen-packets"`
}

type Clients struct {
	ClientMAC   string   `json:"client-mac"`
	ClientManuf string   `json:"client-manuf"`
	Channel     int      `json:"channel"`
	FreqMHZ     string   `json:"freqmhz"`
	MaxSeenRate int      `json:"maxseenrate"`
	Packets     Packets  `json:"packets"`
	DataSize    int      `json:"datasize"`
	SNRInfo     SNRInfo  `json:"snr-info"`
	SeenCard    SeenCard `json:"snr-info"`
}

// Parse takes a byte array of sslscan xml data and unmarshals it into an
// KismetRun struct. All elements are returned as strings, it is up to the caller
// to check and cast them to the proper type.
func Parse(content []byte) (*KismetRun, error) {
	r := &KismetRun{}
	err := xml.Unmarshal(content, r)
	if err != nil {
		return r, err
	}
	return r, nil
}
