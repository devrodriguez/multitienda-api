# Mongo DB
- Select database
use <database_name>

- Insert document
db.categories.insert({"name":"Carnicos", "description": "Carnes y embutidos"})

# Heroku
- Deploy
## Create Heroku app
heroku create -a "project-name"

## Deploy app
git push heroku master


- Heroku repository
https://multitienda.herokuapp.com/

# Google Cloud
- Verificar proyecto predeterminado
gcloud config list

- Set project
gcloud config set project [YOUR_PROJECT_ID]