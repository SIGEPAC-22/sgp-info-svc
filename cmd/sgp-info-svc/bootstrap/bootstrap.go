package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sgp-info-svc/internal/getComorbidity/getComorbidityService"
	handler2 "sgp-info-svc/internal/getComorbidity/platform/handler"
	mysql2 "sgp-info-svc/internal/getComorbidity/platform/storage/mysql"
	"sgp-info-svc/internal/getCountry/getCountryService"
	handler8 "sgp-info-svc/internal/getCountry/platform/handler"
	mysql8 "sgp-info-svc/internal/getCountry/platform/storage/mysql"
	"sgp-info-svc/internal/getDataHistorical/getDataHistorical"
	handler12 "sgp-info-svc/internal/getDataHistorical/platform/handler"
	mysql12 "sgp-info-svc/internal/getDataHistorical/platform/storage/mysql"
	"sgp-info-svc/internal/getDepartment/getDepartmentService"
	handler9 "sgp-info-svc/internal/getDepartment/platform/handler"
	mysql9 "sgp-info-svc/internal/getDepartment/platform/storage/mysql"
	"sgp-info-svc/internal/getInfoPatient/getInfoPatientService"
	"sgp-info-svc/internal/getInfoPatient/platform/handler"
	"sgp-info-svc/internal/getInfoPatient/platform/storage/mysql"
	getInfoPatientFileFileSvc "sgp-info-svc/internal/getInfoPatientFile/getInfoPatientFileSvc"
	handler13 "sgp-info-svc/internal/getInfoPatientFile/platform/handler"
	mysql13 "sgp-info-svc/internal/getInfoPatientFile/platform/storage/mysql"
	handler6 "sgp-info-svc/internal/getOneComorbidity/platform/handler"
	mysql6 "sgp-info-svc/internal/getOneComorbidity/platform/storage/mysql"
	service2 "sgp-info-svc/internal/getOneComorbidity/service"
	"sgp-info-svc/internal/getOneInfoPatient/getOneInfoPatientSvc"
	handler5 "sgp-info-svc/internal/getOneInfoPatient/platform/handler"
	mysql5 "sgp-info-svc/internal/getOneInfoPatient/platform/storage/mysql"
	"sgp-info-svc/internal/getOnePatientFile/getOnePatientFileService"
	handler14 "sgp-info-svc/internal/getOnePatientFile/platform/handler"
	mysql14 "sgp-info-svc/internal/getOnePatientFile/platform/storage/mysql"
	handler4 "sgp-info-svc/internal/getOneSymptom/platform/handler"
	mysql4 "sgp-info-svc/internal/getOneSymptom/platform/storage/mysql"
	"sgp-info-svc/internal/getOneSymptom/service"
	"sgp-info-svc/internal/getSex/getSexService"
	handler10 "sgp-info-svc/internal/getSex/platform/handler"
	mysql10 "sgp-info-svc/internal/getSex/platform/storage/mysql"
	"sgp-info-svc/internal/getStatePatient/getStatePatientService"
	handler11 "sgp-info-svc/internal/getStatePatient/platform/handler"
	mysql11 "sgp-info-svc/internal/getStatePatient/platform/storage/mysql"
	"sgp-info-svc/internal/getSymptom/getSymptomService"
	handler3 "sgp-info-svc/internal/getSymptom/platform/handler"
	mysql3 "sgp-info-svc/internal/getSymptom/platform/storage/mysql"
	"sgp-info-svc/internal/getTypeDocument/getTypeDocumentService"
	handler7 "sgp-info-svc/internal/getTypeDocument/platform/handler"
	mysql7 "sgp-info-svc/internal/getTypeDocument/platform/storage/mysql"

	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)
	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)
	////////////////////////CORS///////////////////////////////////////////

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	/////////////////////GET INFO PATIENT/////////////////////
	repoGetInfoPatient := mysql.NewGetInfoPatientRepo(db, kitlogger)
	serviceGetInfoPatient := getInfoPatientService.NewService(repoGetInfoPatient, kitlogger)
	endpointGetInfoPatient := handler.MakeInfoPatientEndpoints(serviceGetInfoPatient)
	endpointGetInfoPatient = handler.GetInfoPatientTransportMiddleware(kitlogger)(endpointGetInfoPatient)
	transportGetInfoPatient := handler.NewHttpGetInfoPatientHandler(config.GetString("paths.getInfoPatient"), endpointGetInfoPatient)

	/////////////////////GET ONE INFO PATIENT/////////////////////
	repoGetOneInfoPatient := mysql5.NewGetOneInfoPatientRepo(db, kitlogger)
	serviceGetOneInfoPatient := getOneInfoPatientSvc.NewService(repoGetOneInfoPatient, kitlogger)
	endpointGetOneInfoPatient := handler5.MakeGetOneInfoPatientEndpoints(serviceGetOneInfoPatient)
	endpointGetOneInfoPatient = handler5.GetOneInfoPatientTransportMiddleware(kitlogger)(endpointGetOneInfoPatient)
	transportGetOneInfoPatient := handler5.NewHttpGetOneInfoPatientHandler(config.GetString("paths.getOneInfoPatient"), endpointGetOneInfoPatient)
	/////////////////////GET ONE INFO PATIENT/////////////////////

	/////////////////////GET INFO PATIENT FILE/////////////////////
	repoGetPatientFile := mysql13.NewGetInfoPatientFileRepo(db, kitlogger)
	serviceGetPatientFile := getInfoPatientFileFileSvc.NewService(repoGetPatientFile, kitlogger)
	endpointGetPatientFile := handler13.MakeInfoPatientFileEndpoints(serviceGetPatientFile)
	endpointGetPatientFile = handler13.GetInfoPatientFileTransportMiddleware(kitlogger)(endpointGetPatientFile)
	transportGetPatientFile := handler13.NewHttpGetInfoPatientFileHandler(config.GetString("paths.getInfoPatientFile"), endpointGetPatientFile)
	/////////////////////GET INFO PATIENT FILE/////////////////////

	/////////////////////GET ONE PATIENT FILE/////////////////////
	repoGetOnePatientFile := mysql14.NewGetOnePatientFileRepo(db, kitlogger)
	serviceGetOnePatientFile := getOnePatientFileService.NewService(repoGetOnePatientFile, kitlogger)
	endpointGetOnePatientFile := handler14.MakeGetOnePatientFileEndpoint(serviceGetOnePatientFile)
	endpointGetOnePatientFile = handler14.GetOnePatientFileTransportMiddleware(kitlogger)(endpointGetOnePatientFile)
	transportGetOnePatientFile := handler14.NewGetOnePatientFileHandler(config.GetString("paths.getOnePatientFile"), endpointGetOnePatientFile)
	/////////////////////GET ONE PATIENT FILE/////////////////////

	/////////////////////GET CONMORBIDITY/////////////////////
	repoGetComorbidity := mysql2.NewGetComorbidityRepo(db, kitlogger)
	serviceGetComorbidity := getComorbidityService.NewService(repoGetComorbidity, kitlogger)
	endpointGetComorbidity := handler2.MakeComorbidityEndpoints(serviceGetComorbidity)
	endpointGetComorbidity = handler2.GetComorbidityTransportMiddleware(kitlogger)(endpointGetComorbidity)
	transportGetComorbidity := handler2.NewHttpGetComorbidityHandler(config.GetString("paths.getComorbidity"), endpointGetComorbidity)

	/////////////////////GET STATE PATIENT/////////////////////
	repoGetStatePatient := mysql11.NewGetStatePatientRepo(db, kitlogger)
	serviceGetStatePatient := getStatePatientService.NewService(repoGetStatePatient, kitlogger)
	endpointGetStatePatient := handler11.MakeGetStatePatientEndpoints(serviceGetStatePatient)
	endpointGetStatePatient = handler11.GetStatePatientTransportMiddleware(kitlogger)(endpointGetStatePatient)
	transportGetStatePatient := handler11.NewHttpGetStatePatientHandler(config.GetString("paths.getStatePatient"), endpointGetStatePatient)

	/////////////////////GET TYPE DOCUMENT/////////////////////
	repoGetTypeDocument := mysql7.NewGetTypeDocumentRepo(db, kitlogger)
	serviceGetTypeDocument := getTypeDocumentService.NewService(repoGetTypeDocument, kitlogger)
	endpointGetTypeDocument := handler7.MakeGetTypeDocumentEndpoints(serviceGetTypeDocument)
	endpointGetTypeDocument = handler7.GetTypeDocumentTransportMiddleware(kitlogger)(endpointGetTypeDocument)
	transportGetTypeDocument := handler7.NewHttpGetTypeDocumentHandler(config.GetString("paths.getTypeDocument"), endpointGetTypeDocument)

	/////////////////////GET TYPE DEPARTMENT/////////////////////
	repoGetCountry := mysql8.NewGetCountryRepo(db, kitlogger)
	serviceGetCountry := getCountryService.NewService(repoGetCountry, kitlogger)
	endpointGetCountry := handler8.MakeGetCountryEndpoints(serviceGetCountry)
	endpointGetCountry = handler8.GetCountryTransportMiddleware(kitlogger)(endpointGetCountry)
	transportGetCountry := handler8.NewHttpGetCountryHandler(config.GetString("paths.getCountry"), endpointGetCountry)

	/////////////////////GET TYPE DEPARTMENT/////////////////////
	repoGetSex := mysql10.NewGetSexRepo(db, kitlogger)
	serviceGetSex := getSexService.NewService(repoGetSex, kitlogger)
	endpointGetSex := handler10.MakeGetSexEndpoints(serviceGetSex)
	endpointGetSex = handler10.GetSexTransportMiddleware(kitlogger)(endpointGetSex)
	transportGetSex := handler10.NewHttpGetSexHandler(config.GetString("paths.getSex"), endpointGetSex)

	/////////////////////GET TYPE DOCUMENT/////////////////////
	repoGetDepartment := mysql9.NewGetDepartmentRepo(db, kitlogger)
	serviceGetDepartment := getDepartmentService.NewService(repoGetDepartment, kitlogger)
	endpointGetDepartment := handler9.MakeGetDepartmentEndpoints(serviceGetDepartment)
	endpointGetDepartment = handler9.GetDepartmentTransportMiddleware(kitlogger)(endpointGetDepartment)
	transportGetDepartment := handler9.NewHttpGetDepartmentHandler(config.GetString("paths.getDepartment"), endpointGetDepartment)

	/////////////////////GET SYMPTOM/////////////////////
	repoGetSymptom := mysql3.NewGetSymptomRepo(db, kitlogger)
	serviceGetSymptom := getSymptomService.NewService(repoGetSymptom, kitlogger)
	endpointGetSymptom := handler3.MakeSymtpomEndpoints(serviceGetSymptom)
	endpointGetSymptom = handler3.GetSymptomTransportMiddleware(kitlogger)(endpointGetSymptom)
	transportGetSymptom := handler3.NewHttpGetSymptomHandler(config.GetString("paths.getSymptom"), endpointGetSymptom)

	//////////////////////GET ONE CONMORBILITY////////////////////////////////////////////////
	getOneComorbidityRepo := mysql6.NewGetOneComorbidityRepo(db, kitlogger)
	getOneComorbidityService := service2.NewGetOneComorbiditySvc(getOneComorbidityRepo, kitlogger)
	getOneComorbidityEndpoint := handler6.MakeGetOneComorbidityEndpoint(getOneComorbidityService)
	getOneComorbidityEndpoint = handler6.GetOneComorbidityTransportMiddleware(kitlogger)(getOneComorbidityEndpoint)
	getOneComorbidityHandler := handler6.NewGetOneComorbidityHandler(config.GetString("paths.getOneComorbidity"), getOneComorbidityEndpoint)
	//////////////////////GET ONE CONMORBILITY////////////////////////////////////////////////

	//////////////////////GET ONE SYMPTOM////////////////////////////////////////////////
	getOneSymptomRepo := mysql4.NewGetOneSymptomRepo(db, kitlogger)
	getOneSymptomService := service.NewGetOneSymptomSvc(getOneSymptomRepo, kitlogger)
	getOneSymptomEndpoint := handler4.MakeGetOneSymptomEndpoint(getOneSymptomService)
	getOneSymptomEndpoint = handler4.GetOneSymptomTransportMiddleware(kitlogger)(getOneSymptomEndpoint)
	getOneSymptomHandler := handler4.NewGetOneSymptomHandler(config.GetString("paths.getOneSymptom"), getOneSymptomEndpoint)
	//////////////////////GET ONE SYMPTOM////////////////////////////////////////////////

	//////////////////////GET ONE SYMPTOM////////////////////////////////////////////////
	getDataHistoricalRepo := mysql12.NewGetDataHistoricalRepo(db, kitlogger)
	getDataHistoricalService := getDataHistorical.NewService(getDataHistoricalRepo, kitlogger)
	getDataHistoricalEndpoint := handler12.MakeGetDataHistoricalEndpoints(getDataHistoricalService)
	getDataHistoricalEndpoint = handler12.GetDataHistoricalTransportMiddleware(kitlogger)(getDataHistoricalEndpoint)
	transportGetDataHistorical := handler12.NewHttpGetDataHistoricalHandler(config.GetString("paths.getDataHistorical"), getDataHistoricalEndpoint)
	//////////////////////GET ONE SYMPTOM////////////////////////////////////////////////

	mux.Handle(config.GetString("paths.getInfoPatientFile"), transportGetPatientFile)
	mux.Handle(config.GetString("paths.getOnePatientFile"), transportGetOnePatientFile)
	mux.Handle(config.GetString("paths.getDataHistorical"), transportGetDataHistorical)
	mux.Handle(config.GetString("paths.getStatePatient"), transportGetStatePatient)
	mux.Handle(config.GetString("paths.getSex"), transportGetSex)
	mux.Handle(config.GetString("paths.getDepartment"), transportGetDepartment)
	mux.Handle(config.GetString("paths.getCountry"), transportGetCountry)
	mux.Handle(config.GetString("paths.getTypeDocument"), transportGetTypeDocument)
	mux.Handle(config.GetString("paths.getInfoPatient"), transportGetInfoPatient)
	mux.Handle(config.GetString("paths.getOneInfoPatient"), transportGetOneInfoPatient)
	mux.Handle(config.GetString("paths.getComorbidity"), transportGetComorbidity)
	mux.Handle(config.GetString("paths.getSymptom"), transportGetSymptom)
	mux.Handle(config.GetString("paths.getOneComorbidity"), getOneComorbidityHandler)
	mux.Handle(config.GetString("paths.getOneSymptom"), getOneSymptomHandler)
	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
