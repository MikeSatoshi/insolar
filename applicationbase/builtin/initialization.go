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

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package builtin

import (
	"github.com/pkg/errors"

	nodedomain "github.com/insolar/insolar/applicationbase/builtin/contract/nodedomain"
	noderecord "github.com/insolar/insolar/applicationbase/builtin/contract/noderecord"

	XXX_insolar "github.com/insolar/insolar/insolar"
	XXX_artifacts "github.com/insolar/insolar/logicrunner/artifacts"
)

func InitializeContractMethods() map[string]XXX_insolar.ContractWrapper {
	return map[string]XXX_insolar.ContractWrapper{
		"nodedomain": nodedomain.Initialize(),
		"noderecord": noderecord.Initialize(),
	}
}

func shouldLoadRef(strRef string) XXX_insolar.Reference {
	ref, err := XXX_insolar.NewReferenceFromString(strRef)
	if err != nil {
		panic(errors.Wrap(err, "Unexpected error, bailing out"))
	}
	return *ref
}

func InitializeCodeRefs() map[XXX_insolar.Reference]string {
	rv := make(map[XXX_insolar.Reference]string, 2)

	rv[shouldLoadRef("insolar:0AAABAq5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record")] = "nodedomain"
	rv[shouldLoadRef("insolar:0AAABAvLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record")] = "noderecord"

	return rv
}

func InitializePrototypeRefs() map[XXX_insolar.Reference]string {
	rv := make(map[XXX_insolar.Reference]string, 2)

	rv[shouldLoadRef("insolar:0AAABAkocNP8SpY6g890ZsRwVOqLADBviGimy2cm_x60")] = "nodedomain"
	rv[shouldLoadRef("insolar:0AAABAgXJhmV8uwhpxIEfL7hqjD1wQUGg8SArUa0VOAc")] = "noderecord"

	return rv
}

func InitializeCodeDescriptors() []XXX_artifacts.CodeDescriptor {
	rv := make([]XXX_artifacts.CodeDescriptor, 0, 2)

	// nodedomain
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAABAq5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record"),
	))
	// noderecord
	rv = append(rv, XXX_artifacts.NewCodeDescriptor(
		/* code:        */ nil,
		/* machineType: */ XXX_insolar.MachineTypeBuiltin,
		/* ref:         */ shouldLoadRef("insolar:0AAABAvLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record"),
	))

	return rv
}

func InitializePrototypeDescriptors() []XXX_artifacts.PrototypeDescriptor {
	rv := make([]XXX_artifacts.PrototypeDescriptor, 0, 2)

	{ // nodedomain
		pRef := shouldLoadRef("insolar:0AAABAkocNP8SpY6g890ZsRwVOqLADBviGimy2cm_x60")
		cRef := shouldLoadRef("insolar:0AAABAq5GWKE7v1W8gHxS2BzsokOe1vgl-WaKyOMLQhs.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	{ // noderecord
		pRef := shouldLoadRef("insolar:0AAABAgXJhmV8uwhpxIEfL7hqjD1wQUGg8SArUa0VOAc")
		cRef := shouldLoadRef("insolar:0AAABAvLOOIFkH6ikCcIZLil_HvpvwXFMxHvvyDwq8ls.record")
		rv = append(rv, XXX_artifacts.NewPrototypeDescriptor(
			/* head:         */ pRef,
			/* state:        */ *pRef.GetLocal(),
			/* code:         */ cRef,
		))
	}

	return rv
}
