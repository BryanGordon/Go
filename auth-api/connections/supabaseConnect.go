package connections

import (
	"log"

	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func SupaConnec() {
	var err error
	Client, err = supabase.NewClient("", "", &supabase.ClientOptions{})

	if err != nil {
		log.Fatal("No se ha podido conectar al servidor")
	} else {
		log.Print("Conectado a supabase")
	}
}
