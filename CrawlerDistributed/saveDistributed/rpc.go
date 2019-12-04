package saveDistributed

import (
	"NewsSpider/engine"
	"NewsSpider/newsSave"
	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := newsSave.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
