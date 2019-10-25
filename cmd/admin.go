package cmd

import (
	"adm/app/models"
	"errors"
	"github.com/spf13/cobra"
	"log"
)

var name, email, password string

func init() {
	RootCmd.AddCommand(adminCmd)
	adminCmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	adminCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
	adminCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Manage the admin user",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if name == "" || email == "" || password == "" {
			return errors.New("name, email and password are required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		admin := models.AdminUser(name, email, password)
		err := admin.Create()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Admin user <%s> created successfully.", name)
	},
}
