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

package server

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"io"
	message "model/message"
)

func GetMessage(r io.Reader) (ret message.GenericMessage, err error) {
	var rawmessage json.RawMessage
	var m message.GenericMessageJSON

	if err = json.NewDecoder(r).Decode(&rawmessage); err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawmessage, &m)
	if err != nil {
		return nil, err
	}
	switch m.Action {
	case "authorize":
		ret = new(message.AuthorizationMessage)
	case "secret.get":
		ret = new(message.SecretGetMessage)
	case "secret.put":
		ret = new(message.SecretPutMessage)
	case "secret.list":
		ret = new(message.SecretListMessage)
	case "group.list":
		ret = new(message.GroupListMessage)
	case "group.create":
		ret = new(message.GroupCreateMessage)
	case "group.delete":
		ret = new(message.GroupDeleteMessage)
	case "group.member_list":
		ret = new(message.GroupMemberListMessage)
	case "group.member_add":
		ret = new(message.GroupMemberAddMessage)
	case "group.member_remove":
		ret = new(message.GroupMemberRemoveMessage)
	case "principal.list":
		ret = new(message.PrincipalListMessage)
	case "principal.create":
		ret = new(message.PrincipalCreateMessage)
	case "principal.delete":
		ret = new(message.PrincipalDeleteMessage)
	case "acl.get":
		ret = new(message.AclGetMessage)
	case "acl.set":
		ret = new(message.AclSetMessage)
	case "enrol":
		ret = new(message.EnrolMessage)
	default:
		// XXX: handle this more gracefully
		panic("Unknown message type")
	}
	err = json.Unmarshal(rawmessage, &ret)
	spew.Dump(rawmessage, m)
	return ret, err
}

func SendReply(w io.Writer, reply message.GenericReply) (err error) {
	enc := json.NewEncoder(w)
	spew.Dump(reply)
	err = enc.Encode(reply)
	if err != nil {
		return err
	}
	return nil
}

func SendReplySimpleStatus(w io.Writer, status string) (err error) {
	reply := message.GenericReplyJSON{Status: status}
	return SendReply(w, reply)
}

func SendReplySimpleOK(w io.Writer) (err error) {
	reply := message.GenericReplyJSON{Status: "ok"}
	spew.Dump(reply)
	return SendReply(w, reply)
}

func SendReplySimpleError(w io.Writer, reason string) (err error) {
	reply := message.GenericReplyJSON{Status: "error", Reason: reason}
	return SendReply(w, reply)
}
