package main

import (
	"context"
	"flag"
	"os"
	"shop-go/internal/config"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/minio"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// Define command-line flags
	configPath := flag.String("config", "configs/config.yaml", "Path to configuration file")
	singleFile := flag.String("file", "", "Migrate a single file (optional)")
	flag.Parse()

	// Load configuration
	viper.SetConfigFile(*configPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	var cfg config.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic("Failed to parse configuration: " + err.Error())
	}

	// Initialize logger
	logConfig := &logger.LogConfig{
		Level:      cfg.Logger.Level,
		Encoding:   cfg.Logger.Encoding,
		OutputPath: cfg.Logger.OutputPath,
	}
	logger.Init(logConfig)
	defer logger.Sync()

	// Initialize MinIO client
	_, err = minio.InitClient(&cfg.MinIO)
	if err != nil {
		logger.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	// Run migration
	if *singleFile != "" {
		// Migrate a single file
		logger.Infof("Migrating file: %s", *singleFile)
		url, err := minio.MigrateFile(ctx, *singleFile)
		if err != nil {
			logger.Fatalf("Failed to migrate file: %v", err)
		}
		logger.Infof("File migrated successfully. URL: %s", url)
	} else {
		// Migrate all files
		logger.Info("Starting migration of all files from local storage to MinIO...")
		result, err := minio.MigrateLocalToMinIO(ctx)
		if err != nil {
			logger.Fatalf("Migration failed: %v", err)
		}

		// Print summary
		logger.Infof("Migration completed: Total files: %d, Successfully migrated: %d, Failed: %d",
			result.TotalFiles, result.SuccessCount, result.FailedCount)

		// Log failed files
		if result.FailedCount > 0 {
			logger.Warn("Some files failed to migrate")

			for _, file := range result.FailedFiles {
				logger.Warnf("Failed to migrate: %s", file)
			}

			// Write failed files to a log file
			logFile, err := os.Create("migration_failures.log")
			if err == nil {
				defer logFile.Close()
				for _, file := range result.FailedFiles {
					logFile.WriteString(file + "\n")
				}
				logger.Infof("Failed files have been written to migration_failures.log")
			}
		}
	}
}
