package accounts

import "testing"

func TestGetRoleForAccount(t *testing.T) {
	type args struct {
		acct string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoleForAccount(tt.args.acct)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoleForAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRoleForAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
