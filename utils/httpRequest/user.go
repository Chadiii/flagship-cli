package httprequest

import (
	"net/http"

	"github.com/Chadiii/flagship/models"
	"github.com/Chadiii/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListUsers() ([]models.User, error) {
	return HTTPGetAllPages[models.User](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/users")
}

func HTTPBatchUpdateUsers(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPut, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users", []byte(data))
}

func HTTPDeleteUsers(email string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+email, nil)
	return err
}