// Copyright 2018 The ChubaoFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package objectnode

import (
	"fmt"
	"testing"
	"time"
)

func TestXmlMarshal_CopyObjectResult(t *testing.T) {
	result := &CopyObjectResult{
		ETag:         "etag_value",
		LastModified: time.Now().Format("2006-01-02 15:04:05"),
	}
	if marshaled, marshalErr := MarshalXMLEntity(result); marshalErr != nil {
		t.Fatalf("marshal fail cause: %v", marshalErr)
	} else {
		t.Logf("marshal result: %v", string(marshaled))
	}
}

func TestXmlMarshal_DeleteResult(t *testing.T) {
	result := DeleteResult{
		Deleted: []Deleted{
			{Key: "sample1.txt"},
		},
		Error: []Error{
			{Key: "sample2.txt", Code: "AccessDenied", Message: "Access Denied"},
		},
	}
	if marshaled, marshalErr := MarshalXMLEntity(result); marshalErr != nil {
		t.Fatalf("marshal fail cause: %v", marshalErr)
	} else {
		t.Logf("marshal result: %v", string(marshaled))
	}
}

func TestXmlMarshal_InitMultipartUpload(t *testing.T) {
	initResult := InitMultipartResult{
		Bucket:   "ngwCloud1oss",
		Key:      "4989/txt/5678.txt",
		UploadId: "UZSEAFV367N8BZP5LQXVHFDACLTX9HXX",
	}

	if bytes, err := MarshalXMLEntity(initResult); err != nil {
		t.Fatalf("marshal fail cause: %v", err)
	} else {
		t.Logf("marshal result: %v", string(bytes))
	}
}

func TestXMLMarshal_ListPartsResult(t *testing.T) {
	owner := &BucketOwner{
		ID:          "YLWBsakx5hJK4cO4NcwyE72hA9KTGQQ3",
		DisplayName: "YLWBsakx5hJK4cO4NcwyE72hA9KTGQQ3",
	}
	parts := []*Part{
		{
			PartNumber:   1,
			LastModified: "Tue, 20 Aug 2019 07:29:33 GMT",
			ETag:         "d8e2155e77cebd8fa1c3ab77da7c2ca8",
			Size:         12582912,
		},
		{
			PartNumber:   2,
			LastModified: "Tue, 20 Aug 2019 16:09:15 GMT",
			ETag:         "9a7909810df6cde3dedaa06966db5a56",
			Size:         12582912,
		},
	}

	listPartsResult := ListPartsResult{
		Bucket:           "ngwCloud1oss",
		Key:              "4989/txt/5678.txt",
		UploadId:         "UZSEAFV367N8BZP5LQXVHFDACLTX9HXX",
		StorageClass:     "Standard",
		PartNumberMarker: 1,
		NextMarker:       3,
		MaxParts:         2,
		IsTruncated:      true,
		Parts:            parts,
		Owner:            owner,
	}

	if bytes, err := MarshalXMLEntity(listPartsResult); err != nil {
		t.Fatalf("marshal fail cause: %v", err)
	} else {
		t.Logf("marshal result: %v", string(bytes))
	}
}

func TestXMLMarshal_ListUploadsResult(t *testing.T) {
	accessKey := "YLWBsakx5hJK4cO4NcwyE72hA9KTGQQ3"

	fsUploads := []*FSUpload{
		{
			Key:          "mybatis-11.pdf",
			UploadId:     "8378DFB508AE393AAAXGXTYU28",
			StorageClass: "Standard",
			Initiated:    "2018-09-30T19:08:42.000Z",
		},
		{
			Key:          "books/mybatis/mybatis-11.pdf",
			UploadId:     "B86949F2376C6B8BKOINTQ123E",
			StorageClass: "Standard",
			Initiated:    "2018-09-30T19:08:42.000Z",
		},
	}
	uploads := NewUploads(fsUploads, accessKey)

	listUploadsResult := ListUploadsResult{
		Bucket:             "ngwCloud1oss",
		KeyMarker:          "",
		UploadIdMarker:     "",
		NextKeyMarker:      "4989/txt/5678.txt",
		NextUploadIdMarker: "UZSEAFV367N8BZP5LQXVHFDACLTX9HXX",
		Delimiter:          "",
		Prefix:             "4789",
		MaxUploads:         1000,
		IsTruncated:        false,
		Uploads:            uploads,
		CommonPrefixes:     nil,
	}

	if bytes, err := MarshalXMLEntity(listUploadsResult); err != nil {
		t.Fatalf("marshal fail cause: %v", err)
	} else {
		t.Logf("marshal result: %v", string(bytes))
	}
}

func TestNewDeleteRequest(t *testing.T) {
	objectRequest := []Object{
		{
			Key: "jvsTest001_1",
			//VersionId:"v0001",
		},
		{
			Key: "jvsTest001_2",
			//VersionId:"v0001",
		},
		{
			Key: "jvsTest001_3",
			//VersionId:"v0001",
		},
	}

	deleteRequest := DeleteRequest{
		Objects: objectRequest,
	}

	bytes, err := MarshalXMLEntity(deleteRequest)
	if err != nil {
		t.Fatalf("marshal fail cause: %v", err)
	} else {
		t.Logf("marshal result: %v", string(bytes))
	}
}

func TestUnmarshalDeleteRequest(t *testing.T) {
	source := `
<Delete>
  <Object>
    <Key>jvsTest001_1</Key>
  </Object>
  <Object>
    <Key>jvsTest001_2</Key>
  </Object>
  <Object>
    <Key>jvsTest001_3</Key>
  </Object>
</Delete>
	`
	deleteReq := DeleteRequest{}
	err := UnmarshalXMLEntity([]byte(source), &deleteReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("----", deleteReq.Objects)
}

func TestMarshalTagging(t *testing.T) {
	tagging := &Tagging{
		TagSet: []*Tag{
			{
				Key:   "tag1",
				Value: "val1",
			},
			{
				Key:   "tag2",
				Value: "val2",
			},
		},
	}
	marshaled, err := MarshalXMLEntity(tagging)
	if err != nil {
		t.Fatalf("marshal tagging fail: err(%v)", err)
	}
	t.Logf("marshal tagging:\n%v", string(marshaled))
}