# lyant - Learn You A Next Thing

Bored? Want to learn something?

Added github login, to run the project just open `app->controller->app.go`

``` func GetDefaults() Defaults {
	defaults := Defaults{}
	defaults.github_client_id = "GITHUB CLIENT ID"      //SET YOUR GITHUB CLIENT ID
	defaults.github_client_secret = "GITHUB SECRET KEY" //SET SECRET KEY FOR GITHUB ACCOUNT
	defaults.github_redirect_url = "http://localhost:9000/callback"
	defaults.github_oAuth_url = "https://github.com/login/oauth/access_token"
	defaults.github_api_user = "https://api.github.com/user"
	defaults.github_login_url = "https://github.com/login/oauth/authorize?scope=user:email&client_id="
	return defaults
} ```

Enter your creadentials and run the project , open `localhost:9000` and check console after login with github to check your username
