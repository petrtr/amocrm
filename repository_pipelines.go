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
	"io"
	"net/http"
)

type PipelineEmbedded struct {
	Statuses []PipelineStatus `json:"statuses,omitempty"`
}

type PipelineStatus struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Sort       int    `json:"sort,omitempty"`
	IsEditable bool   `json:"is_editable,omitempty"`
	PipelineId int    `json:"pipeline_id,omitempty"`
	Color      string `json:"color,omitempty"`
	Type       int    `json:"type,omitempty"`
	AccountId  int    `json:"account_id,omitempty"`
}

type Pipeline struct {
	Id           int               `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Sort         int               `json:"price,omitempty"`
	IsMain       bool              `json:"is_main,omitempty"`        //Defines whether the pipeline is main for the account
	IsUnsortedOn bool              `json:"is_unsorted_on,omitempty"` //Defines whether Incoming Leads are enabled
	IsArchive    bool              `json:"is_archive,omitempty"`     //Defines whether Incoming Leads are enabled
	AccountId    int               `json:"account_id,omitempty"`
	Embedded     *PipelineEmbedded `json:"_embedded,omitempty"` //Данные вложенных сущностей
}

// Pipelines describes methods available for Pipelines entity.
type Pipelines interface {
	List() ([]Pipeline, error)
}

// Verify interface compliance.
var _ Pipelines = pipelines{}

type pipelines struct {
	api *api
}

func newPipelines(api *api) Pipelines {
	return pipelines{api: api}
}

// Pipelines list
func (a pipelines) List() ([]Pipeline, error) {

	resp, rErr := a.api.do(pipelinesEndpoint, http.MethodGet, nil, nil, nil)
	if rErr != nil {
		return nil, fmt.Errorf("get pipelines: %w", rErr)
	}

	var res struct {
		Embedded struct {
			Pipelines []Pipeline `json:"pipelines"`
		} `json:"_embedded"`
	}
	err := a.api.read(resp, &res)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return res.Embedded.Pipelines, nil
}
