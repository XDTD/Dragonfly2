/*
 *     Copyright 2022 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package objectstorage

import "mime/multipart"

type ObjectParams struct {
	ID        string `uri:"id" binding:"required"`
	ObjectKey string `uri:"object_key" binding:"required"`
}

type PutObjectRequset struct {
	Mode        uint                  `form:"mode,default=0" binding:"omitempty,gte=0,lte=2"`
	Filter      string                `form:"filter" binding:"omitempty"`
	MaxReplicas int                   `form:"maxReplicas" binding:"omitempty,gt=0,lte=100"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
}

type GetObjectQuery struct {
	Filter string `form:"filter" binding:"omitempty"`
}

type GetObjectsListQuery struct {
	// A delimiter is a character you use to group keys.
	Delimiter string `form:"delimiter" binding:"omitempty"`

	// Marker is where you want  to start listing from. Marker can be any key in the bucket.
	Marker string `form:"marker" binding:"omitempty"`

	// Sets the maximum number of keys returned in the response.
	Limit int64 `form:"limit" binding:"omitempty"`

	// Limits the response to keys that begin with the specified prefix.
	Prefix string `form:"prefix" binding:"omitempty"`
}

type CopyObjectRequest struct {
	// Source is a required field
	Source string `form:"file" binding:"required"`

	// CopySource is a required field
	Destination string `form:"file" binding:"required"`
}
