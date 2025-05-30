## Migrações de Banco de Dados

Este projeto utiliza o `golang-migrate` para gerenciar as alterações no esquema do banco de dados. As migrações são arquivos SQL localizados no diretório `database/migrations`.

### Instalação da CLI `golang-migrate`

Você precisa instalar a ferramenta CLI `golang-migrate` para criar e gerenciar arquivos de migração manualmente.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Verifique a instalação:
```bash
migrate -version
```

### Criando Novos Arquivos de Migração

Para criar um novo conjunto de arquivos de migração (por exemplo, para adicionar uma nova tabela ou alterar uma existente), use o comando `migrate create` a partir da raiz do projeto:

```bash
migrate create -ext sql -dir database/migrations -seq <nome_da_migracao>
```

Substitua `<nome_da_migracao>` por um nome descritivo para sua migração (ex: `adicionar_coluna_telefone_usuario`). Isso gerará dois arquivos no diretório `database/migrations`:

*   `YYYYMMDDHHMMSS_<nome_da_migracao>.up.sql`: Contém as instruções SQL para aplicar a migração.
*   `YYYYMMDDHHMMSS_<nome_da_migracao>.down.sql`: Contém as instruções SQL para reverter (rollback) a migração.

**Exemplo:**
```bash
migrate create -ext sql -dir database/migrations -seq criar_tabela_clientes
```
Isso criará (assumindo que esta é a primeira migração ou a próxima na sequência):
*   `database/migrations/000001_criar_tabela_clientes.up.sql`
*   `database/migrations/000001_criar_tabela_clientes.down.sql`

Edite esses arquivos com as instruções SQL DDL desejadas.

### Executando Migrações Manualmente (Levantar e Derrubar)

A aplicação tenta executar as migrações pendentes na inicialização, mas você também pode executá-las manualmente usando a CLI. Certifique-se de que seu arquivo `.env` esteja configurado corretamente ou forneça a URL do banco de dados diretamente.

As variáveis de ambiente do arquivo `.env` usadas para construir a URL do banco de dados para as migrações são:
*   `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_SSLMODE`

O caminho das migrações é definido por `MIGRATIONS_PATH` no `.env`.

**Para aplicar (levantar) todas as migrações "up" pendentes:**
```bash
# Exemplo usando os valores típicos de um .env:
migrate -database "postgres://youruser:yourpassword@localhost:5432/yourdbname?sslmode=disable" -path database/migrations up
```

**Para reverter (derrubar) a última migração aplicada:**
```bash
migrate -database "postgres://youruser:yourpassword@localhost:5432/yourdbname?sslmode=disable" -path database/migrations down 1
```

**Para reverter (derrubar) todas as migrações:**
```bash
migrate -database "postgres://youruser:yourpassword@localhost:5432/yourdbname?sslmode=disable" -path database/migrations down -all
```

### Migrações Automáticas na Inicialização da Aplicação

O arquivo `main.go` está configurado para executar automaticamente quaisquer migrações "up" pendentes quando a aplicação inicia. Ele utiliza os detalhes de conexão do banco de dados e o `MIGRATIONS_PATH` especificados no arquivo `.env`. Se as migrações forem aplicadas com sucesso ou se não houver alterações, a aplicação prosseguirá para iniciar o servidor. Se as migrações falharem, a aplicação registrará um erro fatal e será encerrada.
