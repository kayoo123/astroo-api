# astroo-api

Ce projet a pour but de générer un json permettant d'obtenir l'horoscope astrologique de chaque signe du zodiaque.

#### Planification:
> Les actions sont exécutées chaque jour à 03h20 (UTC), cf: `20 3 * * *`.

#### Urls :
- [x] Aujourd'hui : https://kayoo123.github.io/astroo-api/jour.json
- [x] Semaine : https://kayoo123.github.io/astroo-api/hebdomadaire.json
- [ ] Mois : https://kayoo123.github.io/astroo-api/mensuel.json


#### Calendrier :
| Zodiac  | Naissance | 1er Décan | 2nd Décan | 3eme Décan | 
| :--------------- |:---------------:|:----:|:----:|:----:|
| :aries: Bélier | 21 mars – 20 avril | 21/03 ~ 01/04 | 02/04 ~ 11/04 | 12/04 ~ 20/04|
| :taurus: Taureau | 21 avril – 20 mai | 21/04 ~ 01/05 | 02/05 ~ 11/05 | 12/05 ~ 20/05 |
| :gemini: Gémeaux | 21 mai – 21 juin | 21/05 ~ 31/05 | 01/06 ~ 10/06 | 11/06 ~ 21/06 |
| :cancer: Cancer | 22 juin – 22 juillet | 22/06 ~ 01/07 | 02/07 ~ 12/07 | 13/07 ~ 22/07 |
| :leo: Lion | 23 juillet – 22 août | 23/07 ~ 02/08 | 03/08 ~ 13/08 | 14/08 ~ 22/08 |
| :virgo: Vierge | 23 août – 22 septembre | 23/08 ~ 03/09 | 04/09 ~ 13/09 | 14/09 ~ 22/09 |
| :libra: Balance | 23 septembre – 22 octobre | 23/09 ~ 03/10 | 04/10 ~ 14/10 | 15/10 ~ 22/10 |
| :scorpius: Scorpion | 23 octobre – 22 novembre | 23/10 ~ 03/11 | 04/11 ~ 13/11 | 14/11 ~ 22/11 |
| :sagittarius: Sagittaire | 23 novembre – 21 décembre | 23/11 ~ 02/12 | 03/12 ~ 12/12 | 13/12 ~ 21/12 |
| :capricorn: Capricorne | 22 décembre – 20 janvier | 22/12 ~ 02/01 | 03/01 ~ 11/01 | 12/01 ~ 20/01 |
| :aquarius: Verseau | 21 janvier – 18 février | 21/01 ~ 31/01 | 01/02 ~ 10/02 | 11/02 ~ 18/02 |
| :pisces: Poissons | 19 février – 20 mars | 19/02 ~ 28/02 | 01/03 ~ 10/03 | 11/03 ~ 20/03 |

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
