# Project root                                  .
# Main application entry point                  ├── cmd                         
# Command-line interface (CLI) subpackage       │   └── gator                   
# Initializes and runs the CLI                  │       └── main.go             
# Non-exported code for your application        ├── internal                    
# CLI command handlers                          │   ├── commands                
# Handles HTTP client setup                     │   │   ├── client.go           
# Maps and runs CLI commands                    │   │   ├── commands.go         
# Logic for `addfeed` command                   │   │   ├── handler_addfeed.go  
# Logic for aggregating feeds                   │   │   ├── handler_agg.go      
# Logic for feed-related actions                │   │   ├── handler_feeds.go    
# Logic for handling follow actions             │   │   ├── handler_follow.go   
# Logic for managing following feeds            │   │   ├── handler_following.go
# Logic for retrieving users                    │   │   ├── handler_getusers.go 
# Logic for user login                          │   │   ├── handler_login.go    
# Logic for user registration                   │   │   ├── handler_register.go 
# Logic for resetting the database              │   │   ├── handler_reset.go    
# Helper functions to clean the handlers        │   │   └── middleware.go       
# Configuration-related functionality           ├── config                      
# Loads and manages app configuration           │   └── config.go              
# Database-related functionality                ├── database                    
# SQL files for migrations and queries          │   └── sql                     
# Auto-generated database code (via sqlc)       │       ├── gen                 
# General database connection setup             │       │   ├── db.go           
# Go code generated from SQL                    │       │   ├── feed_follows.sql.go
# Go code generated from SQL                    │       │   ├── feeds.sql.go    
# Definitions for database models               │       │   ├── models.go       
# Go code generated from SQLC                   │       │   └── users.sql.go    
# Application-specific SQL queries              │       ├── queries             
# SQL queries for feed follow                   │       │   ├── feed_follows.sql 
# SQL queries for feed                          │       │   ├── feeds.sql       
# SQL queries for user                          │       │   └── users.sql       
# Database schema migrations                    │       ├── schema              
# Migration to create the `users` table         │       │   ├── 001_users.sql    
# Migration to create the `feeds` table         │       │   ├── 002_feeds.sql    
# Migration to create the `feed_follows` table  │       │   └── 003_feed_follows.sql
# Configuration file for SQLC code generation   │       └── sqlc.yaml           
# RSS feed-related functionality                ├── rss                         
# Logic for fetching RSS feeds                  │   ├── client.go             
# Core RSS parsing and processing logic         │   └── rss.go                
# Keeps track of in-memory application state    ├── state                       
# Manages global app state                      │   └── state.go              
# Build automation and development workflow     ├── Makefile                   
# Project documentation                         ├── README.md                  
# Go module file (dependencies and modules)     ├── go.mod                      
# Checksums for module dependencies             ├── go.sum                      
# This file - documents project structure       └── readme_structure.txt      
