# load_database_go
# Description
### This project aims to read the file, process the data contained in it and save it in a relational database.

- load_database_go
  - Main Languange: Golang
  - Database: Postgresql

# Depedencies
- Docker
- Git
- Go

# Database
### For this project, the database and the table are create when you run:
```console
docker-compose up
```
### The colunms are based on requested file

Id   | cpf | private | imcompleto | data_da_ultima_compra | ticket_medio | ticket_da_ultima_compra | loja_mais_frequente | loja_da_ultima_compra |
--------- | ------------- | ------ | ------ | ---------- | -- | -- | ------------- | ------------- |
1         | Example Cpf   | false  | false  | 2023-01-01 | 0  | 0  | Example Cnpj  | Example Cnpj  |

# Type of Values
- cpf : string
- private : boolean
- imcompleto : boolean
- data_da_ultima_compra: string
- ticket_medio: float
- ticket_da_ultima_compra: float
- loja_mais_frequente: string
- loja_da_ultima_compra: string

# Run the Project
### The project need a file to process and save in database
### In main directory of github, example to run:

```console
go run main.go base_teste.txt
```
### This command start the read and process the data in file
### If success the output is:
```console
Successufully entered data within 37.135665543s
```

- Time may vary depending on hardware
- Test environment
- - VirtualBox 6.1
- - Ubuntu 22.04.2 LTS
- - Memory: 4,9 GiB
- - Processor: Intel® Core™ i5-6200U CPU @ 2.30GHz × 2


### If you want see the data entered in database, follow the next steps:
# Data visualization
### In docker compose project have a image of pgAdmin and it's run on: <http://localhost:5050/>
## Login page
### In login page you must entry with the credentials:
- email: admin@email.com
- password: admin

### this credentials are in docker-compose.yml
## Main Page
### In main page you must click in "Add New Server"

### With window opened go to "General"
- Name : postgres

### Go to "Connection"

- Host name/address : postgres
- Port: 5432
- Maintenance database: postgres
- Username : postres
- Password : 1234

### Click in Save

## The Server connected
### Navegate to: 

```
├─ Servers 
    └ postgres
      └ client
        └ Schemas
          └ public
            └ Tables
              └ client_infos
```
### right click in 'client_infos'
- View / Edit Data > All Rows

### With this you will be able to visualize the processed data in the database.



