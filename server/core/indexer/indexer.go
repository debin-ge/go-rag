package indexer

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino-ext/components/indexer/es8"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/wangle201210/go-rag/server/core/common"
	"github.com/wangle201210/go-rag/server/core/config"
)

// newIndexer component initialization function of node 'Indexer2' in graph 'rag'
func newIndexer(ctx context.Context, conf *config.Config) (idr indexer.Indexer, err error) {
	indexerConfig := &es8.IndexerConfig{
		Client:    conf.Client,
		Index:     conf.IndexName,
		BatchSize: 10,
		DocumentToFields: func(ctx context.Context, doc *schema.Document) (field2Value map[string]es8.FieldValue, err error) {
			var knowledgeName string
			if value, ok := ctx.Value(common.KnowledgeName).(string); ok {
				knowledgeName = value
			} else {
				err = fmt.Errorf("必须提供知识库名称")
				return
			}
			// 没传入才需要生成
			if len(doc.ID) == 0 {
				doc.ID = uuid.New().String()
			}
			if doc.MetaData != nil {
				// 存储ext数据
				marshal, _ := sonic.Marshal(getExtData(doc))
				doc.MetaData[common.FieldExtra] = string(marshal)
			}
			return map[string]es8.FieldValue{
				common.FieldContent: {
					Value:    doc.Content,
					EmbedKey: common.FieldContentVector, // 这里也可以考虑注释掉，等后续异步执行
				},
				common.FieldExtra: {
					Value: doc.MetaData[common.FieldExtra],
				},
				common.KnowledgeName: {
					Value: knowledgeName,
				},
				// common.FieldQAContent: {
				// 	Value:    doc.MetaData[common.FieldQAContent],
				// 	EmbedKey: common.FieldQAContentVector,
				// },
			}, nil
		},
	}
	embeddingIns11, err := common.NewEmbedding(ctx, conf)
	if err != nil {
		return nil, err
	}
	indexerConfig.Embedding = embeddingIns11
	idr, err = es8.NewIndexer(ctx, indexerConfig)
	if err != nil {
		return nil, err
	}
	return idr, nil
}

func getMdContentWithTitle(doc *schema.Document) string {
	if doc.MetaData == nil {
		return doc.Content
	}
	title := ""
	list := []string{"h1", "h2", "h3", "h4", "h5", "h6"}
	for _, v := range list {
		if d, e := doc.MetaData[v].(string); e && len(d) > 0 {
			title += fmt.Sprintf("%s:%s ", v, d)
		}
	}
	if len(title) == 0 {
		return doc.Content
	}
	return title + "\n" + doc.Content
}

func getExtData(doc *schema.Document) map[string]any {
	if doc.MetaData == nil {
		return nil
	}
	res := make(map[string]any)
	for _, key := range common.ExtKeys {
		if v, e := doc.MetaData[key]; e {
			res[key] = v
		}
	}
	return res
}
