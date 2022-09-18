package provider

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"	

	"github.com/snowflakedb/gosnowflake"
)

type SnowflakeClient struct {
	Database *sql.DB	
	AWSId string
	AWSFederatedToken string
}

type Credentials struct {
	Account string
	User string
	Password string
	Region string
}

type PrivatelinkConfig struct {
	Privatelink_account_name           string   `json:"privatelink-account-name"`
	Privatelink_internal_stage		   string   `json:"privatelink-internal-stage"`
	Privatelink_vpce_id                string   `json:"privatelink-vpce-id"`
	Privatelink_account_url            string   `json:"privatelink-account-url"`
	Regionless_privatelink_account_url string   `json:"regionless-privatelink-account-url"`
	Privatelink_ocsp_url               string   `json:"privatelink_ocsp-url"`
	Privatelink_connection_urls        any 		`json:"privatelink-connection-urls"`
}

func (c *SnowflakeClient) NewSnowflakeClient(cred Credentials)(*SnowflakeClient, error) {
	dns, err := gosnowflake.DSN(&gosnowflake.Config{
		Account:   cred.Account,
		User:      cred.User,
		Password:  cred.Password,
		Role:      "accountadmin",
		Region:    cred.Region,
	})
	
	if err != nil {
		log.Fatal("Error while DNS string: ", err)
		
		return nil, err
	}

	db, err := sql.Open("snowflake", dns)
	if err != nil {
		log.Fatal("Error while open DB: ", err)
		return nil, err
	}
	c = &SnowflakeClient{}
	c.Database = db
	return c, nil
}

func (c *SnowflakeClient) EnableInternalStagesForPrivatelink(enableInternalStage bool)(error) {
	db := *c.Database
	defer db.Close()
	sql := fmt.Sprintf("alter account set enable_internal_stages_privatelink = %t;", enableInternalStage)
	_, err := db.Exec(sql)
	
	return err
}

 func (c *SnowflakeClient) GetPrivatelinkConfig()(*PrivatelinkConfig, error) {
	db := *c.Database
	defer db.Close()
	rows, err := db.Query("select SYSTEM$GET_PRIVATELINK_CONFIG();")
	if err != nil {
		log.Fatal("Error executing open 'select SYSTEM$GET_PRIVATELINK_CONFIG()': ", err)
		return nil, err
	}
	var s sql.NullString
	for rows.Next() {
		err := rows.Scan(&s)
		if err != nil {
			log.Fatalf("Failed to scan. err: %v", err)
		}
		if !s.Valid {
			fmt.Println("Retrieved value: NULL")
		}
	}
	privatelink_config_string := s.String
	privatelink_config := PrivatelinkConfig{}
	err = json.Unmarshal([]byte(privatelink_config_string), &privatelink_config)

	if err != nil {
		log.Fatal("Could not unmarshall json results.",err)
	}
	return &privatelink_config, err
}


func (c *SnowflakeClient) GetPrivatelink()(string, error) {
	result := ""
	db := *c.Database
	defer db.Close()
	sql := fmt.Sprintf("select SYSTEM$GET_PRIVATELINK('%s','%s')", c.AWSId, c.AWSFederatedToken)
	
	row := db.QueryRow(sql)

	if err := row.Scan(&result); err != nil {
	    return result, err
    }
	return result, nil
}

func (c *SnowflakeClient) RevokePrivatelink()(error) {
	db := *c.Database
	defer db.Close()
	sql := fmt.Sprintf("select SYSTEM$REVOKE_PRIVATELINK('%s','%s')", c.AWSId, c.AWSFederatedToken)
	_, err := db.Exec(sql)
	
	return err
}

func (c *SnowflakeClient) AuthorizePrivatelink()(error) {
	db := *c.Database
	defer db.Close()
	sql := fmt.Sprintf("select SYSTEM$AUTHORIZE_PRIVATELINK('%s','%s')",c.AWSId, c.AWSFederatedToken)
	_, err := db.Exec(sql)
	
	return err
}