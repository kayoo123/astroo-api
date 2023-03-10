# astroo-api

Ce projet a pour but de générer un json permettant d'obtenir l'horoscope astrologique de chaque signe du zodiaque.

#### Planification:
> Les actions sont exécutées chaque jour à 03h20 (UTC), cf: `20 3 * * *`.

#### Urls :
- [x] Aujourd'hui : https://kayoo123.github.io/astroo-api/jour.json
- [x] Semaine : https://kayoo123.github.io/astroo-api/hebdomadaire.json
- [ ] Mois : https://kayoo123.github.io/astroo-api/mensuel.json


#### Calendrier :
| Zodiac  | Naissance          | Cardan 1 | Cardan 2 | Cardan 3 | 
| :--------------- |:---------------:|:----:|:----:|:----:|
| Bélier | 21 mars – 19 avril |
| Taureau | 20 avril – 20 mai |
| Gémeaux | 21 mai – 21 juin |
| Cancer | 22 juin – 22 juillet |
| Lion | 23 juillet – 22 août |
| Vierge | 23 août – 22 septembre |
| Balance | 23 septembre – 23 octobre |
| Scorpion | 24 octobre – 22 novembre |
| Sagittaire | 23 novembre – 22 décembre |
| Capricorne | 23 décembre – 20 janvier |
| Verseau | 21 janvier – 19 février |
| Poissons | 20 février – 20 mars |

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
