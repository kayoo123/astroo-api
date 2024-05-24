package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
    "reflect"
    "strings"
)

// Définir la structure de la réponse JSON
type Response struct {
    Lion        string `json:"lion"`
    Vierge      string `json:"vierge"`
    Balance     string `json:"balance"`
    Scorpion    string `json:"scorpion"`
    Sagittaire  string `json:"sagittaire"`
    Capricorne  string `json:"capricorne"`
    Verseau     string `json:"verseau"`
    Poissons    string `json:"poissons"`
    Belier      string `json:"belier"`
    Taureau     string `json:"taureau"`
    Gemeaux     string `json:"gemeaux"`
    Cancer      string `json:"cancer"`
}

func main() {
    // Désactiver les indicateurs de date et d'heure dans les messages de log
    log.SetFlags(0)

    // Définir le drapeau (flag) pour le signe du zodiaque
    zodiacSign := flag.String("z", "", "Le signe du zodiaque (lion, vierge, balance, etc...)")
    flag.Parse()

    // Vérifiez si le signe du zodiaque est fourni
    if *zodiacSign == "" {
        log.Fatalf("Vous devez spécifier un signe du zodiaque avec l'option -z\n - Exemple : ./astroo-cli -z vierge")
    }

    // Convertir l'argument en minuscule pour le rendre insensible à la casse
    sign := strings.ToLower(*zodiacSign)

    url := "https://kayoo123.github.io/astroo-api/jour.json"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Erreur lors de la requête HTTP: %v", err)
    }
    defer resp.Body.Close()

    // Décodez la réponse JSON dans la structure
    var response Response
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        log.Fatalf("Erreur lors du décodage du JSON: %v", err)
    }

    // Utilisez la réflexion pour accéder dynamiquement au champ correspondant
    responseValue := reflect.ValueOf(response)
    field := responseValue.FieldByName(strings.Title(sign))

    if field.IsValid() && field.Kind() == reflect.String {
        fmt.Printf("%s: %s\n", strings.Title(sign), field.String())
    } else {
        log.Fatalf("Signe du zodiaque inconnu ou non pris en charge: %s", *zodiacSign)
    }
}
