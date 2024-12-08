package day4

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_fiveZeroMD5(t *testing.T) {
	type args struct {
		secretKey string
		prefix    string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				secretKey: "abcdef",
				prefix:    "00000",
			},
			want: 609043,
		},
		{
			args: args{
				secretKey: "pqrstuv",
				prefix:    "00000",
			},
			want: 1048970,
		},
		{
			args: args{
				secretKey: common.Input(),
				prefix:    "00000",
			},
			want: 117946,
		},
		{
			args: args{
				secretKey: common.Input(),
				prefix:    "000000",
			},
			want: 3938038,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := md5WithPrefix(tt.args.secretKey, tt.args.prefix); got != tt.want {
				t.Errorf("md5WithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
