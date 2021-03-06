/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sync_client

import (
	"context"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/baselib/grpc_util/service_discovery"
	"github.com/nebulaim/telegramd/proto/mtproto"
)

type syncClient struct {
	client mtproto.RPCSyncClient
}

var (
	syncInstance = &syncClient{}
)

func GetSyncClient() *syncClient {
	return syncInstance
}

func InstallSyncClient(discovery *service_discovery.ServiceDiscoveryClientConfig) {
	conn, err := grpc_util.NewRPCClientByServiceDiscovery(discovery)

	if err != nil {
		glog.Error(err)
		panic(err)
	}

	syncInstance.client = mtproto.NewRPCSyncClient(conn)
}

func (c *syncClient) SyncOneUpdateData2(serverId int32, authKeyId, sessionId int64, pushUserId int32, clientMsgId int64, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:    mtproto.SyncType_SYNC_TYPE_RPC_RESULT,
		ServerId:    serverId,
		AuthKeyId:   authKeyId,
		SessionId:   sessionId,
		PushUserId:  pushUserId,
		ClientMsgId: clientMsgId,
		Updates:     updates.To_Updates(),
		RpcResult: &mtproto.RpcResultData{
			AffectedMessages: mtproto.NewTLMessagesAffectedMessages(),
		},
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncOneUpdateData3(serverId int32, authKeyId, sessionId int64, pushUserId int32, clientMsgId int64, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:    mtproto.SyncType_SYNC_TYPE_RPC_RESULT,
		ServerId:    serverId,
		AuthKeyId:   authKeyId,
		SessionId:   sessionId,
		PushUserId:  pushUserId,
		ClientMsgId: clientMsgId,
		Updates:     updates.To_Updates(),
		RpcResult: &mtproto.RpcResultData{
			AffectedHistory: mtproto.NewTLMessagesAffectedHistory(),
		},
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.ClientUpdatesState, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeOneUpdateData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserOneUpdateData(pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{update},
	}}

	m := &mtproto.UpdatesRequest{
		PushType: mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeUpdateShortData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeUpdateShortData(authKeyId, sessionId int64, pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserUpdateShortData(pushUserId int32, update *mtproto.Update) (reply *mtproto.VoidRsp, err error) {
	updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
		Update: update,
	}}

	m := &mtproto.UpdatesRequest{
		PushType: mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates.To_Updates(),
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) SyncUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.ClientUpdatesState, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.SyncUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserNotMeUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_NOTME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserMeUpdatesData(authKeyId, sessionId int64, pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType:   mtproto.SyncType_SYNC_TYPE_USER_ME,
		AuthKeyId:  authKeyId,
		SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) PushToUserUpdatesData(pushUserId int32, updates *mtproto.Updates) (reply *mtproto.VoidRsp, err error) {
	m := &mtproto.UpdatesRequest{
		PushType: mtproto.SyncType_SYNC_TYPE_USER,
		// AuthKeyId:  authKeyId,
		// SessionId:  sessionId,
		PushUserId: pushUserId,
		Updates:    updates,
	}
	reply, err = c.client.PushUpdatesData(context.Background(), m)
	return
}

func (c *syncClient) GetCurrentChannelPts(channelId int32) (pts int32, err error) {
	req := &mtproto.ChannelPtsRequest{
		ChannelId: channelId,
	}
	var ptsId *mtproto.SeqId
	ptsId, err = c.client.GetCurrentChannelPts(context.Background(), req)
	if err == nil {
		pts = ptsId.Pts
	}
	return
}

func (c *syncClient) GetUpdateListByGtPts(userId, pts int32) (updateList []*mtproto.Update, err error) {
	req := &mtproto.UserGtPtsUpdatesRequest{
		UserId: userId,
		Pts:    pts,
	}

	var updates *mtproto.Updates
	updates, err = c.client.GetUserGtPtsUpdatesData(context.Background(), req)
	if err == nil {
		updateList = updates.GetData2().GetUpdates()
	}
	return
}

func (c *syncClient) GetChannelUpdateListByGtPts(channelId, pts int32) (updateList []*mtproto.Update, err error) {
	req := &mtproto.ChannelGtPtsUpdatesRequest{
		ChannelId: channelId,
		Pts:       pts,
	}

	var updates *mtproto.Updates
	updates, err = c.client.GetChannelGtPtsUpdatesData(context.Background(), req)
	if err == nil {
		updateList = updates.GetData2().GetUpdates()
	}
	return
}

func (c *syncClient) GetServerUpdatesState(authKeyId int64, userId int32) (state *mtproto.TLUpdatesState, err error) {
	req := &mtproto.UpdatesStateRequest{
		AuthKeyId: authKeyId,
		UserId:    userId,
	}

	var state2 *mtproto.Updates_State
	state2, err = c.client.GetServerUpdatesState(context.Background(), req)
	if err == nil {
		state = state2.To_UpdatesState()
	}
	return
}

func (c *syncClient) UpdateAuthStateSeq(authKeyId int64, pts, qts int32) (err error) {
	req := &mtproto.UpdatesStateRequest{
		AuthKeyId: authKeyId,
		Pts:       pts,
		Qts:       qts,
	}
	_, err = c.client.UpdateUpdatesState(context.Background(), req)
	return
}

