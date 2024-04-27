package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// DBAdapter Estructura para el adaptador.
type DBAdapter struct {
	db *gorm.DB
}

// NewDBAdapter Constructor para el adaptador.
func NewDBAdapter() *DBAdapter {
	db, err := gorm.Open("sqlite3", "./report.sqlite3")
	if err != nil {
		panic("Error al conectar con la base de datos")
	}

	// Migrar el esquema
	// db.AutoMigrate(&Report{}) // Migrar la tabla de informes

	return &DBAdapter{db}
}

// Report Definición de la estructura del reporte
type Report struct {
	ID          uint `gorm:"primary_key"`
	Image       string
	Description string
	User        string
	Lat         float64
	Lng         float64
	DateTime    string
}

// FindReportByID Método para buscar un reporte por ID
func (adapter *DBAdapter) FindReportByID(id uint) (*Report, error) {
	var report Report
	if err := adapter.db.First(&report, id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func main() {
	adapter := NewDBAdapter()
	fmt.Println(time.Now().Format(time.RFC3339))
	// Buscar un reporte por ID
	report, err := adapter.FindReportByID(1)
	if err != nil {
		fmt.Println("Error al buscar el reporte:", err)
	} else {
		fmt.Println("Reporte encontrado:", report)
	}
}
