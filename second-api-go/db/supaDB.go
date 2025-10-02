package db

import (
	"log"

	"github.com/supabase-community/supabase-go"
)

var api_url = "https://mjrqcemffmprtozbqmhi.supabase.co"
var SupaCli *supabase.Client

func SupaConnect() {
	var err error
	SupaCli, err = supabase.NewClient(api_url, api_key, &supabase.ClientOptions{})

	if err != nil {
		log.Fatal("No se ha podido conectar a supabase... ", err)
	} else {
		log.Print("Se ha conectado a supabase...")
	}

}
