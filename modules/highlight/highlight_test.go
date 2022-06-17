// Copyright 2021 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package highlight

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {

	tests := []struct {
		name string
		code string
		want []string
	}{
		{
			name: "empty.py",
			code: "",
			want: []string{""},
		},
		{
			name: "tags.txt",
			code: "<>",
			want: []string{"&lt;&gt;"},
		},
		{
			name: "tags.py",
			code: "<>",
			want: []string{`<span class="o">&lt;</span><span class="o">&gt;</span>`},
		},
		{
			name: "eol-no.py",
			code: "a=1",
			want: []string{`<span class="n">a</span><span class="o">=</span><span class="mi">1</span>`},
		},
		{
			name: "eol-newline1.py",
			code: "a=1\n",
			want: []string{
				`<span class="n">a</span><span class="o">=</span><span class="mi">1</span>&#10;`,
			},
		},
		{
			name: "eol-newline2.py",
			code: "a=1\n\n",
			want: []string{
				`<span class="n">a</span><span class="o">=</span><span class="mi">1</span>&#10;`,
				`&#10;`,
			},
		},
		{
			name: "empty-line.py",
			code: strings.TrimSpace(`
a=1

b=''
			`),
			want: []string{
				`<span class="n">a</span><span class="o">=</span><span class="mi">1</span>&#10;`,
				`&#10;`,
				`<span class="n">b</span><span class="o">=</span><span class="sa"></span><span class="s1">&#39;</span><span class="s1">&#39;</span>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := File(tt.name, "", []byte(tt.code))
			assert.NoError(t, err)
			assert.EqualValues(t, tt.want, lines)
		})
	}
}

func TestPlainText(t *testing.T) {

	tests := []struct {
		name string
		code string
		want []string
	}{
		{
			name: "empty.py",
			code: "",
			want: []string{""},
		},
		{
			name: "tags.py",
			code: "<>",
			want: []string{"&lt;&gt;"},
		},
		{
			name: "eol-no.py",
			code: "a=1",
			want: []string{`a=1`},
		},
		{
			name: "eol-newline1.py",
			code: "a=1\n",
			want: []string{
				`a=1&#10;`,
			},
		},
		{
			name: "eol-newline2.py",
			code: "a=1\n\n",
			want: []string{
				`a=1&#10;`,
				`&#10;`,
			},
		},
		{
			name: "empty-line.py",
			code: strings.TrimSpace(`
a=1

b=''
			`),
			want: []string{
				`a=1&#10;`,
				`&#10;`,
				`b=&#39;&#39;`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := PlainText([]byte(tt.code))
			assert.EqualValues(t, tt.want, lines)
		})
	}
}
