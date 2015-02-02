package v1beta1

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("v1beta1",
		&Role{},
		&RoleBinding{},
		&Policy{},
		&PolicyBinding{},
		&PolicyList{},
		&PolicyBindingList{},
	)
}

func (*Role) IsAnAPIObject()              {}
func (*Policy) IsAnAPIObject()            {}
func (*PolicyBinding) IsAnAPIObject()     {}
func (*RoleBinding) IsAnAPIObject()       {}
func (*PolicyList) IsAnAPIObject()        {}
func (*PolicyBindingList) IsAnAPIObject() {}
