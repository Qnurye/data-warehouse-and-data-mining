package data

import (
	"data-mining/pkg/base"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"reflect"
	"testing"
)

func Test_loadTransactions(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		want        []base.Transaction
		wantErr     bool
	}{
		{
			name:        "Test LoadTransactions",
			fileContent: "19 41 48 16430\n39 41 9150 10542\n48 592 766 8685 9925",
			want: []base.Transaction{
				{Set: mapset.NewSet("19", "41", "48", "16430")},
				{Set: mapset.NewSet("39", "41", "9150", "10542")},
				{Set: mapset.NewSet("48", "592", "766", "8685", "9925")},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				t.Errorf("Error creating file: %v", err)
			}
			_, err = file.WriteString(tt.fileContent)
			if err != nil {
				t.Errorf("Error writing to file: %v", err)
			}
			err = file.Close()
			if err != nil {
				t.Errorf("Error closing file: %v", err)
			}
			got, err := LoadTransactions("test.txt")
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadTransactions() got = %v, want %v", got, tt.want)
			}
			err = os.Remove("test.txt")
			if err != nil {
				t.Errorf("Error removing file: %v", err)
			}
		})
	}
}
