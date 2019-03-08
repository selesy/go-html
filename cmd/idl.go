package main

import (
	"net/http"

	"github.com/selesy/go-robot/filter"
	"github.com/selesy/go-robot/scrape"
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
	Description string
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

	var webidlurls []MozFolder
	for _, n := range trows {
		path := n.FirstChild
		desc := path.NextSibling
		date := desc.NextSibling
		size := date.NextSibling
		f := MozFolder{
			Path:        scrape.Text(path),
			Description: scrape.Text(desc),
			Date:        scrape.Text(date),
			Size:        scrape.Text(size),
		}
		webidlurls = append(webidlurls, f)
	}
	log.Info(webidlurls)

	// tbody, tr, td/td/td/td =

	return nil
}
