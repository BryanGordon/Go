package connections

import (
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func SupaConnec() {
	var err error
	api_url := os.Getenv("API_URL")
	api_key := os.Getenv("API_KEY")
	Client, err = supabase.NewClient(api_url, api_key, &supabase.ClientOptions{})

	if err != nil {
		log.Fatal("No se ha podido conectar al servidor", err)
	} else {
		log.Print("Conectado a supabase")
	}
}
