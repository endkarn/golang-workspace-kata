package selling_ticket

import (
	"testing"
)

func TestBuyTicket(t *testing.T) {
	type args struct {
		ticketType int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "ซื้อขัตรเด็ก ในวันจันทร์ที่ 30 เมษายน เวลา 10:30",
			args: args{
				ticketType: 1,
			},
			want:    199,
			wantErr: false,
		},
		{
			name: "ซื้อขัตรผู้ใหญ่ ในวันอาทิตย์ที่ 3 พฤษภาคม เวลา 13:30",
			args: args{
				ticketType: 2,
			},
			want:    259,
			wantErr: false,
		},
		{
			name: "ซื้อขัตรผู้ใหญ่ ในวันอเสาร์ที่ 2 พฤษภาคม เวลา 11:30",
			args: args{
				ticketType: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "ซื้อขัตรเด็ก ในวันอังคารที่ 14 เมษายน เวลา 12:30",
			args: args{
				ticketType: 1,
			},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuyTicket(tt.args.ticketType)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuyTicket() got = %v, want %v", got, tt.want)
			}
		})
	}
}
