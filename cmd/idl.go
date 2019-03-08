package main

import (
	"net/http"

	"github.com/selesy/go-html/filter"
	"github.com/selesy/go-html/scrape"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const w3c = "http://w3c.github.io/html/single-page.html"
const whatwg = "https://html.spec.whatwg.org/#the-body-element"

const mozFolderURL = "https://dxr.mozilla.org/mozilla-central/source/dom/webidl/"
const mozTableClass = "folder-content"

type MozFolder struct {
	Path        string
	Pescription string
	Date        string
	Size        string
}

func mozHTML() error {
	resp, err := http.Get(mozFolderURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info(resp.StatusCode)

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	table, ok := scrape.Find(doc, filter.Type(html.ElementNode), filter.Tag(atom.Table), filter.Class(mozTableClass))
	if !ok {
		log.Error("Failed to find the Mozilla data table")
		return nil
	}

	tbody, ok := scrape.Find(table, filter.Tag(atom.Tbody))
	if !ok {
		log.Error("Failed to find the Mozilla data table body")
		return nil
	}

	trows := scrape.FindAll(tbody, filter.Tag(atom.Tr))

	log.Info("Row count:", len(trows))

	// tbody, tr, td/td/td/td =

	return nil
}
