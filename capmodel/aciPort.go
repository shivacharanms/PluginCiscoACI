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
	"encoding/json"
	"fmt"

	dmtf "github.com/ODIM-Project/ODIM/lib-dmtf/model"
	"github.com/ODIM-Project/PluginCiscoACI/db"
)

//PortCollectionResponse ...
type PortCollectionResponse struct {
	TotalCount string                 `json:"totalCount"`
	IMData     []PortCollectionIMData `json:"imdata"`
}

// PortCollectionIMData ...
type PortCollectionIMData struct {
	PhysicalInterface PhysicalInterface `json:"l1PhysIf"`
}

// PhysicalInterface ...
type PhysicalInterface struct {
	Attributes map[string]interface{} `json:"attributes"`
}

// PortInfoResponse ...
type PortInfoResponse struct {
	TotalCount string           `json:"totalCount"`
	IMData     []PortInfoIMData `json:"imdata"`
}

//PortInfoIMData ...
type PortInfoIMData struct {
	PhysicalInterface PhysicalInterface `json:"ethpmPhysIf"`
}

// GetPort collects the port data from the DB
func GetPort(portID string) (*dmtf.Port, error) {
	var port dmtf.Port
	data, err := db.Connector.Get(TablePort, portID)
	if err != nil {
		return nil, fmt.Errorf("while trying to collect port data, got: %w", err)
	}
	err = json.Unmarshal([]byte(data), &port)
	if err != nil {
		return nil, fmt.Errorf("while trying to unmarshal port data, got: %v", err)
	}
	return &port, nil
}
