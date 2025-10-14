package db

import (
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
)

var SupaCli *supabase.Client

func SupaConnect() {
	var err error
	api_url := os.Getenv("API_URL")
	api_key := os.Getenv("API_KEY")

	SupaCli, err = supabase.NewClient(api_url, api_key, &supabase.ClientOptions{})

	if err != nil {
		log.Fatal("No se ha podido conectar a supabase... ", err)
	} else {
		log.Print("Se ha conectado a supabase...")
	}

}
