package chromedphelper

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/huynhminhtruong/go-instagram-downloader/internal/config"
)

func StartSessionWithCookies(postURL string, cookies []config.Cookie) ([]string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageLinks []string

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if e, ok := ev.(*network.EventResponseReceived); ok {
			url := e.Response.URL
			contentType := e.Response.MimeType

			if strings.HasPrefix(contentType, "image/") {
				imageLinks = append(imageLinks, url)
			}
		}
	})

	if err := chromedp.Run(ctx, network.Enable()); err != nil {
		return nil, err
	}

	// Add cookies to session
	var cookieParams []*network.CookieParam
	for _, c := range cookies {
		log.Print("Adding cookie: ", c.Name, "=", c.Value, " for domain: ", c.Domain)
		cookieParams = append(cookieParams, &network.CookieParam{
			Name:   c.Name,
			Value:  c.Value,
			Domain: c.Domain,
			Path:   c.Path,
		})
	}
	if err := chromedp.Run(ctx, network.SetCookies(cookieParams)); err != nil {
		return nil, err
	}

	// Navigate and wait
	if err := chromedp.Run(ctx,
		chromedp.Navigate(postURL),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		return nil, err
	}

	return imageLinks, nil
}
