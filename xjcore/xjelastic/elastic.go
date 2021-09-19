package xjelastic

import (
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ElasticInit(conf xjtypes.ElasticSearch) (*elastic.Client, error) {
	EsClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(conf.Addr),
		elastic.SetBasicAuth(conf.Username, conf.Password))
	if err != nil {
		log.Println(err)
	}
	return EsClient, err
}
