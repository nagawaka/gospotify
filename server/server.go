package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"github.com/spf13/viper"
	"spoti/crawl/auth"
)

// func completer(d prompt.Document) []prompt.Suggest {
// 	s := []prompt.Suggest{
// 		{Text: "users", Description: "Store the username and age"},
// 		{Text: "articles", Description: "Store the article text posted by user"},
// 		{Text: "comments", Description: "Store the text commented to articles"},
// 	}
// 	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
// }

func Start() {
	http.HandleFunc("/", auth.Redirect)
	http.HandleFunc("/callback", auth.Callback)

	viper.SetDefault("CLIENT_STATE", "BpLnfgDsc2WD8F2q")

	err := http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("PORT")), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}