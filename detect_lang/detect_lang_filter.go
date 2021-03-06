//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package detect_lang

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"

	"github.com/blevesearch/cld2"
)

const FilterName = "detect_lang"

type DetectLangFilter struct {
}

func NewDetectLangFilter() *DetectLangFilter {
	return &DetectLangFilter{}
}

func (f *DetectLangFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0, len(input))

	offset := 0
	for _, token := range input {
		token.Term = []byte(cld2.Detect(string(token.Term)))
		token.Start = offset
		token.End = token.Start + len(token.Term)
		token.Type = analysis.AlphaNumeric
		rv = append(rv, token)
		offset = token.End + 1
	}

	return rv
}

func DetectLangFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	return NewDetectLangFilter(), nil
}

func init() {
	registry.RegisterTokenFilter(FilterName, DetectLangFilterConstructor)
}
