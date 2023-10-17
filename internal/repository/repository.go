package repository

import (
	"github.com/sirupsen/logrus"
	_interface "github.com/xvbnm48/bookStore/internal/repository/interface"
	"github.com/xvbnm48/bookStore/internal/repository/mysql"
	"os"
)

var log = logrus.New()

type DB struct {
	URL, Port, User, Schema, Password string
}

type RepoConfigs struct {
	DBConf DB
}

type Repo struct {
	DBReadWriter _interface.ReadWriter
}

func NewBookStoreRepo(rc RepoConfigs) (*Repo, error) {
	log.Out = os.Stdout
	readWriter, err := mysql.NewDBReadWriter(rc.DBConf.URL, rc.DBConf.Port, rc.DBConf.Schema, rc.DBConf.User, rc.DBConf.Password)
	if err != nil {
		return nil, err
	}

	return &Repo{DBReadWriter: readWriter}, nil
}

func (r *Repo) Close() {
	log.Infof("closing erpo")
	r.DBReadWriter.Close()
}
