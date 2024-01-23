package dependencies

import "encoding/xml"

type Data struct {
	Option Option `json:"options"`
	Sucess bool   `json:"success"`
}
type Option struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type MPD struct {
	XMLName                   xml.Name           `xml:"MPD"`
	MinBufferTime             string             `xml:"minBufferTime,attr"`
	Type                      string             `xml:"type,attr"`
	MediaPresentationDuration string             `xml:"mediaPresentationDuration,attr"`
	MaxSegmentDuration        string             `xml:"maxSegmentDuration,attr"`
	Profiles                  string             `xml:"profiles,attr"`
	ProgramInformation        ProgramInformation `xml:"ProgramInformation"`
	Periods                   []Period           `xml:"Period"`
}

type ProgramInformation struct {
	MoreInformationURL string `xml:"moreInformationURL,attr"`
	Title              string `xml:"Title"`
}

type Period struct {
	Duration       string          `xml:"duration,attr"`
	AdaptationSets []AdaptationSet `xml:"AdaptationSet"`
}

type AdaptationSet struct {
	SegmentAlignment string           `xml:"segmentAlignment,attr"`
	Group            string           `xml:"group,attr"`
	MaxWidth         string           `xml:"maxWidth,attr,omitempty"`
	MaxHeight        string           `xml:"maxHeight,attr,omitempty"`
	MaxFrameRate     string           `xml:"maxFrameRate,attr,omitempty"`
	Par              string           `xml:"par,attr,omitempty"`
	Lang             string           `xml:"lang,attr,omitempty"`
	Representations  []Representation `xml:"Representation"`
}

type Representation struct {
	ID           string       `xml:"id,attr"`
	MimeType     string       `xml:"mimeType,attr"`
	Codecs       string       `xml:"codecs,attr"`
	Width        string       `xml:"width,attr,omitempty"`
	Height       string       `xml:"height,attr,omitempty"`
	FrameRate    string       `xml:"frameRate,attr,omitempty"`
	Sar          string       `xml:"sar,attr,omitempty"`
	StartWithSAP string       `xml:"startWithSAP,attr,omitempty"`
	Bandwidth    string       `xml:"bandwidth,attr,omitempty"`
	BaseURL      string       `xml:"BaseURL"`
	SegmentList  SegmentList  `xml:"SegmentList"`
	AudioConfig  *AudioConfig `xml:"AudioChannelConfiguration,omitempty"`
}

type SegmentList struct {
	Timescale      string         `xml:"timescale,attr"`
	Duration       string         `xml:"duration,attr"`
	Initialization Initialization `xml:"Initialization"`
	SegmentURLs    []SegmentURL   `xml:"SegmentURL"`
}

type Initialization struct {
	Range string `xml:"range,attr"`
}

type SegmentURL struct {
	MediaRange string `xml:"mediaRange,attr"`
	IndexRange string `xml:"indexRange,attr,omitempty"`
}

type AudioConfig struct {
	SchemeIdUri string `xml:"schemeIdUri,attr"`
	Value       string `xml:"value,attr"`
}
