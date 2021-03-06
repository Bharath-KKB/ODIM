//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

package services

import (
	"context"
	eventsproto "github.com/ODIM-Project/ODIM/lib-utilities/proto/events"
	"log"
)

// SubscribeToEMB method will subscribe to respective  event queue of the plugin
func SubscribeToEMB(pluginID string, queueList []string) {
	log.Println("info: subscribing to EMB for plugin ", pluginID)
	events := eventsproto.NewEventsService(Events, Service.Client())
	_, err := events.SubsribeEMB(context.TODO(), &eventsproto.SubscribeEMBRequest{
		PluginID:     pluginID,
		EMBQueueName: queueList,
	})
	if err != nil {
		log.Printf("error subscribing to EMB  %v", err)
	}
	return
}

// DeleteSubscription  calls the event service and delete all subscription realated to that server
func DeleteSubscription(uuid string) (*eventsproto.EventSubResponse, error) {
	req := eventsproto.EventRequest{
		UUID: uuid,
	}
	events := eventsproto.NewEventsService(Events, Service.Client())

	return events.DeleteEventSubscription(context.TODO(), &req)
}
