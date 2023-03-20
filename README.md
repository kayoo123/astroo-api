# astroo-api

Ce projet a pour but de générer un json permettant d'obtenir l'horoscope astrologique de chaque signe du zodiaque.

#### Planification:
> Les actions sont exécutées chaque jour à 03h20 (UTC), cf: `20 3 * * *`.

#### Urls :
- [x] Aujourd'hui : https://kayoo123.github.io/astroo-api/jour.json
- [x] Semaine : https://kayoo123.github.io/astroo-api/hebdomadaire.json
- [ ] Mois : https://kayoo123.github.io/astroo-api/mensuel.json


#### Calendrier :
| Zodiac  | Naissance | Cardan 1 | Cardan 2 | Cardan 3 | 
| :--------------- |:---------------:|:----:|:----:|:----:|
| :aries: Bélier | 21 mars – 19 avril | 21/03 ~ 31/03 | 31/03 ~ 10/04 | 10/03 ~ 21/04|
| :taurus: Taureau | 20 avril – 20 mai | 21/04 ~ 30/04 | 30/04 ~ 10/05 | 10/05 ~ 21/05 |
| :gemini: Gémeaux | 21 mai – 21 juin | 21/05 ~ 31/05 | 31/05 ~ 10/06 | 10/06 ~ 21/06 |
| :cancer: Cancer | 22 juin – 22 juillet | 21/06 ~ 2/07 | 02/07 ~ 12/07 | 12/07 ~ 23/07 |
| :leo: Lion | 23 juillet – 22 août | 23/07 ~ 02/08 | 02/08 ~ 12/08 | 12/08 ~ 23/08 |
| :virgo: Vierge | 23 août – 22 septembre | 23/08 ~ 02/09 | 02/09 ~ 12/09 | 12/09 ~ 23/09 |
| :libra: Balance | 23 septembre – 23 octobre | 23/09 ~ 02/10 | 02/10 ~ 12/10 | 12/10 ~ 23/10 |
| :scorpius: Scorpion | 24 octobre – 22 novembre | 23/10 ~ 02/11 | 02/11 ~ 12/11 | 12/11 ~ 22/11 |
| :sagittarius: Sagittaire | 23 novembre – 22 décembre | 22/11 ~ 02/12 | 02/12 ~ 12/12 | 12/12 ~ 21/12 |
| :capricorn: Capricorne | 23 décembre – 20 janvier | 21/12 ~ 01/01 | 01/01 ~ 10/01 | 10/01 ~ 20/01 |
| :aquarius: Verseau | 21 janvier – 19 février | 20/01 ~ 30/01 | 30/01 ~ 8/02 | 8/02 ~ 19/02 |
| :pisces: Poissons | 20 février – 20 mars | 19/02 ~ 01/03 | 01/03 ~ 11/03 | 11/03 ~ 21/03 |

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
