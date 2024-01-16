package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rhinosc/03-web-challenge/internal/handler"
	"github.com/rhinosc/03-web-challenge/internal/loader"
	"github.com/rhinosc/03-web-challenge/internal/repository"
	"github.com/rhinosc/03-web-challenge/internal/service"
)

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db, err := loader.NewLoaderTicketCSV(a.dbFile).Load()
	if err != nil {
		return
	}

	rp := repository.NewRepositoryTicketMap(db, len(db))
	// service ...
	sv := service.NewServiceTicketDefault(rp)
	// handler ...
	hd := handler.NewTicketDefault(sv)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	a.rt.Get("/ticket", hd.Get())
	a.rt.Get("/ticket/getByCountry/{dest}", hd.GetByDestinationCountry())
	a.rt.Get("/ticket/getAverage/:{dest}", hd.GetPercentageByDestinationCountry())

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
