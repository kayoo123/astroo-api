# astroo-api

Ce projet a pour but de générer un json permettant d'obtenir l'horoscope astrologique de chaque signe du zodiaque.

#### Planification:
> Les actions sont exécutées chaque jour à 03h20 (UTC), cf: `20 3 * * *`.

#### Urls :
- [x] Aujourd'hui : https://kayoo123.github.io/astroo-api/jour.json
- [x] Semaine : https://kayoo123.github.io/astroo-api/hebdomadaire.json
- [ ] Mois : https://kayoo123.github.io/astroo-api/mensuel.json

#### Usage : 
- Avec `Curl` et `jq`:
```bash
curl -s https://kayoo123.github.io/astroo-api/hebdomadaire.json | jq .vierge[2]
```

Source : www.astroo.com
