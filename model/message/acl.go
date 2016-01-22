/*
   Copyright 2015  Fastly Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package message

type AclGetMessage struct {
	Action string   `json:"action"`
	Key    []string `json:"key"`
}

func NewAclGetMessage(key []string) AclGetMessage {
	m := AclGetMessage{Action: "acl.get", Key: key}
	return m
}

type AclGetReplyMessage struct {
	Action string              `json:"action"`
	Status string              `json:"status"`
	Reason string              `json:"reason,omitempty"`
	Groups map[string][]string `json:"groups"`
}

func NewAclGetReplyMessage(status string, groups map[string][]string) AclGetReplyMessage {
	m := AclGetReplyMessage{Action: "acl.get", Status: status, Groups: groups}
	return m
}

type AclSetMessage struct {
	Action      string   `json:"action"`
	Key         []string `json:"key"`
	Group       string   `json:"group"`
	Permissions []string `json:"permissions"`
}

func NewAclSetMessage(key []string, group string, permissions []string) AclSetMessage {
	m := AclSetMessage{Action: "acl.set", Key: key, Group: group, Permissions: permissions}
	return m
}

type AclSetReplyMessage struct {
	Action string `json:"action"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func NewAclSetReplyMessage(status string) AclSetReplyMessage {
	m := AclSetReplyMessage{Action: "acl.set", Status: status}
	return m
}
