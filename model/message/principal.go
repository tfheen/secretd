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

type PrincipalListMessage struct {
	Action string `json:"action"`
}

func NewPrincipalListMessage() PrincipalListMessage {
	m := PrincipalListMessage{Action: "principal.list"}
	return m
}

type PrincipalListReplyMessage struct {
	Action     string   `json:"action"`
	Status     string   `json:"status"`
	Reason     string   `json:"reason,omitempty"`
	Principals []string `json:"principals"`
}

func NewPrincipalListReplyMessage(status string, principals []string) PrincipalListReplyMessage {
	m := PrincipalListReplyMessage{Action: "principal.list", Status: status, Principals: principals}
	return m
}

type PrincipalCreateMessage struct {
	Action      string `json:"action"`
	Principal   string `json:"principal"`
	SSHKey      string `json:"key"`
	Provisioned bool   `json:"provisioned"`
}

func NewPrincipalCreateMessage(principal string, sshKey string, provisioned bool) PrincipalCreateMessage {
	m := PrincipalCreateMessage{Action: "principal.create", Principal: principal, SSHKey: sshKey, Provisioned: provisioned}
	return m
}

type PrincipalCreateReplyMessage struct {
	Action     string   `json:"action"`
	Status     string   `json:"status"`
	Reason     string   `json:"reason,omitempty"`
	Principals []string `json:"principals"`
}

func NewPrincipalCreateReplyMessage(status string) PrincipalCreateReplyMessage {
	m := PrincipalCreateReplyMessage{Action: "principal.create", Status: status}
	return m
}

type PrincipalDeleteMessage struct {
	Action    string `json:"action"`
	Principal string `json:"principal"`
}

func NewPrincipalDeleteMessage(principal string) PrincipalDeleteMessage {
	m := PrincipalDeleteMessage{Action: "principal.delete", Principal: principal}
	return m
}

type PrincipalDeleteReplyMessage struct {
	Action string `json:"action"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

func NewPrincipalDeleteReplyMessage(status string) PrincipalDeleteReplyMessage {
	m := PrincipalDeleteReplyMessage{Action: "principal.delete", Status: status}
	return m
}
