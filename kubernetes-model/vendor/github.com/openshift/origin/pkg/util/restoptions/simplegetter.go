/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package restoptions

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage/storagebackend"
)

type simpleGetter struct {
	storage *storagebackend.Config
}

func NewSimpleGetter(storage *storagebackend.Config) Getter {
	return &simpleGetter{storage: storage}
}

func (s *simpleGetter) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	return generic.RESTOptions{
		StorageConfig:           s.storage,
		Decorator:               generic.UndecoratedStorage,
		DeleteCollectionWorkers: 1,
		ResourcePrefix:          resource.Resource,
	}, nil
}
