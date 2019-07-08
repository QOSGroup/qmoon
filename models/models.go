package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/QOSGroup/qmoon/config"
	"github.com/QOSGroup/qmoon/models/migrations"
	"github.com/QOSGroup/qmoon/testdata"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Engine represents a XORM engine or session.
type Engine interface {
	Delete(interface{}) (int64, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Sql(string, ...interface{}) *xorm.Session
	Table(interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
}

var (
	basex *xorm.Engine
	xs    map[string]*xorm.Engine

	qmoonTables []interface{}
	nodeTables  []interface{}

	dbCfg dbconfig
)

type dbconfig struct {
	Type, Host, Name, User, Passwd string
}

func (dc dbconfig) DataSource(name string) string {
	if name == "" {
		name = dc.Name
	}

	switch dc.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&autocommit=true", dc.User, dc.Passwd, dc.Host, name)
	case "postgres":
		return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dc.User, dc.Passwd, dc.Host, name)
	default:
		return ""
	}
}

func init() {
	xs = make(map[string]*xorm.Engine)

	qmoonTables = append(qmoonTables,
		new(Account), new(Node), new(App), new(LoginStatus), new(NodeRoute), new(QmoonStatus),
		new(VerifyCode),
	)

	nodeTables = append(nodeTables,
		new(AtmIpRecord), new(AtmRecord),
		new(Validator), new(BlockValidator), new(Peer),
		new(Block), new(Tx), new(ConsensusState), new(Genesis),
		new(TmBlock), new(TxTransfer),
		new(Fee), new(Evidence), new(Missing),
		new(Network),
	)

	gonicNames := []string{"SSL"}
	for _, name := range gonicNames {
		core.LintGonicMapper[name] = true
	}
}

func checkDriverSupport(d string) bool {
	if d == "mysql" {
		return true
	}

	if d == "postgres" {
		return true
	}

	return false
}

func InitDb(cfg *config.DBConfig) (err error) {
	if !checkDriverSupport(cfg.DriverName) {
		return errors.New("unsupported db driver")
	}

	dbCfg.Type = cfg.DriverName
	dbCfg.Host = cfg.Addr
	dbCfg.Name = cfg.Database
	dbCfg.User = cfg.User
	dbCfg.Passwd = cfg.Password

	err = newEngine()
	if err != nil {
		return err
	}

	nodes, err := Nodes()
	inited := make(map[string]bool)
	if err == nil {
		for _, node := range nodes {
			if inited[node.ChainId] {
				continue
			}

			if err := newNode(node.ChainId); err != nil {
				return err
			}
			inited[node.ChainId] = true
		}
	}

	return nil
}

func getEngine(name string) (*xorm.Engine, error) {
	if x, ok := xs[name]; ok {
		return x, nil
	}

	x, err := xorm.NewEngine(dbCfg.Type, dbCfg.DataSource(name))
	if err != nil {
		return nil, fmt.Errorf("Fail to connect to database: %v ", err)
	}
	x.SetMapper(core.GonicMapper{})
	x.SetLogger(xorm.NewSimpleLogger(os.Stdout))
	//x.SetLogLevel(core.LOG_INFO)
	x.SetLogLevel(core.LOG_DEBUG)
	//x.ShowSQL(true)
	x.SetTZLocation(time.Local)
	xs[name] = x

	return x, nil
}

func newEngine() (err error) {
	basex, err = getEngine(dbCfg.Name)
	if err != nil {
		return err
	}

	if err = migrations.Migrate(basex); err != nil {
		return fmt.Errorf("migrate: %v", err)
	}

	if err = basex.Sync2(qmoonTables...); err != nil {
		return fmt.Errorf("sync database struct error: %v\n", err)
	}

	return nil
}

func nodeDbname(chainId string) string {

	return dbCfg.Name + "_" + strings.Replace(chainId, "-", "_", -1)
}

func GetNodeEngine(name string) (*xorm.Engine, error) {
	name = nodeDbname(name)
	return getEngine(name)
}

func newNode(name string) (err error) {
	name = nodeDbname(name)
	exist, err := databaseExist(basex.DB().DB, name)
	if err != nil {
		return err
	}

	if !exist {
		err := createDatabase(basex.DB().DB, name)
		if err != nil {
			return err
		}
	}

	x, err := getEngine(name)
	if err != nil {
		return err
	}

	if err = x.Sync2(nodeTables...); err != nil {
		return fmt.Errorf("sync database struct error: %v\n", err)
	}

	return nil
}

// databaseExist 检查数据库是否存在
func databaseExist(db *sql.DB, dbName string) (bool, error) {
	s := fmt.Sprintf("select count(*) from pg_catalog.pg_database where datname='%s';", dbName)
	var count int
	err := db.QueryRow(s).Scan(&count)

	//fmt.Printf("databaseExist:%v, err:%v\n", count == 1, err)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// createDatabase 创建数据库
func createDatabase(db *sql.DB, dbName string) error {
	s := fmt.Sprintf("create database %s ENCODING 'UTF8' TEMPLATE template0;", dbName)

	_, err := db.Query(s)

	//fmt.Printf("createDatabase:%v, err:%v\n", dbName, err)
	return err
}

// DropDatabase 删除数据库
func DropDatabase(name string) error {
	dbName := nodeDbname(name)
	var err error

	closeConn := fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity "+
		"WHERE datname='%s' AND pid<>pg_backend_pid();", dbName)
	_, err = basex.Query(closeConn)
	if err != nil {
		return err
	}

	s := fmt.Sprintf("drop database %s;", dbName)
	_, err = basex.Query(s)

	//fmt.Printf("dropDatabase:%v, err:%v\n", dbName, err)

	return err
}

// DropDatabase 删除数据库
func dropDatabase(db *sql.DB, dbName string) error {
	var err error

	closeConn := fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity "+
		"WHERE datname='%s' AND pid<>pg_backend_pid();", dbName)
	_, err = db.Query(closeConn)
	if err != nil {
		return err
	}

	s := fmt.Sprintf("drop database %s;", dbName)
	_, err = db.Query(s)

	//fmt.Printf("dropDatabase:%v, err:%v\n", dbName, err)

	return err
}

type testEngine struct {
	db        *sql.DB
	databases []string
}

func (te testEngine) Close() {
	for _, v := range te.databases {
		dropDatabase(te.db, v)
	}
	te.db.Close()
}

func NewTestEngine(cfg *config.DBConfig) (*testEngine, error) {
	te := new(testEngine)
	if !checkDriverSupport(cfg.DriverName) {
		return nil, errors.New("unsupported db driver")
	}
	dbCfg.Name = ""
	dbCfg.Type = cfg.DriverName
	dbCfg.Host = cfg.Addr
	dbCfg.User = cfg.User
	dbCfg.Passwd = cfg.Password
	db, err := sql.Open(dbCfg.Type, dbCfg.DataSource(""))
	if err != nil {
		return nil, err
	}
	te.db = db

	dbCfg.Name = cfg.Database + fmt.Sprintf("%d", time.Now().UnixNano())
	exist, err := databaseExist(te.db, dbCfg.Name)
	if err != nil {
		return nil, err
	}

	if !exist {
		err := createDatabase(te.db, dbCfg.Name)
		if err != nil {
			return nil, err
		}
	}
	err = newEngine()
	if err != nil {
		return nil, err
	}
	te.databases = append(te.databases, dbCfg.Name)

	if err := newNode(testdata.ChainID); err != nil {
		te.Close()
		return nil, err
	}
	te.databases = append(te.databases, nodeDbname(testdata.ChainID))

	return te, nil
}
