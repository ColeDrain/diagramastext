package core

import (
	"reflect"
	"testing"
)

func TestCode2Path(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: `@startuml
    a -> b
@enduml`,
			args: args{
				s: `@startuml
    a -> b
@enduml`,
			},
			want:    "SoWkIImgAStDuL80WaG5NJk592w7rBmKe100",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got, _ := code2Path(tt.args.s); got != tt.want {
					t.Errorf("code2Path() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_compress(t *testing.T) {
	type args struct {
		v []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "foo",
			args: args{
				v: []byte("foo"),
			},
			want:    []byte{75, 203, 207, 7, 0},
			wantErr: false,
		},
		{
			name: "foobar",
			args: args{
				v: []byte("foobar"),
			},
			want:    []byte{75, 203, 207, 79, 74, 44, 2, 0},
			wantErr: false,
		},
		{
			name: "@startuml",
			args: args{
				v: []byte(`@startuml`),
			},
			want:    []byte{115, 40, 46, 73, 44, 42, 41, 205, 205, 1, 0},
			wantErr: false,
		},
		{
			name: `foo
bar`,
			args: args{
				v: []byte(`foo
bar`),
			},
			want:    []byte{75, 203, 207, 231, 74, 74, 44, 2, 0},
			wantErr: false,
		},
		{
			name: "->",
			args: args{
				v: []byte(`->`),
			},
			want:    []byte{211, 181, 3, 0},
			wantErr: false,
		},
		{
			name: "a->b",
			args: args{
				v: []byte("a->b"),
			},
			want:    []byte{75, 212, 181, 75, 2, 0},
			wantErr: false,
		},
		{
			name: "a -> b",
			args: args{
				v: []byte("a -> b"),
			},
			want:    []byte{75, 84, 208, 181, 83, 72, 2, 0},
			wantErr: false,
		},
		{
			name: `@startuml
    a -> b
@enduml`,
			args: args{
				v: []byte(`@startuml
    a -> b
@enduml`),
			},
			want: []byte{
				115, 40, 46, 73, 44, 42, 41, 205, 205, 225, 82, 0, 130, 68, 5, 93, 59, 133, 36, 46, 135, 212, 188, 20,
				160, 16, 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := compress(tt.args.v)
				if (err != nil) != tt.wantErr {
					t.Errorf("compress() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("compress() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
