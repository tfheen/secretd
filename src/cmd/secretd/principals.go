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

package main

import (
	"database/sql"
	"log"
)

func listPrincipals(db *sql.DB, principal string) (principals []string, err error) {
	if err = CheckAclNonHierarchical(db, principal, "principal_manage"); err != nil {
		return
	}
	rows, err := db.Query("SELECT name FROM principals")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var principal string
		err = rows.Scan(&principal)
		if err != nil {
			log.Fatal(err)
		}
		principals = append(principals, principal)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func createPrincipal(db *sql.DB, principal, newPrincipal, SSHKey string, provisioned bool) (err error) {
	if err = CheckAclNonHierarchical(db, principal, "principal_manage"); err != nil {
		return
	}

	_, err = db.Exec("INSERT INTO principals(name, ssh_key, provisioned) VALUES ($1, $2, $3)", newPrincipal, SSHKey, provisioned)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func deletePrincipal(db *sql.DB, principal, deletePrincipal string) (err error) {
	if err = CheckAclNonHierarchical(db, principal, "principal_manage"); err != nil {
		return
	}

	_, err = db.Exec("DELETE FROM principals WHERE name = $1", deletePrincipal)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
