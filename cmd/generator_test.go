package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	type args struct {
		data data
		args arguments
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				data: data{
					fields: fields{
						{Accessor: "Names", Name: "Name", Type: "string"},
						{Accessor: "Ages", Name: "Age", Type: "int64"},
					},
					pkgName:   "user",
					sliceName: "Users",
				},
				args: arguments{
					importPaths: []importPath{
						{path: "time", alias: "alias_time"},
					},
				},
			},
			want: `// Code generated by "go-gen-slice-accessors"; DO NOT EDIT.
// Based on information from https://github.com/snamiki1212/go-gen-slice-accessors

package user

import (
  alias_time "time"
)

// Names
func (xs Users) Names() []string {
	sli := make([]string, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Name)
	}
	return sli
}

// Ages
func (xs Users) Ages() []int64 {
	sli := make([]int64, 0, len(xs))
	for i := range xs {
		sli = append(sli, xs[i].Age)
	}
	return sli
}
`,
		},
		{
			name: "ok: empty",
			args: args{
				data: data{
					fields:    fields{},
					pkgName:   "user",
					sliceName: "Users",
				},
			},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate(tt.args.data, tt.args.args)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
