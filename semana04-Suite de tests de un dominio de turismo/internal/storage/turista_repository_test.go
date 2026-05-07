package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestTuristaMemoria_Guardar(t *testing.T) {
	casos := []struct {
		nombre      string
		turista     models.Turista
		errEsperado error
	}{
		{
			nombre: "Guardado exitoso",
			turista: models.Turista{
				ID: 1, Nombre: "John Erick", Nacionalidad: "Ecuador", IdiomaPreferido: "es",
			},
			errEsperado: nil,
		},
		{
			nombre: "Error: ID negativo",
			turista: models.Turista{
				ID: -5, Nombre: "John", Nacionalidad: "Ecuador",
			},
			errEsperado: errs.ErrDatosInvalidos,
		},
		{
			nombre: "Error: Nombre vacío",
			turista: models.Turista{
				ID: 2, Nombre: "", Nacionalidad: "Ecuador",
			},
			errEsperado: errs.ErrDatosInvalidos,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			repo := NewTuristaMemoria() // Repo limpio por cada caso [cite: 234]
			err := repo.Guardar(c.turista)

			if !errors.Is(err, c.errEsperado) { // Uso obligatorio de errors.Is [cite: 236, 250]
				t.Errorf("esperaba error %v, obtuvo %v", c.errEsperado, err)
			}
		})
	}
}

func TestTuristaMemoria_BuscarPorID(t *testing.T) {
	repo := NewTuristaMemoria()

	t1 := models.Turista{
		ID:              1,
		Nombre:          "John",
		Nacionalidad:    "Ecuador",
		IdiomaPreferido: "es",
	}

	if err := repo.Guardar(t1); err != nil {
		t.Fatalf("setup de búsqueda falló: %v", err)
	}

	t.Run("Encuentra turista existente", func(t *testing.T) {
		encontrado, err := repo.BuscarPorID(1)
		if err != nil {
			t.Errorf("no esperaba error, obtuvo: %v", err)
		}
		if encontrado.Nombre != "John" {
			t.Errorf("Nombre: esperaba John, obtuvo %s", encontrado.Nombre)
		}
	})

	t.Run("Error: ID inexistente", func(t *testing.T) {
		_, err := repo.BuscarPorID(99)
		if !errors.Is(err, errs.ErrNoEncontrado) {
			t.Errorf("esperaba ErrNoEncontrado, obtuvo %v", err)
		}
	})
}

func TestTuristaMemoria_Listar(t *testing.T) {
	repo := NewTuristaMemoria()
	repo.Guardar(models.Turista{ID: 1, Nombre: "Turista A", Nacionalidad: "EC", IdiomaPreferido: "es"})
	repo.Guardar(models.Turista{ID: 2, Nombre: "Turista B", Nacionalidad: "EC", IdiomaPreferido: "es"})

	lista := repo.Listar()

	if len(lista) != 2 {
		t.Errorf("esperaba 2 turistas, obtuvo %d", len(lista))
	}
}
