//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

package capmodel

import (
	"reflect"
	"testing"

	"github.com/ODIM-Project/PluginCiscoACI/capdata"
	"github.com/ODIM-Project/PluginCiscoACI/db"
)

func TestGetZone(t *testing.T) {
	db.Connector = mockConnector{}
	type args struct {
		zoneID string
	}
	tests := []struct {
		name    string
		args    args
		want    *capdata.ZoneData
		wantErr bool
	}{
		{
			name: "successful get on port",
			args: args{
				zoneID: "validID",
			},
			want: &capdata.ZoneData{
				FabricID: "validID",
			},
			wantErr: false,
		},
		{
			name: "failed get on port",
			args: args{
				zoneID: "invalidID",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetZone(tt.args.zoneID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetZone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllZones(t *testing.T) {
	db.Connector = mockConnector{}
	tests := []struct {
		name    string
		want    []capdata.ZoneData
		wantErr bool
	}{
		{
			name: "successful get on zone collection",
			want: []capdata.ZoneData{
				capdata.ZoneData{
					FabricID: "validID",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllZones("somePattern")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllZones() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllZones() = %v, want %v", got, tt.want)
			}
		})
	}
}
