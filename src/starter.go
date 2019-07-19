package main

import (
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/polaris-team/test-report-server/src/graphql/gqlgen"
	"github.com/polaris-team/test-report-server/src/graphql/resolvers"
	"github.com/polaris-team/test-report-server/pkg/config"
	"github.com/polaris-team/test-report-server/pkg/db/mysql"
	"github.com/rs/cors"
	"net/http"
	"os"
	"strconv"
)

func init(){
	//配置snowflake workId
	os.Setenv("WORK_ID", "1")
	//配置文件
	config.LoadConfig("configs", "application")
	//打印db log
	os.Setenv("UPPERIO_DB_DEBUG", "1")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler.Playground("GraphQL playground", "/report"))
	mux.Handle("/report", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolvers.Resolver{}})))
	handler := cors.Default().Handler(mux)

	serverConfig := config.GetServerConfig()
	port := strconv.Itoa(serverConfig.Port)
	fmt.Println("connect to http://localhost:" + port + "/ for GraphQL playground")
	http.ListenAndServe(":"+port, handler)
}

func GenerateStruct(savePath string, tables ...string) error{
	return mysql.Generate(savePath, tables)
}



