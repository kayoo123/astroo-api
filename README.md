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
- ...et en rajoutant un `Cowsay`:
```
curl -s https://kayoo123.github.io/astroo-api/jour.json | jq -r .vierge |cowsay
 _________________________________________
/ Depuis quelque temps, vous vous ouvrez  \
| aux autres et vous découvrez l'écoute,  |
| le sens du partage. Cela vous permet de |
\ faire quelques prises de conscience...  /
 -----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
Source : www.astroo.com
