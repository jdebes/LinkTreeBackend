package model

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// TODO Change the types to use their descriptive strings instead of enum value.
type LinkAsset struct {
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	PlatformID int64  `json:"platformId"`
}

type Link struct {
	Type       int         `json:"linkType"`
	LinkAssets []LinkAsset `json:"linkAssets"`
}

func (s Link) Valid() error {
	// TODO Implement better and extendable validation based on matching the type text
	if s.Type > 0 && s.Type < 2 {
		return errors.New("a link must have at least 1 resource")
	}

	// TODO enhance validation to ensure different link types have LinkAssets of expected type.
	// E.g we expect a classic link to have Resource type asset.
	// E.g we expected music link to have an Embedded type asset.
	if len(s.LinkAssets) == 0 {
		return errors.New("a link must have at least 1 resource")
	}

	for i, value := range s.LinkAssets {
		name := strings.TrimSpace(value.Name)
		if len(name) == 0 || utf8.RuneCountInString(name) > 144 {
			return errors.New("name must be less than 144 characters and cannot be blank")
		}
		s.LinkAssets[i].Name = name

		// TODO better validation to url to ensure it has valid form with url.Parse() or similar.
		url := strings.TrimSpace(value.Url)
		if len(url) == 0 || utf8.RuneCountInString(url) > 2048 {
			return errors.New("name must be less than 144 characters and cannot be blank")
		}
		s.LinkAssets[i].Url = url

		// TODO Implement better and extendable validation based on matching the type text
		if value.Type > 0 && value.Type < 2 {
			return errors.New("a link must have at least 1 resource")
		}
	}

	return nil
}
