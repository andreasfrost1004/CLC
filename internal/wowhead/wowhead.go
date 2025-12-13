package wowhead

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WowheadItemXML struct {
	Item struct {
		ID   int    `xml:"id,attr"`
		Name string `xml:"name"`
	} `xml:"item"`
}

func FetchItemXML(itemID int) (*WowheadItemXML, error) {
	url := fmt.Sprintf("https://classic.wowhead.com/item=%d&xml", itemID)

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wowhead returned %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data WowheadItemXML
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
