//go:generate protoc --go_out=plugins=grpc,paths=source_relative:. -I .:$GOPATH/src/gitlab.okta-solutions.com/mashroom/backend:$GOPATH/src/gitlab.okta-solutions.com/mashroom/backend/common/protobuf:vendor/gitlab.okta-solutions.com/mashroom/backend:vendor/gitlab.okta-solutions.com/mashroom/backend/common/protobuf api.proto
package zoopla

import (
	_ "gitlab.okta-solutions.com/mashroom/backend/common/protobuf/google/protobuf"
)