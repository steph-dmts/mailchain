// nolint: dupl
package settings

import (
	"github.com/mailchain/mailchain"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/output"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/values"
	"github.com/mailchain/mailchain/internal/protocols/ethereum"
	"github.com/mailchain/mailchain/nameservice"
	"github.com/pkg/errors"
)

func domainNameServices(s values.Store) *DomainNameServices {
	return &DomainNameServices{
		clients: map[string]NameServiceDomainClient{
			mailchain.Mailchain: mailchainDomainNameServices(s),
		},
	}
}

type DomainNameServices struct {
	clients map[string]NameServiceDomainClient
}

func (s DomainNameServices) Output() output.Element {
	elements := []output.Element{}
	for _, c := range s.clients {
		elements = append(elements, c.Output())
	}
	return output.Element{
		FullName: "nameservice-domain-name",
		Elements: elements,
	}
}

func (s DomainNameServices) Produce(client string) (nameservice.ForwardLookup, error) {
	if client == "" {
		return nil, nil
	}
	m, ok := s.clients[client]
	if !ok {
		return nil, errors.Errorf("%s not a supported address name service", client)
	}
	return m.Produce()
}

func mailchainDomainNameServices(s values.Store) *MailchainDomainNameServices {
	enabledNetworks := []string{}
	for _, n := range ethereum.Networks() {
		enabledNetworks = append(enabledNetworks, "ethereum/"+n)
	}
	return &MailchainDomainNameServices{
		BaseURL: values.NewDefaultString("https://ns.mailchain.xyz/", s, "nameservice-domain-name.base-url"),
		EnabledProtocolNetworks: values.NewDefaultStringSlice(
			enabledNetworks,
			s,
			"nameservice-domain-name.mailchain.enabled-networks",
		),
	}
}

type MailchainDomainNameServices struct {
	BaseURL                 values.String
	EnabledProtocolNetworks values.StringSlice
}

func (s MailchainDomainNameServices) Produce() (nameservice.ForwardLookup, error) {
	return nameservice.NewLookupService(s.BaseURL.Get()), nil
}

func (s MailchainDomainNameServices) Supports() map[string]bool {
	m := map[string]bool{}
	for _, np := range s.EnabledProtocolNetworks.Get() {
		m[np] = true
	}
	return m
}

func (s MailchainDomainNameServices) Output() output.Element {
	return output.Element{
		FullName: "nameservice-domain-name.mailchain",
		Attributes: []output.Attribute{
			s.BaseURL.Attribute(),
			s.EnabledProtocolNetworks.Attribute(),
		},
	}
}
