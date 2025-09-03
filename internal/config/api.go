package config

import "github.com/gin-gonic/gin"

// type APIServer struct {
// 	Name   string
// 	Addr   string
// 	Router *gin.Engine

// 	log        *zap.SugaredLogger
// 	tp         *trace.TracerProvider
// 	logFlusher func()
// }

// func NewAPIServer(name, addr string) *APIServer {
// 	cfg := config.LoadConfig()

// 	logger, flusher, err := logger.NewLogger(name, cfg.LoggerLevel)
// 	if err != nil {
// 		panic(err)
// 	}

// 	zap.ReplaceGlobals(logger)
// 	log := logger.Sugar()

// 	log.Debugw("Starting application...")

// 	tp, err := tracing.SetTraceProvider(tracing.TraceConfig{
// 		Name:            name,
// 		TracingEndPoint: cfg.TracingEndPoint,
// 		OtelState:       cfg.OtelState,
// 		OtelSampler:     cfg.OtelSampler,
// 		UptraceDsn:      cfg.UptraceDsn,
// 	})

// 	if err != nil {
// 		log.Panic("failed to set trace provider", "error", err)
// 	}

// 	router := gin.Default()

// 	p := ginprom.New(
// 		ginprom.Engine(router),
// 		ginprom.Path("/metrics"),
// 	)
// 	router.Use(p.Instrument())

// 	// init request id middleware
// 	router.Use(requestid.New())

// 	router.Use(otelgin.Middleware(name))
// 	router.Use(middleware.TraceRequestIDMiddleware())

// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"*"}
// 	config.AllowHeaders = []string{"*"}
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
// 	config.AllowCredentials = true

// 	router.Use(cors.New(config))

// 	router.GET("/healthz", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "ok"})
// 	})

// 	// Register Swagger documentation
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	server := &APIServer{
// 		Name:       name,
// 		Addr:       addr,
// 		Router:     router,
// 		log:        log,
// 		logFlusher: flusher,
// 		tp:         tp,
// 	}

// 	return server
// }

// func (s *APIServer) Run() {
// 	srv := &http.Server{
// 		Addr:    s.Addr,
// 		Handler: s.Router,
// 	}

// 	g := &run.Group{}

// 	g.Add(func() error {
// 		s.log.Infof("Server starting at %s", s.Addr)
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			s.log.Fatalf("Failed to listen and serve: %v", err)
// 			return err
// 		}
// 		return nil
// 	}, func(err error) {
// 		if err := srv.Shutdown(context.Background()); err != nil {
// 			s.log.Fatalf("Server forced to shutdown: %v", err)
// 		}
// 		s.log.Errorf("API %s listen happens error for: %v", s.Name, err)
// 		s.logFlusher()
// 		if err := s.tp.Shutdown(context.Background()); err != nil {
// 			s.log.Errorw("failed to shutdown trace provider", "error", err)
// 		}
// 		s.log.Info("Server exiting")
// 	})

// 	g.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

// 	if err := g.Run(); err != nil {
// 		s.log.Errorf("Error when runing http server: %v", err)
// 		os.Exit(1)
// 	}
// }

type APIServer struct {
	Name string
	Addr string
}

func NewAPIServer(name, addr string) *APIServer {
	return &APIServer{
		Name: name,
		Addr: addr,
	}
}

func (s *APIServer) Run() {
	gin.New().Run(s.Addr)
}
