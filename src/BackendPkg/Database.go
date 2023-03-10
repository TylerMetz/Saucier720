import (
    "database/sql"
    "github.com/mattn/go-sqlite3"
)

func main(){
	database, _ := sql.Opeen("sqllite3", "./Publix.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS FoodItems (
		Name TEXT,
		StoreCost REAL,
		OnSale INTEGER,
		SalePrice REAL,
		SaleDetails TEXT,
		Quantity INTEGER
	)")
	statement.Exec();
}