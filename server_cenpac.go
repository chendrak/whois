package whois

import (
	"net/url"
	"strings"
)

type Cenpac struct{}

func (c *Cenpac) Resolve(req *Request) error {
	labels := strings.SplitN(req.Query, ".", 2)
	values := url.Values{}
	values.Set("subdomain", labels[0])
	values.Set("tld", labels[1])
	// "http://cenpac.net.nr/dns/whois.html?subdomain=domai&tld=nr"
	req.URL = "http://cenpac.net.nr/dns/whois.html?" + values.Encode()
	req.Body = ""
	return nil
}

func init() {
	register(&Cenpac{}, "cenpac.net.nr")
}
