package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func setupRepos(t *testing.T) (TuristaRepository, NegocioRepository, *CheckInMemoria) {
	t.Helper()
	repoT := NewTuristaMemoria()
	repoN := NewNegocioMemoria()
	repoC := NewCheckInMemoria(repoT, repoN)

	repoT.Guardar(models.Turista{ID: 1, Nombre: "John Erick", Nacionalidad: "Ecuador", IdiomaPreferido: "es"})
	repoN.Guardar(models.Negocio{ID: 1, Nombre: "Café", Tipo: "restaurante", Ciudad: "Manta", IdiomasHablados: []string{"es"}})

	return repoT, repoN, repoC
}

func TestCheckInMemoria_Guardar(t *testing.T) {
	casos := []struct {
		nombre      string
		entrada     models.CheckIn
		errEsperado error
	}{
		{"Checkin válido", models.CheckIn{ID: 10, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5}, nil},
		{"Fecha vacía", models.CheckIn{ID: 11, TuristaID: 1, NegocioID: 1, Fecha: "", Calificacion: 5}, errs.ErrDatosInvalidos},
		{"Calificación = 0", models.CheckIn{ID: 12, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 0}, errs.ErrDatosInvalidos},
		{"Calificación = 6", models.CheckIn{ID: 13, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 6}, errs.ErrDatosInvalidos},
		{"Turista NO existe", models.CheckIn{ID: 14, TuristaID: 99, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5}, errs.ErrNoEncontrado},
		{"Negocio NO existe", models.CheckIn{ID: 15, TuristaID: 1, NegocioID: 99, Fecha: "2026-04-10", Calificacion: 5}, errs.ErrNoEncontrado},
		{"ID ya usado", models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5}, errs.ErrYaExiste},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			_, _, checkins := setupRepos(t)

			if c.nombre == "ID ya usado" {
				checkins.Guardar(models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5})
			}

			err := checkins.Guardar(c.entrada)
			if !errors.Is(err, c.errEsperado) {
				t.Errorf("esperaba %v, obtuvo %v", c.errEsperado, err)
			}
		})
	}
}

func TestCheckInMemoria_BuscarPorTurista(t *testing.T) {
	t.Run("Turista con múltiples visitas", func(t *testing.T) {
		_, _, checkins := setupRepos(t)
		checkins.Guardar(models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5})
		checkins.Guardar(models.CheckIn{ID: 2, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-11", Calificacion: 4})

		visitas, err := checkins.BuscarPorTurista(1)

		if err != nil {
			t.Errorf("no esperaba error, obtuvo %v", err)
		}
		if len(visitas) != 2 {
			t.Errorf("esperaba 2 visitas, obtuvo %d", len(visitas))
		}
	})

	t.Run("Turista existe pero sin visitas", func(t *testing.T) {
		_, _, checkins := setupRepos(t)
		visitas, err := checkins.BuscarPorTurista(1) // El turista 1 existe pero no tiene check-ins aún
		if err != nil || len(visitas) != 0 {
			t.Errorf("debería devolver slice vacío y nil, obtuvo len=%d y err=%v", len(visitas), err)
		}
	})

	t.Run("Turista no existe en el sistema", func(t *testing.T) {
		_, _, checkins := setupRepos(t)
		visitas, _ := checkins.BuscarPorTurista(999)
		if len(visitas) != 0 {
			t.Errorf("esperaba 0 resultados para turista inexistente")
		}
	})
}

func TestCheckInMemoria_Listar(t *testing.T) {
	_, _, checkins := setupRepos(t)
	checkins.Guardar(models.CheckIn{ID: 1, TuristaID: 1, NegocioID: 1, Fecha: "2026-04-10", Calificacion: 5})

	lista := checkins.Listar()
	if len(lista) != 1 {
		t.Errorf("esperaba 1, obtuvo %d", len(lista))
	}
}
