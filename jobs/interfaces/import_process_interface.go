// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package interfaces

import (
	"github.com/tetrafolium/mattermost-server/model"
)

type ImportProcessInterface interface {
	MakeWorker() model.Worker
}
