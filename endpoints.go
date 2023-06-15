// The MIT License (MIT)
//
// Copyright (c) 2021 Alexey Khan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package amocrm

import (
	"fmt"
	"strings"
)

type endpoint string

func (e endpoint) path() string {
	if e == eventsV2endpoint {
		return fmt.Sprintf("/api/v2/%s/", e)
	}

	return fmt.Sprintf("/api/v%d/%s", apiVersion, e)
}

const (
	accountsEndpoint endpoint = "accounts"

	leadsEndpoint     endpoint = "leads"
	leadEndpoint      endpoint = "leads/{id}"
	pipelinesEndpoint endpoint = "leads/pipelines"

	contactsEndpoint endpoint = "contacts"
	contactEndpoint  endpoint = "contacts/{id}"
)

func leadEndPoint(leadId int) endpoint {
	return endpoint(strings.Replace(string(leadEndpoint), "{id}", fmt.Sprintf("%d", leadId), 1))
}

func contactEndPoint(contactId int) endpoint {
	return endpoint(strings.Replace(string(contactEndpoint), "{id}", fmt.Sprintf("%d", contactId), 1))
}
