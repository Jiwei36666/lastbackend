//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package types

import (
	"encoding/json"
)

type AppList []*App

type App struct {
	Meta Meta `json:"meta"`
}

type AppCreateSpec struct {
	Meta Meta `json:"meta"`
}

type AppUpdateSpec struct {
	Meta Meta `json:"meta"`
}

func (p *App) ToJson() ([]byte, error) {
	buf, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (p *AppList) ToJson() ([]byte, error) {

	if p == nil {
		return []byte("[]"), nil
	}

	buf, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
