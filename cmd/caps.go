/*
Copyright © 2022 Tarmo Katmuk <tarmo.katmuk@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/spf13/cobra"
)

// capsCmd represents the caps command
var (
	capsCmd = &cobra.Command{
		Use:   "caps",
		Short: "User Capabilities operations",
		Long:  `User Capabilities operations`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add user capabilities",
		Long: `Add user capabilities in form 

"users|buckets=*|read|write|read,write"

Add multiple capabilities to user:

--caps "buckets=*,users=read"`,
		Run: func(cmd *cobra.Command, args []string) {
			access := &Access{
				Ceph:         cephHost,
				AccessKey:    cephAccessKey,
				AccessSecret: cephAccessSecret,
			}

			user := &User{
				ID:          userName,
				DisplayName: userFullname,
				Email:       userEmail,
				UserCaps:    userCaps,
			}
			if user.ID == "" || user.UserCaps == "" {
				cmd.Help()
				os.Exit(1)
			}

			addUserCaps(*access, *user)
		},
	}
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove user capabilities",
		Long: `Remove user capabilities in form 

"users|buckets=*|read|write|read,write"

Remove multiple capabilities to user:

--caps "buckets=*,users=read"`,
		Run: func(cmd *cobra.Command, args []string) {
			access := &Access{
				Ceph:         cephHost,
				AccessKey:    cephAccessKey,
				AccessSecret: cephAccessSecret,
			}

			user := &User{
				ID:          userName,
				DisplayName: userFullname,
				Email:       userEmail,
				UserCaps:    userCaps,
			}

			if user.ID == "" || user.UserCaps == "" {
				cmd.Help()
				os.Exit(1)
			}

			removeUserCaps(*access, *user)
		},
	}
)

func init() {
	userCmd.AddCommand(capsCmd)
	capsCmd.AddCommand(addCmd)
	capsCmd.AddCommand(removeCmd)
	userCmd.MarkFlagRequired("user")
	addCmd.MarkFlagRequired("user")
	removeCmd.MarkFlagRequired("user")
	userCmd.MarkFlagRequired("caps")
}

func addUserCaps(ceph Access, user User) {
	c, err := admin.New(ceph.Ceph, ceph.AccessKey, ceph.AccessSecret, nil)
	if err != nil {
		panic(err)
	}

	userCaps, err := c.AddUserCap(context.Background(), user.ID, user.UserCaps)

	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Println(userCaps)
}

func removeUserCaps(ceph Access, user User) {
	c, err := admin.New(ceph.Ceph, ceph.AccessKey, ceph.AccessSecret, nil)
	if err != nil {
		panic(err)
	}

	userCaps, err := c.RemoveUserCap(context.Background(), user.ID, user.UserCaps)

	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Println(userCaps)

}
