// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	sdkclient "go.temporal.io/sdk/client"
	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common/persistence/visibility/manager"
	"go.temporal.io/server/common/resource"
	"go.temporal.io/server/service/history/configs"
	"go.temporal.io/server/service/history/events"
	"go.temporal.io/server/service/history/queues"
	"go.temporal.io/server/service/history/shard"
	"go.temporal.io/server/service/history/workflow"
	"go.temporal.io/server/service/worker/archiver"
	"go.uber.org/fx"
)

type (
	HistoryEngineFactoryParams struct {
		fx.In

		VisibilityMgr           manager.VisibilityManager
		MatchingClient          resource.MatchingClient
		HistoryClient           historyservice.HistoryServiceClient
		PublicClient            sdkclient.Client
		EventNotifier           events.Notifier
		Config                  *configs.Config
		ReplicationTaskFetchers ReplicationTaskFetchers
		RawMatchingClient       resource.MatchingRawClient
		NewCacheFn              workflow.NewCacheFn
		ArchivalClient          archiver.Client
		QueueProcessorFactories []queues.ProcessorFactory `group:"queueProcessorFactory"`
	}

	historyEngineFactory struct {
		HistoryEngineFactoryParams
	}
)

func (f *historyEngineFactory) CreateEngine(
	context shard.Context,
) shard.Engine {
	return NewEngineWithShardContext(
		context,
		f.VisibilityMgr,
		f.MatchingClient,
		f.HistoryClient,
		f.PublicClient,
		f.EventNotifier,
		f.Config,
		f.ReplicationTaskFetchers,
		f.RawMatchingClient,
		f.NewCacheFn,
		f.ArchivalClient,
		f.QueueProcessorFactories,
	)
}
