// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package rag

import (
	"context"

	"github.com/wangle201210/go-rag/server/api/rag/v1"
)

type IRagV1 interface {
	Chat(ctx context.Context, req *v1.ChatReq) (res *v1.ChatRes, err error)
	ChatStream(ctx context.Context, req *v1.ChatStreamReq) (res *v1.ChatStreamRes, err error)
	ChunksList(ctx context.Context, req *v1.ChunksListReq) (res *v1.ChunksListRes, err error)
	ChunkDelete(ctx context.Context, req *v1.ChunkDeleteReq) (res *v1.ChunkDeleteRes, err error)
	UpdateChunk(ctx context.Context, req *v1.UpdateChunkReq) (res *v1.UpdateChunkRes, err error)
	UpdateChunkContent(ctx context.Context, req *v1.UpdateChunkContentReq) (res *v1.UpdateChunkContentRes, err error)
	DocumentsList(ctx context.Context, req *v1.DocumentsListReq) (res *v1.DocumentsListRes, err error)
	DocumentsDelete(ctx context.Context, req *v1.DocumentsDeleteReq) (res *v1.DocumentsDeleteRes, err error)
	Indexer(ctx context.Context, req *v1.IndexerReq) (res *v1.IndexerRes, err error)
	KBCreate(ctx context.Context, req *v1.KBCreateReq) (res *v1.KBCreateRes, err error)
	KBUpdate(ctx context.Context, req *v1.KBUpdateReq) (res *v1.KBUpdateRes, err error)
	KBDelete(ctx context.Context, req *v1.KBDeleteReq) (res *v1.KBDeleteRes, err error)
	KBGetOne(ctx context.Context, req *v1.KBGetOneReq) (res *v1.KBGetOneRes, err error)
	KBGetList(ctx context.Context, req *v1.KBGetListReq) (res *v1.KBGetListRes, err error)
	Retriever(ctx context.Context, req *v1.RetrieverReq) (res *v1.RetrieverRes, err error)
	RetrieverDify(ctx context.Context, req *v1.RetrieverDifyReq) (res *v1.RetrieverDifyRes, err error)
}
