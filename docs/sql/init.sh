HOST="localhost"
PORT="3306"
DB_NAME="investod_db"
USERNAME="root"
PASSWORD="password"
PROTOCOL="tcp"

echo "DROP DATABASE IF EXISTS ${DB_NAME};\n" >> temp.sql
echo "CREATE DATABASE IF NOT EXISTS ${DB_NAME};\n" >> temp.sql
echo "USE ${DB_NAME};\n" >> temp.sql
echo "SET foreign_key_checks = 0;\n" >> temp.sql
cat ./docs/sql/schema.sql >> temp.sql
cat ./docs/sql/seed.sql >> temp.sql
echo "SET foreign_key_checks = 1;\n"

mysql -h ${HOST} -P ${PORT} -u ${USERNAME} --password=${PASSWORD} --protocol=${PROTOCOL} < temp.sql
rm temp.sql