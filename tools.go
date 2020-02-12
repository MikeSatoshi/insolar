// Copyright 2020 Insolar Network Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build tools

package tools

import _ "golang.org/x/tools/cmd/stringer"
import _ "github.com/gogo/protobuf/protoc-gen-gogoslick"
import _ "github.com/golang/protobuf/protoc-gen-go"
import _ "github.com/dgraph-io/badger/badger"
import _ "github.com/gojuno/minimock/v3/cmd/minimock"
